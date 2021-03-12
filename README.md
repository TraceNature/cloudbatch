# cloudbatch

## 快速开始

以批量添加waf域名为例:
1. 在家目录中创建`cloudbatch/config.yaml`配置文件
```yaml
accessKey: xxx # accessKey
secretKey: xxx # secretKey
scheme: https # 调用openapi的协议类型: https/http
timeout: 20 # 调用openapi超时时间
sdkLogLevel: 2 # 调用openapi的SDK日志级别: 
internal: false # 是否通过内网调用openapi
logConfig:
  loglevel: info # cloudbatch日志
  output: ./output.log # 日志文件路径
```
2. 编辑Excel配置文件，模板参考config_example/waf.xlsx
3. cmd命令下执行cloudbatch命令
```bash
cloudbatch.exe waf adddomains --regionId cn-norht-1 --file waf.xlsx
```