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
	UpdateOne(id int, event model.Event, userId int64) error
	DeleteOne(id int, userId int64) error
	RegisterEvent(eventId int, userId int64) (*model.Registration, error)
	UnregisterEvent(eventId int, userId int64) error
}

func (r eventRepository) UnregisterEvent(eventId int, userId int64) error {
	err := r.db.Where("event_id = ? and user_id = ?", eventId, userId).Delete(&model.Registration{}).Error

	return err
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

func (e *eventRepository) UpdateOne(id int, eventToUpdate model.Event, userId int64) error {
	var event model.Event
	err := e.db.Where("id = ? and user_id = ?", id, userId).First(&event).Error
	if err != nil {
		return fmt.Errorf("could not find your event with id %d", id)
	}

	event.Name = eventToUpdate.Name
	event.Description = eventToUpdate.Description
	event.Location = eventToUpdate.Location
	event.DateTime = eventToUpdate.DateTime
	err = e.db.Save(&event).Error
	if err != nil {
		return fmt.Errorf("could not update your event with id %d", id)
	}

	return nil
}

func (e *eventRepository) DeleteOne(id int, userId int64) error {
	var event model.Event
	err := e.db.Where("id = ? and user_id = ?", id, userId).First(&event).Error
	if err != nil {
		return fmt.Errorf("could not find your event with id %d", id)
	}

	err = e.db.Delete(&event).Error
	if err != nil {
		return fmt.Errorf("could not delete your event with id %d", id)
	}

	return nil
}

func (r eventRepository) RegisterEvent(eventId int, userId int64) (*model.Registration, error) {
	registration := model.Registration{EventId: eventId, UserId: userId}
	err := r.db.Create(&registration).Error
	if err != nil {
		return nil, fmt.Errorf("could not register event with id %d", eventId)
	}
	return &registration, nil
}
