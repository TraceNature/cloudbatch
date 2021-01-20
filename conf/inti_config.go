/**
  @author: xinyulu
  @date: 2021/1/19 10:40
  @note: 
**/
package conf

import (
    "cloudbatch/commons"
    "cloudbatch/log"
    "fmt"
    "github.com/spf13/viper"
    "os"
    "path/filepath"
)

var v *viper.Viper

func init() {
    CloudBatchConfigInit()
    logInit()
}

func GetCloudBatchConfig() *viper.Viper {
    return v
}

// log init
func logInit() {
    // 设置日志级别
    loglevel := v.GetString("log_config.loglevel")
    if loglevel != "" {
        log.SetLevel(loglevel)
    }
    // 将log写入指定文件
    logFile := v.GetString("log_config.output")
    if logFile != "" {
        file, err := os.OpenFile(logFile, os.O_CREATE|os.O_APPEND|os.O_WRONLY, 0644)
        if err != nil {
            fmt.Printf("open log file '%s' error:%s\n", logFile, err)
            return
        }
        log.SetOutput(file)
    }
}

// cloudbatch config init
func CloudBatchConfigInit() {
    v = viper.New()
    homeDir, err := commons.HomeDir()
    if err != nil {
        homeDir = "."
    }
    configDir := filepath.Join(homeDir, "cloudbatch")
    configName := "config.yaml"
    configAbs := filepath.Join(configDir, configName)

    v.AddConfigPath(configDir)
    v.SetConfigName(configName)
    v.SetConfigType("yaml")

    // 判断配置文件路径是否存在
    if !commons.FileExists(configDir) {
        // 如果目录不存在，则新建目录
        err = os.MkdirAll(configDir, 0755)
        if err != nil {
            log.Error(configDir + ": NotFound", log.Field("err", err))
            return
        }
        // 新建配置文件
        _, err := os.Create(configAbs)
        if err != nil {
            log.Error(configAbs + " Create Failed", log.Field("err", err))
            return
        }
    } else if !commons.FileExists(configAbs){
        // 新建配置文件
        _, err := os.Create(configAbs)
        if err != nil {
            log.Error(configAbs + " Create Failed", log.Field("err", err))
            return
        }
    }
    if err := v.ReadInConfig(); err != nil {
        if _, ok := err.(viper.ConfigFileNotFoundError); ok {
            log.Error("ConfigFileNotFoundError", log.Field("err", err))
            return
        } else {
            log.Error("other Error", log.Field("err", err))
            return
        }
    }
}
