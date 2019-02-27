package models

import (
	"github.com/ory/hydra/sdk/go/hydra/swagger"
	"go.uber.org/zap/zapcore"
)

type Oauth2LoginForm struct {
	Challenge string `query:"login_challenge" form:"login_challenge" validate:"required"`
}

type Oauth2LoginSubmitForm struct {
	Csrf      string `query:"_csrf" form:"_csrf" validate:"required"`
	Challenge string `query:"challenge" form:"challenge" validate:"required"`
	Email     string `query:"email" form:"email" validate:"required"`
	Password  string `query:"password" form:"password" validate:"required"`
	Remember  bool   `query:"remember" form:"remember"`
}

type Oauth2ConsentForm struct {
	Challenge string `query:"consent_challenge" form:"consent_challenge" validate:"required"`
}

type Oauth2ConsentSubmitForm struct {
	Csrf      string   `query:"_csrf" form:"_csrf" validate:"required"`
	Challenge string   `query:"challenge" form:"challenge" validate:"required"`
	Scope     []string `query:"scope" form:"scope" validate:"required"`
}

type Oauth2CallbackForm struct {
	Code  string `query:"code" form:"code" validate:"required"`
	State string `query:"state" form:"state" validate:"required"`
	Scope string `query:"scope" form:"scope" validate:"false"`
}

func (a *Oauth2LoginSubmitForm) MarshalLogObject(enc zapcore.ObjectEncoder) error {
	enc.AddString("Challenge", a.Challenge)
	enc.AddString("Email", a.Email)
	enc.AddString("Password", "[HIDDEN]")
	enc.AddString("Csrf", a.Csrf)

	return nil
}

type Oauth2IntrospectForm struct {
	ClientID string `query:"client_id" form:"client_id" validate:"required"`
	Secret   string `query:"secret" form:"secret" validate:"required"`
	Token    string `query:"token" form:"token" validate:"required"`
}

func (a *Oauth2IntrospectForm) MarshalLogObject(enc zapcore.ObjectEncoder) error {
	enc.AddString("ClientID", a.ClientID)
	enc.AddString("Token", a.Token)
	enc.AddString("Secret", "[HIDDEN]")

	return nil
}

type Oauth2TokenIntrospection struct {
	*swagger.OAuth2TokenIntrospection
}