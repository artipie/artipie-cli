package config

// Auth represents authentification options.
// TODO Auth options should generate valid header for request
type Auth interface {
	validate() error
}

type errInvalidAuth struct{}

func (e *errInvalidAuth) Error() string {
	return "auth options must be either user+password or token"
}

// AuthBasic is for authentification via username and password.
type AuthBasic struct {
	UserName string `yaml:"username"`
	Password string `yaml:"password"`
}

func (a AuthBasic) validate() error {
	if a.UserName == "" {
		return &errConfigFieldEmpty{"username"}
	}
	if a.Password == "" {
		return &errConfigFieldEmpty{"password"}
	}
	return nil
}

// AuthToken is for authentification via token.
type AuthToken struct {
	Token string `yaml:"token"`
}

func (a AuthToken) validate() error {
	if a.Token == "" {
		return &errConfigFieldEmpty{"token"}
	}
	return nil
}
