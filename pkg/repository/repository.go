package repository

import (
	"context"
	"go.mongodb.org/mongo-driver/mongo"
	"src/pkg/domain"
)

type Database struct {
	DB *mongo.Collection
}

func NewRepository(db *mongo.Collection) Database {
	return Database{
		DB: db,
	}
}

func (db Database) GetTweetsRepo() string {
	return "tweet retornado"
}

func (db Database) CreateTweetRepo(tweetUser domain.UserTweet) *mongo.InsertOneResult {
	result, _ := db.DB.InsertOne(context.TODO(), tweetUser)

	return result
}
