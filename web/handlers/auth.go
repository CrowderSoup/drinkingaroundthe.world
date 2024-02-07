package handlers

import (
	"net/http"

	"github.com/CrowderSoup/drinkingaroundthe.world/services"
	"github.com/CrowderSoup/drinkingaroundthe.world/services/email"
	"github.com/labstack/echo/v4"
)

type LoginForm struct {
	Email string `form:"email"`
}

func initAuthHandlerGroup(e *echo.Echo, path string) {
	group := e.Group(path)

	group.GET("", getLogin)
	group.POST("/login", handleLoginSubmit)
	group.GET("/verify", handleLoginVerify)
}

func getLogin(c echo.Context) error {
	return c.Render(http.StatusOK, "auth/login", echo.Map{})
}

func handleLoginSubmit(c echo.Context) error {
	var form LoginForm
	err := c.Bind(&form)
	if err != nil {
		return c.String(http.StatusBadRequest, "bad request")
	}

	// TODO: create session, set sessionId as a claim in the JWT
	mailgunService := email.NewMailgun()
	jwtService := services.NewJwtService()

	tokenString, err := jwtService.CreateLoginToken()
	if err != nil {
		return c.String(http.StatusInternalServerError, "server error while creating login token")
	}

	mailgunService.SendMagicLink(form.Email, tokenString)

	return c.Render(http.StatusOK, "auth/login-email-sent.html", echo.Map{
		"email": form.Email,
	})
}

func handleLoginVerify(c echo.Context) error {
	// TODO: get session and JWT, validate JWT, ensure sessionId matches claim
	return nil
}
