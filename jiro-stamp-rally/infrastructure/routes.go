package infrastructure

import (
	"fmt"
	"net/http"
	"time"

	dbconfig "github.com/fesInformatics/jiro-stamp-rally/infrastructure/config"
	"github.com/fesInformatics/jiro-stamp-rally/infrastructure/db"
	repository "github.com/fesInformatics/jiro-stamp-rally/infrastructure/gateway"
	"github.com/fesInformatics/jiro-stamp-rally/interface/context"
	"github.com/fesInformatics/jiro-stamp-rally/interface/controller"
	"github.com/fesInformatics/jiro-stamp-rally/usecase/interactor"
)

func Run() {
	http.HandleFunc("/health", healthCheck)
	http.HandleFunc("/user/register", createUser)

	s := http.Server{
		Addr: ":8080",
		ReadTimeout: 30 * time.Second,
		WriteTimeout: 90 * time.Second,
		IdleTimeout: 90 * time.Second,
	}

	fmt.Println("Server Start")

	err := s.ListenAndServe()
	if err != nil {
		if err != http.ErrServerClosed {
			panic(err)
		}
	}
}

func healthCheck(w http.ResponseWriter, r *http.Request) {
	controller := controller.NewHealthcheckController()
	ctx := context.NewContext(w, r)
	controller.Get(ctx)
}

func createUser(w http.ResponseWriter, r *http.Request) {
	controller := controller.NewUserController(interactor.NewUserInteactor(repository.NewUserRespository(client.NewDbClient(dbconfig.NewDBConfig()))))
	ctx := context.NewContext(w, r)
	controller.Create(ctx)
}
