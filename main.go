package main

import (
	"log"
	"net/http"
	"time"

	appValidator "github.com/0xTatsu/g-api/handler/validator"
	"github.com/go-chi/chi/v5"
	"github.com/go-chi/chi/v5/middleware"
	"github.com/go-chi/render"
	"go.uber.org/zap"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"

	"github.com/0xTatsu/g-api/config"
	"github.com/0xTatsu/g-api/handler"
	"github.com/0xTatsu/g-api/jwt"
	"github.com/0xTatsu/g-api/model"
	"github.com/0xTatsu/g-api/repo"
)

func main() {
	_, undoReplaceGlobalLog := initLogger()
	defer undoReplaceGlobalLog()

	envCfg, err := config.New()
	if err != nil {
		log.Print("cannot load config:", err)
		return
	}

	db := initDB(envCfg.DBURL)

	r := chi.NewRouter()
	r.Use(middleware.RequestID)
	r.Use(middleware.RealIP)
	r.Use(middleware.Logger)
	r.Use(middleware.Recoverer)
	r.Use(middleware.Heartbeat("/ping"))
	r.Use(middleware.Timeout(time.Second * time.Duration(envCfg.ServerTimeout)))
	r.Use(render.SetContentType(render.ContentTypeJSON))

	authJWT := jwt.NewJWT(envCfg)
	userRepo := repo.NewUser(db)
	authAPI := handler.NewAuth(authJWT, userRepo, envCfg, appValidator.New())

	r.Mount("/auth", authAPI.Router(r))
	r.Group(func(r chi.Router) {
		r.Use(authJWT.Verifier())
		r.Use(jwt.Authenticator)

		// r.Get("/admin", func(w http.ResponseWriter, r *http.Request) {
		// 	_, claims, _ := jwtauth.FromContext(r.Context())
		// 	fmt.Println(claims) // nolint
		// 	render.JSON(w, r, http.NoBody)
		// })
	})

	if err := http.ListenAndServe(":"+envCfg.ServerPort, r); err != nil {
		log.Panicf("cannot start server: %s", err)
	}
}

func initLogger() (*zap.Logger, func()) {
	logger, errZapLog := zap.NewDevelopment()
	if errZapLog != nil {
		log.Fatalf("cannot init ZAP log: %s", errZapLog)
	}

	undoReplaceGlobalLog := zap.ReplaceGlobals(logger)

	return logger, undoReplaceGlobalLog
}

func initDB(dsn string) *gorm.DB {
	db, err := gorm.Open(postgres.Open(dsn), &gorm.Config{})
	if err != nil {
		log.Fatal("cannot connect db: ", err)
	}

	if err = db.AutoMigrate(&model.User{}); err != nil {
		log.Fatal("failed to migrate db: ", err)
	}

	return db
}
