package repositories

import (
	"context"
	"fmt"
	"strconv"

	"github.com/mathvaillant/ticket-booking-project-v0/models"
	"gorm.io/gorm"
)

type TicketRepository struct {
	db *gorm.DB
}

func (r *TicketRepository) GetMany(ctx context.Context) ([]*models.Ticket, error) {
	tickets := []*models.Ticket{}

	res := r.db.Model(&models.Ticket{}).Preload("Event").Find(&tickets)

	if res.Error != nil {
		return nil, res.Error
	}

	return tickets, nil
}

func (r *TicketRepository) GetOne(ctx context.Context, ticketId string) (*models.Ticket, error) {
	id, _ := strconv.Atoi(ticketId)
	ticket := &models.Ticket{}

	res := r.db.Model(ticket).Where("id = ?", uint(id)).Preload("Event").First(ticket)

	if res.Error != nil {
		return nil, res.Error
	}

	return ticket, nil
}

func (r *TicketRepository) CreateOne(ctx context.Context, ticket *models.Ticket) (*models.Ticket, error) {
	res := r.db.Model(ticket).Create(ticket)

	if res.Error != nil {
		return nil, res.Error
	}

	return r.GetOne(ctx, fmt.Sprint(ticket.ID))
}

func (r *TicketRepository) UpdateOne(ctx context.Context, ticketId string, updateData map[string]interface{}) (*models.Ticket, error) {
	id, _ := strconv.Atoi(ticketId)
	ticket := &models.Ticket{}

	updateRes := r.db.Model(ticket).Where("id = ?", uint(id)).Updates(updateData)

	if updateRes.Error != nil {
		return nil, updateRes.Error
	}

	return r.GetOne(ctx, ticketId)
}

func NewTicketRepository(db *gorm.DB) models.TicketRepository {
	return &TicketRepository{
		db: db,
	}
}
