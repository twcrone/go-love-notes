package service

import "github.com/heroku/go-getting-started/database"

type Note struct {
	Id      int    `json:"id"`
	Message string `json:"message"`
}

func GetAllNotes() ([]Note, error) {
	results, err := database.DbConn.Query(`select id, message from notes`)
	if err != nil {
		return nil, err
	}
	defer results.Close()
	notes := make([]Note, 0)
	for results.Next() {
		var note Note
		_ = results.Scan(&note.Id, &note.Message)
		notes = append(notes, note)
	}
	return notes, nil
}
