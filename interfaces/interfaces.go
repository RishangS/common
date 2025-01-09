package interfaces

type AppData struct {
	ID                string       `json:"id"`
	Locale            string       `json:"locale"`
	Country           string       `json:"country"`
	URL               string       `json:"url"`
	FullURL           string       `json:"fullUrl"`
	Name              string       `json:"name"`
	Description       string       `json:"description"`
	Summary           string       `json:"summary"`
	Icon              string       `json:"icon"`
	Screenshots       []string     `json:"screenshots"`
	Score             float64      `json:"score"`
	PriceText         *string      `json:"priceText"`
	IsFree            bool         `json:"isFree"`
	Installs          int64        `json:"installs"`
	InstallsText      string       `json:"installsText"`
	AppVersion        string       `json:"appVersion"`
	AndroidVersion    string       `json:"androidVersion"`
	MinAndroidVersion string       `json:"minAndroidVersion"`
	Size              *string      `json:"size"`
	ContentRating     string       `json:"contentRating"`
	PrivacyPolicyUrl  string       `json:"privacyPolicyUrl"`
	Category          AppCategory  `json:"category"`
	HistogramRating   AppHistogram `json:"histogramRating"`
	Reviews           []AppReview  `json:"reviews"`
	NumberVoters      int64        `json:"numberVoters"`
	NumberReviews     int          `json:"numberReviews"`
	RecentChanges     string       `json:"recentChanges"`
	EditorsChoice     bool         `json:"editorsChoice"`
	Released          string       `json:"released"`
	ReleasedTimestamp int64        `json:"releasedTimestamp"`
	Updated           string       `json:"updated"`
	UpdatedTimestamp  int64        `json:"updatedTimestamp"`
	Developer         AppDeveloper `json:"developer"`
	DeveloperName     string       `json:"developerName"`
	AvailableLocales  []string     `json:"availableLocales"`
}

type AppCategory struct {
	ID                    string `json:"id"`
	Name                  string `json:"name"`
	IsApplicationCategory bool   `json:"isApplicationCategory"`
	IsGamesCategory       bool   `json:"isGamesCategory"`
}

type AppHistogram struct {
	Five  int `json:"five"`
	Four  int `json:"four"`
	Three int `json:"three"`
	Two   int `json:"two"`
	One   int `json:"one"`
}

type AppReview struct {
	ID         string    `json:"id"`
	UserName   string    `json:"userName"`
	Text       string    `json:"text"`
	Avatar     string    `json:"avatar"`
	AppVersion string    `json:"appVersion"`
	Date       string    `json:"date"`
	Timestamp  int64     `json:"timestamp"`
	Score      int       `json:"score"`
	CountLikes int       `json:"countLikes"`
	Reply      *AppReply `json:"reply,omitempty"`
}

type AppReply struct {
	Text      string `json:"text"`
	Date      string `json:"date"`
	Timestamp int64  `json:"timestamp"`
}

type AppDeveloper struct {
	ID      string  `json:"id"`
	URL     string  `json:"url"`
	Name    string  `json:"name"`
	Website string  `json:"website"`
	Email   string  `json:"email"`
	Address string  `json:"address"`
	Icon    *string `json:"icon"`
	Cover   *string `json:"cover"`
}

type FailedAppID struct {
	AppID    string `bson:"app_id"`
	ErrorMsg string `bson:"error_message"`
	Status   string `bson:"status"`
}
