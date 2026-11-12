package db

import (
	"encoding/json"
)

type SearchResult struct {
	Command     string
	Description string
}

func (d *Database) InsertCommand(command, description string, embedding []float32) error {
	// 1. Insert the text
	res, err := d.db.Exec(`INSERT INTO tol_entries (command, description) VALUES (?, ?)`, command, description)
	if err != nil {
		return err
	}
	
	id, err := res.LastInsertId()
	if err != nil {
		return err
	}

	embJSON, _ := json.Marshal(embedding)
	_, err = d.db.Exec(`INSERT INTO vec_tol_entries (rowid, embedding) VALUES (?, ?)`, id, string(embJSON))
	return err
}

func (d *Database) SearchKNN(embedding []float32, limit int) (SearchResult, error) {
	embJSON, _ := json.Marshal(embedding)

	query := `
	SELECT t.command, t.description 
	FROM tol_entries t
	JOIN vec_tol_entries v ON t.id = v.rowid
	WHERE v.embedding MATCH ? AND k = ?
	ORDER BY v.distance`
	
	row := d.db.QueryRow(query, string(embJSON), limit)
	
	var res SearchResult
	err := row.Scan(&res.Command, &res.Description)
	if err != nil {
		if err.Error() == "sql: no rows in result set" {
			return res, nil
		}
		return res, err
	}
	return res, nil
}