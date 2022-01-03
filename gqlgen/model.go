// Code generated by github.com/99designs/gqlgen, DO NOT EDIT.

package gqlgen

type Artist struct {
	ID          int32    `json:"id"`
	UserID      int      `json:"user_id"`
	DisplayName *string  `json:"display_name"`
	Avatar      *string  `json:"avatar"`
	ShortBio    *string  `json:"short_bio"`
	Events      []*Event `json:"events"`
}

type Category struct {
	ID          int32  `json:"id"`
	Description string `json:"description"`
	Image       string `json:"image"`
	Status      int    `json:"status"`
}

type Event struct {
	ID          int32      `json:"id"`
	Title       string     `json:"title"`
	Description string     `json:"description"`
	BannerImage string     `json:"banner_image"`
	StartDate   string     `json:"startDate"`
	EndDate     string     `json:"endDate"`
	Venue       int        `json:"venue"`
	Type        int        `json:"type"`
	UserID      int32      `json:"user_id"`
	Category    int        `json:"category"`
	Sponsors    []*Sponsor `json:"sponsors"`
	HostID      int        `json:"hostID"`
	Tickets     []*Ticket  `json:"tickets"`
	Status      int        `json:"status"`
	Images      []*Image   `json:"images"`
	Videos      []*Video   `json:"videos"`
}

type EventFavorite struct {
	ID      int32 `json:"id"`
	EventID int   `json:"event_id"`
	UserID  int   `json:"user_id"`
}

type EventType struct {
	ID          int32  `json:"id"`
	Description string `json:"description"`
	Status      int    `json:"status"`
}

type GetEvent struct {
	Category int `json:"category"`
	PageSize int `json:"pageSize"`
	Offset   int `json:"offset"`
	Status   int `json:"status"`
}

type GetEventByLocation struct {
	Latitude  float64 `json:"latitude"`
	Longitude float64 `json:"longitude"`
	Miles     int     `json:"miles"`
}

type Host struct {
	ID          int32    `json:"id"`
	UserID      int      `json:"user_id"`
	DisplayName *string  `json:"display_name"`
	Avatar      *string  `json:"avatar"`
	ShortBio    *string  `json:"short_bio"`
	Events      []*Event `json:"events"`
}

type Image struct {
	ID      int32   `json:"id"`
	EventID int32   `json:"event_id"`
	Name    *string `json:"name"`
	URL     string  `json:"url"`
}

type Login struct {
	Username string `json:"username"`
	Password string `json:"password"`
}

type LoginResponse struct {
	Token   *string `json:"token"`
	User    *User   `json:"user"`
	Success bool    `json:"success"`
	Message *string `json:"message"`
}

type NewArtist struct {
	UserID  int `json:"user_id"`
	EventID int `json:"event_id"`
}

type NewCategory struct {
	Description string  `json:"description"`
	Image       *string `json:"image"`
	Status      int     `json:"status"`
}

type NewEvent struct {
	Title       string       `json:"title"`
	Description string       `json:"description"`
	BannerImage string       `json:"banner_image"`
	StartDate   string       `json:"startDate"`
	EndDate     string       `json:"endDate"`
	Venue       int          `json:"venue"`
	Type        int          `json:"type"`
	UserID      int32        `json:"user_id"`
	Category    int          `json:"category"`
	Status      int          `json:"status"`
	Images      []*NewImage  `json:"images"`
	Videos      []*NewVideo  `json:"videos"`
	Tickets     []*NewTicket `json:"tickets"`
}

type NewEventFavorite struct {
	EventID int `json:"event_id"`
	UserID  int `json:"user_id"`
}

type NewHost struct {
	UserID  int `json:"user_id"`
	EventID int `json:"event_id"`
}

type NewImage struct {
	Name *string `json:"name"`
	URL  string  `json:"url"`
}

type NewSponsor struct {
	UserID  int `json:"user_id"`
	EventID int `json:"event_id"`
}

type NewTicket struct {
	Name        string  `json:"name"`
	Price       float64 `json:"price"`
	EventID     int     `json:"event_id"`
	Currency    string  `json:"currency"`
	Description *string `json:"description"`
}

type NewUser struct {
	Phone     *string `json:"phone"`
	FirstName string  `json:"first_name"`
	LastName  string  `json:"last_name"`
	Email     string  `json:"email"`
	Username  string  `json:"username"`
	Password  string  `json:"password"`
	Usertype  int     `json:"usertype"`
	Avatar    *string `json:"avatar"`
}

type NewVenue struct {
	Name        string   `json:"name"`
	Address     *string  `json:"address"`
	PostalCode  *string  `json:"postal_code"`
	City        *string  `json:"city"`
	Province    *string  `json:"province"`
	CountryCode *string  `json:"country_code"`
	VenueOwner  int      `json:"venue_owner"`
	BannerImage *string  `json:"banner_image"`
	Longitude   *float64 `json:"longitude"`
	Latitude    *float64 `json:"latitude"`
	Rating      *int     `json:"rating"`
	Status      int      `json:"status"`
}

type NewVenueFavorite struct {
	VenueID int `json:"venue_id"`
	UserID  int `json:"user_id"`
}

type NewVideo struct {
	Name *string `json:"name"`
	URL  string  `json:"url"`
}

type RefreshTokenInput struct {
	Token string `json:"token"`
}

type Sponsor struct {
	ID          int32    `json:"id"`
	UserID      int      `json:"user_id"`
	DisplayName *string  `json:"display_name"`
	Avatar      *string  `json:"avatar"`
	ShortBio    *string  `json:"short_bio"`
	Events      []*Event `json:"events"`
}

type Ticket struct {
	ID          int32   `json:"id"`
	Name        string  `json:"name"`
	Price       float64 `json:"price"`
	EventID     int     `json:"event_id"`
	Currency    string  `json:"currency"`
	Description *string `json:"description"`
}

type UpdateArtist struct {
	ID          int     `json:"id"`
	Avatar      *string `json:"avatar"`
	DisplayName *string `json:"display_name"`
	ShortBio    *string `json:"short_bio"`
}

type UpdateEvent struct {
	ID          int         `json:"id"`
	Title       *string     `json:"title"`
	Description *string     `json:"description"`
	BannerImage *string     `json:"banner_image"`
	StartDate   *string     `json:"startDate"`
	EndDate     *string     `json:"endDate"`
	Venue       *int        `json:"venue"`
	Type        *int        `json:"type"`
	Category    *int        `json:"category"`
	Status      *int        `json:"status"`
	Images      []*NewImage `json:"images"`
	Videos      []*NewVideo `json:"videos"`
}

type UpdateEventState struct {
	EventID     int `json:"event_id"`
	EventStatus int `json:"event_status"`
}

type UpdateEventStatus struct {
	EventID     int `json:"event_id"`
	EventStatus int `json:"event_status"`
}

type UpdateHost struct {
	ID          int     `json:"id"`
	Avatar      *string `json:"avatar"`
	DisplayName *string `json:"display_name"`
	ShortBio    *string `json:"short_bio"`
}

type UpdateSponsor struct {
	ID          int     `json:"id"`
	Avatar      *string `json:"avatar"`
	DisplayName *string `json:"display_name"`
	ShortBio    *string `json:"short_bio"`
}

type UpdateVenue struct {
	ID          int      `json:"id"`
	Name        *string  `json:"name"`
	Address     *string  `json:"address"`
	PostalCode  *string  `json:"postal_code"`
	City        *string  `json:"city"`
	Province    *string  `json:"province"`
	CountryCode *string  `json:"country_code"`
	URL         *string  `json:"url"`
	Longitude   *float64 `json:"longitude"`
	Latitude    *float64 `json:"latitude"`
	Rating      *int     `json:"rating"`
	Status      *int     `json:"status"`
}

type User struct {
	ID              int32    `json:"id"`
	Phone           *string  `json:"phone"`
	FirstName       string   `json:"first_name"`
	LastName        string   `json:"last_name"`
	Email           string   `json:"email"`
	Username        string   `json:"username"`
	Password        string   `json:"password"`
	Usertype        int      `json:"usertype"`
	Avatar          string   `json:"avatar"`
	FavoritesVenues []*Venue `json:"favorites_venues"`
	FavoritesEvents []*Event `json:"favorites_events"`
}

type Usertype struct {
	ID          int32  `json:"id"`
	Description string `json:"description"`
	Status      *int   `json:"status"`
}

type Venue struct {
	ID          int32    `json:"id"`
	Name        string   `json:"name"`
	Address     *string  `json:"address"`
	PostalCode  *string  `json:"postal_code"`
	City        *string  `json:"city"`
	Province    *string  `json:"province"`
	CountryCode *string  `json:"country_code"`
	Longitude   *float64 `json:"Longitude"`
	Latitude    *float64 `json:"Latitude"`
	BannerImage *string  `json:"banner_image"`
	VenueOwner  int      `json:"venue_owner"`
	Rating      *int     `json:"rating"`
	Status      int      `json:"status"`
}

type VenueFavorite struct {
	ID      int32 `json:"id"`
	VenueID int   `json:"venue_id"`
	UserID  int   `json:"user_id"`
}

type Video struct {
	ID      int32   `json:"id"`
	EventID int32   `json:"event_id"`
	Name    *string `json:"name"`
	URL     string  `json:"url"`
}
