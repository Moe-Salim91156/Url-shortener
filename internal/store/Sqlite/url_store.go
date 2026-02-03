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

	query := `SELECT short_code , long_url, owner_id , creation_time FROM urls WHERE short_code = ? `
	rows, err := s.db.Query(query)
	if err != nil {
		return nil, err
	}
	defer rows.Close()
	var timeStr string
	for rows.Next() {
		err = rows.Scan(
			&data.ShortCode,
			&data.LongUrl,
			&data.OwnerID,
			&timeStr,
		)
	}
	data.CreationTime, _ = time.Parse(time.RFC3339, timeStr)
	return &data, nil
}

func GetByOwner(OwnerID string) ([]models.UrlData, error) {
	// loop through urls table
	// any url matches the ownerID provided, append it to resuls slice
	// return a slice with all urls that this owner OWNS
	query := `SELECT FROM urls WHERE owner_id = ?`
	return nil, nil
}

func Delete(shortCode string) error {
	return nil
}
