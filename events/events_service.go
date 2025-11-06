package events

import "github.com/RonanConway/eventsRestAPI/models"

type EventsService interface {
	Save(*models.Event) error
	Update(*models.Event) error
	Delete(*models.Event) error
	GetAllEvents() ([]models.Event, error)
}

type eventServiceImpl struct{}

func (s eventServiceImpl) Save(event *models.Event) error {
	return event.Save()
}

func (s eventServiceImpl) Update(event *models.Event) error {
	return event.Update()
}

func (s eventServiceImpl) Delete(event *models.Event) error {
	return event.Delete()
}

func (s eventServiceImpl) GetAllEvents() ([]models.Event, error) {
	return models.GetAllEvents()
}
