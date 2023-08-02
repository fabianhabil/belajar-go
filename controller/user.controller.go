package controller

import (
	"belajar-go/database"
	"belajar-go/model/dto"
	"belajar-go/model/entity"
	"belajar-go/utils"
	"log"

	"github.com/go-playground/validator/v10"
	"github.com/gofiber/fiber/v2"
)

func GetAllUser(ctx *fiber.Ctx) error {
	var user []entity.User

	data := database.DB.Find(&user)
	if data.Error != nil {
		log.Println(data.Error)
	}

	return ctx.JSON(user)
}

func CreateUser(ctx *fiber.Ctx) error {
	user := new(dto.UserRegister)

	if err := ctx.BodyParser(user); err != nil {
		log.Println(err)
	}

	validate := validator.New()

	if err := validate.Struct(user); err != nil {
		return ctx.Status(400).JSON(fiber.Map{
			"message": "Wrong Data Type",
			"info":    err.Error(),
		})
	}

	newUser := entity.User{
		Username: user.Username,
		Email:    user.Email,
	}

	hashPassword, err := utils.HashPassword(user.Password)

	if err != nil {
		return ctx.Status(fiber.StatusInternalServerError).SendString("Internal Server Error")
	}

	newUser.Password = hashPassword

	// Add to database
	if add := database.DB.Create(&newUser).Error; add != nil {
		return ctx.Status(fiber.StatusInternalServerError).JSON(fiber.Map{
			"message": add,
		})
	}

	return ctx.JSON(fiber.Map{
		"messaage": "User Added",
		"data":     newUser,
	})
}
