package config

import (
	"fmt"
	"strings"
)

type Config struct {
	DbURL           string `json:"db_url"`
	CurrentUserName string `json:"current_user_name"`
}

func (c Config) SetUser(user_name string) {
	c.CurrentUserName = user_name
	write(c)
}

func (c Config) SPrint() string {
	var result []string

	result = append(result, "Config:")
	result = append(result, "-------")
	result = append(result, fmt.Sprintf("db_url: %s", c.DbURL))
	result = append(result, fmt.Sprintf("current_user_name: %s", c.CurrentUserName))

	return strings.Join(result, "\n")
}
