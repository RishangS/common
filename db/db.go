package db

import (
	"database/sql"
	"fmt"
	"log"

	"github.com/lib/pq"
	_ "github.com/lib/pq" // PostgreSQL driver

	"github.com/RishangS/common/interfaces"
)

// DBService struct holds the database connection pool
type DBService struct {
	DB *sql.DB
}

// NewDBService initializes the DBService and connects to the database
func NewDBService(user, password, dbname string) (*DBService, error) {
	connStr := fmt.Sprintf("postgres://%s:%s@localhost:5432/%s?sslmode=disable", user, password, dbname)

	// Open connection to the database and check for connection errors
	db, err := sql.Open("postgres", connStr)
	if err != nil {
		return nil, fmt.Errorf("error opening database connection: %v", err)
	}

	// Test connection
	if err := db.Ping(); err != nil {
		return nil, fmt.Errorf("error pinging the database: %v", err)
	}

	log.Println("Successfully connected to the PostgreSQL database!")

	return &DBService{DB: db}, nil
}

// GetDB returns the current database connection
func (db *DBService) GetDB() *sql.DB {
	return db.DB
}

// Close closes the database connection pool
func (db *DBService) Close() {
	if err := db.DB.Close(); err != nil {
		log.Fatalf("Error closing database connection: %v", err)
	}
	log.Println("Database connection closed.")
}

// // Helper function to execute an SQL query and handle errors
// func (ds *DBService) executeQuery(query string, args ...interface{}) (*sql.Rows, error) {
// 	rows, err := ds.DB.Query(query, args...)
// 	if err != nil {
// 		return nil, fmt.Errorf("error executing query: %v", err)
// 	}
// 	return rows, nil
// }

// InsertApp inserts new app data into the database
// InsertApp inserts new app data into the database
func (ds *DBService) InsertApp(app *interfaces.AppData) error {
	query := `INSERT INTO apps 
		(name, locale, country, url, full_url, description, summary, icon, score, 
		price_text, is_free, installs, installs_text, app_version, android_version, min_android_version, 
		size, content_rating, privacy_policy_url, category_id, histogram_rating_id, developer_id, 
		number_voters, number_reviews, recent_changes, editors_choice, released, updated, developer_name, 
		available_locales) 
		VALUES 
		($1, $2, $3, $4, $5, $6, $7, $8, $9, $10, $11, $12, $13, $14, $15, 
		$16, $17, $18, $19, $20, $21, $22, $23, $24, $25, $26, $27, $28, $29) 
		RETURNING id`

	err := ds.DB.QueryRow(query,
		app.Name, app.Locale, app.Country, app.URL, app.FullURL, app.Description, app.Summary, app.Icon, app.Score,
		app.PriceText, app.IsFree, app.Installs, app.InstallsText, app.AppVersion, app.AndroidVersion, app.MinAndroidVersion,
		app.Size, app.ContentRating, app.PrivacyPolicyUrl, app.Category.ID, 0, app.Developer.ID,
		app.NumberVoters, app.NumberReviews, app.RecentChanges, app.EditorsChoice, app.Released, app.Updated, app.DeveloperName,
		// Ensure available_locales is a valid PostgreSQL array
		pq.Array(app.AvailableLocales),
	).Scan(&app.ID)

	if err != nil {
		return fmt.Errorf("error inserting new app data: %v", err)
	}

	log.Println("New app data inserted successfully")
	return nil
}

// GetApp retrieves an app by ID
func (ds *DBService) GetApp(id int64) (*interfaces.AppData, error) {
	query := `SELECT id, name, locale, country, url, full_url, description, summary, icon, score, price_text, 
		is_free, installs, installs_text, app_version, android_version, min_android_version, size, content_rating, 
		privacy_policy_url, category_id, histogram_rating_id, developer_id, number_voters, number_reviews, recent_changes, 
		editors_choice, released, updated, developer_name, available_locales FROM apps WHERE id = $1`

	app := &interfaces.AppData{}
	var availableLocales string

	err := ds.DB.QueryRow(query, id).Scan(
		&app.ID, &app.Name, &app.Locale, &app.Country, &app.URL, &app.FullURL, &app.Description, &app.Summary, &app.Icon,
		&app.Score, &app.PriceText, &app.IsFree, &app.Installs, &app.InstallsText, &app.AppVersion, &app.AndroidVersion,
		&app.MinAndroidVersion, &app.Size, &app.ContentRating, &app.PrivacyPolicyUrl, &app.Category.ID,
		0, &app.Developer.ID, &app.NumberVoters, &app.NumberReviews, &app.RecentChanges,
		&app.EditorsChoice, &app.Released, &app.Updated, &app.DeveloperName, &availableLocales,
	)

	if err != nil {
		if err == sql.ErrNoRows {
			return nil, fmt.Errorf("no app found with ID %d", id)
		}
		return nil, fmt.Errorf("error fetching app data: %v", err)
	}

	// Convert available_locales string to array
	app.AvailableLocales = []string{availableLocales}

	return app, nil
}

// UpdateApp updates the app data in the database
func (ds *DBService) UpdateApp(app *interfaces.AppData) error {
	query := `UPDATE apps SET name = $1, locale = $2, country = $3, url = $4, full_url = $5, description = $6, 
		summary = $7, icon = $8, score = $9, price_text = $10, is_free = $11, installs = $12, installs_text = $13, 
		app_version = $14, android_version = $15, min_android_version = $16, size = $17, content_rating = $18, 
		privacy_policy_url = $19, category_id = $20, histogram_rating_id = $21, developer_id = $22, 
		number_voters = $23, number_reviews = $24, recent_changes = $25, editors_choice = $26, 
		released = $27, updated = $28, developer_name = $29, available_locales = $30 WHERE id = $31`

	_, err := ds.DB.Exec(query,
		app.Name, app.Locale, app.Country, app.URL, app.FullURL, app.Description, app.Summary, app.Icon, app.Score,
		app.PriceText, app.IsFree, app.Installs, app.InstallsText, app.AppVersion, app.AndroidVersion, app.MinAndroidVersion,
		app.Size, app.ContentRating, app.PrivacyPolicyUrl, app.Category.ID, 0, app.Developer.ID,
		app.NumberVoters, app.NumberReviews, app.RecentChanges, app.EditorsChoice, app.Released, app.Updated, app.DeveloperName,
		app.AvailableLocales, app.ID,
	)

	if err != nil {
		return fmt.Errorf("error updating app data: %v", err)
	}

	log.Println("App data updated successfully")
	return nil
}

// DeleteApp deletes an app by ID
func (ds *DBService) DeleteApp(id int64) error {
	query := `DELETE FROM apps WHERE id = $1`

	_, err := ds.DB.Exec(query, id)
	if err != nil {
		return fmt.Errorf("error deleting app data: %v", err)
	}

	log.Println("App data deleted successfully")
	return nil
}
