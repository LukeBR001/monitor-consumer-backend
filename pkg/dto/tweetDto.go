package dto

import (
	"src/pkg/domain"
	"time"
)

type TweetDto struct {
	Id                      string
	Name                    string
	ScreenName              string
	Location                time.Location
	Description             string
	ProfileImageURL         string
	MiniProfileImageURL     string
	OriginalProfileImageURL string
	FriendsCount            int
	CreatedAt               string
	TimeZone                string
}

func (t TweetDto) toTweetDomain() domain.UserTweet {
	return domain.UserTweet{
		Id:                      t.Id,
		Name:                    t.Name,
		ScreenName:              t.ScreenName,
		Location:                t.Location,
		Description:             t.Description,
		ProfileImageURL:         t.ProfileImageURL,
		MiniProfileImageURL:     t.MiniProfileImageURL,
		OriginalProfileImageURL: t.OriginalProfileImageURL,
		FriendsCount:            t.FriendsCount,
		CreatedAt:               t.CreatedAt,
		TimeZone:                t.TimeZone,
	}
}
