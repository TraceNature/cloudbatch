/**
  @author: xinyulu
  @date: 2021/1/21 11:02
  @note: 
**/
package client

import (
    "cloudbatch/conf"
    lbClient "github.com/jdcloud-api/jdcloud-sdk-go/services/lb/client"
    vmClient "github.com/jdcloud-api/jdcloud-sdk-go/services/vm/client"
    wafClient "github.com/jdcloud-api/jdcloud-sdk-go/services/waf/client"
    cdnClient "github.com/jdcloud-api/jdcloud-sdk-go/services/cdn/client"
    "sync"
)

var vmclient *vmClient.VmClient
var lbclient *lbClient.LbClient
var wafclient *wafClient.WafClient
var cdnclient *cdnClient.CdnClient
var once sync.Once

// 获取vm client
func GetVmClient() *vmClient.VmClient {
    once.Do(func() {
        vmclient = vmClient.NewVmClient(credentials)
    })
    // 如果配置为true，则可以通过VPC内网调用
    if conf.GetInternal() {
        config.SetEndpoint(VM_ENDPOINT_INTERNAL)
    }
    vmclient.SetConfig(config)
    vmclient.SetLogger(logger)
    return vmclient
}

// 获取waf client
func GetWafClient() *wafClient.WafClient {
    once.Do(func() {
        wafclient = wafClient.NewWafClient(credentials)
    })
    if conf.GetInternal() {
        config.SetEndpoint(WAF_ENDPOINT_INTERNAL)
    }
    wafclient.SetConfig(config)
    wafclient.SetLogger(logger)
    return wafclient
}

// 获取Lb client
func GetLbClient() *lbClient.LbClient {
    once.Do(func() {
        lbclient = lbClient.NewLbClient(credentials)
    })
    if conf.GetInternal() {
        config.SetEndpoint(LB_ENDPOINT_INTERNAL)
    }
    lbclient.SetConfig(config)
    lbclient.SetLogger(logger)
    return lbclient
}

// 获取cdn client
func GetCdnClient() *cdnClient.CdnClient {
    once.Do(func() {
        cdnclient = cdnClient.NewCdnClient(credentials)
    })
    if conf.GetInternal() {
        config.SetEndpoint(LB_ENDPOINT_INTERNAL)
    }
    cdnclient.SetConfig(config)
    cdnclient.SetLogger(logger)
    return cdnclient
}