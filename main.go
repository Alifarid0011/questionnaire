package main

import (
	"fmt"
	"github.com/Alifarid0011/questionnaire-back-end/config"
	"github.com/Alifarid0011/questionnaire-back-end/router"
	"github.com/Alifarid0011/questionnaire-back-end/wire"
	"log"
	"os"
)

func init() {
	config.ExposeConfig(os.Getenv("APP_ENV"))
}
func main() {
	app, errInitializeApp := wire.InitializeApp()
	if errInitializeApp != nil {
		log.Fatalf("Failed to initialize app: %v", errInitializeApp)
	}
	r := router.SetupRouter(app)
	r.RemoveExtraSlash = true
	if err := r.Run(fmt.Sprintf(":%v", config.Get.App.Port)); err != nil {
		log.Fatalf("Failed to start server: %v", err)
	}
}
