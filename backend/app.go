package backend

import (
	"coecss/backend/database"
	"coecss/backend/models"
	"context"
	"log"

	"github.com/wailsapp/wails/v2/pkg/runtime"
	"go.etcd.io/bbolt"
)

// App struct
type App struct {
	ctx     context.Context
	DB      *bbolt.DB
	Session *models.Session
}

// NewApp creates a new App application struct
func NewApp() *App {

	return &App{}
}

func (a *App) InitializeDB(dbPath string) error {
	db, err := bbolt.Open(dbPath, 0600, nil)
	if err != nil {
		return err
	}
	a.DB = db

	return a.DB.Update(func(tx *bbolt.Tx) error {
		for _, name := range database.BucketNames {

			_, err := tx.CreateBucketIfNotExists([]byte(name))
			if err != nil {
				log.Printf("Error creating bucket %s: %v", name, err)
				return err
			}
		}

		return nil
	})

}

func (a *App) SetSession(session *models.Session) {
	a.Session = session
	runtime.EventsEmit(a.ctx, "session-changed", a.Session)
}

func (a *App) CloseDB() {
	if err := a.DB.Close(); err != nil {
		log.Printf("Error closing database: %v", err)
	}
}

func (a *App) Startup(ctx context.Context) {
	a.ctx = ctx

	if !database.SetupDone(a.DB) {
		runtime.EventsEmit(a.ctx, "needs-setup")
	}
}

func (a *App) CheckIfNeedsSetup() bool {
	return !database.SetupDone(a.DB)
}
