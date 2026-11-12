package db

func (d *Database) migrate() error {
	// Create the standard text table
	_, err := d.db.Exec(`
	CREATE TABLE IF NOT EXISTS tol_entries (
		id INTEGER PRIMARY KEY AUTOINCREMENT,
		command TEXT NOT NULL,
		description TEXT NOT NULL
	);
	`)
	if err != nil {
		return err
	}

	// Create the vector shadow table
	_, err = d.db.Exec(`
	CREATE VIRTUAL TABLE IF NOT EXISTS vec_tol_entries USING vec0(
		embedding float[768] distance_metric=cosine
	);
	`)
	return err
}
