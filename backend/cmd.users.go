package backend

import (
	"coecss/backend/database"
	"coecss/backend/models"
	"fmt"
	"time"

	"github.com/google/uuid"
)

// RegisterUserCommand registers a new user
func (a *App) RegisterUserCommand(username, password string) error {
	user := &models.User{Username: username, Password: password, Role: "Student"}
	return database.RegisterUser(a.DB, user)
}

// LoginUserCommand logs in a user
func (a *App) LoginUserCommand(username, password string) (*models.User, error) {
	user, err := database.LoginUser(a.DB, username, password)
	fmt.Println(err)
	if err != nil {
		return nil, err
	}

	session := &models.Session{ID: uuid.NewString(), UserID: user.Username, Active: true, Created: time.Now()}

	sessionAddErr := database.AddItem(a.DB, database.SessionsBucket, session)
	if sessionAddErr != nil {
		fmt.Println("Session add error", sessionAddErr)
	}

	a.SetSession(session)

	return user, nil
}

func (a *App) Logout() {
	a.Session.Active = false
	a.Session.Terminated = time.Now()
}

func (a *App) GetUsers() ([]models.User, error) {
	return database.GetAllItems[models.User](a.DB, database.UsersBucket)
}
