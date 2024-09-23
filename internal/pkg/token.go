package pkg

import (
	"errors"
	"time"

	"github.com/golang-jwt/jwt/v5"
)

type TokenJWT struct {
	AccessToken  string `json:"access_token"`
	RefreshToken string `json:"refresh_token"`
}

type TokenManager interface {
	NewJWTtoken(user_id string) (TokenJWT, error)
	ParseAccessToken(accessToken string) (string, error)
	RefreshToken(refreshToken string) (TokenJWT, error)
}

type InfoToken struct {
	UserId string
	Typ    string
}

const (
	jwtKey          = "daskmfgafjfjilaj3225242#dsadisasadjl"
	accessTokenTTL  = 15 * time.Minute
	refreshTokenTTL = 7 * 24 * time.Hour
	refreshTyp      = "refresh"
	accessTyp       = "access"
)

type TokenJWTManager struct {
}

func NewTokenJWTManager() *TokenJWTManager {
	return &TokenJWTManager{}
}

func (t *TokenJWTManager) NewJWTtoken(user_id string) (TokenJWT, error) {
	accessToken, err := t.generateToken(user_id, accessTokenTTL, accessTyp)
	if err != nil {
		return TokenJWT{}, err
	}
	refreshToken, err := t.generateToken(user_id, refreshTokenTTL, refreshTyp)
	if err != nil {
		return TokenJWT{}, err
	}
	return TokenJWT{
		AccessToken:  accessToken,
		RefreshToken: refreshToken,
	}, nil

}

func (t *TokenJWTManager) ParseAccessToken(accessToken string) (string, error) {
	tokenInfo, err := t.parseToken(accessToken)
	if err != nil {
		return "", err
	}
	if tokenInfo.Typ != accessTyp {
		return "", errors.New("token must be access")
	}
	return tokenInfo.UserId, nil
}

func (t *TokenJWTManager) RefreshToken(refreshToken string) (TokenJWT, error) {
	tokenInfo, err := t.parseToken(refreshToken)
	if err != nil {
		return TokenJWT{}, err
	}
	if tokenInfo.Typ != refreshTyp {
		return TokenJWT{}, errors.New("token must be refresh")
	}

	return t.NewJWTtoken(tokenInfo.UserId)

}

func (t *TokenJWTManager) parseToken(tokenJwt string) (InfoToken, error) {
	token, err := t.validateToken(tokenJwt)
	if err != nil {
		return InfoToken{}, err
	}
	claims, ok := token.Claims.(jwt.MapClaims)
	if !ok {
		return InfoToken{}, errors.New("token claims are not of type")
	}
	typeToken := claims["typ"]
	userId, err := claims.GetSubject()
	if err != nil {
		return InfoToken{}, err
	}

	return InfoToken{
		Typ:    typeToken.(string),
		UserId: userId,
	}, nil
}

func (t *TokenJWTManager) generateToken(user_id string, timeExp time.Duration, typeToken string) (string, error) {
	token := jwt.NewWithClaims(jwt.SigningMethodHS256,
		&jwt.MapClaims{
			"sub": user_id,
			"exp": time.Now().Add(timeExp).Unix(),
			"iat": time.Now(),
			"typ": typeToken,
		},
	)
	tokenString, err := token.SignedString([]byte(jwtKey))
	if err != nil {
		return "", err
	}
	return tokenString, nil
}

func (t *TokenJWTManager) validateToken(jwtToken string) (*jwt.Token, error) {
	token, err := jwt.ParseWithClaims(jwtToken, jwt.MapClaims{}, func(accessToken *jwt.Token) (interface{}, error) {
		if _, ok := accessToken.Method.(*jwt.SigningMethodHMAC); !ok {
			return nil, errors.New("invalid signing method")
		}
		return []byte(jwtKey), nil
	})
	if err != nil {
		return nil, err
	}
	return token, nil
}
