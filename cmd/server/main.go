package main

import (
	"fmt"
	"log"

	"github.com/gofiber/fiber/v2"
	"github.com/gofiber/fiber/v2/middleware/cors"
	"github.com/gofiber/swagger"
	"github.com/joho/godotenv"
	_ "github.com/qhmd/gitforgits/cmd/server/docs"
	handler "github.com/qhmd/gitforgits/internal/delivery/http"
	repo "github.com/qhmd/gitforgits/internal/repository"
	useCase "github.com/qhmd/gitforgits/internal/usecase"
	"github.com/qhmd/gitforgits/pkg/database"
	"github.com/qhmd/gitforgits/utils"
)

func init() {
	if err := godotenv.Load(); err != nil {
		log.Fatal("Error loading .env file")
	}
}

// @title           GitForGits API
// @version         1.0
// @description     API documentation for project GitForGits
// @termsOfService  http://swagger.io/terms/

// @securityDefinitions.apikey  BearerAuth
// @in                         header
// @name                       Authorization
// @description                Type "Bearer" followed by a space and JWT token.

func main() {
	app := fiber.New()
	app.Use(cors.New())

	utils.InitValidator()
	fmt.Print("listen in port 8080...")
	db := database.InitMySQL()
	database.RunMigration(db)
	repoBook := repo.NewMySQLBookRepository(db)
	repoAuth := repo.NewMySQLAuthRepository(db)
	repoUsers := repo.NewUserMySqlRepo(db)

	ucBook := useCase.NewBookUsecase(repoBook)
	ucAuth := useCase.NewAuthUsecase(repoAuth)
	ucUsers := useCase.NewUsersUseCase(repoUsers)

	app.Get("/swagger/*", swagger.HandlerDefault)
	handler.NewBookHandler(app, ucBook)
	handler.NewAuthHandler(app, ucAuth)
	handler.NewHandlerUser(app, ucUsers)
	app.Listen(":8080")
}
