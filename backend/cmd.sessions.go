package backend

import (
	"coecss/backend/database"
	"coecss/backend/models"
	"fmt"
	"log"
	"time"

	"go.etcd.io/bbolt"
)

func updateSession(db *bbolt.DB, s *models.Session) {
	err := database.UpdateItem(db, database.SessionsBucket, s)
	if err != nil {
		log.Fatal(err)
	}
}

func (a *App) SessionActivate() {
	a.Session.Active = true
	updateSession(a.DB, a.Session)
}

func (a *App) SessionDeactivate() {
	a.Session.Active = false
	updateSession(a.DB, a.Session)
}

func (a *App) SessionState() bool {
	return a.Session.Active
}

func (a *App) SessionTerminate() {
	fmt.Println(a.Session)
	if a.Session != nil {
		a.Session.Active = false
		a.Session.Terminated = time.Now()
		updateSession(a.DB, a.Session)
		a.SetSession(nil)
	}
}

func (a *App) GetSession(id string) (*models.Session, error) {
	return database.GetItem[models.Session](a.DB, database.SessionsBucket, id)
}

func (a *App) GetSessions() ([]models.Session, error) {
	return database.GetAllItems[models.Session](a.DB, database.SessionsBucket)
}
