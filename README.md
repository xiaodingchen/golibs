## Go组件
### Config 配置组件
`Config` 组件支持多配置文件场景，具体可以参考 `examples`目录下的使用示例。
配置组件使用了`viper`库，在加载完配置后，可以直接通过库本身提供的方法读取配置项，不建议对库再次封装，它已经做得够好了。组件所支持的文件格式：`toml`, `json`, `yml`, `properties`

### Logger 日志组件
logger组件是对 `zerolog`的简单封装，实现了`grpclog.LoggerV2`接口。具体可以参考 `examples`目录下的使用示例。文档地址：https://godoc.org/github.com/xiaodingchen/golibs/logger

