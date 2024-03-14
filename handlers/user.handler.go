package handlers

import (
	"os"

	"github.com/mesxx/Fiber_User_Management_API/helpers"
	"github.com/mesxx/Fiber_User_Management_API/models"
	"github.com/mesxx/Fiber_User_Management_API/usecases"

	"github.com/go-playground/validator/v10"
	"github.com/gofiber/fiber/v2"
)

type (
	UserHandler interface {
		Create(c *fiber.Ctx) error
		Login(c *fiber.Ctx) error
		GetAll(c *fiber.Ctx) error
		GetByID(c *fiber.Ctx) error
		UpdateByID(c *fiber.Ctx) error
		ForgotPassword(c *fiber.Ctx) error
		ResetPassword(c *fiber.Ctx) error
		DeleteByID(c *fiber.Ctx) error
		DeleteAll(c *fiber.Ctx) error
	}

	userHandler struct {
		UserUsecase usecases.UserUsecase
	}
)

func NewUserHandler(usecase usecases.UserUsecase) UserHandler {
	return &userHandler{
		UserUsecase: usecase,
	}
}

func (handler userHandler) Create(c *fiber.Ctx) error {
	// START request
	var requestCreateUser models.RequestCreateUser
	if err := c.BodyParser(&requestCreateUser); err != nil {
		return fiber.NewError(fiber.StatusBadRequest, err.Error())
	}
	// END request

	// START validator
	validate := validator.New()
	if err := validate.Struct(requestCreateUser); err != nil {
		return fiber.NewError(fiber.StatusBadRequest, err.Error())
	}
	// END validator

	// START hash password
	hashedPassword, err1 := helpers.HashPassword(requestCreateUser.Password)
	if err1 != nil {
		return fiber.NewError(fiber.StatusBadRequest, err1.Error())
	}

	requestCreateUser.Password = hashedPassword
	// END hash password

	res, err2 := handler.UserUsecase.Create(&requestCreateUser)
	if err2 != nil {
		return fiber.NewError(fiber.StatusBadRequest, err2.Error())
	}
	return c.Status(fiber.StatusCreated).JSON(helpers.GetResponseData(fiber.StatusCreated, "success", res))
}

func (handler userHandler) Login(c *fiber.Ctx) error {
	// START request
	var requestLoginUser models.RequestLoginUser
	if err := c.BodyParser(&requestLoginUser); err != nil {
		return fiber.NewError(fiber.StatusBadRequest, err.Error())
	}
	// END request

	// START validator
	validate := validator.New()
	if err := validate.Struct(requestLoginUser); err != nil {
		return fiber.NewError(fiber.StatusBadRequest, err.Error())
	}
	// END validator

	user, err1 := handler.UserUsecase.GetByEmail(requestLoginUser.Email)
	if err1 != nil {
		return fiber.NewError(fiber.StatusBadRequest, err1.Error())
	}

	// START check hash password
	if err := helpers.CheckHashPassword(requestLoginUser.Password, user.Password); err != nil {
		return fiber.NewError(fiber.StatusBadRequest, err.Error())
	}
	// END check hash password

	// START generate jwt
	jwtToken, err2 := helpers.GenerateJWT(user)
	if err2 != nil {
		return fiber.NewError(fiber.StatusBadRequest, err2.Error())
	}
	// END generate jwt

	res := fiber.Map{
		"user":  user,
		"token": jwtToken,
	}
	return c.Status(fiber.StatusOK).JSON(helpers.GetResponseData(fiber.StatusOK, "success", res))
}

func (handler userHandler) GetAll(c *fiber.Ctx) error {
	users, err := handler.UserUsecase.GetAll()
	if err != nil {
		return fiber.NewError(fiber.StatusBadRequest, err.Error())
	}
	return c.Status(fiber.StatusOK).JSON(helpers.GetResponseData(fiber.StatusOK, "success", users))
}

func (handler userHandler) GetByID(c *fiber.Ctx) error {
	userSigned := c.Locals("user").(*models.CustomClaims)
	user, err := handler.UserUsecase.GetByID(userSigned.ID)
	if err != nil {
		return fiber.NewError(fiber.StatusBadRequest, err.Error())
	}
	return c.Status(fiber.StatusOK).JSON(helpers.GetResponseData(fiber.StatusOK, "success", user))
}

func (handler userHandler) UpdateByID(c *fiber.Ctx) error {
	userSigned := c.Locals("user").(*models.CustomClaims)

	user, err1 := handler.UserUsecase.GetByID(userSigned.ID)
	if err1 != nil {
		return fiber.NewError(fiber.StatusBadRequest, err1.Error())
	}

	// START request
	var requestUpdateByIDUser models.User
	if err := c.BodyParser(&requestUpdateByIDUser); err != nil {
		return fiber.NewError(fiber.StatusBadRequest, err.Error())
	}
	// END request

	if requestUpdateByIDUser.Name != "" {
		user.Name = requestUpdateByIDUser.Name
	}

	if requestUpdateByIDUser.Email != "" {
		user.Email = requestUpdateByIDUser.Email
	}

	if requestUpdateByIDUser.Password != "" {
		// START hash password
		hashedPassword, err := helpers.HashPassword(requestUpdateByIDUser.Password)
		if err != nil {
			return fiber.NewError(fiber.StatusBadRequest, err.Error())
		}
		// END hash password

		user.Password = hashedPassword
	}

	updateByID, err2 := handler.UserUsecase.UpdateByID(user)
	if err2 != nil {
		return fiber.NewError(fiber.StatusBadRequest, err2.Error())
	}
	return c.Status(fiber.StatusOK).JSON(helpers.GetResponseData(fiber.StatusOK, "success", updateByID))
}

func (handler userHandler) ForgotPassword(c *fiber.Ctx) error {
	// START request
	var requestForgotPasswordUser models.RequestForgotPasswordUser
	if err := c.BodyParser(&requestForgotPasswordUser); err != nil {
		return fiber.NewError(fiber.StatusBadRequest, err.Error())
	}
	// END request

	// START validator
	validate := validator.New()
	if err := validate.Struct(requestForgotPasswordUser); err != nil {
		return fiber.NewError(fiber.StatusBadRequest, err.Error())
	}
	// END validator

	user, err1 := handler.UserUsecase.GetByEmail(requestForgotPasswordUser.Email)
	if err1 != nil {
		return fiber.NewError(fiber.StatusBadRequest, err1.Error())
	}

	// START encrypt
	encryptedEmail, err2 := helpers.EncryptText(user.Email)
	if err2 != nil {
		return fiber.NewError(fiber.StatusBadRequest, err2.Error())
	}
	// END encrypt

	res := fiber.Map{
		"link": "http://localhost:" + os.Getenv("PORT") + "/api/user/password/reset?hash=" + encryptedEmail,
	}
	return c.Status(fiber.StatusOK).JSON(helpers.GetResponseData(fiber.StatusOK, "success", res))
}

func (handler userHandler) ResetPassword(c *fiber.Ctx) error {
	hashQuery := c.Query("hash")

	// START decrypt
	userEmail, err1 := helpers.DecryptText(hashQuery)
	if err1 != nil {
		return fiber.NewError(fiber.StatusBadRequest, err1.Error())
	}
	// END decrypt

	user, err2 := handler.UserUsecase.GetByEmail(userEmail)
	if err2 != nil {
		return fiber.NewError(fiber.StatusBadRequest, err2.Error())
	}

	// START request
	var requestResetPasswordUser models.RequestResetPasswordUser
	if err := c.BodyParser(&requestResetPasswordUser); err != nil {
		return fiber.NewError(fiber.StatusBadRequest, err.Error())
	}
	// END request

	// START validator
	validate := validator.New()
	if err := validate.Struct(requestResetPasswordUser); err != nil {
		return fiber.NewError(fiber.StatusBadRequest, err.Error())
	}
	// END validator

	// START hash password
	hashedPassword, err3 := helpers.HashPassword(requestResetPasswordUser.Password)
	if err3 != nil {
		return fiber.NewError(fiber.StatusBadRequest, err3.Error())
	}
	// END hash password

	user.Password = hashedPassword
	updateByID, err4 := handler.UserUsecase.UpdateByID(user)
	if err4 != nil {
		return fiber.NewError(fiber.StatusBadRequest, err4.Error())
	}
	return c.Status(fiber.StatusOK).JSON(helpers.GetResponseData(fiber.StatusOK, "success", updateByID))
}

func (handler userHandler) DeleteByID(c *fiber.Ctx) error {
	userSigned := c.Locals("user").(*models.CustomClaims)

	user, err1 := handler.UserUsecase.GetByID(userSigned.ID)
	if err1 != nil {
		return fiber.NewError(fiber.StatusBadRequest, err1.Error())
	}

	deleteByID, err2 := handler.UserUsecase.DeleteByID(user)
	if err2 != nil {
		return fiber.NewError(fiber.StatusBadRequest, err2.Error())
	}
	return c.Status(fiber.StatusOK).JSON(helpers.GetResponseData(fiber.StatusOK, "success", deleteByID))
}

func (handler userHandler) DeleteAll(c *fiber.Ctx) error {
	if err := handler.UserUsecase.DeleteAll(); err != nil {
		return fiber.NewError(fiber.StatusBadRequest, err.Error())
	}
	return c.Status(fiber.StatusOK).JSON(helpers.GetResponse(fiber.StatusOK, "success"))
}
