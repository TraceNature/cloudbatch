/**
  @author: xinyulu
  @date: 2021/3/1 17:35
  @note:
**/
package cmd

import (
	"cloudbatch/client"
	"cloudbatch/log"
	"strings"

	"github.com/360EntSecGroup-Skylar/excelize/v2"
	"github.com/jdcloud-api/jdcloud-sdk-go/services/waf/apis"
	"github.com/jdcloud-api/jdcloud-sdk-go/services/waf/models"
	"github.com/spf13/cobra"
)

// waf域名配置
type DomainsConfig struct {
    Domains []models.AddDomain `json:"spec"`
}

func NewWafCommand() *cobra.Command {
    cmd := &cobra.Command{
        Use:   "waf",
        Short: "waf batch",
    }
    cmd.AddCommand(NewAddDomain())
    return cmd
}

func NewAddDomain() *cobra.Command {
    cmd := &cobra.Command{
        Use:   "adddomains",
        Short: "add waf domains",
        Run:   addDomainFunc,
    }
    cmd.Flags().String("file", "resource_config.xlsx", "resource configure file path")
    cmd.Flags().String("regionId", "cn-north-1", "waf instances region")
    return cmd
}

// waf域名配置参数个数
const addWafDomainMaxParameter = 13

// 将excel行数据映射为底层接口数据
func excelWaf2interface(row []string) *models.AddDomain {
    // 判断row中的元素个数是否和定义的Excel字段数相等
    if len(row) != addWafDomainMaxParameter {
        log.Error("row length is invalid")
        return nil
    }
    domain := &models.AddDomain{
        RsConfig: &models.RsConfig{},
    }
    // 映射waf实例ID
    domain.WafInstanceId = row[0]
    // 映射域名
    domain.Domain = row[1]
    // 映射协议类型
    protocols := strings.Split(row[2], ",")
    domain.Protocols = protocols
    // 映射开启https强制跳转
    switch row[3] {
    case "是":
        httpsRedirect := 1
        domain.HttpsRedirect = &httpsRedirect
    case "否":
        httpsRedirect := 0
        domain.HttpsRedirect = &httpsRedirect
    }
    // 映射ssl协议类型
    sslProtocols := strings.Split(row[4], ",")
    domain.SslProtocols = sslProtocols
    // 映射加密套件等级
    switch row[5] {
    case "中级":
        suiteLevel := 0
        domain.SuiteLevel = &suiteLevel
    case "高级":
        suiteLevel := 1
        domain.SuiteLevel = &suiteLevel
    case "低级":
        suiteLevel := 2
        domain.SuiteLevel = &suiteLevel
    }
    // 映射回源服务器地址
    rsAddr := strings.Split(row[6], ",")
    domain.RsConfig.RsAddr = rsAddr
    // 映射负载均衡算法
    domain.LbType = row[7]
    // 映射回源服务器端口
    httpsRsPort := strings.Split(row[8], ",")
    domain.RsConfig.HttpsRsPort = httpsRsPort
    // 映射是否已使用代理
    switch row[9] {
    case "是":
        pureClient := 1
        domain.PureClient = &pureClient
    case "否":
        pureClient := 0
        domain.PureClient = &pureClient
    default:
        pureClient := 0
        domain.PureClient = &pureClient
    }
    // 映射是否开启回源长连接
    switch row[10] {
    case "是":
        enableKeepalive := 1
        domain.EnableKeepalive = &enableKeepalive
    case "否":
        enableKeepalive := 0
        domain.EnableKeepalive = &enableKeepalive
    }
    // 映射请求头支持下划线
    switch row[11] {
    case "开启":
        enableUnderscores := 0
        domain.EnableUnderscores = &enableUnderscores
    case "关闭":
        enableUnderscores := 1
        domain.EnableUnderscores = &enableUnderscores
    }
    // 映射回源类型
    switch row[12] {
    case "IP":
        rsType := 0
        domain.RsConfig.RsType = &rsType
    case "域名":
        rsType := 1
        domain.RsConfig.RsType = &rsType
    }
    return domain
}

func addDomainFunc(cmd *cobra.Command, args []string) {
    filePath, _ := cmd.Flags().GetString("file")
    regionId, _ := cmd.Flags().GetString("regionId")
    domainConfig := &DomainsConfig{}
    domainConfig.Domains = make([]models.AddDomain, 0)
    // 读取Excel
    f, err := excelize.OpenFile(filePath)
    if err != nil {
        log.Error(err.Error())
        return
    }
    rows, err := f.GetRows("WAF")
    if err != nil {
        log.Error(err.Error())
        return
    }
    for index, row := range rows {
        // 跳过第一行表头
        if index == 0 {
            continue
        }
        // 将Excel中数据映射为接口数据
        domain := excelWaf2interface(row)
        if domain == nil {
            return
        }
        domainConfig.Domains = append(domainConfig.Domains, *domain)
    }
    // 调用接口
    wafAddDomainRequestSend(regionId, domainConfig)
    // 输出报表
}

func wafAddDomainRequestSend(regionId string, domainConfig *DomainsConfig) {
    wafclient := client.GetWafClient()
    for _, domain := range domainConfig.Domains {
        req := apis.NewAddDomainRequest(regionId, domain.WafInstanceId, &domain)
        log.Info("", log.Field("req", req))
        resp, err := wafclient.AddDomain(req)
        if err != nil {
            log.Error(err.Error())
            return
        }
        if resp.Error.Code != 200 {
            log.Error("", log.Field("error", resp.Error))
            return
        } else {
            log.Info("", log.Field("resp", resp))
        }
    }
}