package models

import (
	"github.com/google/uuid"
	"time"
)

/**
* Shared entities that are imported by more than one package.
**/
type User struct {
	BaseDBDateModel
	Email    string `db:"email" json:"email"`
	Name     string `db:"name" json:"name"`
	Password string `db:"password" json:"password,omitempty"`
}

type Booking struct {
	BaseDBUserDateModel
	// FK to User, One to Many Relation
	UserID    uuid.UUID `db:"user_id" json:"userId"`
	Status    string    `db:"status" json:"status"`
	StartDate time.Time `db:"start_date" json:"startDate"`
	EndDate   time.Time `db:"end_date" json:"endDate"`
}

/**
* Base models for default table columns.
**/

type BaseDBUserDateModel struct {
	BaseDBModel
	BaseDBUserModel
	BaseDBTimeModel
}

type BaseDBUserModel struct {
	UpdatedUser uuid.UUID `db:"updated_user" json:"updatedUser"`
	CreatedUser uuid.UUID `db:"created_user" json:"createdUser"`
}

type BaseDBDateModel struct {
	BaseDBModel
	BaseDBTimeModel
}

type BaseDBTimeModel struct {
	CreatedAt time.Time `db:"created_at" json:"created_at"`
	UpdatedAt time.Time `db:"updated_at" json:"updated_at"`
}

type BaseDBModel struct {
	ID uuid.UUID `db:"id" json:"id"`
}
