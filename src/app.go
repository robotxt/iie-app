package src

import (
	"context"
	"net/http"
	"os"
	"os/signal"
	"strings"
	"syscall"
	"time"

	"github/robotxt/iie-app/src/api"
	"github/robotxt/iie-app/src/logging"
	repo "github/robotxt/iie-app/src/repo/firebase"

	"github.com/gorilla/mux"
)

type App struct {
	Router *mux.Router
}

var env string = strings.ToUpper(os.Getenv("ENVIRONMENT"))
var router *mux.Router = mux.NewRouter().StrictSlash(false)

func (a *App) Initialize() {

	port := os.Getenv("PORT")
	if port == "" {
		// set default port if empty
		port = "8080"
	}

	logging.Info("port: ", port)
	ctx := context.Background()

	// Initialize firebase
	firebase := &repo.FirebaseApp{}
	if env == "PRODUCTION" {
		_ = firebase.InitializeProdFirebase(ctx)
	} else {
		_ = firebase.InitializeFirebase(ctx)
	}

	firebaseApp := firebase.InitializeAllFirebaseService(ctx)

	apiv1 := api.ApiV1{}
	apiv1.Router = router
	apiv1.Firebase = firebaseApp
	apiv1.Ctx = ctx
	apiv1.SetRouters()

	srv := &http.Server{
		Handler:      router,
		Addr:         ":" + port,
		ReadTimeout:  10 * time.Second, // timeout in data read
		WriteTimeout: 10 * time.Second, // timeout for response
	}

	go func() {
		logging.Info("Starting Server")
		if err := srv.ListenAndServe(); err != nil {
			logging.Fatal(err)
		}
	}()
	waitForShutdown(srv)

}

func checkErr(err error) {
	if err != nil {
		logging.Fatal(err)
	}
}

func waitForShutdown(srv *http.Server) {
	interruptChan := make(chan os.Signal, 1)
	signal.Notify(interruptChan, os.Interrupt, syscall.SIGINT, syscall.SIGTERM)

	// Block until we receive our signal.
	<-interruptChan

	// Create a deadline to wait for.
	ctx, cancel := context.WithTimeout(context.Background(), time.Second*10)
	defer cancel()
	srv.Shutdown(ctx)

	logging.Info("Shutting down")
	os.Exit(0)
}
