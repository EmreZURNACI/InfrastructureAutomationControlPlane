package auth

import (
	"github.com/EmreZURNACI/InfrastructureAutomationControlPlaneProxy/app/auth"
	"github.com/EmreZURNACI/InfrastructureAutomationControlPlaneProxy/controllers"
	"github.com/gofiber/fiber/v2"
)

func (srv *AuthService) SignIn(c *fiber.Ctx) error {

	var req auth.SignInRequest
	if err := c.BodyParser(&req); err != nil {
		return c.JSON(controllers.FailureResponse(400, err.Error(), err.Error()))
	}

	if err := controllers.Validate(req); err != nil {
		return c.JSON(controllers.FailureResponse(400, err.Error(), err.Error()))
	}

	token, err := srv.SignInService.Execute(c.UserContext(), &req)
	if err != nil {
		return c.JSON(controllers.FailureResponse(400, err.Error(), err.Error()))
	}

	return c.JSON(controllers.SuccessResponse(200, "login successfully", *token.Token))

}
