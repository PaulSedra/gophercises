package db

import (
	"fmt"
	"github.com/boltdb/bolt"
)

// InitializeDB initializes a BoltDB database and creates the necessary bucket
func InitializeDB(dbPath string) (*bolt.DB, error) {

	db, err := bolt.Open(dbPath, 0600, nil)
	if err != nil {
		return nil, fmt.Errorf("error opening database: %v", err)
	}

	err = db.Update(func(tx *bolt.Tx) error {
		_, err := tx.CreateBucketIfNotExists([]byte("Redirects"))
		return err
	})
	if err != nil {
		CloseDB(db)
		return nil, fmt.Errorf("error creating bucket: %v", err)
	}

	return db, nil
}

// CloseDB closes the BoltDB database
func CloseDB(db *bolt.DB) {

	err := db.Close()
	if err != nil {
		return
	}
}

// AddRedirect adds a new redirect entry to the BoltDB database
func AddRedirect(db *bolt.DB, path, url string) error {

	err := db.Update(func(tx *bolt.Tx) error {
		bucket := tx.Bucket([]byte("Redirects"))
		if bucket == nil {
			return fmt.Errorf("bucket not found")
		}

		err := bucket.Put([]byte(path), []byte(url))
		if err != nil {
			return fmt.Errorf("error adding redirect to database: %v", err)
		}

		return nil
	})

	return err
}

// GetRedirect retrieves the URL for a given path from the BoltDB database
func GetRedirect(db *bolt.DB, path string) (string, error) {

	var url string

	err := db.View(func(tx *bolt.Tx) error {
		bucket := tx.Bucket([]byte("Redirects"))
		if bucket == nil {
			return fmt.Errorf("bucket not found")
		}

		urlBytes := bucket.Get([]byte(path))
		if urlBytes == nil {
			return fmt.Errorf("path not found in the database")
		}

		url = string(urlBytes)
		return nil
	})

	return url, err
}

// UpdateRedirect updates the URL for a given path in the BoltDB database
func UpdateRedirect(db *bolt.DB, path, newURL string) error {

	err := db.Update(func(tx *bolt.Tx) error {
		bucket := tx.Bucket([]byte("Redirects"))
		if bucket == nil {
			return fmt.Errorf("bucket not found")
		}

		err := bucket.Put([]byte(path), []byte(newURL))
		if err != nil {
			return fmt.Errorf("error updating redirect in the database: %v", err)
		}

		return nil
	})

	return err
}

// DeleteRedirect deletes a redirect entry for a given path from the BoltDB database
func DeleteRedirect(db *bolt.DB, path string) error {

	err := db.Update(func(tx *bolt.Tx) error {
		bucket := tx.Bucket([]byte("Redirects"))
		if bucket == nil {
			return fmt.Errorf("bucket not found")
		}

		err := bucket.Delete([]byte(path))
		if err != nil {
			return fmt.Errorf("error deleting redirect from the database: %v", err)
		}

		return nil
	})

	return err
}
