package authrepository

import (
	"errors"
	"fmt"

	"github.com/golang-jwt/jwt/v4"
)

func ParseAccessToken(jwt_secret string, access_token string) (*AccessTokenPayload, error) {
	token, err := jwt.Parse(access_token, func(token *jwt.Token) (interface{}, error) {
		if _, ok := token.Method.(*jwt.SigningMethodHMAC); !ok {
			return nil, errors.New(fmt.Sprintf("unexpected signing method: %v", token.Header["alg"]))
		}

		return []byte(jwt_secret), nil
	})
	if err != nil {
		return nil, err
	}

	claims, ok := token.Claims.(jwt.MapClaims)
	if !ok || !token.Valid {
		return nil, errors.New("invalid token")
	}

	identifier, ok := claims["sub"].(string)
	if !ok {
		return nil, errors.New("invalid token")
	}

	var client_type TokenOwnerType
	token_client_type, ok := claims["sub_typ"].(string)
	if token_client_type == string(UserTokenOwnerType) {
		client_type = UserTokenOwnerType
	} else if token_client_type == string(ModuleTokenOwnerType) {
		client_type = ModuleTokenOwnerType
	} else {
		return nil, errors.New("invalid token")
	}

	var role *string = nil
	token_role, ok := claims["rol"].(string)
	if ok {
		role = &token_role
	}

	return &AccessTokenPayload{
		Identifier: identifier,
		ClientType: client_type,
		Role:       role,
	}, nil
}
