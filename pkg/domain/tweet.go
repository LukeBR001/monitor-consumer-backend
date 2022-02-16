package domain

import "time"

type UserTweet struct {
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
