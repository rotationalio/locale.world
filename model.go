package main

import (
	"database/sql"
)

type territory struct {
	ID          int        `json:"id"`
	Code        string     `json:"code"`
	Name        string     `json:"name"`
	Gdp         string     `json:"gdp"`
	LiteracyPct float32    `json:"literacyPct"`
	Population  int        `json:"population"`
	Languages   [][]string `json:"languages"`
}

func (tr *territory) getTerritory(db *sql.DB) error {
	return db.QueryRow(
		"SELECT code, name, gdp, literacyPct, population, languages FROM territories WHERE id=$1",
		tr.ID).Scan(&tr.Code, &tr.Name, &tr.Gdp, &tr.LiteracyPct, &tr.Population, &tr.Languages)
}

func (tr *territory) updateTerritory(db *sql.DB) error {
	_, err :=
		db.Exec(
			"UPDATE territories SET code=$1, name=$2, gdp=$3, literacyPct=$4, population=$5, languages=$6 WHERE id=$7",
			tr.Code, tr.Name, tr.Gdp, tr.LiteracyPct, tr.Population, tr.Languages, tr.ID)

	return err
}

func (tr *territory) deleteTerritory(db *sql.DB) error {
	_, err := db.Exec("DELETE FROM territories WHERE id=$1", tr.ID)

	return err
}

func (tr *territory) createTerritory(db *sql.DB) error {
	err := db.QueryRow(
		"INSERT INTO territories(code, name, gdp, literacyPct, population, languages) VALUES($1, $2, $3, $4, $5, $6) RETURNING id",
		tr.Code, tr.Name, tr.Gdp, tr.LiteracyPct, tr.Population, tr.Languages).Scan(&tr.ID)

	if err != nil {
		return err
	}

	return nil
}

func getTerritories(db *sql.DB, start, count int) ([]territory, error) {
	rows, err := db.Query(
		"SELECT code, name, gdp, literacyPct, population, languages FROM territories LIMIT $1 OFFSET $2",
		count, start)

	if err != nil {
		return nil, err
	}

	defer rows.Close()

	territories := []territory{}

	for rows.Next() {
		var tr territory
		if err := rows.Scan(&tr.ID, &tr.Code, &tr.Name); err != nil {
			return nil, err
		}
		territories = append(territories, tr)
	}

	return territories, nil
}
