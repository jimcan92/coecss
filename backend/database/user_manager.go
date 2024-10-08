package database

import (
	"coecss/backend/models"
	"encoding/json"
	"fmt"

	"go.etcd.io/bbolt"
)

func RegisterUser(db *bbolt.DB, user *models.User) error {
	return db.Update(func(tx *bbolt.Tx) error {
		bucket := tx.Bucket([]byte(UsersBucket))
		if bucket == nil {
			return fmt.Errorf("bucket-not-found")
		}

		existingUserData := bucket.Get([]byte(user.Username))
		if existingUserData != nil {
			return fmt.Errorf("user-exists")
		}

		if err := user.HashPassword(); err != nil {
			return fmt.Errorf("error-hashing-password")
		}

		data, err := json.Marshal(user)
		if err != nil {
			return fmt.Errorf("serializing-error")
		}

		return bucket.Put([]byte(user.Username), data)
	})
}

func LoginUser(db *bbolt.DB, username, password string) (*models.User, error) {
	var user models.User

	err := db.View(func(tx *bbolt.Tx) error {
		bucket := tx.Bucket([]byte(UsersBucket))
		if bucket == nil {
			return fmt.Errorf("bucket-not-found")
		}

		data := bucket.Get([]byte(username))
		if data == nil {
			return fmt.Errorf("no-user-exists")
		}

		err := json.Unmarshal(data, &user)
		if err != nil {
			return fmt.Errorf("deserializing-error")
		}

		return nil
	})

	if err != nil {
		return nil, err
	}

	if !user.CheckPassword(password) {
		return nil, fmt.Errorf("wrong-password")
	}

	return &user, nil
}
