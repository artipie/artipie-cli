package config

import (
	"errors"
	"net/http"
)

// Auth represents authentification options.
type Auth interface {
	validate() error
	SetAuthHeader(req *http.Request)
}

// ErrInvalidAuth occurs when authentification options is invalid.
var ErrInvalidAuth = errors.New("auth options must be either user+password or token")

// AuthBasic is for authentification via username and password.
type AuthBasic struct {
	UserName string `yaml:"username"`
	Password string `yaml:"password"`
}

func (a *AuthBasic) validate() error {
	if a.UserName == "" {
		return &ErrConfigFieldEmpty{"username"}
	}
	if a.Password == "" {
		return &ErrConfigFieldEmpty{"password"}
	}
	return nil
}

func (a *AuthBasic) SetAuthHeader(req *http.Request) {
	req.SetBasicAuth(a.UserName, a.Password)
}

// AuthToken is for authentification via token.
type AuthToken struct {
	Token string `yaml:"token"`
}

func (a *AuthToken) validate() error {
	if a.Token == "" {
		return &ErrConfigFieldEmpty{"token"}
	}
	return nil
}

func (a *AuthToken) SetAuthHeader(req *http.Request) {
	var bearer = "Bearer " + a.Token
	req.Header.Add("Authorization", bearer)
}
