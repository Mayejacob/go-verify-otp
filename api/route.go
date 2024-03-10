package api

import (
	"github.com/gin-gonic/gin"
)

type Config struct {
	Router *gin.Engine
}

func (app *Config) Routes() {
	app.Router.POST("/api/v1/otp", app.sendSMS())
	app.Router.POST("/api/v1/verify_otp", app.verifySMS())
}
