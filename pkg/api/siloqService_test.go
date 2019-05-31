package api

import (
	"os"
	"net/http/httptest"
	"net/http"
	"testing"

	"github.com/gorilla/mux"

	"siloqcrud/pkg/db"
	"siloqcrud/pkg/config"
)

func ReadJokeRouter(db db.DataStore, config config.Config) *mux.Router {
        r := mux.NewRouter()
        r.Handle("/", GetNamedJoke(db, config)).Methods("GET")
        return r
}

func TestHealthCheck(t *testing.T) {
        req := httptest.NewRequest("GET", "/health", nil)
        w := httptest.NewRecorder()

        HealthCheck(w, req)

        if w.Body.String() != "OK" || w.Code != 200 {
                t.Error("healthcheck did not return 200 OK as expected.")
        }
}

func TestGetNamedJoke(t *testing.T) {
        dao := db.InitializeMock()
	config, err := config.GetConfig(os.Getenv("CONFIG_FILE_PATH"))
	//Include params if necessary
        r, err := http.NewRequest("GET", "/", nil)
        if err != nil {
                t.Error(err.Error())
        }

        w := httptest.NewRecorder()
        ReadJokeRouter(dao, config).ServeHTTP(w, r)
        if w.Code != http.StatusOK {
                t.Errorf("expected %v, got %v", http.StatusOK, w.Code)
        }
}
