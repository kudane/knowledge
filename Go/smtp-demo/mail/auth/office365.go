// credit: from gist in github, Thank you bro.

package auth

import (
	"errors"
	"net/smtp"
)

type credential struct {
	Username string
	Password string
}

func NewAuthOffeice365(username, password, host, port string) (string, smtp.Auth) {
	server := host + port
	auth := &credential{username, password}
	return server, auth
}

func OutlookAuth(username, password string) smtp.Auth {
	return &credential{username, password}
}

func (a *credential) Start(server *smtp.ServerInfo) (string, []byte, error) {
	return "LOGIN", []byte(a.Username), nil
}

func (a *credential) Next(fromServer []byte, more bool) ([]byte, error) {
	if more {
		switch string(fromServer) {
		case "Username:":
			return []byte(a.Username), nil
		case "Password:":
			return []byte(a.Password), nil
		default:
			return nil, errors.ErrUnsupported
		}
	}
	return nil, nil
}
