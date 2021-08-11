package topup

import (
	"ticketing/business/topup"
	"time"
)

type Topup struct {
	ID        int       `json:"id"`
	User_ID   int       `json:"user_id"`
	Name      string    `json:"name"`
	Balance   float64   `json:"balance"`
	CreatedAt time.Time `json:"created_at"`
	UpdatedAt time.Time `json:"updated_at"`
}

func (rec *Topup) toDomain() topup.Domain {
	return topup.Domain{
		ID:        rec.ID,
		User_ID:   rec.User_ID,
		Name:      rec.Name,
		Balance:   rec.Balance,
		CreatedAt: rec.CreatedAt,
		UpdatedAt: rec.UpdatedAt,
	}
}

func fromDomain(topupDomain topup.Domain) *Topup {
	return &Topup{
		ID:        topupDomain.ID,
		User_ID:   topupDomain.User_ID,
		Name:      topupDomain.Name,
		Balance:   topupDomain.Balance,
		CreatedAt: topupDomain.CreatedAt,
		UpdatedAt: topupDomain.UpdatedAt,
	}
}
