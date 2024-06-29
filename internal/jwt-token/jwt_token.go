package jwttoken

import (
	"fmt"
	"time"

	"github.com/gin-gonic/gin"
	jwt "github.com/golang-jwt/jwt/v5"
)

type AccountInfo struct {
	Role string
	Id   int32
	// AdditionalData interface{}
}

type CustomClaims struct {
	AccountInfo    AccountInfo
	AdditionalInfo interface{}
	jwt.RegisteredClaims
}


type AddlInfoType interface {
	isCustomClaimType()
	getWithLabel() additionalInfoWithLabel
	getFromInfoMap(infoMap map[string]interface{}) (*ExtraPurposeInfo, error)
	SetContext(c *gin.Context)
}
type additionalInfoWithLabel struct {
	Label string
	Info  interface{}
}
func getCustomClaim(accInfo AccountInfo, addlInfo AddlInfoType, expiryTime time.Duration) *CustomClaims {
	claims := &CustomClaims{
		AccountInfo: accInfo,
		RegisteredClaims: jwt.RegisteredClaims{
			ExpiresAt: jwt.NewNumericDate(time.Now().Add(time.Duration(expiryTime))),
		},
	}
	if addlInfo != nil {
		claims.AdditionalInfo = addlInfo.getWithLabel()
	}
	return claims
}
func (g *TokenGenerator) GenerateToken(accInfo AccountInfo, additionalInfo AddlInfoType, expiryTime time.Duration) (*string, error) {
	claims := getCustomClaim(accInfo, additionalInfo, expiryTime)

	token := jwt.NewWithClaims(jwt.SigningMethodRS256, claims)

	// Sign and get the complete encoded token as a string using the secret
	tokenString, err := token.SignedString(g.privateKey)

	// fmt.Println(tokenString, err)
	return &tokenString, err
}

func (v *TokenVerifier) ValidateToken(tokenString string) (bool, *AccountInfo, AddlInfoType, error) {
	token, err := jwt.ParseWithClaims(tokenString, &CustomClaims{}, func(token *jwt.Token) (interface{}, error) {
		return v.publicKey, nil
	})
	if err != nil {
		return false, nil, nil, fmt.Errorf("error in parsing jwt token:%v", err)
	}

	if !token.Valid {
		return false, nil, nil, fmt.Errorf("token is invalid")
	}

	claims, ok := token.Claims.(*CustomClaims)
	if !ok {
		return false, nil, nil, fmt.Errorf("error while decoding token into custom claims")
	}

	//check if token expired
	if claims.ExpiresAt.Time.Before(time.Now()) {
		return false, nil, nil, fmt.Errorf("token expired")
	}

	if claims.AdditionalInfo == nil {
		return true, &claims.AccountInfo, nil, nil
	}

	addlInfoWithLabel := claims.AdditionalInfo.(map[string]interface{})
	addlInfo := getAddlInfo(addlInfoWithLabel)
	return true, &claims.AccountInfo, addlInfo, nil
}
func getAddlInfo(addlInfo map[string]interface{}) AddlInfoType {
	fmt.Println("in getAddlInfo, addlInfo:", addlInfo)
	switch addlInfo["Label"] {
	case "ExtraPurposeInfo":
		extraPurposeInfo := &ExtraPurposeInfo{}
		extraPurposeInfo, err := extraPurposeInfo.getFromInfoMap(addlInfo)
		if err != nil {
			fmt.Println("error in getting extraPurposeInfo from infoMap:", err)
			return nil
		}
		return extraPurposeInfo
	default:
		panic("Invalid label")	//unimplemented labels will panic
	}
}
