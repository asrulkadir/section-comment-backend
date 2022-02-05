package main

import (
	// "database/sql"
	// "fmt"
	// "log"
	"net/http"
	"os"

	"github.com/labstack/echo/v4"
	"github.com/labstack/echo/v4/middleware"
	_ "github.com/lib/pq"

	// "github.com/asrulkadir/section-comment-backend/comments"
	// "github.com/asrulkadir/section-comment-backend/config"
	"github.com/asrulkadir/section-comment-backend/pkg/validator"
)

func main() {
	// configViper, err := config.InitViper()
	// if err != nil {
	// 	log.Fatal(err)
	// }

	// dsn := fmt.Sprintf("postgres://%s:%s@%s:%s/%s?sslmode=disable",
	// 	configViper.Database.DBUser,
	// 	configViper.Database.DBPassword,
	// 	configViper.Database.DBHost,
	// 	configViper.Database.DBPort,
	// 	configViper.Database.DBName,
	// )

	// conn, err := sql.Open("postgres", dsn)
	// if err != nil {
	// 	log.Fatal(err)
	// }

	// if err := conn.Ping(); err != nil {
	// 	panic(err)
	// } else {
	// 	fmt.Println("DB Connected...")
	// }

	e := echo.New()
	validator := validator.New()
	e.Validator = validator
	e.Use(middleware.CORSWithConfig(middleware.CORSConfig{
		AllowOrigins: []string{"*"},
		AllowMethods: []string{http.MethodGet, http.MethodPut, http.MethodPost, http.MethodDelete},
	}))
	e.GET("/", func(c echo.Context) error {
		return c.String(http.StatusOK, "Hello, World!")
	})

	//repository
	// commentRepo := comments.NewRepositoryComments(conn)

	// //service
	// commentService := comments.NewServiceComment(commentRepo)

	// //controller
	// comments.NewController(e, commentService, validator)
	port := os.Getenv("PORT")

	e.Logger.Fatal(e.Start(":" + port))
}
