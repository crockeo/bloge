package models

import (
	"errors"
	"net/http"
	"strings"
)

type Auth struct {
	Id       int64
	Username string
	Password string
}

func AuthFromString(input string) (Auth, error) {
	sinput := strings.Split(input, "|")

	if len(sinput) == 2 {
		return Auth{
			Username: sinput[0],
			Password: sinput[1],
		}, nil
	} else {
		return Auth{}, errors.New("Could not parse.")
	}
}

func (this *Auth) String() string {
	return this.Username + "|" + this.Password
}

func (this *Auth) Equals(other Auth) bool {
	return this.Username == other.Username &&
		this.Password == other.Password
}

func (this *Auth) MakeCookie() *http.Cookie {
	return &http.Cookie{
		Name:  "auth",
		Value: this.String(),
	}
}
