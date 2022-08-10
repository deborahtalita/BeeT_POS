package service

import (
	"beet_pos/entity"
	"fmt"
	"log"
	"time"

	"github.com/dgrijalva/jwt-go"
)

type (
	// CustomKey contains to 2 string to create hmac
	CustomKey struct {
		Secret string
		Data   string
	}

	// RefreshTokenCustomClaims specifies the claims for refresh token
	RefreshTokenCustomClaims struct {
		UserData  string
		CustomKey string
		KeyType   string
		jwt.StandardClaims
	}

	// AccessTokenCustomClaims specifies the claims for access token
	AccessTokenCustomClaims struct {
		Username  string
		Role      string
		Outlet_id string
		KeyType   string
		jwt.StandardClaims
	}
)

type JWTService interface {
	GenerateAccessToken(userData entity.User) string
	//GenerateRefreshToken
	ValidateToken(token string) (*jwt.Token, error)
}

type jwtService struct {
	secretKey string
	issuer    string
}

func NewJWTService() JWTService {
	return &jwtService{
		issuer:    "beetpos",
		secretKey: "beetpos",
	}
}

// GenerateAccessToken implements JWTService
func (js *jwtService) GenerateAccessToken(userData entity.User) string {
	log.Printf("JWTService : GenerateAccessToken")
	v := userData
	claims := &AccessTokenCustomClaims{
		v.User_name,
		v.User_role,
		v.Outlet_id,
		"access",
		jwt.StandardClaims{
			ExpiresAt: time.Now().AddDate(1, 0, 0).Unix(),
			Issuer:    js.issuer,
		},
	}
	token := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)
	t, err := token.SignedString([]byte(js.secretKey))
	if err != nil {
		panic(err)
	}
	return t
}

// ValidateToken implements JWTService
func (js *jwtService) ValidateToken(token string) (*jwt.Token, error) {
	return jwt.Parse(token, func(t_ *jwt.Token) (interface{}, error) {
		if _, ok := t_.Method.(*jwt.SigningMethodHMAC); !ok {
			return nil, fmt.Errorf("unexpected signing method %v", t_.Header["alg"])
		}
		return []byte(js.secretKey), nil
	})
}
