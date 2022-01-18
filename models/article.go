package models

import (
	"context"
	"time"

	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
)

type Article struct {
	Id         primitive.ObjectID `bson:"_id,omitempty" json:"id" `
	Tag        []string           `bson:"tag" json:"tag"`
	Title      string             `bson:"title" json:"title" binding:"required"`
	Content    string             `bson:"content" json:"content" binding:"required"`
	CreateTime time.Time          `bson:"create_time" json:"create_time"`
	ModifyTime time.Time          `bson:"modify_time" json:"modify_time"`
	PageView   int64              `bson:"page_view" json:"page_view"`
}

type Articles []Article

func (article *Article) FindByID(id string) error {
	p, _ := primitive.ObjectIDFromHex(id)
	err := db.Article.FindOne(context.TODO(), bson.M{"_id": p}).Decode(&article)
	return err
}

func (article *Article) Insert() (*mongo.InsertOneResult, error) {
	article.CreateTime = time.Now()
	article.ModifyTime = time.Now()
	return db.Article.InsertOne(context.TODO(), article)
}

func (article *Article) Delete() (*mongo.DeleteResult, error) {
	return db.Article.DeleteOne(context.TODO(), bson.M{"_id": article.Id})
}

func (article *Article) Update() (*mongo.UpdateResult, error) {
	article.ModifyTime = time.Now()
	return db.Article.UpdateByID(context.TODO(), article.Id, bson.M{"$set": article})
}

func (articles *Articles) FindByPage(options *options.FindOptions) error {
	cursor, err := db.Article.Find(context.TODO(), bson.M{}, options)
	if err != nil {
		return err
	}
	err = cursor.All(context.TODO(), articles)
	return err
}
