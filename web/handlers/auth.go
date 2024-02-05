package handlers

import (
	"net/http"

	"github.com/CrowderSoup/drinkingaroundthe.world/services/email"
	"github.com/labstack/echo/v4"
	"github.com/spf13/viper"
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

	// TODO: create JWT & session, generate token, set token in JWT as a cliam and in session, send email with JWT
	mailgunApiKey := viper.GetString("mailgun_api_key")
	mailgunDomain := viper.GetString("mailgun_domain")
	mailgunSendingAddress := viper.GetString("mailgun_sending_address")
	mailgunService := email.NewMailgun(mailgunApiKey, mailgunDomain, mailgunSendingAddress)

	mailgunService.SendMagicLink("aaron@crowder.cloud", "123456")

	return c.Render(http.StatusOK, "auth/login-email-sent.html", echo.Map{
		"email": form.Email,
	})
}

func handleLoginVerify(c echo.Context) error {
	// TODO: get session and JWT, if JWT is valid and token claim matches token in session login is valid
	return nil
}
