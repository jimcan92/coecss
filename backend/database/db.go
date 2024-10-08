package database

import (
	"coecss/backend/models"
	"encoding/json"
	"fmt"

	"go.etcd.io/bbolt"
)

type BucketName string

const (
	UsersBucket       BucketName = "Users"
	SessionsBucket    BucketName = "Sessions"
	InstructorsBucket BucketName = "Instructors"
	SubjectsBucket    BucketName = "Subjects"
	SectionsBucket    BucketName = "Sections"
	RoomsBucket       BucketName = "Rooms"
)

var BucketNames = []BucketName{
	UsersBucket,
	SessionsBucket,
	InstructorsBucket,
	SubjectsBucket,
	SectionsBucket,
	RoomsBucket,
}

func GetItem[T models.Identifiable](db *bbolt.DB, bucketName BucketName, id string) (*T, error) {
	var item T

	err := db.View(func(tx *bbolt.Tx) error {
		bucket := tx.Bucket([]byte(bucketName))
		if bucket == nil {
			return fmt.Errorf("bucket not found: %s", bucketName)
		}

		// Get the item data from the bucket using its ID
		data := bucket.Get([]byte(id))
		if data == nil {
			return fmt.Errorf("item not found: %s", id)
		}

		// Unmarshal the JSON data into the item
		return json.Unmarshal(data, &item)
	})

	if err != nil {
		return nil, err
	}

	return &item, nil
}

func GetAllItems[T models.Identifiable](db *bbolt.DB, bucketName BucketName) ([]T, error) {
	var items []T

	err := db.View(func(tx *bbolt.Tx) error {
		bucket := tx.Bucket([]byte(bucketName))
		if bucket == nil {
			return fmt.Errorf("bucket not found: %s", bucketName)
		}

		// Iterate over all items in the bucket
		return bucket.ForEach(func(k, v []byte) error {
			var item T
			// Unmarshal each item into the slice
			if err := json.Unmarshal(v, &item); err != nil {
				return err // Return error if unmarshalling fails
			}
			items = append(items, item) // Add the item to the slice
			return nil
		})
	})

	if err != nil {
		return nil, err // Return any error that occurred during the view
	}

	return items, nil // Return the slice of items
}

func AddItem[T models.Identifiable](db *bbolt.DB, bucketName BucketName, item T) error {
	return db.Update(func(tx *bbolt.Tx) error {
		bucket := tx.Bucket([]byte(bucketName))
		if bucket == nil {
			return fmt.Errorf("bucket not found: %s", bucketName)
		}

		data, err := json.Marshal(item)
		if err != nil {
			return err
		}

		fmt.Println(item.GetID())

		return bucket.Put([]byte(item.GetID()), data)
	})
}

func UpdateItem[T models.Identifiable](db *bbolt.DB, bucketName BucketName, item T) error {
	id := item.GetID()

	return db.Update(func(tx *bbolt.Tx) error {
		bucket := tx.Bucket([]byte(bucketName))
		if bucket == nil {
			return fmt.Errorf("bucket not found: %s", bucketName)
		}

		// Marshal the item to JSON
		data, err := json.Marshal(item)
		if err != nil {
			return err // Return error if marshaling fails
		}

		// Update the item in the bucket
		return bucket.Put([]byte(id), data)
	})
}

func DeleteItem[T models.Identifiable](db *bbolt.DB, bucketName BucketName, id string) error {
	return db.Update(func(tx *bbolt.Tx) error {
		bucket := tx.Bucket([]byte(bucketName))
		if bucket == nil {
			return fmt.Errorf("bucket not found: %s", bucketName)
		}

		// Delete the item from the bucket using its ID
		if err := bucket.Delete([]byte(id)); err != nil {
			return err // Return error if deletion fails
		}

		return nil // Successful deletion
	})
}

func GetLastInserted[T any](db *bbolt.DB, bucketName BucketName) ([]byte, *T, error) {
	var lastKey []byte
	var lastValue []byte
	var result T

	// Read-only transaction to retrieve data
	err := db.View(func(tx *bbolt.Tx) error {
		bucket := tx.Bucket([]byte(bucketName))
		if bucket == nil {
			return fmt.Errorf("bucket %s not found", bucketName)
		}

		// Create a cursor to traverse the bucket
		cursor := bucket.Cursor()

		// Move to the last key-value pair
		lastKey, lastValue = cursor.Last()

		// If there are no entries
		if lastKey == nil {
			return fmt.Errorf("no data found in bucket %s", bucketName)
		}

		// Unmarshal the last value into the expected generic type
		if err := json.Unmarshal(lastValue, &result); err != nil {
			return fmt.Errorf("failed to unmarshal data: %v", err)
		}

		return nil
	})

	if err != nil {
		return nil, nil, err
	}

	return lastKey, &result, nil
}

func SetupDone(db *bbolt.DB) bool {

	err := db.View(func(tx *bbolt.Tx) error {
		bucket := tx.Bucket([]byte(UsersBucket))
		if bucket == nil {
			return fmt.Errorf("bucket-not-found")
		}

		cursor := bucket.Cursor()
		k, _ := cursor.First()
		if k == nil {
			return fmt.Errorf("users-bucket-empty")
		}
		return nil
	})

	fmt.Println(err == nil)

	return err == nil
}
