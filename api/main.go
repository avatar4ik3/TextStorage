package main

import (
	handlers "avatar4ik3/TextStorage/api/handlers"
	textHandlers "avatar4ik3/TextStorage/api/handlers/text"
	"avatar4ik3/TextStorage/api/middleware"
	"avatar4ik3/TextStorage/api/models"
	"fmt"
	"os"
	"runtime"

	"github.com/gin-contrib/cors"
	"github.com/gin-gonic/gin"
	"github.com/joho/godotenv"
	"github.com/sirupsen/logrus"
	"go.uber.org/fx"
)

func AsRoute(f any) any {
	return fx.Annotate(
		f,
		fx.As(new(handlers.Route)),
		fx.ResultTags(`group:"routes"`),
	)
}

func AsMiddleware(f any) any {
	return fx.Annotate(
		f,
		fx.ResultTags(`group:"middleware"`),
	)
}

func RegisterRoutes(r *gin.Engine, routes []handlers.Route) {
	for _, route := range routes {
		handler := route.Handle()
		r.Handle(handler.Method, handler.Path, handler.Func)
	}
}

func RegisterMiddlewares(r *gin.Engine, mws []gin.HandlerFunc) {
	fmt.Printf("using mws %d \n", len(mws))
	for _, mw := range mws {
		r.Use(mw)
	}
}

func CreateLogger() *logrus.Logger {
	log := logrus.New()
	log.SetLevel(logrus.DebugLevel)
	// log.AddHook(logruseq.NewSeqHook(SEQ_connectionString()))
	log.Error("Logger Created!")
	return log
}

func SEQ_connectionString() string {
	return "http://" + os.Getenv("SEQ_NAME") + ":" + os.Getenv("SEQ_PORT")
}

func PSG_connectionString() string {
	return "postgresql://" + os.Getenv("POSTGRES_NAME") + ":" +
		os.Getenv("POSTGRES_PORT") + "/" +
		os.Getenv("POSTGRES_DB") +
		"?user=" + os.Getenv("POSTGRES_USER") +
		"&password=" + os.Getenv("POSTGRES_PASSWORD") +
		"&sslmode=disable"
}

func RunWithPort(r *gin.Engine, logger *logrus.Logger) {
	r.Run(":" + os.Getenv("APP_PORT"))
}

func LoadEnv(logger *logrus.Logger) {
	if runtime.GOOS == "windows" {
		logger.Info("parsing .env file")
		if err := godotenv.Load("../.debug.env"); err != nil {
			logger.Error(err.Error())
			os.Exit(1)
		}

	}
	if runtime.GOOS != "windows" {
		logger.Info("using real env variables")
	}
}

func CreateEngine(mws []gin.HandlerFunc) *gin.Engine {
	r := gin.Default()
	//todo разобратька как регистрировать мвшки после сборки fx'a
	RegisterMiddlewares(r, mws)
	return r
}

func CreateStore() *models.Store {
	store, err := models.NewStore(PSG_connectionString())
	if err != nil {
		panic(err.Error())
	}
	return store
}

func main() {

	fx.
		New(
			fx.Provide(
				CreateLogger,
				CreateStore,

				models.NewRepository,

				fx.Annotate(
					CreateEngine,
					fx.ResultTags(`name:"engine"`),
					fx.ParamTags(`group:"middleware"`),
				),

				AsRoute(handlers.NewEchoHandler),
				AsRoute(handlers.NewPingPongHandler),
				AsRoute(textHandlers.NewAddTextHandler),
				AsRoute(textHandlers.NewGetAllTextsHandler),
				AsRoute(textHandlers.NewRemoveTextHandler),
				AsRoute(textHandlers.NewRemoveTextByIdHandler),

				AsMiddleware(cors.Default),
				AsMiddleware(middleware.NewErrorHandler),
			),
			fx.Invoke(
				LoadEnv,
				fx.Annotate(
					RegisterRoutes,
					fx.ParamTags(`name:"engine"`, `group:"routes"`),
				),
				fx.Annotate(
					RunWithPort,
					fx.ParamTags(`name:"engine"`),
				),
			),
		).
		Run()

}
