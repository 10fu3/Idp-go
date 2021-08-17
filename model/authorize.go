package model

import (
	"github.com/google/uuid"
	"strings"
)

type AuthorizeRequest struct {
	Uuid string
	ClientId string
	RedirectUrl string
	State string
	Scope ServiceScope
	CodeChallenge string
}

func ToAuthorizeRequest(query map[string]string) AuthorizeRequest {

	scopes := strings.Split(query["scope"]," ")

	scope := ServiceScope{}

	for _, v := range scopes {
		switch v {
			case "openid":
				scope.OpenID = true
			case "profile":
				scope.Profile = true
			case "mail":
				scope.Mail = true
		}
	}

	return AuthorizeRequest{
		Uuid: uuid.NewString(),
		ClientId:      query["client_id"],
		RedirectUrl:   query["redirect_url"],
		State:         query["state"],
		CodeChallenge: query["code_challenge"],
		Scope: scope,
	}
}

