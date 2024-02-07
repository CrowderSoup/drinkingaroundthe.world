package handlers

import (
	"net/http"

	"github.com/CrowderSoup/drinkingaroundthe.world/services"
	"github.com/CrowderSoup/drinkingaroundthe.world/services/email"
	"github.com/CrowderSoup/drinkingaroundthe.world/web/middleware"
	"github.com/golang-jwt/jwt"
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

	drinksContext := c.(*middleware.DrinksContext)
	session := drinksContext.Get("session").(*services.Session)

	mailgunService := email.NewMailgun()
	jwtService := services.NewJwtService()

	tokenString, err := jwtService.CreateLoginToken(jwt.MapClaims{
		"sessionId": session.Internal.ID,
	})
	if err != nil {
		c.Logger().Print("error getting token")
		return c.String(http.StatusInternalServerError, "server error while creating login token")
	}

	// TODO: Mailgun is return a 401... need to figure that out tomorrow
	mailgunService.SendMagicLink(form.Email, tokenString)

	return c.Render(http.StatusOK, "auth/login-email-sent.html", echo.Map{
		"email": form.Email,
	})
}

func handleLoginVerify(c echo.Context) error {
	// TODO: get session and JWT, validate JWT, ensure sessionId matches claim
	return nil
}
