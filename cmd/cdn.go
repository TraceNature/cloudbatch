package cmd

import (
	"cloudbatch/client"
	"cloudbatch/log"
	"strconv"
	"strings"

	"github.com/360EntSecGroup-Skylar/excelize/v2"
	cdnApis "github.com/jdcloud-api/jdcloud-sdk-go/services/cdn/apis"
	cdnModels "github.com/jdcloud-api/jdcloud-sdk-go/services/cdn/models"
	"github.com/spf13/cobra"
)


func NewCdnCommand() *cobra.Command {
    cmd := &cobra.Command{
        Use:   "cdn",
        Short: "cdn batch",
    }
    cmd.AddCommand(NewCreateCdnDomain())
    return cmd
}

func NewCreateCdnDomain() *cobra.Command {
    cmd := &cobra.Command{
        Use:   "createdomains",
        Short: "create cdn domains",
        Run:   createCdnDomainFunc,
    }
    cmd.Flags().String("file", "resource_config.xlsx", "cdn domain file path")
    return cmd
}

// 批量添加CDN域名参数列表长度
const createCdnDomainMaxParameter = 7

// 将CDN添加域名参数转换为接口参数
func excelCdn2interface(row []string) *cdnApis.CreateDomainRequest {
    if len(row) != createCdnDomainMaxParameter {
        log.Error("row length is invalid")
        return nil
    }
    domain := row[1]
    req := cdnApis.NewCreateDomainRequest(domain)
    // 映射业务类型
    switch row[0] {
    case "图片小文件":
        req.SetCdnType("web")
    case "大文件下载":
        req.SetCdnType("download")
    case "视频文件":
        req.SetCdnType("vod")
    default:
        req.SetCdnType("web")
    }
    // 映射回源方式
    switch row[2] {
    case "OSS回源":
        req.SetSourceType("oss")
    case "域名回源":
        req.SetSourceType("domain")
    case "IP回源":
        req.SetSourceType("ips")
    default:
        req.SetSourceType("oss")
    }
    // 映射回源地址
    switch *req.SourceType {
    case "oss":
        req.SetOssSource(row[3])
    case "domain":
        // 将第四列数据按分号分隔，每一组数据为domainSource接口参数的一个数组元素
        domainSourceInfos := strings.Split(row[3], ";")
        iDomainSourceInfos := make([]cdnModels.DomainSourceInfo, 0)
        for _, domainSourceInfo := range domainSourceInfos {
            paras := strings.Split(domainSourceInfo, ",")
            if len(paras) == 3 {
                priority, err := strconv.Atoi(paras[0])
                if err != nil {
                    log.Error("priority is invalid")
                    return nil
                }
                domain := paras[1]
                sourceHost := paras[2]
                iDomainSourceInfos = append(iDomainSourceInfos, cdnModels.DomainSourceInfo{Priority: priority, SourceHost: sourceHost, Domain: domain})
            } else if len(paras) == 2 {
                priority, err := strconv.Atoi(paras[0])
                if err != nil {
                    log.Error("priority is invalid")
                    return nil
                }
                domain := paras[1]
                iDomainSourceInfos = append(iDomainSourceInfos, cdnModels.DomainSourceInfo{Priority: priority, Domain: domain})

            } else {
                log.Error("回源地址参数不合法，请检查回源地址字段")
                return nil
            }
        }
        req.SetDomainSource(iDomainSourceInfos)
    case "ips":
        ipSourceInfos := strings.Split(row[3], ";")
        iIpSourceInfos := make([]cdnModels.IpSourceInfo, 0)
        for _, ipSourceInfo := range ipSourceInfos {
            paras := strings.Split(ipSourceInfo, ",")
            if len(paras) != 3 {
                log.Error("回源地址参数不合法，请检查回源地址字段")
                return nil
            }
            master, err := strconv.Atoi(paras[0])
            if err != nil {
                log.Error("master is invalid")
                return nil
            }
            ip := paras[1]
            ratio, err := strconv.Atoi(paras[2])
            if err != nil {
                log.Error("ratio is invalid")
                return nil
            }
            iIpSourceInfos = append(iIpSourceInfos, cdnModels.IpSourceInfo{Master: master, Ip: ip, Ratio: float64(ratio)})
        }
        req.SetIpSource(iIpSourceInfos)
    default:
        req.SetOssSource(row[3])
    }
    // 映射默认回源Host
    if len(row[4]) != 0 {
        req.SetDefaultSourceHost(row[4])
    }
    // 映射源站端口
    switch row[5] {
    case "443":
        req.SetBackSourceType("https")
    case "80":
        req.SetBackSourceType("http")
    }
    
    // 映射加速区域
    switch row[6] {
    case "全球":
        req.SetAccelerateRegion("all")
    case "中国境内":
        req.SetAccelerateRegion("mainLand")
    case "中国境外":
        req.SetAccelerateRegion("nonMainLand")
    default:
        req.SetAccelerateRegion("mainLand")
    }
    req.SetHttpType("http")
    
    return req
}

func cdnCreateDomainRequestSend(req *cdnApis.CreateDomainRequest) bool {
    log.Info("", log.Field("req", req))
    cdnclient := client.GetCdnClient()
    resp, err := cdnclient.CreateDomain(req)
    if err != nil {
        log.Error(err.Error())
        return false
    }
    if resp.Error.Code != 0 {
        log.Error("", log.Field("resp", resp))
        return false
    } else {
        log.Info("", log.Field("resp", resp))
    }
    return true
}

func createCdnDomainFunc(cmd *cobra.Command, args []string) {
    // 获取命令行传入的xlsx文件
    filePath, _ := cmd.Flags().GetString("file")
    // 读取xlsx文件
    f, err := excelize.OpenFile(filePath)
    if err != nil {
        log.Error(err.Error())
        return
    }
    rows, err := f.GetRows("CDN")
    if err != nil {
        log.Error(err.Error())
        return
    }

    // 参数映射
    for index, row := range rows {
        if index == 0 {
            continue
        }
        req := excelCdn2interface(row)
        success := cdnCreateDomainRequestSend(req)
        if success {
            log.Info(row[1] + ": success!")
            continue
        } else {
            log.Error(row[1] + ": failed")
            return
        } 
    }
}