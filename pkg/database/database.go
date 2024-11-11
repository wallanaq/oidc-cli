package database

import (
	"fmt"
	"path/filepath"
	"time"

	"github.com/wallanaq/oidc-cli/pkg/paths"
	"go.etcd.io/bbolt"
)

var db *bbolt.DB

// Open initializes the connection to the local database.
// It sets up necessary resources and prepares the database for use.
//
// Returns:
//
//	error - if the database fails to open
func Open() error {

	configDir, err := paths.GetConfigDir()

	if err != nil {
		return fmt.Errorf("failed to get configuration directory: %w", err)
	}

	dbPath := filepath.Join(configDir, "db")

	if db, err = bbolt.Open(dbPath, 0600, &bbolt.Options{Timeout: 1 * time.Second}); err != nil {
		return fmt.Errorf("")
	}

	return nil
}

// Close terminates the connection to the local database and releases any
// allocated resources.
//
// Returns:
//
//	error - if there is an issue closing the database
func Close() error {

	if db == nil {
		return fmt.Errorf("database is not open")
	}

	return db.Close()

}

// Set stores a value in the specified bucket within the database using the
// given key. It creates the bucket if it does not exist.
//
// Example:
//
//	err := Set("user_data", "username", "exampleUser")
//	if err != nil {
//	    log.Fatal(err)
//	}
//
// Parameters:
//
//	bucketName - the name of the bucket where the key-value pair is stored
//	key - the key associated with the value
//	value - the value to be stored
//
// Returns:
//
//	error - if unable to set the value
func Set(bucketName string, key string, value string) error {

	return db.Update(func(tx *bbolt.Tx) error {

		bucket, err := tx.CreateBucketIfNotExists([]byte(bucketName))

		if err != nil {
			return fmt.Errorf("failed to create or retrieve bucket %s: %w", bucketName, err)
		}

		if err := bucket.Put([]byte(key), []byte(value)); err != nil {
			return fmt.Errorf("failed to set key %s in bucket %s: %w", key, bucketName, err)
		}

		return nil

	})

}

// Get retrieves the value associated with the specified key from the given
// bucket in the database.
//
// Example:
//
//	value, err := Get("user_data", "username")
//	if err != nil {
//	    log.Fatal(err)
//	}
//
// Parameters:
//
//	bucketName - the name of the bucket where the key is stored
//	key - the key associated with the value
//
// Returns:
//
//	string - the retrieved value
//	error - if unable to retrieve the value
func Get(bucketName string, key string) (value string, err error) {

	err = db.View(func(tx *bbolt.Tx) error {

		bucket := tx.Bucket([]byte(bucketName))

		if bucket == nil {
			return fmt.Errorf("bucket %s not found", bucketName)
		}

		bValue := bucket.Get([]byte(key))

		if bValue == nil {
			return fmt.Errorf("key %s not found in bucket %s", key, bucketName)
		}

		value = string(bValue)

		return nil

	})

	if err != nil {
		return "", fmt.Errorf("failed to get key %s from bucket %s: %w", key, bucketName, err)
	}

	return
}

// Delete removes the key-value pair associated with the specified key from
// the given bucket in the database.
//
// Example:
//
//	err := Delete("user_data", "username")
//	if err != nil {
//	    log.Fatal(err)
//	}
//
// Parameters:
//
//	bucketName - the name of the bucket containing the key-value pair
//	key - the key to delete
//
// Returns:
//
//	error - if unable to delete the key-value pair
func Delete(bucketName string, key string) error {

	return db.Update(func(tx *bbolt.Tx) error {

		bucket := tx.Bucket([]byte(bucketName))

		if bucket == nil {
			return fmt.Errorf("bucket %s not found", bucketName)
		}

		if err := bucket.Delete([]byte(key)); err != nil {
			return fmt.Errorf("failed to delete key %s from bucket %s: %w", key, bucketName, err)
		}

		return nil

	})

}
