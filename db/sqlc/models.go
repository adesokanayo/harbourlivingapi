// Code generated by sqlc. DO NOT EDIT.

package db

import (
	"database/sql"
	"time"
)

type Artist struct {
	ID          int32          `json:"id"`
	UserID      int32          `json:"user_id"`
	DisplayName sql.NullString `json:"display_name"`
	AvatarUrl   sql.NullString `json:"avatar_url"`
	ShortBio    sql.NullString `json:"short_bio"`
	CreatedAt   time.Time      `json:"created_at"`
}

type Category struct {
	ID          int32          `json:"id"`
	Description string         `json:"description"`
	Image       sql.NullString `json:"image"`
	Status      int32          `json:"status"`
	CreatedAt   sql.NullTime   `json:"created_at"`
}

type Dayplan struct {
	ID            int32          `json:"id"`
	StartTime     time.Time      `json:"start_time"`
	EndTime       time.Time      `json:"end_time"`
	ScheduleID    int32          `json:"schedule_id"`
	Title         sql.NullString `json:"title"`
	Description   sql.NullString `json:"description"`
	PerformerName sql.NullString `json:"performer_name"`
	CreatedAt     time.Time      `json:"created_at"`
}

type Event struct {
	ID          int32         `json:"id"`
	Title       string        `json:"title"`
	Description string        `json:"description"`
	BannerImage string        `json:"banner_image"`
	StartDate   time.Time     `json:"start_date"`
	EndDate     time.Time     `json:"end_date"`
	Venue       int32         `json:"venue"`
	Type        int32         `json:"type"`
	UserID      int32         `json:"user_id"`
	Category    int32         `json:"category"`
	TicketID    sql.NullInt32 `json:"ticket_id"`
	Recurring   sql.NullBool  `json:"recurring"`
	Status      int32         `json:"status"`
	CreatedAt   sql.NullTime  `json:"created_at"`
}

type EventsArtist struct {
	ID        int32     `json:"id"`
	EventID   int32     `json:"event_id"`
	ArtistID  int32     `json:"artist_id"`
	CreatedAt time.Time `json:"created_at"`
}

type EventsFavorite struct {
	ID        int32     `json:"id"`
	EventID   int32     `json:"event_id"`
	UserID    int32     `json:"user_id"`
	CreatedAt time.Time `json:"created_at"`
}

type EventsHost struct {
	ID        int32     `json:"id"`
	EventID   int32     `json:"event_id"`
	HostID    int32     `json:"host_id"`
	CreatedAt time.Time `json:"created_at"`
}

type EventsImage struct {
	ID        int32     `json:"id"`
	EventID   int32     `json:"event_id"`
	ImageID   int32     `json:"image_id"`
	CreatedAt time.Time `json:"created_at"`
}

type EventsSponsor struct {
	ID        int32     `json:"id"`
	EventID   int32     `json:"event_id"`
	SponsorID int32     `json:"sponsor_id"`
	CreatedAt time.Time `json:"created_at"`
}

type EventsStatus struct {
	ID          int32          `json:"id"`
	Description sql.NullString `json:"description"`
	CreatedAt   time.Time      `json:"created_at"`
}

type EventsType struct {
	ID          int32  `json:"id"`
	Description string `json:"description"`
	Status      int32  `json:"status"`
}

type EventsVideo struct {
	ID        int32     `json:"id"`
	EventID   int32     `json:"event_id"`
	VideoID   int32     `json:"video_id"`
	CreatedAt time.Time `json:"created_at"`
}

type EventsView struct {
	ID        int32     `json:"id"`
	EventID   int32     `json:"event_id"`
	UserID    int32     `json:"user_id"`
	CreatedAt time.Time `json:"created_at"`
}

type Host struct {
	ID          int32          `json:"id"`
	UserID      int32          `json:"user_id"`
	DisplayName sql.NullString `json:"display_name"`
	AvatarUrl   sql.NullString `json:"avatar_url"`
	ShortBio    sql.NullString `json:"short_bio"`
	CreatedAt   time.Time      `json:"created_at"`
}

type Image struct {
	ID   int32          `json:"id"`
	Name sql.NullString `json:"name"`
	Url  string         `json:"url"`
}

type News struct {
	ID           int32          `json:"id"`
	Title        string         `json:"title"`
	Description  string         `json:"description"`
	FeatureImage sql.NullString `json:"feature_image"`
	Body         string         `json:"body"`
	UserID       int32          `json:"user_id"`
	PublishDate  time.Time      `json:"publish_date"`
	Tags         sql.NullString `json:"tags"`
	CreatedAt    time.Time      `json:"created_at"`
}

type Plan struct {
	ID          int32     `json:"id"`
	Name        string    `json:"name"`
	Description string    `json:"description"`
	Price       float64   `json:"price"`
	NoOfDays    int32     `json:"no_of_days"`
	CreatedAt   time.Time `json:"created_at"`
}

type Promotion struct {
	ID        int32     `json:"id"`
	EventID   int32     `json:"event_id"`
	UserID    int32     `json:"user_id"`
	PlanID    int32     `json:"plan_id"`
	StartDate time.Time `json:"start_date"`
	EndDate   time.Time `json:"end_date"`
	CreatedAt time.Time `json:"created_at"`
}

type Schedule struct {
	ID        int32     `json:"id"`
	EventID   int32     `json:"event_id"`
	Date      time.Time `json:"date"`
	StartTime time.Time `json:"start_time"`
	EndTime   time.Time `json:"end_time"`
	CreatedAt time.Time `json:"created_at"`
}

type Sponsor struct {
	ID          int32          `json:"id"`
	UserID      int32          `json:"user_id"`
	DisplayName sql.NullString `json:"display_name"`
	AvatarUrl   sql.NullString `json:"avatar_url"`
	ShortBio    sql.NullString `json:"short_bio"`
	CreatedAt   time.Time      `json:"created_at"`
}

type Ticket struct {
	ID          int32          `json:"id"`
	Name        string         `json:"name"`
	EventID     int32          `json:"event_id"`
	Price       float64        `json:"price"`
	Currency    string         `json:"currency"`
	Description sql.NullString `json:"description"`
}

type User struct {
	ID                int32          `json:"id"`
	Phone             sql.NullString `json:"phone"`
	FirstName         string         `json:"first_name"`
	LastName          string         `json:"last_name"`
	Email             string         `json:"email"`
	Username          string         `json:"username"`
	Password          string         `json:"password"`
	PasswordChangedAt sql.NullTime   `json:"password_changed_at"`
	Usertype          int32          `json:"usertype"`
	AvatarUrl         sql.NullString `json:"avatar_url"`
	DateOfBirth       time.Time      `json:"date_of_birth"`
	CreatedAt         time.Time      `json:"created_at"`
}

type UsersTicket struct {
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

type UsersType struct {
	ID          int32  `json:"id"`
	Description string `json:"description"`
	Status      int32  `json:"status"`
}

type Venue struct {
	ID          int32           `json:"id"`
	Name        string          `json:"name"`
	Address     sql.NullString  `json:"address"`
	PostalCode  sql.NullString  `json:"postal_code"`
	City        sql.NullString  `json:"city"`
	Province    sql.NullString  `json:"province"`
	CountryCode sql.NullString  `json:"country_code"`
	VenueOwner  int32           `json:"venue_owner"`
	BannerImage sql.NullString  `json:"banner_image"`
	Rating      sql.NullFloat64 `json:"rating"`
	Longitude   sql.NullFloat64 `json:"longitude"`
	Latitude    sql.NullFloat64 `json:"latitude"`
	Status      int32           `json:"status"`
	CreatedAt   sql.NullTime    `json:"created_at"`
}

type VenuesFavorite struct {
	ID        int32     `json:"id"`
	VenueID   int32     `json:"venue_id"`
	UserID    int32     `json:"user_id"`
	CreatedAt time.Time `json:"created_at"`
}

type VenuesStatus struct {
	ID          int32          `json:"id"`
	Description sql.NullString `json:"description"`
	CreatedAt   time.Time      `json:"created_at"`
}

type Video struct {
	ID   int32          `json:"id"`
	Name sql.NullString `json:"name"`
	Url  string         `json:"url"`
}
