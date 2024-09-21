package models

import (
	"database/sql"

	"go-web/pkg/models/snippets"
)

type PostModel struct {
	DB *sql.DB
}

func (m *PostModel) Insert(title, content string) int {
	return 0
}

func (m *PostModel) Get(id int) (*snippets.Post, error) {
	return nil, nil
}
