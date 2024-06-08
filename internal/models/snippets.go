package models

import (
	"database/sql"
	"time"
)

type Snippet struct {
	ID      int
	Title   string
	Content string
	Created time.Time
	Expires time.Time
}

type SnippetModel struct {
	DB *sql.DB
}

func (m *SnippetModel) Insert(title string, content string, expires int) (int, error) {

	// Write the SQL statement we want to execute.
	stmt := `INSERT INTO snippets (title, content, created, expires)
	VALUES($1, $2, CURRENT_TIMESTAMP, CURRENT_TIMESTAMP + INTERVAL '1 DAY' * $3)
	RETURNING id`

	// Use the Exec() method on the embedded connection pool to execute the
	// statement. The first parameter is the SQL statement, followed by the
	// values for the placeholder parameters: title, content and expiry in
	// that order. This method returns a sql.Result type, which contains some
	// basic information about what happened when the statement was executed.
	var id int
	err := m.DB.QueryRow(stmt, title, content, expires).Scan(&id)
	if err != nil {
		return 0, err
	}

	return id, nil

}

func (m *SnippetModel) Get(id int) (*Snippet, error) {
	return &Snippet{}, nil
}

func (m *SnippetModel) Latest() ([]*Snippet, error) {
	return nil, nil
}
