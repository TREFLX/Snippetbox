package mysql

import (
	"database/sql"
	"errors"
	"fmt"

	"golangify.com/snippetbox/pkg/models"
)

type NotesModel struct {
	DB *sql.DB
}

// Insert - А method adds the received data to a new row in the database
func (m *NotesModel) Insert(title, content, expires string) (int, error) {
	stmt := `INSERT INTO snippets (title, content, created, expires)
    VALUES(?, ?, UTC_TIMESTAMP(), DATE_ADD(UTC_TIMESTAMP(), INTERVAL ? DAY))`
	result, err := m.DB.Exec(stmt, title, content, expires)
	if err != nil {
		return 0, err
	}
	id, err := result.LastInsertId()
	if err != nil {
		return 0, err
	}
	return int(id), nil

}

// Get - A method for returning note data by its ID.
func (m *NotesModel) Get(id int) (*models.Notes, error) {
	stmt := `SELECT id, title, content, created, expires FROM snippets
    WHERE expires > UTC_TIMESTAMP() AND id = ?`
	row := m.DB.QueryRow(stmt, id)
	s := &models.Notes{}
	err := row.Scan(&s.ID, &s.Title, &s.Content, &s.Created, &s.Expires)
	if err != nil {
		if errors.Is(err, sql.ErrNoRows) {
			return nil, models.ErrNoRecord
		} else {
			return nil, err
		}
	}
	return s, nil
}

// Latest - The method returns the 10 most frequently used notes.
func (m *NotesModel) Latest() ([]*models.Notes, error) {
	stmt := `SELECT id, title, content, created, expires FROM snippets
    WHERE expires > UTC_TIMESTAMP() ORDER BY created DESC LIMIT 10`

	rows, err := m.DB.Query(stmt)
	if err != nil {
		return nil, err
	}
	defer rows.Close()
	var snippets []*models.Notes
	for rows.Next() {
		s := &models.Notes{}
		err = rows.Scan(&s.ID, &s.Title, &s.Content, &s.Created, &s.Expires)
		if err != nil {
			return nil, err
		}
		snippets = append(snippets, s)
	}
	if err = rows.Err(); err != nil {
		return nil, err
	}
	return snippets, nil
}

// Delete - Checks if there is an entry in the database by id , if there is, it deletes the entry, if there is none, it outputs an error
func (m *NotesModel) Delete(id int) (int, error) {
	var exists bool
	err := m.DB.QueryRow("SELECT EXISTS(SELECT 1 FROM snippets WHERE id = ?)", id).Scan(&exists)
	if err != nil {
		return 0, err
	}
	if !exists {
		return 0, fmt.Errorf("запись с id %d не найдена", id)
	}
	stmt := `DELETE FROM snippets WHERE id = ?`
	result, err := m.DB.Exec(stmt, id)
	if err != nil {
		return 0, err
	}

	rowsAffected, err := result.RowsAffected()
	if err != nil {
		return 0, err
	}

	return int(rowsAffected), nil
}
