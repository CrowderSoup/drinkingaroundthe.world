package handlers

import (
	"net/http"

	"github.com/CrowderSoup/drinkingaroundthe.world/services"
	"github.com/CrowderSoup/drinkingaroundthe.world/services/email"
	"github.com/CrowderSoup/drinkingaroundthe.world/web/middleware"
	"github.com/golang-jwt/jwt"
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

	drinksContext := c.(*middleware.DrinksContext)
	session := drinksContext.Get("session").(*services.Session)

	mailgunDomain := viper.GetString("mailgun_domain")
	mailgunApiKey := viper.GetString("mailgun_api_key")
	mailgunSendingAddress := viper.GetString("mailgun_sending_address")
	c.Logger().Printf("Domain: %s, APIKey: %s, SendingAddress: %s", mailgunDomain, mailgunApiKey, mailgunSendingAddress)
	mailgunService := email.NewMailgun(mailgunDomain, mailgunApiKey, mailgunSendingAddress)
	jwtService := services.NewJwtService()

	tokenString, err := jwtService.CreateLoginToken(jwt.MapClaims{
		"sessionId": session.GetValue("ID"),
	})
	if err != nil {
		c.Logger().Print("error getting token")
		return c.String(http.StatusInternalServerError, "server error while creating login token")
	}

	err = mailgunService.SendMagicLink(form.Email, tokenString)
	if err != nil {
		return c.String(http.StatusInternalServerError, "server error while sending mailgun email")
	}

	return c.Render(http.StatusOK, "auth/login-email-sent.html", echo.Map{
		"email": form.Email,
	})
}

func handleLoginVerify(c echo.Context) error {
	// TODO: get session and JWT, validate JWT, ensure sessionId matches claim
	return nil
}
