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

type LoginVerifyQueryParams struct {
	Token string `query:"m"`
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

	session.SetValue("Email", form.Email, true)

	mailgunDomain := viper.GetString("mailgun_domain")
	mailgunApiKey := viper.GetString("mailgun_api_key")
	mailgunSendingAddress := viper.GetString("mailgun_sending_address")
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
	var query LoginVerifyQueryParams
	err := c.Bind(&query)
	if err != nil {
		return c.Render(http.StatusBadRequest, "auth/invalid", echo.Map{})
	}

	drinksContext := c.(*middleware.DrinksContext)
	session := drinksContext.Get("session").(*services.Session)

	jwtService := services.NewJwtService()
	_, err = jwtService.ValidateToken(query.Token, jwt.MapClaims{
		"sessionId": session.GetValue("ID"),
	})
	if err != nil {
		return c.Render(http.StatusBadRequest, "auth/invalid", echo.Map{})
	}

	session.SetValue("LoggedIn", true, true)

	// TODO: Now that we know this is a valid user we should create an
	// entry for them in the `Users` table

	return c.Render(http.StatusOK, "auth/valid", echo.Map{})
}
