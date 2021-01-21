/**
  @author: xinyulu
  @date: 2021/1/21 11:18
  @note: 
**/
package client

import (
    "cloudbatch/conf"
    "github.com/jdcloud-api/jdcloud-sdk-go/core"
)

var credentials *core.Credential
var config *core.Config
var logger core.Logger

// 初始化jdcloud-sdk client
func clientInit() {
    credentials = core.NewCredentials(conf.GetAccessKey(), conf.GetSecretKey())
    config = core.NewConfig()
    config.SetTimeout(conf.GetTimeout())
    config.SetScheme(conf.GetScheme())

    logger = core.NewDefaultLogger(conf.GetSdkLogLevel())
}

func init() {
    clientInit()
}
