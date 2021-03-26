package cmd

import (
	"fmt"
	"testing"

	"github.com/jdcloud-api/jdcloud-sdk-go/core"
	"github.com/jdcloud-api/jdcloud-sdk-go/services/cdn/apis"
	"github.com/jdcloud-api/jdcloud-sdk-go/services/cdn/client"
)

func TestCreateCdnDomain(t *testing.T) {
	accessKey := ""
	secretKey := ""
	credential := core.NewCredentials(accessKey, secretKey)
	cdnclient := client.NewCdnClient(credential)
	cdnclient.DisableLogger()
	domain := "www.xinyulu3344.cn"
	const OK = 0
	req := apis.NewCreateDomainRequest(domain)
	req.SetCdnType("web")
	req.SetSourceType("oss")
	req.SetOssSource("xinyulu3344.s3.cn-north-1.jdcloud-oss.com")
	req.SetDefaultSourceHost("xinyulu3344.s3.cn-north-1.jdcloud-oss.com")
	req.SetBackSourceType("https")
	req.SetAccelerateRegion("all")
	req.SetHttpType("http")
	fmt.Println("开始执行!")
	resp, err := cdnclient.CreateDomain(req)
	fmt.Println("执行完毕! ")
	if err != nil {
		t.Log("error: ", err)
		return
	}
	if resp.Error.Code != OK {
		fmt.Println("Error Code: ", resp.Error.Code, resp.Error.Message, resp.Error.Status)
		return
	}
	fmt.Println("success")
}
