package defaultmux

// Credentials is a struct to read the username and password from the request body
type Credentials struct {
	Username string `json:"login"`
	Password string `json:"password"`
}

// TokenPair represents JWT tokens given to authenticated user
type TokenPair struct {
	AccessToken  string `json:"accessToken"`
	RefreshToken string `json:"refreshToken"`
}
