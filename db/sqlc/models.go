// Code generated by sqlc. DO NOT EDIT.

package db

import (
	"database/sql"
	"time"
)

type Category struct {
	ID     int32          `json:"id"`
	Desc   sql.NullString `json:"desc"`
	Status sql.NullInt32  `json:"status"`
}

type Event struct {
	ID          int32          `json:"id"`
	Title       string         `json:"title"`
	Description string         `json:"description"`
	BannerImage string         `json:"banner_image"`
	StartDate   time.Time      `json:"start_date"`
	EndDate     time.Time      `json:"end_date"`
	Venue       int32          `json:"venue"`
	Type        int32          `json:"type"`
	UserID      int32          `json:"user_id"`
	Category    int32          `json:"category"`
	Subcategory int32          `json:"subcategory"`
	TicketID    sql.NullInt32  `json:"ticket_id"`
	Recurring   sql.NullBool   `json:"recurring"`
	Status      sql.NullString `json:"status"`
	Image1      sql.NullString `json:"image1"`
	Image2      sql.NullString `json:"image2"`
	Image3      sql.NullString `json:"image3"`
	Video1      sql.NullString `json:"video1"`
	Video2      sql.NullString `json:"video2"`
	CreatedAt   sql.NullTime   `json:"created_at"`
}

type EventType struct {
	ID     int32          `json:"id"`
	Desc   sql.NullString `json:"desc"`
	Status sql.NullInt32  `json:"status"`
}

type EventsSponsor struct {
	ID        int32     `json:"id"`
	EventID   int32     `json:"event_id"`
	SponsorID int32     `json:"sponsor_id"`
	CreatedAt time.Time `json:"created_at"`
}

type Sponsor struct {
	ID        int32     `json:"id"`
	UserID    int32     `json:"user_id"`
	CreatedAt time.Time `json:"created_at"`
}

type Subcategory struct {
	ID     int32          `json:"id"`
	Desc   sql.NullString `json:"desc"`
	Status sql.NullInt32  `json:"status"`
}

type Ticket struct {
	ID       int32   `json:"id"`
	Name     string  `json:"name"`
	EventID  int32   `json:"event_id"`
	Price    float64 `json:"price"`
	Quantity int32   `json:"quantity"`
	Status   int32   `json:"status"`
}

type TicketStatus struct {
	ID     int32          `json:"id"`
	Desc   sql.NullString `json:"desc"`
	Status int32          `json:"status"`
}

type User struct {
	ID                int32        `json:"id"`
	Title             string       `json:"title"`
	FirstName         string       `json:"first_name"`
	LastName          string       `json:"last_name"`
	Email             string       `json:"email"`
	Username          string       `json:"username"`
	Password          string       `json:"password"`
	PasswordChangedAt sql.NullTime `json:"password_changed_at"`
	Usertype          int32        `json:"usertype"`
	DateOfBirth       time.Time    `json:"date_of_birth"`
	CreatedAt         time.Time    `json:"created_at"`
}

type UserTicket struct {
	ID            int32           `json:"id"`
	UserID        int32           `json:"user_id"`
	TicketID      int32           `json:"ticket_id"`
	Quantity      sql.NullInt32   `json:"quantity"`
	TotalCost     sql.NullFloat64 `json:"total_cost"`
	Paid          sql.NullBool    `json:"paid"`
	PaymentRef    sql.NullString  `json:"payment_ref"`
	PaymentMethod sql.NullString  `json:"payment_method"`
	CreatedAt     time.Time       `json:"created_at"`
}

type Usertype struct {
	ID     int32          `json:"id"`
	Desc   sql.NullString `json:"desc"`
	Status sql.NullInt32  `json:"status"`
}

type Venue struct {
	ID          int32  `json:"id"`
	Name        string `json:"name"`
	Address     string `json:"address"`
	PostalCode  string `json:"postal_code"`
	City        string `json:"city"`
	Province    string `json:"province"`
	CountryCode string `json:"country_code"`
}
