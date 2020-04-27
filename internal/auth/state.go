package auth

import (
	"encoding/base64"
	"encoding/json"

	"go.uber.org/zap/zapcore"
)

type State struct {
	Redirect string `json:"redirect"`
	Code     string `json:"code,omitempty"`
}

func (s State) MarshalLogObject(e zapcore.ObjectEncoder) error {
	e.AddString("redirect", s.Redirect)
	e.AddString("code", s.Code)
	return nil
}

func EncodeState(v interface{}) (string, error) {
	d, err := json.Marshal(v)
	if err != nil {
		return "", err
	}
	return base64.RawURLEncoding.EncodeToString(d), nil
}

func DecodeState(state string, v interface{}) error {
	d, err := base64.RawURLEncoding.DecodeString(state)
	if err != nil {
		return err
	}
	return json.Unmarshal(d, v)
}
