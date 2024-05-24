package config

import (
	"os"

	"golang.org/x/oauth2"
	"golang.org/x/oauth2/google"
	"gopkg.in/yaml.v3"
)

type Config struct {
	Port  string `yaml:"port"`
	Auths struct {
		Google struct {
			ClientID     string        `yaml:"client_id"`
			ClientSecret string        `yaml:"client_secret"`
			RedirectURL  string        `yaml:"redirect_url"`
			OAuthConfig  oauth2.Config `yaml:"-"`
		} `yaml:"google"`
	} `yaml:"auths"`
}

func parseFile(path string) Config {
	c := Config{}
	contentByte, err := os.ReadFile(path)
	if err != nil {
		panic("Failed to open file")
	}

	err = yaml.Unmarshal([]byte(contentByte), &c)
	if err != nil {
		panic("Failed to unmarschal config")
	}
	return c
}

func (c Config) withGoogleAuth() Config {
	c.Auths.Google.OAuthConfig = oauth2.Config{
		ClientID:     c.Auths.Google.ClientID,
		ClientSecret: c.Auths.Google.ClientSecret,
		RedirectURL:  c.Auths.Google.RedirectURL,
		Scopes: []string{
			"https://www.googleapis.com/auth/userinfo.email",
		},
		Endpoint: google.Endpoint,
	}

	return c
}

func NewConfig(path string) Config {
	c := parseFile(path)
	c = c.withGoogleAuth()
	return c
}
