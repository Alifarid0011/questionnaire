package main

import (
	"context"
	"fmt"
	"github.com/Alifarid0011/questionnaire-back-end/config"
	"github.com/Alifarid0011/questionnaire-back-end/constant"
	_ "github.com/Alifarid0011/questionnaire-back-end/docs"
	"github.com/Alifarid0011/questionnaire-back-end/internal/dto"
	"github.com/Alifarid0011/questionnaire-back-end/internal/repository"
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
	ctx := context.Background()
	errEnsureAllIndexes := repository.EnsureAllIndexes(ctx, []repository.IndexEnsurer{
		app.CommentRepo,
		app.QuizRepo,
		app.UserAnswerRepo,
		app.UserRepo,
	})
	defaultPermissions(app)
	if errEnsureAllIndexes != nil {
		log.Fatalf("Failed to index on mongo: %v", errEnsureAllIndexes)
	}
	err := ensureSuperAdmin(ctx, app)
	if err != nil {
		log.Println(err)
	}
	if err := r.Run(fmt.Sprintf(":%v", config.Get.App.Port)); err != nil {
		log.Fatalf("Failed to start server: %v", err)
	}
}

func defaultPermissions(app *wire.App) {
	for i, permission := range constant.DefaultPermissions {
		_, err := app.CasbinRepo.AddPolicy(permission.Subject, permission.Object, permission.Action, permission.AllowOrDeny)
		if err != nil {
			log.Println(i, err)
		}
	}
}
func ensureSuperAdmin(ctx context.Context, app *wire.App) error {
	superAdminReq := dto.CreateUserRequest{
		Username: config.Get.SuperUser.Username,
		FullName: config.Get.SuperUser.FullName,
		Email:    config.Get.SuperUser.Email,
		Mobile:   config.Get.SuperUser.Mobile,
		Password: config.Get.SuperUser.Password,
	}
	superUser, err := app.UserService.CreateUser(superAdminReq, ctx)
	if err.Error() == "username already exists" {
		return nil
	}
	_, err = app.CasbinRepo.AddGroupingPolicy(superUser.UID.Hex(), constant.RoleSuperAdmin)
	return err
}
