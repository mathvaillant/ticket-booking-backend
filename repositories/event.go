package repositories

import (
	"context"
	"strconv"

	"github.com/mathvaillant/ticket-booking-project-v0/models"
	"gorm.io/gorm"
)

type EventRepository struct {
	db *gorm.DB
}

func (r *EventRepository) GetMany(ctx context.Context) ([]*models.Event, error) {
	events := []*models.Event{}

	res := r.db.Model(&models.Event{}).Find(&events)

	if res.Error != nil {
		return nil, res.Error
	}

	return events, nil
}

func (r *EventRepository) GetOne(ctx context.Context, eventId string) (*models.Event, error) {
	id, _ := strconv.Atoi(eventId)
	event := &models.Event{}

	res := r.db.Model(event).Where("id = ?", uint(id)).First(event)

	if res.Error != nil {
		return nil, res.Error
	}

	return event, nil
}

func (r *EventRepository) CreateOne(ctx context.Context, event *models.Event) (*models.Event, error) {
	res := r.db.Model(event).Create(event)

	if res.Error != nil {
		return nil, res.Error
	}

	return event, nil
}

func (r *EventRepository) UpdateOne(ctx context.Context, eventId string, updateData map[string]interface{}) (*models.Event, error) {
	id, _ := strconv.Atoi(eventId)
	event := &models.Event{}

	updateRes := r.db.Model(event).Where("id = ?", uint(id)).Updates(updateData)

	if updateRes.Error != nil {
		return nil, updateRes.Error
	}

	getRes := r.db.Model(event).Where("id = ?", uint(id)).First(event)

	if getRes.Error != nil {
		return nil, getRes.Error
	}

	return event, nil
}

func (r *EventRepository) DeleteOne(ctx context.Context, eventId string) error {
	res := r.db.Delete(&models.Event{}, eventId)
	return res.Error
}

func NewEventRepository(db *gorm.DB) models.EventRepository {
	return &EventRepository{
		db: db,
	}
}
