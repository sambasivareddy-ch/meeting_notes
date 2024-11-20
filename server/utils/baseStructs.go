package utils

type OAuthResponse struct {
	AccessToken  string  `json:"access_token"`
	ExpiresIn    float64 `json:"expires_in"`
	IdToken      string  `json:"id_token"`
	RefreshToken string  `json:"refresh_token"`
	Scope        string  `json:"scope"`
	TokenType    string  `json:"token_type"`
}
