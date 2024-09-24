package controllers

import (
	requestmodels "github.com/ZiplEix/pixel-espion/request_models"
	"github.com/ZiplEix/pixel-espion/services"
	"github.com/ZiplEix/pixel-espion/validation"
	"github.com/gofiber/fiber/v2"
)

// Login godoc
// @Summary User login
// @Description Logs in a user and returns a JWT token and user information
// @Tags auth
// @Accept json
// @Produce json
// @Param login body requestmodels.LoginReq true "Login Request"
// @Success 200 {object} fiber.Map{token=string,user=fiber.Map{email=string,name=string,id=uint}} "Login successful"
// @Failure 400 {object} errorResponse "Bad Request"
// @Failure 401 {object} errorResponse "Unauthorized"
// @Failure 404 {object} errorResponse "User Not Found"
// @Failure 500 {object} errorResponse "Internal Server Error"
// @Router /login [post]
func Login(c *fiber.Ctx) error {
	var req requestmodels.LoginReq
	if err := c.BodyParser(&req); err != nil {
		return c.Status(fiber.StatusBadRequest).JSON(&errorResponse{Error: err.Error()})
	}

	err := validation.Login(req)
	if err != nil {
		return c.Status(fiber.StatusBadRequest).JSON(&errorResponse{Error: err.Error()})
	}

	token, user, err := services.Login(req)
	if err != nil {
		return c.Status(err.(services.ServiceError).Code).JSON(errorResponse{
			Error: err.Error(),
		})
	}

	// set the token on the cookies
	c.Cookie(&fiber.Cookie{
		Name:  "jwt",
		Value: token,
	})
	c.Cookie(&fiber.Cookie{
		Name:  "user",
		Value: user.Name,
	})

	return c.Status(fiber.StatusOK).JSON(fiber.Map{
		"token": token,
		"user": fiber.Map{
			"email": user.Email,
			"name":  user.Name,
			"id":    user.ID,
		},
	})
}

// Register godoc
// @Summary User registration
// @Description Registers a new user and returns a JWT token and user information
// @Tags auth
// @Accept json
// @Produce json
// @Param register body requestmodels.RegisterReq true "Register Request"
// @Success 201 {object} fiber.Map{token=string,user=fiber.Map{email=string,name=string,id=uint}} "Registration successful"
// @Failure 400 {object} errorResponse "Bad Request"
// @Failure 500 {object} errorResponse "Internal Server Error"
// @Router /register [post]
func Register(c *fiber.Ctx) error {
	var req requestmodels.RegisterReq
	if err := c.BodyParser(&req); err != nil {
		return c.Status(fiber.StatusBadRequest).JSON(&errorResponse{Error: err.Error()})
	}

	err := validation.Register(req)
	if err != nil {
		return c.Status(fiber.StatusBadRequest).JSON(&errorResponse{Error: err.Error()})
	}

	token, user, err := services.Register(req)
	if err != nil {
		return c.Status(err.(services.ServiceError).Code).JSON(errorResponse{
			Error: err.Error(),
		})
	}

	// set the token on the cookies
	c.Cookie(&fiber.Cookie{
		Name:  "jwt",
		Value: token,
	})
	c.Cookie(&fiber.Cookie{
		Name:  "user",
		Value: user.Name,
	})

	return c.Status(fiber.StatusCreated).JSON(fiber.Map{
		"token": token,
		"user": fiber.Map{
			"email": user.Email,
			"name":  user.Name,
			"id":    user.ID,
		},
	})
}
