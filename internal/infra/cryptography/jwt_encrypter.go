package cryptography

import (
	"fmt"

	"github.com/lestrrat-go/jwx/v2/jwa"
	"github.com/lestrrat-go/jwx/v2/jwt"
	"github.com/nitoba/go-api/configs"
)

type JWTEncrypter struct{}

func (j *JWTEncrypter) Encrypt(payload map[string]interface{}) string {
	configs := configs.GetConfig()
	tok, err := jwt.NewBuilder().Build()
	if err != nil {
		fmt.Printf("failed to build token: %s\n", err)
		return ""
	}

	for k, v := range payload {
		tok.Set(k, v)
	}

	signed, err := jwt.Sign(tok, jwt.WithKey(jwa.HS256, configs.JWTSecret))

	if err != nil {
		fmt.Printf("failed to sign token: %s\n", err)
		return ""
	}

	return string(signed)
}

func (f *JWTEncrypter) Verify(token string) (map[string]interface{}, error) {
	configs := configs.GetConfig()
	verifiedToken, err := jwt.ParseString(token, jwt.WithKey(jwa.HS256, configs.JWTSecret))

	if err != nil {
		fmt.Printf("failed to parse token: %s\n", err)
		return nil, err
	}

	sub, _ := verifiedToken.Get(jwt.SubjectKey)
	exp, _ := verifiedToken.Get(jwt.ExpirationKey)

	return map[string]interface{}{
		"sub": sub,
		"exp": exp,
	}, err
}
