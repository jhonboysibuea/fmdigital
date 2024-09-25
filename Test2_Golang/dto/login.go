package dto

type LoginResponse struct {
	Status string `json:"status"`
	Login  Login  `json:"result"`
}

type Login struct {
	AccessToken  string `json:"access_token"`
	RefreshToken string `json:"refresh_token"`
}
