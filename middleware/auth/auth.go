package auth

import (
	"errors"
	"fmt"
	"github.com/InnoSoft/task/infrastructure/config"
	jwt "github.com/dgrijalva/jwt-go"
	"net/http"
	"strconv"
	"strings"
	"time"
)

//CreateToken will create a user jwt token
func CreateToken(user_id uint32) (string, error) {
	claims := jwt.MapClaims{}
	claims["authorized"] = true
	claims["user_id"] = strconv.Itoa(int(user_id))
	claims["exp"] = time.Now().Add(time.Hour * 1).Unix() //Token expires after 1 hour
	token := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)
	viper:=config.NewViper()
	return token.SignedString([]byte(viper.App.JwtSecret))

}

func VerifyToken(r *http.Request) (*jwt.Token, error) {
	viper:=config.NewViper()
	tokenString := ExtractToken(r)
	token, err := jwt.Parse(tokenString, func(token *jwt.Token) (interface{}, error) {
		if _, ok := token.Method.(*jwt.SigningMethodHMAC); !ok {
			return nil, fmt.Errorf("unexpected signing method: %v", token.Header["alg"])
		}
		return []byte(viper.App.JwtSecret) , nil
	})
	if err != nil {
		return nil, err
	}
	return token, nil
}

//ExtractToken get the token from the request body
func ExtractToken(r *http.Request) string {
	bearToken := r.Header.Get("Authorization")
	strArr := strings.Split(bearToken, " ")
	if len(strArr) == 2 {
		return strArr[1]
	}
	return ""
}

func Extract(token *jwt.Token) (string, error) {
	claims, ok := token.Claims.(jwt.MapClaims)
	if ok && token.Valid {
		userId, userOk := claims["user_id"].(string)
		if ok == false || userOk == false {
			return "", errors.New("unauthorized")
		} else {
			return userId, nil
		}
	}
	return "", errors.New("something went wrong")
}

func ExtractTokenMetadata(r *http.Request) (string, error) {
	token, err := VerifyToken(r)
	if err != nil {
		return "", err
	}
	acc, err := Extract(token)
	if err != nil {
		return "", err
	}
	fmt.Println("acc ", acc)
	return acc, nil
}