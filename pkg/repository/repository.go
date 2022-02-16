package repository

import (
	"src/pkg/domain"
)

func GetTweetsRepo() string {
	return "tweet retornado"
}

func CreateTweetRepo(tweetUser domain.UserTweet) domain.UserTweet {
	return tweetUser
}
