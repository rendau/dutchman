package entities

// auth

type ProfileAuthReqSt struct {
	Password string `json:"password"`
}

type ProfileAuthRepSt struct {
	RefreshToken string `json:"refresh_token"`
	AccessToken  string `json:"access_token"`
}

// auth by refresh token

type ProfileAuthByRefreshTokenReqSt struct {
	RefreshToken string `json:"refresh_token"`
}

type ProfileAuthByRefreshTokenRepSt struct {
	AccessToken string `json:"access_token"`
}

// profile

type ProfileSt struct {
}
