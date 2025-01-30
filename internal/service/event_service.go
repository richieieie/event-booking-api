package service

import (
	dto "github.com/richieieie/event-booking/internal/DTO"
	"github.com/richieieie/event-booking/internal/model"
	"github.com/richieieie/event-booking/internal/repository"
)

// Use can map Event model to EventResponseDTO and vice versa, but I'm too lazy right now
// type EventResponseDTO struct {

// }

type IEventService interface {
	GetAll() ([]model.Event, error)
	GetById(id int) (model.Event, error)
	CreateOne(eventDto dto.CreateEventDTO) (int, error)
	UpdateOne(id int, eventDto dto.EventUpdateDTO) error
	DeleteOne(id int) error
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
	})
}

func (e *eventService) UpdateOne(id int, eventDto dto.EventUpdateDTO) error {
	// Add business logic here
	// For example: maybe we want to validate the input here
	// I do not have any ideas right now, so I just put the response from repository call here
	return e.iEventRepository.UpdateOne(id, model.Event{
		Name:        eventDto.Name,
		Description: eventDto.Description,
		Location:    eventDto.Location,
		DateTime:    eventDto.DateTime,
	})
}

func (e *eventService) DeleteOne(id int) error {
	// Add business logic here
	// For example: maybe we want to validate the input here
	// I do not have any ideas right now, so I just put the response from repository call here
	return e.iEventRepository.DeleteOne(id)
}
