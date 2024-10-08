package backend

import (
	"coecss/backend/database"
	"coecss/backend/models"
	"fmt"
)

func (a *App) GetInstructor(id string) (*models.Instructor, error) {
	// Call the generic GetItem function
	instructor, err := database.GetItem[models.Instructor](a.DB, database.InstructorsBucket, id)
	if err != nil {
		return nil, err // Return the error if there's an issue
	}
	return instructor, nil // Return the retrieved instructor
}

func (a *App) GetAllInstructors() ([]models.Instructor, error) {
	return database.GetAllItems[models.Instructor](a.DB, database.InstructorsBucket)
}

func (a *App) AddInstructor(name string) (*models.Instructor, error) {
	instructor := models.NewInstructor(name)

	err := database.AddItem(a.DB, database.InstructorsBucket, instructor)

	return instructor, err
}

func (a *App) UpdateInstructor(instructor *models.Instructor) error {
	// Check if the instructor ID is not empty
	if instructor.ID == "" {
		return fmt.Errorf("instructor ID cannot be empty")
	}

	return database.UpdateItem(a.DB, database.InstructorsBucket, *instructor)
}

func (a *App) DeleteInstructor(id string) error {
	if id == "" {
		return fmt.Errorf("instructor ID cannot be empty")
	}

	return database.DeleteItem[models.Instructor](a.DB, database.InstructorsBucket, id)
}
