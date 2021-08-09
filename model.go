package main

import (
	"database/sql"
)

type territory struct {
	Code          string `json:"code"`
	Name          string `json:"name"`
	Timezone      string `json:"timezone"`
	Population    int    `json:"population"`
	Main_language string `json:"main_language"`
	Gallery       string `json:"gallery"`
	Flag          string `json:"flag"`
	Map           string `json:"map"`
	Background    string `json:"background"`
	Location      string `json:"location"`
}

func (tr *territory) getTerritory(db *sql.DB) error {
	return db.QueryRow(
		`SELECT code, name, timezone, population, main_language, gallery, flag, map, background, location
		 FROM territories WHERE code=$1`, tr.Code).Scan(
		&tr.Code, &tr.Name, &tr.Timezone, &tr.Population, &tr.Main_language,
		&tr.Gallery, &tr.Flag, &tr.Map, &tr.Background, &tr.Location)
}

func getTerritories(db *sql.DB, start, count int) ([]territory, error) {
	rows, err := db.Query(
		`SELECT code, name, timezone, population, main_language, gallery, flag, map, background, location
		 FROM territories LIMIT $1 OFFSET $2`,
		count, start)

	if err != nil {
		return nil, err
	}

	defer rows.Close()

	territories := []territory{}

	for rows.Next() {
		var tr territory
		if err := rows.Scan(&tr.Code, &tr.Name, &tr.Timezone, &tr.Population, &tr.Main_language,
			&tr.Gallery, &tr.Flag, &tr.Map, &tr.Background, &tr.Location); err != nil {
			return nil, err
		}
		territories = append(territories, tr)
	}

	return territories, nil
}
