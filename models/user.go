package models

import "time"

type User struct {
	ID        string `json:"id"`
	Name      string `json:"name"`
	Role      string
	UpdatedAt *time.Time
	CreatedAt *time.Time
	CreatedBy string
	UpdatedBy string
}
