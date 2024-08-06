package db

// TrainingSessionDelete deletes a session from the database
func TrainingSessionDelete(session_uuid, email string) (err error) {
	query := `
		DELETE FROM ` + DB.TableTrainingSessions + `
		WHERE session_uuid = $1 AND email = $2;`

	_, err = DB.Database.Exec(query, session_uuid, email)
	if err != nil {
		return
	}

	return
}
