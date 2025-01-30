package repository

import (
	"fmt"

	"github.com/richieieie/event-booking/internal/model"
	"gorm.io/gorm"
)

// With bigger projects, we can create a generic repository, they will perform basic function such as get all, get by id, create, update, delete, etc.

type IEventRepository interface {
	GetAll() ([]model.Event, error)
	GetById(id int) (model.Event, error)
	CreateOne(event model.Event) (int, error)
	UpdateOne(id int, event model.Event) error
	DeleteOne(id int) error
}

type eventRepository struct {
	db *gorm.DB
}

func NewEventRepository(db *gorm.DB) *eventRepository {
	return &eventRepository{db: db}
}

func (e *eventRepository) GetAll() ([]model.Event, error) {
	var events []model.Event
	err := e.db.Find(&events).Error
	if err != nil {
		return nil, fmt.Errorf("could not get events")
	}

	return events, nil
}

func (e *eventRepository) GetById(id int) (model.Event, error) {
	var event model.Event
	err := e.db.First(&event, id).Error
	if err != nil {
		return model.Event{}, fmt.Errorf("could not find event with id %d", id)
	}

	return event, nil
}

func (e *eventRepository) CreateOne(event model.Event) (int, error) {
	err := e.db.Create(&event).Error
	if err != nil {
		return 0, fmt.Errorf("could not create event")
	}

	return event.Id, nil
}

func (e *eventRepository) UpdateOne(id int, eventToUpdate model.Event) error {
	var event model.Event
	err := e.db.First(&event, id).Error
	if err != nil {
		return fmt.Errorf("could not find event with id %d", id)
	}

	event.Name = eventToUpdate.Name
	event.Description = eventToUpdate.Description
	event.Location = eventToUpdate.Location
	event.DateTime = eventToUpdate.DateTime
	err = e.db.Save(&event).Error
	if err != nil {
		return fmt.Errorf("could not update event with id %d", id)
	}

	return nil
}

func (e *eventRepository) DeleteOne(id int) error {
	var event model.Event
	err := e.db.First(&event, id).Error
	if err != nil {
		return fmt.Errorf("could not find event with id %d", id)
	}

	err = e.db.Delete(&event).Error
	if err != nil {
		return fmt.Errorf("could not delete event with id %d", id)
	}

	return nil
}
