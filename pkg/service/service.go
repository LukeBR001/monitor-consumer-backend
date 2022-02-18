package service

import (
	"go.mongodb.org/mongo-driver/mongo"
	"src/pkg/dto"
	"src/pkg/repository"
)

func NewService(r repository.Database) Service {
	return Service{
		repository: r,
	}
}

type Service struct {
	repository repository.Database
}

func (s Service) GetTweetsService() string {
	return s.repository.GetTweetsRepo()
}

func (s Service) CreateTweet(tweet dto.TweetDto) *mongo.InsertOneResult {
	tweetDomainUser := dto.TweetDto.ToTweetDomain(tweet)

	return s.repository.CreateTweetRepo(tweetDomainUser)
}
