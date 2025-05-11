package config

import "github.com/zeromicro/go-zero/zrpc"

type DatabaseConfig struct {
	URI      string
	Host     string
	Port     int
	User     string
	Password string
	Database string
}

type Config struct {
	zrpc.RpcServerConf
	Database DatabaseConfig
}
