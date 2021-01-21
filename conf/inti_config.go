/**
  @author: xinyulu
  @date: 2021/1/19 10:40
  @note: 
**/
package conf

import (
    "cloudbatch/commons"
    "cloudbatch/log"
    "errors"
    "fmt"
    "github.com/spf13/viper"
    "os"
    "path/filepath"
    "time"
)

var v *viper.Viper

func init() {
    err := CloudBatchConfigInit()
    if err != nil {
        os.Exit(1)
    }
    logInit()
}

func GetCloudBatchConfig() *viper.Viper {
    return v
}

// log init
func logInit() {
    // 设置日志级别
    loglevel := v.GetString("logConfig.loglevel")
    if loglevel != "" {
        log.SetLevel(loglevel)
    }
    // 将log写入指定文件
    logFile := v.GetString("logConfig.output")
    if logFile != "" {
        file, err := commons.CreateFile(logFile)
        if err != nil {
            fmt.Printf("open log file '%s' error:%s\n", logFile, err)
            return
        }
        log.SetOutput(file)
    }
}

// cloudbatch config init
func CloudBatchConfigInit() error {
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

    // 判断配置文件是否存在, 不存在则退出
    if !commons.FileExists(configAbs) {
        log.Error("config file not found!! file path: " + configAbs)
        return errors.New("config file not found!! file path: " + configAbs)
    }
    if err := v.ReadInConfig(); err != nil {
        if _, ok := err.(viper.ConfigFileNotFoundError); ok {
            log.Error("ConfigFileNotFoundError", log.Field("err", err))
            return err
        } else {
            log.Error("other Error", log.Field("err", err))
            return err
        }
    }
    return nil
}

func GetAccessKey() string {
    return v.GetString("accessKey")
}

func GetSecretKey() string {
    return v.GetString("secretKey")
}

func GetTimeout() time.Duration {
    return v.GetDuration("timeout") * time.Second
}

func GetSdkLogLevel() int {
    return v.GetInt("sdkLogLevel")
}

func GetScheme() string {
    return v.GetString("scheme")
}

func GetInternal() bool {
    return v.GetBool("internal")
}