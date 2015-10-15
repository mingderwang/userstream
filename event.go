package userstream

import (
	"encoding/json"
	//	"github.com/davecgh/go-spew/spew"
	"github.com/k0kubun/twitter"
	"strings"
)

type Record struct {
	DirectMessage struct {
		ID     int64  `json:"id"`
		IDStr  string `json:"id_str"`
		Text   string `json:"text"`
		Sender struct {
			ID          int    `json:"id"`
			IDStr       string `json:"id_str"`
			Name        string `json:"name"`
			ScreenName  string `json:"screen_name"`
			Location    string `json:"location"`
			Description string `json:"description"`
			URL         string `json:"url"`
			Entities    struct {
				URL struct {
					Urls []struct {
						URL         string `json:"url"`
						ExpandedURL string `json:"expanded_url"`
						DisplayURL  string `json:"display_url"`
						Indices     []int  `json:"indices"`
					} `json:"urls"`
				} `json:"url"`
				Description struct {
					Urls []interface{} `json:"urls"`
				} `json:"description"`
			} `json:"entities"`
			Protected                      bool   `json:"protected"`
			FollowersCount                 int    `json:"followers_count"`
			FriendsCount                   int    `json:"friends_count"`
			ListedCount                    int    `json:"listed_count"`
			CreatedAt                      string `json:"created_at"`
			FavouritesCount                int    `json:"favourites_count"`
			UtcOffset                      int    `json:"utc_offset"`
			TimeZone                       string `json:"time_zone"`
			GeoEnabled                     bool   `json:"geo_enabled"`
			Verified                       bool   `json:"verified"`
			StatusesCount                  int    `json:"statuses_count"`
			Lang                           string `json:"lang"`
			ContributorsEnabled            bool   `json:"contributors_enabled"`
			IsTranslator                   bool   `json:"is_translator"`
			IsTranslationEnabled           bool   `json:"is_translation_enabled"`
			ProfileBackgroundColor         string `json:"profile_background_color"`
			ProfileBackgroundImageURL      string `json:"profile_background_image_url"`
			ProfileBackgroundImageURLHTTPS string `json:"profile_background_image_url_https"`
			ProfileBackgroundTile          bool   `json:"profile_background_tile"`
			ProfileImageURL                string `json:"profile_image_url"`
			ProfileImageURLHTTPS           string `json:"profile_image_url_https"`
			ProfileBannerURL               string `json:"profile_banner_url"`
			ProfileLinkColor               string `json:"profile_link_color"`
			ProfileSidebarBorderColor      string `json:"profile_sidebar_border_color"`
			ProfileSidebarFillColor        string `json:"profile_sidebar_fill_color"`
			ProfileTextColor               string `json:"profile_text_color"`
			ProfileUseBackgroundImage      bool   `json:"profile_use_background_image"`
			HasExtendedProfile             bool   `json:"has_extended_profile"`
			DefaultProfile                 bool   `json:"default_profile"`
			DefaultProfileImage            bool   `json:"default_profile_image"`
			Following                      bool   `json:"following"`
			FollowRequestSent              bool   `json:"follow_request_sent"`
			Notifications                  bool   `json:"notifications"`
		} `json:"sender"`
		SenderID         int    `json:"sender_id"`
		SenderIDStr      string `json:"sender_id_str"`
		SenderScreenName string `json:"sender_screen_name"`
		Recipient        struct {
			ID          int64       `json:"id"`
			IDStr       string      `json:"id_str"`
			Name        string      `json:"name"`
			ScreenName  string      `json:"screen_name"`
			Location    string      `json:"location"`
			Description string      `json:"description"`
			URL         interface{} `json:"url"`
			Entities    struct {
				Description struct {
					Urls []interface{} `json:"urls"`
				} `json:"description"`
			} `json:"entities"`
			Protected                      bool        `json:"protected"`
			FollowersCount                 int         `json:"followers_count"`
			FriendsCount                   int         `json:"friends_count"`
			ListedCount                    int         `json:"listed_count"`
			CreatedAt                      string      `json:"created_at"`
			FavouritesCount                int         `json:"favourites_count"`
			UtcOffset                      interface{} `json:"utc_offset"`
			TimeZone                       interface{} `json:"time_zone"`
			GeoEnabled                     bool        `json:"geo_enabled"`
			Verified                       bool        `json:"verified"`
			StatusesCount                  int         `json:"statuses_count"`
			Lang                           string      `json:"lang"`
			ContributorsEnabled            bool        `json:"contributors_enabled"`
			IsTranslator                   bool        `json:"is_translator"`
			IsTranslationEnabled           bool        `json:"is_translation_enabled"`
			ProfileBackgroundColor         string      `json:"profile_background_color"`
			ProfileBackgroundImageURL      string      `json:"profile_background_image_url"`
			ProfileBackgroundImageURLHTTPS string      `json:"profile_background_image_url_https"`
			ProfileBackgroundTile          bool        `json:"profile_background_tile"`
			ProfileImageURL                string      `json:"profile_image_url"`
			ProfileImageURLHTTPS           string      `json:"profile_image_url_https"`
			ProfileLinkColor               string      `json:"profile_link_color"`
			ProfileSidebarBorderColor      string      `json:"profile_sidebar_border_color"`
			ProfileSidebarFillColor        string      `json:"profile_sidebar_fill_color"`
			ProfileTextColor               string      `json:"profile_text_color"`
			ProfileUseBackgroundImage      bool        `json:"profile_use_background_image"`
			HasExtendedProfile             bool        `json:"has_extended_profile"`
			DefaultProfile                 bool        `json:"default_profile"`
			DefaultProfileImage            bool        `json:"default_profile_image"`
			Following                      bool        `json:"following"`
			FollowRequestSent              bool        `json:"follow_request_sent"`
			Notifications                  bool        `json:"notifications"`
		} `json:"recipient"`
		RecipientID         int64  `json:"recipient_id"`
		RecipientIDStr      string `json:"recipient_id_str"`
		RecipientScreenName string `json:"recipient_screen_name"`
		CreatedAt           string `json:"created_at"`
		Entities            struct {
			Hashtags     []interface{} `json:"hashtags"`
			Symbols      []interface{} `json:"symbols"`
			UserMentions []interface{} `json:"user_mentions"`
			Urls         []interface{} `json:"urls"`
		} `json:"entities"`
	} `json:"direct_message"`
}

type FriendList struct {
	Friends []int64
}

type Tweet twitter.Tweet

type Delete struct {
	Id     int64
	UserId int64 `json:"user_id"`
}

type Favorite struct {
	Source       *twitter.User
	Target       *twitter.User
	TargetObject *twitter.Tweet `json:"target_object"`
}
type Unfavorite Favorite

type Follow struct {
	Source *twitter.User
	Target *twitter.User
}
type Unfollow Follow

type ListMemberAdded struct {
	Source       *twitter.User
	Target       *twitter.User
	TargetObject *twitter.List `json:"target_object"`
}
type ListMemberRemoved ListMemberAdded

func ParseJson(jsonText string) interface{} {
	hash := map[string]string{}
	decoder := json.NewDecoder(strings.NewReader(jsonText))
	decoder.Decode(&hash)

	decoder = json.NewDecoder(strings.NewReader(jsonText))
	if _, hasKey := hash["friends"]; hasKey {
		friendList := FriendList{}
		decoder.Decode(&friendList)
		return &friendList
	} else if _, hasKey := hash["event"]; hasKey {
		return parseEvent(decoder, hash["event"])
	} else if _, hasKey := hash["delete"]; hasKey {
		deleteHash := map[string]map[string]*Delete{}
		decoder.Decode(&deleteHash)
		return deleteHash["delete"]["status"]
	} else if _, hasKey := hash["created_at"]; hasKey {
		tweet := twitter.Tweet{}
		decoder.Decode(&tweet)
		return &tweet
	} else if _, hasKey := hash["direct_message"]; hasKey {
		dm := Record{}
		decoder := json.NewDecoder(strings.NewReader(jsonText))
		decoder.Decode(&dm)
		//spew.Dump(dm)
		return &dm
	}
	return nil
}

func parseEvent(decoder *json.Decoder, eventName string) interface{} {
	switch eventName {
	case "favorite":
		favorite := Favorite{}
		decoder.Decode(&favorite)
		return &favorite
	case "unfavorite":
		unfavorite := Unfavorite{}
		decoder.Decode(&unfavorite)
		return &unfavorite
	case "follow":
		follow := Follow{}
		decoder.Decode(&follow)
		return &follow
	case "unfollow":
		unfollow := Unfollow{}
		decoder.Decode(&unfollow)
		return &unfollow
	case "list_member_added":
		listMemberAdded := ListMemberAdded{}
		decoder.Decode(&listMemberAdded)
		return &listMemberAdded
	case "list_member_removed":
		listMemberRemoved := ListMemberRemoved{}
		decoder.Decode(&listMemberRemoved)
		return &listMemberRemoved
	}
	return nil
}
