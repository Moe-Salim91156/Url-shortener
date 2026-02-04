/* ************************************************************************** */
/*                                                                            */
/*                                                        :::      ::::::::   */
/*   url_store.go                                       :+:      :+:    :+:   */
/*                                                    +:+ +:+         +:+     */
/*   By: moe <marvin@42.fr>                         +#+  +:+       +#+        */
/*                                                +#+#+#+#+#+   +#+           */
/*   Created: 2026/02/04 16:21:54 by moe               #+#    #+#             */
/*   Updated: 2026/02/04 16:23:23 by moe              ###   ########.fr       */
/*                                                                            */
/* ************************************************************************** */

package sqlite

import (
	"Url-shortener/internal/models"
	"database/sql"
	"time"
)

type SQLiteStore struct {
	db *sql.DB
}

func NewSQLiteURLStore(db *sql.DB) *SQLiteStore {
	return &SQLiteStore{db: db}
}

// now with the interfaces functions , CRUD ?
func (s *SQLiteStore) Save(data models.UrlData) error {
	// insert into URL table
	sqlstmt := `INSERT INTO urls (short_code, long_url, owner_id, creation_time) VALUES (?, ?, ?, ?)`

	// formating the timie to a proper format RFC3339
	timeStr := data.CreationTime.Format(time.RFC3339)

	if _, err := s.db.Exec(sqlstmt, data.ShortCode, data.LongUrl, data.OwnerID, timeStr); err != nil {
		return err
	}
	return nil
}
func (s *SQLiteStore) Get(shortCode string) (*models.UrlData, error) {
	var data models.UrlData
	var timeStr string

	query := `SELECT short_code, long_url, owner_id, creation_time FROM urls WHERE short_code = ?`

	err := s.db.QueryRow(query, shortCode).Scan(
		&data.ShortCode,
		&data.LongUrl,
		&data.OwnerID,
		&timeStr,
	)

	if err == sql.ErrNoRows {
		return nil, fmt.Errorf("URL NOT FOUND")
	}
	if err != nil {
		return nil, err
	}

	data.CreationTime, _ = time.Parse(time.RFC3339, timeStr)
	return &data, nil
}
func (s *SQLiteStore) GetByOwner(ownerID string) ([]models.UrlData, error) {
	var url models.UrlData
	var timeStr string

	query := `SELECT short_code, long_url, owner_id, creation_time FROM urls WHERE owner_id = ?`

	rows, err := s.db.Query(query, ownerID)
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	var urls []models.UrlData

	for rows.Next() {
		err := rows.Scan(&url.ShortCode, &url.LongUrl, &url.OwnerID, &timeStr)
		if err != nil {
			return nil, err
		}

		url.CreationTime, _ = time.Parse(time.RFC3339, timeStr)
		urls = append(urls, url)
	}

	return urls, rows.Err()
}

func Delete(shortCode string) error {
	return nil
}
