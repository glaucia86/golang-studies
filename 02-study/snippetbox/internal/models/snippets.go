/**
 * file: internal/models/snippets.go
 * description: file responsible for the model of the application.
 * data: 01/01/2024
 * author: Glaucia Lemos <Twitter: @glaucia_lemos86>
 */

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
	return 0, nil
}

func (m *SnippetModel) Get(id int) (Snippet, error) {
	return Snippet{}, nil
}

func (m *SnippetModel) LatestSnippets() ([]Snippet, error) {
	return nil, nil
}
