package models

type TokensWithUser struct {
	Token *TokenDetails `json:"token"`
	User  *User         `json:"user"`
}

type AccessDetails struct {
	AccessUuid string
	UserId     string
}

type TokenDetails struct {
	AccessToken  string `json:"accessToken"`
	RefreshToken string `json:"refreshToken"`
	AccessUuid   string
	RefreshUuid  string
	AtExpires    int64
	RtExpires    int64
}
