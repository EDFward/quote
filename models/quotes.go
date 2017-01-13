package models

import (
	"fmt"
	"strings"
)

type Quote struct {
	ID      int    `json:"id"`
	Content string `json:"content"`
	Author  string `json:"author"`
	// Nullable fields.
	Source  *string `json:"source,omitempty"`
	Section *string `json:"section,omitempty"`
	// TODO: Generated image.
}

const (
	quoteTable = "quotes"
)

var (
	sqlGetQuotes   string
	sqlInsertQuote string
)

// Prepare querie.
func init() {
	toGet := []string{"id", "content", "author", "source", "section"}
	sqlGetQuotes = fmt.Sprintf(
		"SELECT %s FROM %s WHERE user_id = $1 AND platform = $2",
		strings.Join(toGet, ","), quoteTable)

	toInsert := append([]string{"user_id", "platform"}, toGet[1:]...)
	sqlInsertQuote = fmt.Sprintf(
		"INSERT INTO %s (%s) VALUES ($1, $2, $3, $4, $5, $6) RETURNING ID",
		quoteTable, strings.Join(toInsert, ","))
}

func GetQuotes(user, platform string) ([]Quote, error) {
	rows, err := db.Query(sqlGetQuotes, user, platform)
	if err != nil {
		return nil, err
	}

	defer rows.Close()

	res := make([]Quote, 0)
	for rows.Next() {
		var quote Quote
		err = rows.Scan(
			&quote.ID, &quote.Content, &quote.Author, &quote.Source,
			&quote.Section)
		if err != nil {
			// TODO: Log and skip.
			return nil, err
		}

		res = append(res, quote)
	}
	return res, nil
}

func AddQuote(user, platform, content, author string, source, section *string) (int64, error) {
	var id int64
	err := db.QueryRow(
		sqlInsertQuote, user, platform, content, author, source, section,
	).Scan(&id)
	if err != nil {
		return 0, err
	}
	return id, nil
}
