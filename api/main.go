package main

import (
	handlers "avatar4ik3/TextStorage/api/handlers"
	"avatar4ik3/TextStorage/api/models"
	"fmt"
	"os"
	"runtime"

	"github.com/gin-contrib/cors"
	"github.com/gin-gonic/gin"
	"github.com/joho/godotenv"
	"github.com/nullseed/logruseq"
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

func CreateEngineWithoutLogger(routes ...handlers.Route) *gin.Engine {
	router := gin.Default()
	router.Use(cors.Default())
	for _, r := range routes {
		handler := r.Handle()
		router.Handle(handler.Method, handler.Path, handler.Func)
	}
	return router
}

func CreateLogger() *logrus.Logger {
	log := logrus.New()
	log.SetLevel(logrus.DebugLevel)
	log.AddHook(logruseq.NewSeqHook("http://" + os.Getenv("SEQ_NAME") + ":" + os.Getenv("SEQ_PORT")))
	log.Error("Logger Created!")
	return log
}

func CreateStore(logger *logrus.Logger) *models.Store {
	s, err := models.NewStore(
		"postgresql://" + os.Getenv("POSTGRES_NAME") + ":" +
			os.Getenv("POSTGRES_PORT") + "/" +
			os.Getenv("POSTGRES_DB") +
			"?user=" + os.Getenv("POSTGRES_USER") +
			"&password=" + os.Getenv("POSTGRES_PASSWORD") +
			"&sslmode=disable")
	if err != nil {
		panic(err)
	}
	return s
}

func RunWithPort(r *gin.Engine, logger *logrus.Logger) {
	logger.Error("RRRRRRUN")
	r.Run(":" + os.Getenv("APP_PORT"))
}

func main() {
	if runtime.GOOS == "windows" {
		fmt.Println("parsing .env file")

		if err := godotenv.Load("../.debug.env"); err != nil {
			fmt.Println(err.Error())

			os.Exit(1)
		}

	}
	if runtime.GOOS != "windows" {
		fmt.Println("using real env variables")
	}

	fx.
		New(
			// fx.WithLogger(
			// 	fx.Annotate(
			// 		func(logger *logrus.Logger) fxevent.Logger {
			// 			zaplogger, _ := zap.NewDevelopment()
			// 			hook, _ := logrus_zap_hook.NewZapHook(zaplogger)
			// 			logger.Hooks.Add(hook)
			// 			return &fxevent.ZapLogger{Logger: zaplogger}
			// 		},
			// 		fx.ParamTags(`name:"logger"`),
			// 	),
			// ),
			fx.Provide(
				CreateLogger,
				CreateStore,
				models.NewRepository,
				// fx.Annotate(
				// 	CreateLogger,
				// 	fx.ResultTags(`name:"logger"`),
				// ),
				fx.Annotate(
					CreateEngineWithoutLogger,
					fx.ParamTags(`group:"routes"`),
					fx.ResultTags(`name:"engine"`),
				),

				AsRoute(handlers.NewEchoHandler),
				AsRoute(handlers.NewPingPongHandler),
				AsRoute(handlers.NewAddTextHandler),
				AsRoute(handlers.NewGetAllTextsHandler),
				AsRoute(handlers.NewRemoveTextHandler),
			),
			fx.Invoke(
				fx.Annotate(
					RunWithPort,
					fx.ParamTags(`name:"engine"`),
				),
			),
		).
		Run()

}
