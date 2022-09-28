package handlers

import (
	"encoding/json"
	"io"
	"net/http"

	"github.com/SowinskiBraeden/SulliCartShare/api"
	"github.com/SowinskiBraeden/SulliCartShare/models"
	"github.com/gorilla/mux"
	"go.uber.org/zap"

	"github.com/SowinskiBraeden/SulliCartShare/config"
	"github.com/SowinskiBraeden/SulliCartShare/databases"
)

// App stores the router and db connection so it can be reused
type App struct {
	Router   *mux.Router
	DB       databases.CollectionHelper
	Config   config.Config
	dbHelper databases.DatabaseHelper
}

// New creates a new mux router and all the routes
func (a *App) New() *mux.Router {
	r := mux.NewRouter()
	cow := Cow{DB: databases.NewCowDatabase(a.dbHelper)}

	// healthcheck
	r.HandleFunc("/health", healthCheckHandler)

	apiCreate := r.PathPrefix("/api/v1").Subrouter()

	apiCreate.Handle("/cow/{cow_id}", api.Middleware(http.HandlerFunc(cow.CowByIDHandler))).Methods("GET") // By Object ID not CowCode
	apiCreate.Handle("/cows", api.Middleware(http.HandlerFunc(cow.CowHandler))).Methods("GET")
	apiCreate.Handle("/cows/new", api.Middleware(http.HandlerFunc(cow.NewCowHandler))).Methods("POST")
	apiCreate.Handle("/cows/update/{cow_id}", api.Middleware(http.HandlerFunc(cow.UpdateCowHandler))).Methods("POST") // By Object ID not CowCode

	return r
}

func (a *App) Initialize() error {
	client, err := databases.NewClient(&a.Config)
	if err != nil {
		// if we fail to create a new database client, the kill the pod
		zap.S().With(err).Error("failed to create new client")
		return err
	}

	a.dbHelper = databases.NewDatabase(&a.Config, client)
	err = client.Connect()
	if err != nil {
		// if we fail to connect to the database, the kill the pod
		zap.S().With(err).Error("failed to connect to database")
		return err
	}
	zap.S().Info("SulliCartCheckout has connected to the database")

	// initialize api router
	a.initializeRoutes()
	return nil

}

func (a *App) initializeRoutes() {
	a.Router = a.New()
}

func healthCheckHandler(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusOK)
	b, _ := json.Marshal(models.HealthCheckResponse{
		Alive: true,
	})
	_, _ = io.WriteString(w, string(b))
}