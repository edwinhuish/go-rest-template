package router

import (
	"fmt"
	"io"
	"os"

	"github.com/edwinhuish/go-rest-template/internal/api/controllers"
	"github.com/edwinhuish/go-rest-template/internal/api/gin2"
	"github.com/edwinhuish/go-rest-template/internal/api/middlewares"
	"github.com/gin-gonic/gin"
	swaggerFiles "github.com/swaggo/files"
	ginSwagger "github.com/swaggo/gin-swagger"
)

func Setup() *gin.Engine {
	app := gin.New()

	// Logging to a file.
	f, _ := os.Create("log/api.log")
	gin.DisableConsoleColor()
	gin.DefaultWriter = io.MultiWriter(f)

	// Middlewares
	app.Use(gin.LoggerWithFormatter(func(param gin.LogFormatterParams) string {
		return fmt.Sprintf("%s - - [%s] \"%s %s %s %d %s\" \"%s\" \"%s\"\n",
			param.ClientIP,
			param.TimeStamp.Format("02/Jan/2006:15:04:05 -0700"),
			param.Method,
			param.Path,
			param.Request.Proto,
			param.StatusCode,
			param.Latency,
			param.Request.UserAgent(),
			param.ErrorMessage,
		)
	}))
	app.Use(gin.Recovery())
	app.Use(middlewares.CORS())
	app.NoRoute(middlewares.NoRouteHandler())

	// Routes
	// ================== Login Routes
	authCtrl := controllers.NewAuthController()
	app.POST("/api/login", gin2.Cover(authCtrl.Login))
	// ================== Docs Routes
	app.GET("/docs/*any", ginSwagger.WrapHandler(swaggerFiles.Handler))
	// ================== User Routes
	userCtrl := controllers.NewUserController()
	app.GET("/api/users", gin2.Cover(userCtrl.List))
	app.GET("/api/users/:id", gin2.Cover(userCtrl.Find))
	app.POST("/api/users", gin2.Cover(userCtrl.Create))
	app.PUT("/api/users/:id", gin2.Cover(userCtrl.Update))
	app.DELETE("/api/users/:id", gin2.Cover(userCtrl.Delete))
	// ================== Tasks Routes
	taskCtrl := controllers.NewTaskController()
	app.GET("/api/tasks/:id", gin2.Cover(taskCtrl.Find))
	app.GET("/api/tasks", gin2.Cover(taskCtrl.List))
	app.POST("/api/tasks", gin2.Cover(taskCtrl.Create))
	app.PUT("/api/tasks/:id", gin2.Cover(taskCtrl.Update))
	app.DELETE("/api/tasks/:id", gin2.Cover(taskCtrl.Delete))

	return app
}
