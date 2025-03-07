# Go MSA Arch

Scalable Go Microservice Architecture Demo

```text
.
├── cmd
│   └── server
│       └── main.go               // application entrypoint
├── go.mod
├── go.sum
└── internal
    └── app                       // application package
        ├── app.go                // inits and runs applications
        ├── read_config.go        // reads config using viper
        ├── read_config_test.go
        └── server
            ├── config.go
            ├── factory.go
            └── server.go
```

## Dependency

| Lib | Desc |
| --- | --- |
| [go.uber.org/zap](https://github.com/uber-go/zap) | Fast, Production-ready logger library, subjectively believed to be the best |
| [spf13/viper](https://github.com/spf13/viper) | Complete configuration solution. Helps to parse config file and unmarshal to go struct |
