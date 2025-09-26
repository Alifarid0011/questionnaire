package main

import (
	"fmt"
	"github.com/Alifarid0011/questionnaire-back-end/config"
	"github.com/Alifarid0011/questionnaire-back-end/routers"
	"github.com/Alifarid0011/questionnaire-back-end/wire"
	"log"
	"os"
)

func init() {
	config.ExposeConfig(os.Getenv("APP_ENV"))
}

// @title Swagger Example API
// @version 1.0
// @description This is a sample API for demonstrating Swagger with Bearer Authentication in Go using Gin
// @securityDefinitions.apikey AuthBearer
// @in header
// @name Authorization
func main() {
	app, errInitializeApp := wire.InitializeApp(config.Get.Token.SecretKey)
	if errInitializeApp != nil {
		log.Fatalf("Failed to initialize app: %v", errInitializeApp)
	}
	r := routers.SetupRouter(app)
	r.RemoveExtraSlash = true
	if err := r.Run(fmt.Sprintf(":%v", config.Get.App.Port)); err != nil {
		log.Fatalf("Failed to start server: %v", err)
	}
}
