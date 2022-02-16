package service

import (
	"src/pkg/domain"
	"src/pkg/dto"
	"src/pkg/repository"
)

func GetTweetsService() string {
	return repository.GetTweetsRepo()
}

func CreateTweet(tweet dto.TweetDto) domain.UserTweet {
	tweetDomainUser := dto.TweetDto.ToTweetDomain(tweet)

	return repository.CreateTweetRepo(tweetDomainUser)
}
