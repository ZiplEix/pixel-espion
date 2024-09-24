package services

import (
	"fmt"
	"os"
	"time"

	"github.com/ZiplEix/pixel-espion/database"
	"github.com/ZiplEix/pixel-espion/models"
	requestmodels "github.com/ZiplEix/pixel-espion/request_models"
	"github.com/golang-jwt/jwt/v4"
	"golang.org/x/crypto/bcrypt"
)

func hashPassword(password string) (string, error) {
	bytes, err := bcrypt.GenerateFromPassword([]byte(password), 14)
	return string(bytes), err
}

func checkPaswword(password, hash string) bool {
	err := bcrypt.CompareHashAndPassword([]byte(hash), []byte(password))
	return err == nil
}

func createJWT(user models.User) (string, error) {
	claims := jwt.MapClaims{
		"authorized": true,
		"user_id":    user.ID,
		"email":      user.Email,
		"exp":        time.Now().Add(time.Hour * 24).Unix(),
	}

	token := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)

	return token.SignedString([]byte(os.Getenv("JWT_SECRET")))
}

func Login(req requestmodels.LoginReq) (string, models.User, error) {
	var user models.User
	if err := database.Db.Where("email = ?", req.Email).First(&user).Error; err != nil {
		return "", models.User{}, ServiceError{
			Code:    404,
			Message: "User not found",
		}
	}

	if !checkPaswword(req.Password, user.Password) {
		return "", models.User{}, ServiceError{
			Code:    401,
			Message: "Invalid password",
		}
	}

	token, err := createJWT(user)
	if err != nil {
		return "", models.User{}, ServiceError{
			Code:    500,
			Message: "Error while creating JWT : " + err.Error(),
		}
	}

	return token, user, nil
}

func Register(req requestmodels.RegisterReq) (string, models.User, error) {
	hashedPassword, err := hashPassword(req.Password)
	if err != nil {
		return "", models.User{}, ServiceError{
			Code:    500,
			Message: "Error while hashing password",
		}
	}

	user := models.User{
		Email:    req.Email,
		Password: hashedPassword,
		Name:     req.Name,
	}

	if err := database.Db.Create(&user).Error; err != nil {
		return "", models.User{}, ServiceError{
			Code:    500,
			Message: "Error while creating user: " + err.Error(),
		}
	}

	fmt.Println(user)

	token, err := createJWT(user)
	if err != nil {
		return "", models.User{}, ServiceError{
			Code:    500,
			Message: "Error while creating JWT : " + err.Error(),
		}
	}

	return token, user, nil
}
