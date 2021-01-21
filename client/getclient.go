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
    "sync"
)

var vmclient *vmClient.VmClient
var lbclient *lbClient.LbClient
var once sync.Once


func GetVmClient() *vmClient.VmClient {
    once.Do(func() {
        vmclient = vmClient.NewVmClient(credentials)
    })
    if conf.GetInternal() {
        config.SetEndpoint(VM_ENDPOINT_INTERNAL)
    }
    vmclient.SetConfig(config)
    vmclient.SetLogger(logger)
    return vmclient
}

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
