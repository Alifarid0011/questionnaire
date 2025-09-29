package provider

import (
	"github.com/Alifarid0011/questionnaire-back-end/internal/controller"
	"github.com/Alifarid0011/questionnaire-back-end/internal/repository"
	"github.com/Alifarid0011/questionnaire-back-end/internal/service"
	"go.mongodb.org/mongo-driver/mongo"
)

func CommentController(service service.CommentService) controller.CommentController {
	return controller.NewCommentController(service)
}
func CommentService(repo repository.CommentRepository) service.CommentService {
	return service.NewCommentService(repo)
}

func CommentRepository(db *mongo.Database) repository.CommentRepository {
	return repository.NewCommentRepository(db)
}
