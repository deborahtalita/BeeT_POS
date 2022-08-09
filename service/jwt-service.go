package service

import (
	"beet_pos/entity"
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
		Username string
		Role	 string
		Outlet_ID	string
		KeyType  string
		jwt.StandardClaims
	}
)

type JWTService interface {
	GenerateAccessToken(userData entity.User) string
	//GenerateRefreshToken
	//ValidateToken
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
		v.Username,
		v.User_role,
		v.Outlet_ID,
		"access",
		jwt.StandardClaims{
			ExpiresAt: time.Now().AddDate(1, 0, 0).Unix(),
			Issuer: js.issuer,
		},
	}
	token := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)
	t, err := token.SignedString([]byte(js.secretKey))
	if err != nil {
		panic(err)
	}
	return t
}
