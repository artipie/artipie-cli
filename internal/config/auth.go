package config

import "errors"

// Auth represents authentification options.
// TODO Auth options should generate valid header for request
type Auth interface {
	validate() error
}

// ErrInvalidAuth occurs when authentification options is invalid.
var ErrInvalidAuth = errors.New("auth options must be either user+password or token")

// AuthBasic is for authentification via username and password.
type AuthBasic struct {
	UserName string `yaml:"username"`
	Password string `yaml:"password"`
}

func (a AuthBasic) validate() error {
	if a.UserName == "" {
		return &ErrConfigFieldEmpty{"username"}
	}
	if a.Password == "" {
		return &ErrConfigFieldEmpty{"password"}
	}
	return nil
}

// AuthToken is for authentification via token.
type AuthToken struct {
	Token string `yaml:"token"`
}

func (a AuthToken) validate() error {
	if a.Token == "" {
		return &ErrConfigFieldEmpty{"token"}
	}
	return nil
}
