package api

import (
	"fmt"
	"net/http"

	"github.com/MykytaPopov/try-go-rest-api/internal/app/repository"
	"github.com/gorilla/mux"
	"github.com/sirupsen/logrus"
)

type Api struct {
	config     *Config
	logger     *logrus.Logger
	router     *mux.Router
	repository *repository.Repository
}

func NewApi(config *Config) *Api {
	return &Api{
		config: config,
		logger: logrus.New(),
		router: mux.NewRouter(),
	}
}

func (a *Api) Start() error {
	if err := a.ConfigureLogger(); err != nil {
		return err
	}

	a.ConfigureRouter()

	if err := a.ConfigureRepository(); err != nil {
		return err
	}

	return http.ListenAndServe(a.config.GetAddress(), a.router)
}

func (a *Api) ConfigureLogger() error {
	level, err := logrus.ParseLevel(a.config.LogLevel)
	if err != nil {
		return err
	}

	a.logger.SetLevel(level)

	return nil
}

func (a *Api) ConfigureRouter() {
	a.router.HandleFunc("/hello/{id:[0-9]+}", a.hello)
}

func (a *Api) ConfigureRepository() error {
	r := repository.NewRepository(a.config.Repository)
	if err := r.Open(); err != nil {
		return err
	}

	a.repository = r

	return nil
}

func (a *Api) hello(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	id := vars["id"]
	response := fmt.Sprintf("hello to you %s times", id)
	fmt.Fprint(w, response)
}
