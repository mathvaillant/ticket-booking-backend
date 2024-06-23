package models

import (
	"context"
	"time"

	"gorm.io/gorm"
)

type Event struct {
	ID                    uint      `json:"id" gorm:"primarykey"`
	Name                  string    `json:"name"`
	Location              string    `json:"location"`
	TotalTicketsPurchased int64     `json:"totalTicketsPurchased" gorm:"-"`
	TotalTicketsEntered   int64     `json:"totalTicketsEntered" gorm:"-"`
	Date                  time.Time `json:"date"`
	CreatedAt             time.Time `json:"createdAt"`
	UpdatedAt             time.Time `json:"updatedAt"`
}

type EventRepository interface {
	GetMany(ctx context.Context) ([]*Event, error)
	GetOne(ctx context.Context, eventId uint) (*Event, error)
	CreateOne(ctx context.Context, event *Event) (*Event, error)
	UpdateOne(ctx context.Context, eventId uint, updateData map[string]interface{}) (*Event, error)
	DeleteOne(ctx context.Context, eventId uint) error
}

func (e *Event) AfterFind(db *gorm.DB) (err error) {
	baseQuery := db.Model(&Ticket{}).Where(&Ticket{EventID: e.ID})

	if res := baseQuery.Count(&e.TotalTicketsPurchased); res.Error != nil {
		return res.Error
	}
	if res := baseQuery.Where("entered = ?", true).Count(&e.TotalTicketsEntered); res.Error != nil {
		return res.Error
	}
	return nil
}
