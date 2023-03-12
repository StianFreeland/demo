package mdb

import (
	"demo/services/zlog"
	"github.com/spf13/viper"
	"go.uber.org/zap"
)

var scheme string
var username string
var password string
var address string

func loadConfig() {
	zlog.Warn(moduleName, "load config ...")
	scheme = viper.GetString("mongo.scheme")
	if scheme == "" {
		zlog.Fatal(moduleName, "load config", zap.String("scheme", scheme))
	}
	zlog.Warn(moduleName, "load config", zap.String("scheme", scheme))
	username = viper.GetString("mongo.username")
	if username == "" {
		zlog.Fatal(moduleName, "load config", zap.String("username", username))
	}
	zlog.Warn(moduleName, "load config", zap.String("username", username))
	password = viper.GetString("mongo.password")
	if password == "" {
		zlog.Fatal(moduleName, "load config", zap.String("password", password))
	}
	zlog.Warn(moduleName, "load config", zap.String("password", password))
	address = viper.GetString("mongo.address")
	if address == "" {
		zlog.Fatal(moduleName, "load config", zap.String("address", address))
	}
	zlog.Warn(moduleName, "load config", zap.String("address", address))
}
