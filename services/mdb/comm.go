package mdb

import (
	"context"
	"demo/models"
	"demo/services/zlog"
	"fmt"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
	"go.uber.org/zap"
	"time"
)

const moduleName = "mongo"

var db string
var client *mongo.Client

func Start() {
	zlog.Warn(moduleName, "start ...")
	loadConfig()
	db = models.GetDB()
	uri := fmt.Sprintf("%v://%v:%v@%v", scheme, username, password, address)
	opts := options.Client().ApplyURI(uri)
	if c, err := mongo.Connect(context.Background(), opts); err != nil {
		zlog.Fatal(moduleName, "connect", zap.Error(err))
	} else {
		client = c
	}
	ctx, cancel := context.WithTimeout(context.Background(), time.Second*3)
	defer cancel()
	if err := client.Ping(ctx, nil); err != nil {
		zlog.Fatal(moduleName, "ping", zap.Error(err))
	}
}

func Stop() {
	zlog.Warn(moduleName, "stop ...")
	if err := client.Disconnect(context.Background()); err != nil {
		zlog.Error(moduleName, "disconnect", zap.Error(err))
		return
	}
}

func Collection(name string, opts ...*options.CollectionOptions) *mongo.Collection {
	return client.Database(db).Collection(name, opts...)
}
