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

var (
	store *db.Store
)
var tokenMaker token.Maker

func init() {
	config, err := util.LoadConfig(".")
	if err != nil {
		log.Fatal("cannot find config ", err)
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

	if input.Virtual {
		createVirtual := db.CreateVirtualVenueParams{
			Name: input.Name,
			Url: sql.NullString{
				Valid:  true,
				String: *input.URL,
			},
			Virtual: false,
		}

		venue, err := store.CreateVirtualVenue(ctx, createVirtual)
		if err != nil {
			return nil, err
		}
		return &Venue{Name: venue.Name, ID: venue.ID}, nil
	}

	var result Venue
	var createVenueReq db.CreateVenueParams

	createVenueReq.Name = input.Name

	if input.Address != nil {
		createVenueReq.Address = sql.NullString{
			String: *input.Address,
			Valid:  true,
		}
	}

	if input.City != nil {
		createVenueReq.City = sql.NullString{
			String: *input.City,
			Valid:  true,
		}
	}

	if input.PostalCode != nil {
		createVenueReq.PostalCode = sql.NullString{
			String: *input.PostalCode,
			Valid:  true,
		}
	}

	if input.Province != nil {
		createVenueReq.Province = sql.NullString{
			String: *input.Province,
			Valid:  true,
		}
	}

	if input.CountryCode != nil {
		createVenueReq.CountryCode = sql.NullString{
			String: *input.CountryCode,
			Valid:  true,
		}
	}

	if input.Longitude != nil && input.Latitude != nil {
		createVenueReq.Longitude = sql.NullFloat64{
			Float64: *input.Longitude,
			Valid:   true,
		}
		createVenueReq.Latitude = sql.NullFloat64{
			Float64: *input.Latitude,
			Valid:   true,
		}
	}
	venue, err := store.CreateVenue(ctx, createVenueReq)
	if err != nil {
		return nil, err
	}

	result.ID = venue.ID
	result.Name = venue.Name

	if venue.PostalCode.Valid {
		result.PostalCode = &venue.PostalCode.String
	}
	if venue.City.Valid {
		result.City = &venue.City.String
	}
	if venue.Province.Valid {
		result.Province = &venue.Province.String
	}
	if venue.CountryCode.Valid {
		result.CountryCode = &venue.CountryCode.String
	}
	if venue.Address.Valid {
		result.Address = &venue.Address.String
	}
	if venue.Rating.Valid {
		result.Rating = venue.Rating.Float64
	}
	if venue.Longitude.Valid {
		result.Longitude = &venue.Longitude.Float64
	}
	if venue.Latitude.Valid {
		result.Latitude = &venue.Latitude.Float64
	}

	return &result, nil
}

func (r *mutationResolver) CreateUser(ctx context.Context, input NewUser) (*User, error) {

	var phone sql.NullString
	var avatar sql.NullString

	hashedPassword, err := util.HashPassword(input.Password)
	if err != nil {
		return nil, err
	}

	if input.Phone != nil {
		phone.String = *input.Phone
		phone.Valid = true
	}

	if input.Avatar != nil {
		avatar.String = *input.Avatar
		avatar.Valid = true
	}

	arg := db.CreateUserParams{
		Phone:     phone,
		FirstName: input.FirstName,
		LastName:  input.LastName,
		Email:     input.Email,
		Password:  hashedPassword,
		Username:  input.Username,
		Usertype:  int32(input.Usertype),
		AvatarUrl: avatar,
	}

	user, err := store.CreateUser(context.Background(), arg)
	if err != nil {
		return nil, err
	}
	return &User{
		ID:        user.ID,
		Phone:     &user.Phone.String,
		FirstName: user.FirstName,
		LastName:  user.LastName,
		Email:     user.Email,
		Username:  user.Username,
		Password:  user.Password,
		Usertype:  int(user.Usertype),
		Avatar: user.AvatarUrl.String,
	}, nil
}

func (r *mutationResolver) CreateEvent(ctx context.Context, input NewEvent) (*Event, error) {

	var result *Event
	startdate, err := util.ProcessDateTime("rfc", input.StartDate)
	if err != nil {
		return nil, err
	}
	enddate, err := util.ProcessDateTime("rfc", input.EndDate)
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
		Status:      int32(input.Status),
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

		// link Host with Event
		arg := db.LinkHostToEventParams{
			HostID:  host.ID,
			EventID: event.ID,
		}

		linkedEventHost, err := store.LinkHostToEvent(ctx, arg)
		if err != nil {
			return err
		}

		//create images
		var images []*Image
		if input.Images != nil {
			for _, i := range input.Images {

				arg := db.CreateImageParams{
					Name: sql.NullString{
						String: i.Name,
						Valid:  true},
					Url: *i.URL,
				}

				image, err := store.CreateImage(ctx, arg)
				if err != nil {
					return err
				}

				//link Image to Events
				argImageLink := db.LinkImageToEventParams{
					EventID: event.ID,
					ImageID: image.ID,
				}

				errs := store.LinkImageToEvent(ctx, argImageLink)
				if errs != nil {
					return errs
				}

				images = append(images, &Image{
					ID:      image.ID,
					EventID: event.ID,
					Name:    image.Name.String,
					URL:     &image.Url,
				})

			}
		}

		//create vidoes and link to Event
		var videos []*Video
		if input.Vidoes != nil {
			for _, i := range input.Vidoes {

				arg := db.CreateVideoParams{
					Name: sql.NullString{
						String: i.Name,
						Valid:  true},
					Url: *i.URL,
				}

				video, err := store.CreateVideo(ctx, arg)
				if err != nil {
					return err
				}

				//link Video to Events
				argVideoLink := db.LinkVideoToEventParams{
					EventID: event.ID,
					VideoID: video.ID,
				}

				errs := store.LinkVideoToEvent(ctx, argVideoLink)
				if errs != nil {
					return errs
				}

				videos = append(videos, &Video{
					ID:      video.ID,
					EventID: event.ID,
					Name:    video.Name.String,
					URL:     &video.Url,
				})
			}
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
			Images:      images,
			Videos:      videos,
		}
		return nil
	})

	if err != nil {
		return nil, err
	}
	return result, nil
}

func (r *mutationResolver) UpdateEvent(ctx context.Context, input UpdateEvent) (*Event, error) {

	var result *Event
	//fetch Event

	arg := db.UpdateEventParams{}

	event, err := store.GetEvent(ctx, int32(input.ID))
	if err != nil {
		return nil, err
	}

	if input.Title != nil {
		arg.Title = *input.Title
		arg.TitleToUpdate = true

	}
	if input.Description != nil {
		arg.Description = *input.Description
		arg.DescriptionToUpdate = true

	}

	if input.BannerImage != nil {
		arg.BannerImageToUpdate = true
		arg.BannerImage = *input.BannerImage
	}

	if input.StartDate != nil {
		startdate, err := util.ProcessDateTime("rfc", *input.StartDate)
		if err != nil {
			return nil, err
		}
		arg.StartDateToUpdate = true
		arg.StartDate = *startdate
	}

	if input.EndDate != nil {
		enddate, err := util.ProcessDateTime("rfc", *input.EndDate)
		if err != nil {
			return nil, err
		}
		arg.EndDateToUpdate = true
		arg.EndDate = *enddate
	}

	if input.Venue != nil {
		arg.Venue = int32(*input.Venue)
		arg.VenueToUpdate = true
	}

	if input.Type != nil {
		arg.Type = int32(*input.Type)
		arg.TypeToUpdate = true
	}

	if input.Category != nil {
		arg.Category = int32(*input.Category)
		arg.CategoryToUpdate = true
	}

	if input.Status != nil {
		arg.Status = int32(*input.Status)
		arg.StatusToUpdate = true
	}

	arg.ID = event.ID

	//Use Transaction
	err = store.ExecTx(ctx, func(q *db.Queries) error {

		//update event
		event, err := store.UpdateEvent(ctx, arg)
		if err != nil {
			return err
		}

		//create images
		var images []*Image
		if input.Images != nil {
			for _, i := range input.Images {

				arg := db.CreateImageParams{
					Name: sql.NullString{
						String: i.Name,
						Valid:  true},
					Url: *i.URL,
				}

				image, err := store.CreateImage(ctx, arg)
				if err != nil {
					return err
				}

				//link Image to Events
				argImageLink := db.LinkImageToEventParams{
					EventID: event.ID,
					ImageID: image.ID,
				}

				errs := store.LinkImageToEvent(ctx, argImageLink)
				if errs != nil {
					return errs
				}

				images = append(images, &Image{
					ID:      image.ID,
					EventID: event.ID,
					Name:    image.Name.String,
					URL:     &image.Url,
				})

			}
		}

		//create vidoes and link to Event
		var videos []*Video
		if input.Vidoes != nil {
			for _, i := range input.Vidoes {

				arg := db.CreateVideoParams{
					Name: sql.NullString{
						String: i.Name,
						Valid:  true},
					Url: *i.URL,
				}

				video, err := store.CreateVideo(ctx, arg)
				if err != nil {
					return err
				}

				//link Video to Events
				argVideoLink := db.LinkVideoToEventParams{
					EventID: event.ID,
					VideoID: video.ID,
				}

				errs := store.LinkVideoToEvent(ctx, argVideoLink)
				if errs != nil {
					return errs
				}

				videos = append(videos, &Video{
					ID:      video.ID,
					EventID: event.ID,
					Name:    video.Name.String,
					URL:     &video.Url,
				})
			}
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
			Images:      images,
			Videos:      videos,
			Venue:       int(event.Venue),
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
			Phone:     &user.Phone.String,
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
		Phone:     &user.Phone.String,
		FirstName: user.FirstName,
		LastName:  user.LastName,
		Email:     user.Email,
		Username:  user.Username,
		Password:  user.Password,
		Usertype:  int(user.Usertype),
	}, nil
}

func (r *queryResolver) GetVenue(ctx context.Context, input int32) (*Venue, error) {
	var result Venue
	venue, err := store.GetVenue(ctx, input)
	if err != nil {
		return nil, err
	}

	result.ID = venue.ID
	if venue.Province.Valid {
		result.Province = &venue.Province.String
	}
	if venue.City.Valid {
		result.City = &venue.City.String
	}
	if venue.PostalCode.Valid {
		result.PostalCode = &venue.PostalCode.String
	}
	if venue.Address.Valid {
		result.Address = &venue.Address.String
	}
	if venue.CountryCode.Valid {
		result.CountryCode = &venue.CountryCode.String
	}

	if venue.Latitude.Valid {
		result.Latitude = &venue.Latitude.Float64
	}

	if venue.Longitude.Valid {
		result.Longitude = &venue.Longitude.Float64
	}
	return &result, nil
}

func (r *queryResolver) GetEvent(ctx context.Context, input int32) (*Event, error) {

	var sponsors []*Sponsor
	var images []*Image
	var videos []*Video
	var tickets []*Ticket

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
		sponsors = append(sponsors, &Sponsor{ID: s.SponsorID})
	}

	// fetch images
	eventImages, err := store.GetImagesByEvent(ctx, input)
	if err != nil {
		return nil, err
	}
	for _, i := range eventImages {
		images = append(images, &Image{
			ID:      i.ID,
			EventID: i.EventID,
			Name:    i.Name.String,
			URL:     &i.Url,
		})
	}

	// fetch videos
	eventVideos, err := store.GetVideosByEvent(ctx, input)
	if err != nil {
		return nil, err
	}
	for _, i := range eventVideos {
		videos = append(videos, &Video{
			ID:      i.ID,
			EventID: i.EventID,
			Name:    i.Name.String,
			URL:     &i.Url,
		})
	}

	// fetch tickets
	eventTickets, err := store.GetTicketsByEventID(ctx, event.ID)
	if err != nil {
		return nil, err
	}

	for _, v := range eventTickets {
		tickets = append(tickets, &Ticket{
			ID:       v.ID,
			Name:     v.Name,
			Price:    int(v.Price),
			EventID:  int(v.EventID),
			Quantity: int(v.Quantity),
			Status:   int(v.Status),
		})
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
		Sponsors:    sponsors,
		Status:      int(event.Status),
		Images:      images,
		Videos:      videos,
		Ticket:      tickets,
	}, nil
}

func (r *queryResolver) GetEvents(ctx context.Context, input GetEvent) ([]Event, error) {

	var result []Event
	var images []*Image
	var videos []*Video
	var eventSponsors []*Sponsor

	arg := db.GetEventsParams{
		Category:    int32(input.Category),
		Subcategory: int32(input.Subcategory),
		Status:      int32(input.Status),
		Limit:       int32(input.PageSize),
		Offset:      int32(input.Offset),
	}
	events, err := store.GetEvents(ctx, arg)
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

		// fetch images
		eventImages, err := store.GetImagesByEvent(ctx, event.ID)
		if err != nil {
			return nil, err
		}
		for _, i := range eventImages {
			images = append(images, &Image{
				ID:      i.ID,
				EventID: i.EventID,
				Name:    i.Name.String,
				URL:     &i.Url,
			})
		}

		// fetch videos
		eventVideos, err := store.GetVideosByEvent(ctx, event.ID)
		if err != nil {
			return nil, err
		}
		for _, i := range eventVideos {
			videos = append(videos, &Video{
				ID:      i.ID,
				EventID: i.EventID,
				Name:    i.Name.String,
				URL:     &i.Url,
			})
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
			Images:      images,
			Videos:      videos,
		})
	}

	return result, nil
}

func (r *queryResolver) GetEventsByLocation(ctx context.Context, input GetEventByLocation) ([]Event, error) {

	var result []Event
	var eventSponsors []*Sponsor

	arg := db.GetEventsByLocationParams{
		Point:     input.Latitude,
		Point_2:   input.Longitude,
		Longitude: sql.NullFloat64{Valid: true, Float64: float64(input.Miles)},
	}
	events, err := store.GetEventsByLocation(ctx, arg)
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
			Phone:     &u.Phone.String,
			FirstName: u.FirstName,
			LastName:  u.LastName,
			Email:     u.Email,
			Username:  u.Username,
			Password:  u.Password,
			Usertype:  int(u.Usertype),
			Avatar:    u.AvatarUrl.String,
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

func (r *mutationResolver) UpdateEventStatus(ctx context.Context, input UpdateEventStatus) (*UpdateEventState, error) {
	arg := db.UpdateEventStatusParams{
		ID:     int32(input.EventID),
		Status: int32(input.EventStatus),
	}
	result, err := store.UpdateEventStatus(ctx, arg)
	if err != nil {
		return nil, err
	}
	return &UpdateEventState{
		EventID:     int(result.ID),
		EventStatus: int(result.Status),
	}, nil
}

func (q *queryResolver) GetCategory(ctx context.Context, id int32) (*Category, error) {

	category, err := store.GetCategory(ctx, id)

	if err != nil {
		return nil, err
	}
	return &Category{
		ID:     category.ID,
		Desc:   category.Desc,
		Status: int(category.Status.Int32),
	}, nil
}
func (q *queryResolver) GetCategories(ctx context.Context) ([]Category, error) {

	var result []Category
	categories, err := store.GetCategories(ctx)

	if err != nil {
		return nil, err
	}

	for _, v := range categories {
		result = append(result, Category{
			v.ID,
			v.Desc,
			v.Image.String,
			int(v.Status.Int32),
		})
	}

	return result, nil
}

func (q *queryResolver) GetSubcategory(ctx context.Context, id int32) (*Subcategory, error) {
	subcategory, err := store.GetSubCategory(ctx, id)
	if err != nil {
		return nil, err
	}
	return &Subcategory{
		ID:     subcategory.ID,
		Desc:   subcategory.Desc,
		Status: int(subcategory.Status.Int32),
	}, nil
}

// Mutation returns MutationResolver implementation.
func (r *Resolver) Mutation() MutationResolver { return &mutationResolver{r} }

// Query returns QueryResolver implementation.
func (r *Resolver) Query() QueryResolver { return &queryResolver{r} }

type mutationResolver struct{ *Resolver }
type queryResolver struct{ *Resolver }
