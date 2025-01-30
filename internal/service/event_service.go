package service

import (
	"fmt"

	dto "github.com/richieieie/event-booking/internal/DTO"
	"github.com/richieieie/event-booking/internal/model"
	"github.com/richieieie/event-booking/internal/repository"
)

// In UnregisterEvent method, We can use RegistrationId to delete, but I will use both eventId and userId to delete
type IEventService interface {
	GetAll() ([]model.Event, error)
	GetById(id int) (model.Event, error)
	CreateOne(eventDto dto.CreateEventDTO) (int, error)
	UpdateOne(id int, eventDto dto.EventUpdateDTO, userId int64) error
	DeleteOne(id int, userId int64) error
	RegisterEvent(eventId int, userId int64) (int64, error)
	UnregisterEvent(eventId int, userId int64) error
}

type eventService struct {
	iEventRepository repository.IEventRepository
}

func NewEventService(iEventRepository repository.IEventRepository) *eventService {
	// Add business logic here
	// For example: maybe we want to map to DTO here, or filter some events by a condition by default
	// I do not have any ideas right now, so I just put the response from repository call here
	return &eventService{iEventRepository: iEventRepository}
}

func (e *eventService) GetAll() ([]model.Event, error) {
	// Add business logic here
	// For example: maybe we want to map to DTO here, or filter some events by a condition by default
	// I do not have any ideas right now, so I just put the response from repository call here
	return e.iEventRepository.GetAll()
}

func (e *eventService) GetById(id int) (model.Event, error) {
	// Add business logic here
	// For example: maybe we want to validate the input here
	// I do not have any ideas right now, so I just put the response from repository call here
	return e.iEventRepository.GetById(id)
}

func (e *eventService) CreateOne(eventDto dto.CreateEventDTO) (int, error) {
	// Add business logic here
	// For example: maybe we want to validate the input here
	// I do not have any ideas right now, so I just put the response from repository call here
	return e.iEventRepository.CreateOne(model.Event{
		Id:          0,
		Name:        eventDto.Name,
		Description: eventDto.Description,
		Location:    eventDto.Location,
		DateTime:    eventDto.DateTime,
		UserId:      eventDto.UserId,
	})
}

func (e *eventService) UpdateOne(id int, eventDto dto.EventUpdateDTO, userId int64) error {
	// Add business logic here
	// For example: maybe we want to validate the input here
	// I do not have any ideas right now, so I just put the response from repository call here
	return e.iEventRepository.UpdateOne(id, model.Event{
		Name:        eventDto.Name,
		Description: eventDto.Description,
		Location:    eventDto.Location,
		DateTime:    eventDto.DateTime,
	}, userId)
}

func (e *eventService) DeleteOne(id int, userId int64) error {
	// Add business logic here
	// For example: maybe we want to validate the input here
	// I do not have any ideas right now, so I just put the response from repository call here
	return e.iEventRepository.DeleteOne(id, userId)
}

func (e *eventService) RegisterEvent(eventId int, userId int64) (int64, error) {
	_, err := e.iEventRepository.GetById(eventId)
	if err != nil {
		return 0, fmt.Errorf("could not find event with id %d", eventId)
	}

	registration, err := e.iEventRepository.RegisterEvent(eventId, userId)
	if err != nil {
		return 0, fmt.Errorf("could not register event with id %d", eventId)
	}

	return registration.Id, nil
}

func (e *eventService) UnregisterEvent(eventId int, userId int64) error {
	_, err := e.iEventRepository.GetById(eventId)
	if err != nil {
		return fmt.Errorf("could not find event with id %d", eventId)
	}

	err = e.iEventRepository.UnregisterEvent(eventId, userId)
	if err != nil {
		return fmt.Errorf("could not unregister event with id %d", eventId)
	}

	return nil
}
