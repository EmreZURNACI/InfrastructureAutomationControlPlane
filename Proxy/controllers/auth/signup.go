package auth

import (
	"github.com/EmreZURNACI/InfrastructureAutomationControlPlaneProxy/app/auth"
	"github.com/EmreZURNACI/InfrastructureAutomationControlPlaneProxy/controllers"
	"github.com/gofiber/fiber/v2"
)

func (srv *AuthService) SignUp(c *fiber.Ctx) error {

	var req auth.SignUpRequest

	if err := c.BodyParser(&req); err != nil {
		return c.JSON(controllers.FailureResponse(fiber.ErrBadRequest.Code, fiber.ErrBadRequest.Message, fiber.ErrBadRequest.Message))
	}

	if err := controllers.Validate(req); err != nil {
		return c.JSON(controllers.FailureResponse(fiber.ErrBadRequest.Code, err.Error(), fiber.ErrBadRequest.Message))
	}

	res, err := srv.SignUpService.Execute(c.UserContext(), &req)
	if err != nil {
		return c.JSON(controllers.FailureResponse(fiber.ErrBadRequest.Code, err.Error(), err.Error()))
	}

	return c.JSON(controllers.SuccessResponse(fiber.StatusCreated, res.Message, nil))
}
