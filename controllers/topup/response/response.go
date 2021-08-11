package response

import (
	"ticketing/business/topup"
	"time"
)

type TopUp struct {
	ID        int       `json:"id"`
	User_ID   int       `json:"user_id"`
	Name      string    `json:"name"`
	Balance   float64   `json:"balance"`
	CreatedAt time.Time `json:"created_at"`
	UpdatedAt time.Time `json:"updated_at"`
}

func FromDomain(topupDomain topup.Domain) TopUp {
	return TopUp{
		ID:        topupDomain.ID,
		Name:      topupDomain.Name,
		User_ID:   topupDomain.User_ID,
		CreatedAt: topupDomain.CreatedAt,
		UpdatedAt: topupDomain.UpdatedAt,
	}
}
