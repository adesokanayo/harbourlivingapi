package gqlgen

// THIS CODE IS A STARTING POINT ONLY. IT WILL NOT BE UPDATED WITH SCHEMA CHANGES.

import (
	"context"
	"database/sql"
	"log"
	"time"

	db "github.com/BigListRyRy/harbourlivingapi/db/sqlc"
	token "github.com/BigListRyRy/harbourlivingapi/token"
	"github.com/BigListRyRy/harbourlivingapi/util"
	_ "github.com/lib/pq"
	"github.com/pkg/errors"
)

var store *db.Store
var tokenMaker token.Maker

func init() {
	config, err := util.LoadConfig(".")
	if err != nil {
		log.Fatal("cannot local config", err)
	}

	conn, err := sql.Open(config.DBDriver, config.DBSource)
	if err != nil {
		log.Fatalln("cannot connect to database, ", err)
	}

	store = db.NewStore(conn)
	tokenMaker, err = token.NewJWTMaker(config.TokenSymmetricKey)
	if err != nil {
		log.Fatalln("cannot create a token maker ", err)
	}

}

type Resolver struct {
	repo *db.Store
}

func (r *mutationResolver) CreateVenue(ctx context.Context, input NewVenue) (*Venue, error) {
	createEventReq := db.CreateVenueParams{
		Name:        input.Name,
		Address:     input.Address,
		PostalCode:  input.PostalCode,
		City:        input.City,
		Province:    input.Province,
		CountryCode: input.CountryCode,
	}
	venue, err := store.CreateVenue(ctx, createEventReq)
	if err != nil {
		return nil, err
	}
	return &Venue{
		ID:          venue.ID,
		Name:        venue.Name,
		PostalCode:  venue.PostalCode,
		Address:     venue.Address,
		City:        venue.City,
		Province:    venue.Province,
		CountryCode: venue.CountryCode,
	}, nil
}

func (r *mutationResolver) CreateUser(ctx context.Context, input NewUser) (*User, error) {

	hashedPassword, err := util.HashPassword(input.Password)
	if err != nil {
		return nil, err
	}
	arg := db.CreateUserParams{
		Title:     input.Title,
		FirstName: input.FirstName,
		LastName:  input.LastName,
		Email:     input.Email,
		Password:  hashedPassword,
		Username:  input.Username,
		Usertype:  int32(input.Usertype),
	}
	user, err := store.CreateUser(context.Background(), arg)
	if err != nil {
		return nil, err
	}
	return &User{
		ID:        user.ID,
		Title:     user.Title,
		FirstName: user.FirstName,
		LastName:  user.LastName,
		Email:     user.Email,
		Username:  user.Username,
		Password:  user.Password,
		Usertype:  int(user.Usertype),
	}, nil
}

func (r *mutationResolver) CreateEvent(ctx context.Context, input NewEvent) (*Event, error) {

	var result *Event
	startdate, err := util.ProcessDateTime(input.StartDate)
	if err != nil {
		return nil, err
	}
	enddate, err := util.ProcessDateTime(input.EndDate)
	if err != nil {
		return nil, err
	}
	arg := db.CreateEventParams{
		Title:       input.Title,
		Description: input.Description,
		BannerImage: input.BannerImage,
		StartDate:   *startdate,
		EndDate:     *enddate,
		Venue:       int32(input.Venue),
		Type:        int32(input.Type),
		UserID:      int32(input.UserID),
		Category:    int32(input.Category),
		Subcategory: int32(input.Subcategory),
		Status:      sql.NullString{},
		Image1:      sql.NullString{},
		Image2:      sql.NullString{},
		Image3:      sql.NullString{},
		Video1:      sql.NullString{},
		Video2:      sql.NullString{},
	}

	//Use Transaction
	err = store.ExecTx(ctx, func(q *db.Queries) error {

		//create event
		event, err := store.CreateEvent(ctx, arg)
		if err != nil {
			return err
		}

		//create host
		host, err := store.CreateHost(ctx, input.UserID)
		if err != nil {
			return err
		}

		//Link Host with Event

		arg := db.LinkHostToEventParams{
			HostID:  host.ID,
			EventID: event.ID,
		}

		linkedEventHost, err := store.LinkHostToEvent(ctx, arg)
		if err != nil {
			return nil
		}

		result = &Event{
			ID:          event.ID,
			Title:       event.Title,
			Description: event.Description,
			StartDate:   event.StartDate.String(),
			EndDate:     event.EndDate.String(),
			BannerImage: event.BannerImage,
			Subcategory: int(event.Subcategory),
			Category:    int(event.Category),
			HostID:      int(linkedEventHost.HostID),
		}
		return nil
	})

	if err != nil {
		return nil, err
	}
	return result, nil
}

func (r *mutationResolver) Login(ctx context.Context, input Login) (*LoginResponse, error) {

	user, err := store.GetUsername(ctx, input.Username)
	if err != nil {
		if err == sql.ErrNoRows {
			return nil, errors.New("invalid user")
		}
		return nil, err
	}

	err = util.CheckPassword(input.Password, user.Password)
	if err != nil {
		return nil, errors.New("invalid username & password combination")
	}

	token, err := tokenMaker.CreateToken(user.Username, time.Hour*24)
	if err != nil {
		return nil, errors.New("unable to create token")
	}
	return &LoginResponse{
		Token: &token,
		User: &User{
			ID:        user.ID,
			Email:     user.Email,
			Usertype:  int(user.Usertype),
			Username:  user.Username,
			LastName:  user.LastName,
			Title:     user.Title,
			FirstName: user.FirstName,
		},
		Success: true,
	}, nil
}

func (r *mutationResolver) RefreshToken(ctx context.Context, input RefreshTokenInput) (*string, error) {
	username, err := tokenMaker.ParseToken(input.Token)
	if err != nil {
		return nil, errors.New("invalid token")
	}

	token, err := tokenMaker.CreateToken(username, time.Hour*24)
	if err != nil {
		return nil, err
	}
	return &token, nil

}

func (r *queryResolver) GetUser(ctx context.Context, input int32) (*User, error) {

	user, err := store.GetUser(ctx, input)
	if err != nil {
		if err == sql.ErrNoRows {
			return nil, errors.New("no record found")
		}

		return nil, errors.New("an error occured")
	}

	return &User{
		ID:        user.ID,
		Title:     user.Title,
		FirstName: user.FirstName,
		LastName:  user.LastName,
		Email:     user.Email,
		Username:  user.Username,
		Password:  user.Password,
		Usertype:  int(user.Usertype),
	}, nil
}

func (r *queryResolver) GetVenue(ctx context.Context, input int32) (*Venue, error) {

	venue, err := store.GetVenue(ctx, input)
	if err != nil {
		return nil, err
	}
	return &Venue{
		ID:          venue.ID,
		Name:        venue.Name,
		PostalCode:  venue.PostalCode,
		Address:     venue.Address,
		City:        venue.City,
		Province:    venue.Province,
		CountryCode: venue.CountryCode,
	}, nil
}

func (r *queryResolver) GetEvent(ctx context.Context, input int32) (*Event, error) {

	var result []*Sponsor
	event, err := store.GetEvent(ctx, input)
	if err != nil {
		if err == sql.ErrNoRows {
			return nil, errors.New("no record found")
		}
		return nil, err
	}
	//Fetch  Sponsors

	eventSponsors, err := store.GetSponsorByEvent(ctx, input)
	if err != nil {
		return nil, err
	}

	for _, s := range eventSponsors {
		result = append(result, &Sponsor{ID: s.SponsorID})
	}

	return &Event{
		ID:          event.ID,
		Title:       event.Title,
		Description: event.Description,
		BannerImage: event.BannerImage,
		StartDate:   event.StartDate.String(),
		EndDate:     event.EndDate.String(),
		Venue:       int(event.Venue),
		Type:        int(event.Type),
		UserID:      event.UserID,
		Category:    int(event.Category),
		Subcategory: int(event.Subcategory),
		Sponsors:    result,
		Status:      nil,
		Image1:      nil,
		Image2:      nil,
		Image3:      nil,
		Video1:      nil,
		Video2:      nil,
	}, nil
}

func (r *queryResolver) GetEventByProperties(ctx context.Context, input GetEvent) ([]Event, error) {

	var result []Event
	var eventSponsors []*Sponsor

	arg := db.GetEventsByFilterParams{
		Category:    int32(input.Category),
		Subcategory: int32(input.Subcategory),
		City:        *input.City,
		Province:    *input.Province,
		Limit:       int32(input.PageSize),
		Offset:      int32(input.Offset),
	}
	events, err := store.GetEventsByFilter(ctx, arg)
	if err != nil {
		return nil, err
	}
	// Get all the Sponsors for these events

	for _, event := range events {
		sponsors, err := store.GetSponsorByEvent(ctx, event.ID)

		if err != nil {
			return nil, err
		}

		for _, v := range sponsors {
			eventSponsors = append(eventSponsors, &Sponsor{
				ID: v.ID})
		}

		result = append(result, Event{
			ID:          event.ID,
			Title:       event.Title,
			Description: event.Description,
			StartDate:   event.StartDate.String(),
			EndDate:     event.EndDate.String(),
			Category:    int(event.Category),
			Subcategory: int(event.Subcategory),
			Type:        int(event.Type),
			UserID:      event.UserID,
			Venue:       int(event.Venue),
			BannerImage: event.BannerImage,
			Sponsors:    eventSponsors,
		})
	}

	return result, nil
}
func (r *queryResolver) GetUsers(ctx context.Context) ([]User, error) {
	var AllUsers []User
	users, err := store.GetAllUsers(ctx)
	if err != nil {
		return nil, err
	}

	for _, u := range users {
		AllUsers = append(AllUsers, User{
			ID:        u.ID,
			Title:     u.Title,
			FirstName: u.FirstName,
			LastName:  u.LastName,
			Email:     u.Email,
			Username:  u.Username,
			Password:  u.Password,
			Usertype:  int(u.Usertype),
		})
	}
	return AllUsers, nil
}

func (r *mutationResolver) CreateTicket(ctx context.Context, input NewTicket) (*Ticket, error) {
	arg := db.CreateTicketParams{
		EventID:  int32(input.EventID),
		Quantity: int32(input.Quantity),
		Price:    float64(input.Price),
		Status:   int32(input.Status),
		Name:     input.Name,
	}

	ticket, err := store.CreateTicket(ctx, arg)
	if err != nil {
		return nil, err
	}

	return &Ticket{
		ID:       ticket.ID,
		Name:     ticket.Name,
		Price:    int(ticket.Price),
		EventID:  int(ticket.EventID),
		Quantity: int(ticket.Quantity),
		Status:   int(ticket.Status),
	}, nil
}

func (r *mutationResolver) CreateSponsorForEvent(ctx context.Context, input NewSponsor) (*Sponsor, error) {

	var result *Sponsor

	err := store.ExecTx(ctx, func(q *db.Queries) error {

		// Create a Sponsor based on the provided user
		sponsor, err := store.CreateSponsor(ctx, int32(input.UserID))
		if err != nil {
			return err
		}

		// Link the Sponsor to the Event
		arg := db.LinkSponsorToEventParams{
			SponsorID: sponsor.ID,
			EventID:   int32(input.EventID),
		}

		linkedSponsor, err := store.LinkSponsorToEvent(ctx, arg)
		if err != nil {
			return err
		}

		result = &Sponsor{
			ID:     linkedSponsor.SponsorID,
			UserID: input.UserID,
		}
		return nil
	})
	if err != nil {
		return nil, err
	}
	return result, nil
}

// Mutation returns MutationResolver implementation.
func (r *Resolver) Mutation() MutationResolver { return &mutationResolver{r} }

// Query returns QueryResolver implementation.
func (r *Resolver) Query() QueryResolver { return &queryResolver{r} }

type mutationResolver struct{ *Resolver }
type queryResolver struct{ *Resolver }
