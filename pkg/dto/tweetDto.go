package dto

import (
	"src/pkg/domain"
	"time"
)

type TweetDto struct {
	Id                      string        `json:"id"`
	Name                    string        `json:"name"`
	ScreenName              string        `json:"ScreenName"`
	Location                time.Location `json:"location"`
	Description             string        `json:"description"`
	ProfileImageURL         string
	MiniProfileImageURL     string
	OriginalProfileImageURL string
	FriendsCount            int
	CreatedAt               string
	TimeZone                string
}

func (t TweetDto) ToTweetDomain() domain.UserTweet {
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
