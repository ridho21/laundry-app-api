package common

import (
	"errors"
	"go-enigma-laundry/config"
	"go-enigma-laundry/model"
	"time"

	"github.com/dgrijalva/jwt-go"
)


type JwtClaim struct {
	jwt.StandardClaims
	UserId string `json:"userId"`
	Email string `json:"email"`
	Role string `json:"role"`
}

type JwtToken interface {
	GenerateTokenJwt(userData model.User) (string,error)
	VerifyToken(tokenString string) (jwt.MapClaims, error) 
}
type jwtToken struct {
	config config.TokenConfig
}

func (j *jwtToken) GenerateTokenJwt(userData model.User) (string,error) {
	claims := JwtClaim{
		StandardClaims: jwt.StandardClaims{
			Issuer: j.config.IssuerName,
			ExpiresAt: time.Now().Add(j.config.JwtLifeTime).Unix(),
		},
		UserId : userData.Id,
		Email : userData.Email,
		Role : userData.Role,
	}

	token := jwt.NewWithClaims(jwt.SigningMethodHS256,claims)
	signedToken, err := token.SignedString(j.config.JwtSignatureKey)
	if err != nil {
		return "",err
	}
	return signedToken,nil
}

func (j *jwtToken) VerifyToken(tokenString string) (jwt.MapClaims, error) {
    token, err := jwt.Parse(tokenString, func(token *jwt.Token) (interface{}, error) {		
        return j.config.JwtSignatureKey, nil
    })
    if err != nil {
        return nil, err
    }

    // Convert dari jwt.Claims ke jwt.MapClaims
    claims, ok := token.Claims.(jwt.MapClaims)
    if !ok || !token.Valid {
        return nil, errors.New("Failed to parse map claims or token is not valid")
    }

    // Verifikasi issuer
    if !claims.VerifyIssuer(j.config.IssuerName, true) {
        return nil, errors.New("Failed to verify issuer name")
    }

    // Verifikasi expired
    if !claims.VerifyExpiresAt(time.Now().Unix(), true) {
        return nil, errors.New("Token is expired")
    }

    return claims, nil
}


func NewJwtToken(config config.TokenConfig) JwtToken {
	return &jwtToken{
		config: config,
	}
}