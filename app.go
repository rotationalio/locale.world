package main

import (
	"database/sql"
	"fmt"
	"log"

	"encoding/json"
	"net/http"
	"strconv"

	"github.com/gorilla/mux"
	_ "github.com/lib/pq"
)

type App struct {
	Router *mux.Router
	DB     *sql.DB
}

func (a *App) Initialize(user, password, dbname string) {
	connectionString :=
		fmt.Sprintf("user=%s password=%s dbname=%s sslmode=disable", user, password, dbname)

	var err error
	a.DB, err = sql.Open("postgres", connectionString)
	if err != nil {
		log.Fatal(err)
	}

	a.Router = mux.NewRouter()

	a.initializeRoutes()
}

func (a *App) ensureTableExists() {
	if _, err := a.DB.Exec(tableCreationQuery); err != nil {
		log.Fatal(err)
	}
}

const tableCreationQuery = `CREATE TABLE IF NOT EXISTS territories
(
		id SERIAL,
		code text NOT NULL,
		name text NOT NULL,
		gdp text,
		literacyPct real,
		population int,
		languages text[][3],
		CONSTRAINT territories_pkey PRIMARY KEY (id)
)`

func (a *App) clearTable() {
	a.DB.Exec("DELETE FROM territories")
	a.DB.Exec("ALTER SEQUENCE territories_id_seq RESTART WITH 1")
}

func (a *App) Run(addr string) {
	log.Fatal(http.ListenAndServe(":8010", a.Router))
}

func (a *App) getTerritory(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	id, err := strconv.Atoi(vars["id"])
	if err != nil {
		respondWithError(w, http.StatusBadRequest, "Invalid territory ID")
		return
	}

	tr := territory{ID: id}
	if err := tr.getTerritory(a.DB); err != nil {
		switch err {
		case sql.ErrNoRows:
			respondWithError(w, http.StatusNotFound, "Territory not found")
		default:
			respondWithError(w, http.StatusInternalServerError, err.Error())
		}
		return
	}

	respondWithJSON(w, http.StatusOK, tr)
}

func respondWithError(w http.ResponseWriter, code int, message string) {
	respondWithJSON(w, code, map[string]string{"error": message})
}

func respondWithJSON(w http.ResponseWriter, code int, payload interface{}) {
	response, _ := json.Marshal(payload)

	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(code)
	w.Write(response)
}

func (a *App) getTerritories(w http.ResponseWriter, r *http.Request) {
	count, _ := strconv.Atoi(r.FormValue("count"))
	start, _ := strconv.Atoi(r.FormValue("start"))

	if count > 10 || count < 1 {
		count = 10
	}
	if start < 0 {
		start = 0
	}

	territories, err := getTerritories(a.DB, start, count)
	if err != nil {
		respondWithError(w, http.StatusInternalServerError, err.Error())
		return
	}

	respondWithJSON(w, http.StatusOK, territories)
}

func (a *App) createTerritory(w http.ResponseWriter, r *http.Request) {
	var tr territory
	decoder := json.NewDecoder(r.Body)
	if err := decoder.Decode(&tr); err != nil {
		respondWithError(w, http.StatusBadRequest, "Invalid request payload")
		return
	}
	defer r.Body.Close()

	if err := tr.createTerritory(a.DB); err != nil {
		respondWithError(w, http.StatusInternalServerError, err.Error())
		return
	}

	respondWithJSON(w, http.StatusCreated, tr)
}

func (a *App) updateTerritory(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	id, err := strconv.Atoi(vars["id"])
	if err != nil {
		respondWithError(w, http.StatusBadRequest, "Invalid territory ID")
		return
	}

	var tr territory
	decoder := json.NewDecoder(r.Body)
	if err := decoder.Decode(&tr); err != nil {
		respondWithError(w, http.StatusBadRequest, "Invalid resquest payload")
		return
	}
	defer r.Body.Close()
	tr.ID = id

	if err := tr.updateTerritory(a.DB); err != nil {
		respondWithError(w, http.StatusInternalServerError, err.Error())
		return
	}

	respondWithJSON(w, http.StatusOK, tr)
}

func (a *App) deleteTerritory(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	id, err := strconv.Atoi(vars["id"])
	if err != nil {
		respondWithError(w, http.StatusBadRequest, "Invalid territory ID")
		return
	}

	tr := territory{ID: id}
	if err := tr.deleteTerritory(a.DB); err != nil {
		respondWithError(w, http.StatusInternalServerError, err.Error())
		return
	}

	respondWithJSON(w, http.StatusOK, map[string]string{"result": "success"})
}

func (a *App) initializeRoutes() {
	a.Router.HandleFunc("/territories", a.getTerritories).Methods("GET")
	a.Router.HandleFunc("/territory", a.createTerritory).Methods("POST")
	a.Router.HandleFunc("/territory/{id:[0-9]+}", a.getTerritory).Methods("GET")
	a.Router.HandleFunc("/territory/{id:[0-9]+}", a.updateTerritory).Methods("PUT")
	a.Router.HandleFunc("/territory/{id:[0-9]+}", a.deleteTerritory).Methods("DELETE")
}
