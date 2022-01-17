package models

import (
	"context"
	"fmt"
	"log"

	"blog_backend/pkg/setting"

	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
)

type DataBase struct {
	db      *mongo.Database
	Article *mongo.Collection
}

var db *DataBase

func init() {
	sec, err := setting.Cfg.GetSection("database")
	if err != nil {
		log.Fatal(2, "Fail to get section 'database': %v", err)
	}
	dbType := sec.Key("TYPE").String()
	dbName := sec.Key("NAME").String()
	host := sec.Key("HOST").String()
	uri := fmt.Sprintf("%s://%s", dbType, host)
	clientOptions := options.Client().ApplyURI(uri)
	client, err := mongo.Connect(context.TODO(), clientOptions)
	if err != nil {
		log.Fatal(2, "Fail to connect %v", err)
	}
	database := client.Database(dbName)
	db = &DataBase{
		Article: database.Collection("article"),
		db:      database,
	}
}

func CloseDB() {
	defer db.db.Client().Disconnect(context.TODO())
}
