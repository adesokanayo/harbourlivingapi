package gqlgen

// THIS CODE IS A STARTING POINT ONLY. IT WILL NOT BE UPDATED WITH SCHEMA CHANGES.

import (
	"context"
	"database/sql"
	db "github.com/BigListRyRy/harbourlivingapi/db/sqlc"
	"github.com/BigListRyRy/harbourlivingapi/util"
	_ "github.com/lib/pq"
	"log"
)

var store *db.Store

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

}

type Resolver struct {
	repo *db.Store
}

func (r *mutationResolver) CreateVenue(ctx context.Context, input NewVenue) (*Venue, error) {

	createEventReq := db.CreateVenueParams{
		Name:        input.Name,
		Address:     input.Address,
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
	event, err := store.CreateEvent(ctx, arg)
	if err != nil {
		return nil, err
	}

	return &Event{
		ID:          event.ID,
		Title:       event.Title,
		Description: event.Description,
		StartDate:   event.StartDate.String(),
		EndDate:     event.EndDate.String(),
		BannerImage: event.BannerImage,
		Subcategory: int(event.Subcategory),
		Category:    int(event.Category),
	}, nil
}

func (r *mutationResolver) Login(ctx context.Context, input Login) (string, error) {
	panic("not implemented")
}

func (r *mutationResolver) RefreshToken(ctx context.Context, input RefreshTokenInput) (string, error) {
	panic("not implemented")
}

func (r *queryResolver) GetUser(ctx context.Context, input int32) (*User, error) {

	user, err := store.GetUser(ctx, input)
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

func (r *queryResolver) GetVenue(ctx context.Context, input int32) (*Venue, error) {

	panic("not implemented")
}

func (r *queryResolver) GetEvent(ctx context.Context, input int32) (*Event, error) {

	panic("not implemented")
}


func (r *queryResolver) GetEventByProperties(ctx context.Context, category *int,subcategory *int, city *string, province *string) ([]Event, error) {

	var result []Event

	arg:= db.GetEventsByFilterParams{
		Category:    1,
		Subcategory: 1,
		City:       "Calgary",
		Province:    "AB",
		Limit:      10,
		Offset:      1,
	}
	events , err := store.GetEventsByFilter(ctx, arg)
	if err != nil {
		return nil, err
	}

	 for _, event := range events{
		 result = append(result, Event{
		 	ID: event.ID,
		 	Title: event.Title,
		 	Description: event.Description,
		 	StartDate: event.StartDate.String(),
		 	EndDate:  event.EndDate.String(),
		 	Category: int(event.Category),
		 	Subcategory: int(event.Subcategory),
		 	Type: int(event.Type),
		 	UserID: event.UserID,
		 	Venue: int(event.Venue),
		 	BannerImage: event.BannerImage,
		 })
	 }

	return  result,nil
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

// Mutation returns MutationResolver implementation.
func (r *Resolver) Mutation() MutationResolver { return &mutationResolver{r} }

// Query returns QueryResolver implementation.
func (r *Resolver) Query() QueryResolver { return &queryResolver{r} }

type mutationResolver struct{ *Resolver }
type queryResolver struct{ *Resolver }
