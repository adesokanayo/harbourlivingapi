package graphql

// THIS CODE IS A STARTING POINT ONLY. IT WILL NOT BE UPDATED WITH SCHEMA CHANGES.
//go:generate go run github.com/99designs/gqlgen

import (
	"context"
	"database/sql"
	"log"
	"time"

	db "github.com/BigListRyRy/harbourlivingapi/db/sqlc"
	"github.com/BigListRyRy/harbourlivingapi/middleware"
	token "github.com/BigListRyRy/harbourlivingapi/token"
	"github.com/BigListRyRy/harbourlivingapi/util"
	_ "github.com/lib/pq"
	"github.com/pkg/errors"
)

const (
	AttendeeHarbour HarbourUserTypes = iota
	HostHarbour
	SponsorHarbour
	AdminHarbour
)

var (
	store *db.Store
)
var tokenMaker token.Maker

type HarbourUserTypes int

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

	if input.BannerImage != nil {
		createVenueReq.BannerImage = sql.NullString{
			String: *input.BannerImage,
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
	createVenueReq.Status = int32(input.Status)
	createVenueReq.VenueOwner = int32(input.VenueOwner)

	venue, err := store.CreateVenue(ctx, createVenueReq)
	if err != nil {
		return nil, err
	}

	result.ID = venue.ID
	result.Name = venue.Name
	result.VenueOwner = int(venue.VenueOwner)

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

	if venue.BannerImage.Valid {
		result.Address = &venue.BannerImage.String
	}
	if venue.Longitude.Valid {
		result.Longitude = &venue.Longitude.Float64
	}
	if venue.Latitude.Valid {
		result.Latitude = &venue.Latitude.Float64
	}

	result.Status = int(venue.Status)

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
		Avatar:    user.AvatarUrl.String,
	}, nil
}

func (r *mutationResolver) CreateCategory(ctx context.Context, input NewCategory) (*Category, error) {

	userinfo := middleware.CtxValue(ctx)
	if userinfo.UserType != int(AdminHarbour) {
		return nil, errors.New(util.ErrPermissionDenied)
	}
	arg := db.CreateCategoryParams{}

	if input.Image != nil {
		arg.Image = sql.NullString{String: *input.Image, Valid: true}
	}
	arg.Description = input.Description
	arg.Status = int32(input.Status)

	category, err := store.CreateCategory(ctx, arg)

	if err != nil {
		return nil, err
	}

	return &Category{
		ID:     category.ID,
		Image:  category.Image.String,
		Status: int(category.Status),
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
						String: *i.Name,
						Valid:  true},
					Url: i.URL,
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
					Name:    &image.Name.String,
					URL:     image.Url,
				})

			}
		}

		//create vidoes and link to Event
		var videos []*Video
		if input.Videos != nil {
			for _, i := range input.Videos {

				arg := db.CreateVideoParams{
					Name: sql.NullString{
						String: *i.Name,
						Valid:  true},
					Url: i.URL,
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
					Name:    &video.Name.String,
					URL:     video.Url,
				})
			}
		}

		//create tickets and link to event
		var tickets []*Ticket
		if input.Tickets != nil {
			for _, i := range input.Tickets {

				arg := db.CreateTicketParams{
					Name:     i.Name,
					Currency: i.Currency,
					EventID:  event.ID,
					Price:    float64(i.Price),
					Description: sql.NullString{
						Valid:  true,
						String: *i.Description,
					},
				}

				ticket, err := store.CreateTicket(ctx, arg)
				if err != nil {
					return err
				}

				tickets = append(tickets, &Ticket{
					ID:          ticket.ID,
					EventID:     int(event.ID),
					Name:        ticket.Name,
					Currency:    i.Currency,
					Price:       ticket.Price,
					Description: &ticket.Description.String,
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
			Category:    int(event.Category),
			HostID:      int(linkedEventHost.HostID),
			Images:      images,
			Videos:      videos,
			Status:      int(event.Status),
			Tickets:     tickets,
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
						String: *i.Name,
						Valid:  true},
					Url: i.URL,
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
					Name:    &image.Name.String,
					URL:     image.Url,
				})

			}
		}

		//create videos and link to Event
		var videos []*Video
		if input.Videos != nil {
			for _, i := range input.Videos {

				arg := db.CreateVideoParams{
					Name: sql.NullString{
						String: *i.Name,
						Valid:  true},
					Url: i.URL,
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
					Name:    &video.Name.String,
					URL:     video.Url,
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

func (r *mutationResolver) DeleteEvent(ctx context.Context, input int32) (bool, error) {

	_, err := store.GetEvent(ctx, input)
	if err != nil {
		return false, err
	}

	err = store.DeleteEvent(ctx, input)
	if err != nil {
		return false, err
	}
	return true, nil
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

	userinfo := token.UserInfo{
		UserID:   int(user.ID),
		Username: user.Username,
		Email:    user.Email,
		UserType: int(user.Usertype),
	}

	token, err := tokenMaker.CreateToken(userinfo, time.Hour*24)
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

	user, err := store.GetUsername(ctx, username)
	if err != nil {
		if err == sql.ErrNoRows {
			return nil, errors.New("invalid user")
		}
		return nil, err
	}

	userinfo := token.UserInfo{
		UserID:   int(user.ID),
		Username: user.Username,
		Email:    user.Email,
		UserType: int(user.Usertype),
	}

	token, err := tokenMaker.CreateToken(userinfo, time.Hour*24)
	if err != nil {
		return nil, err
	}
	return &token, nil

}

func (r *queryResolver) GetUser(ctx context.Context, input int32) (*User, error) {
	var events []*Event
	var venues []*Venue
	user, err := store.GetUser(ctx, input)
	if err != nil {
		if err == sql.ErrNoRows {
			return nil, errors.New("no record found")
		}

		return nil, errors.New("an error occured")
	}

	favEvents, err := store.GetFavoriteEvents(ctx, user.ID)
	if err != nil {
		return nil, err
	}

	for _, v := range favEvents {
		f, err := GetEventHelper(ctx, v.EventID)
		if err != nil {
			return nil, err
		}
		events = append(events, f)
	}

	favVenues, err := store.GetFavoriteVenues(ctx, user.ID)
	if err != nil {
		return nil, err
	}

	for _, v := range favVenues {
		f, err := GetVenueHelper(ctx, v.VenueID)
		if err != nil {
			return nil, err
		}
		venues = append(venues, f)
	}

	return &User{
		ID:              user.ID,
		Phone:           &user.Phone.String,
		FirstName:       user.FirstName,
		LastName:        user.LastName,
		Email:           user.Email,
		Username:        user.Username,
		Password:        user.Password,
		Usertype:        int(user.Usertype),
		FavoritesEvents: events,
		FavoritesVenues: venues,
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
	var promoted bool

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
			Name:    &i.Name.String,
			URL:     i.Url,
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
			Name:    &i.Name.String,
			URL:     i.Url,
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
			Price:    v.Price,
			EventID:  int(v.EventID),
			Currency: v.Currency,
		})
	}

	// determine favorites
	eventsFavorites, err := store.GetFavoriteEvents(ctx, event.ID)
	if err != nil {
		return nil, err
	}

	// get views
	eventsViews, err := store.GetViewedEvents(ctx, event.ID)
	if err != nil {
		return nil, err
	}

	// check if event is Promoted
	promotions, err := store.GetPromotionsForEvent(ctx, event.ID)
	if err != nil {
		return nil, err
	}
	if len(promotions) > 0 {
		promoted = true
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
		Sponsors:    sponsors,
		Status:      int(event.Status),
		Images:      images,
		Videos:      videos,
		Tickets:     tickets,
		Meta: &Metadata{
			TotalFavorite: int32(len(eventsFavorites)),
			TotalView:     int32(len(eventsViews)),
		},
		Promoted: promoted,
	}, nil
}

func (r *queryResolver) GetEvents(ctx context.Context, input GetEvent) ([]Event, error) {

	var result []Event
	var images []*Image
	var videos []*Video
	var eventSponsors []*Sponsor
	var promoted bool

	arg := db.GetEventsParams{
		Category: int32(input.Category),
		Status:   int32(input.Status),
		Limit:    int32(input.PageNumber),
		Offset:   int32((input.PageNumber * input.Limit) - input.Limit),
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
				Name:    &i.Name.String,
				URL:     i.Url,
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
				Name:    &i.Name.String,
				URL:     i.Url,
			})
		}

		//check if event is promoted
		promotions, err := store.GetPromotionsForEvent(ctx, event.ID)
		if err != nil {
			return nil, err
		}
		if len(promotions) > 0 {
			promoted = true
		}

		result = append(result, Event{
			ID:          event.ID,
			Title:       event.Title,
			Description: event.Description,
			StartDate:   event.StartDate.String(),
			EndDate:     event.EndDate.String(),
			Category:    int(event.Category),
			Type:        int(event.Type),
			UserID:      event.UserID,
			Venue:       int(event.Venue),
			BannerImage: event.BannerImage,
			Sponsors:    eventSponsors,
			Images:      images,
			Videos:      videos,
			Promoted:    promoted,
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
			Type:        int(event.Type),
			UserID:      event.UserID,
			Venue:       int(event.Venue),
			//BannerImage: *event.BannerImage,
			Sponsors: eventSponsors,
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
		Price:    float64(input.Price),
		Name:     input.Name,
		Currency: input.Currency,
		Description: sql.NullString{
			String: *input.Description,
			Valid:  true,
		},
	}

	ticket, err := store.CreateTicket(ctx, arg)
	if err != nil {
		return nil, err
	}

	return &Ticket{
		ID:          ticket.ID,
		Name:        ticket.Name,
		Price:       ticket.Price,
		EventID:     int(ticket.EventID),
		Currency:    ticket.Currency,
		Description: &ticket.Description.String,
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

func (r *mutationResolver) UpdateHost(ctx context.Context, input UpdateHost) (*Host, error) {

	arg := db.UpdateHostParams{}
	arg.ID = int32(input.ID)
	if input.Avatar != nil {
		arg.AvatarUrlToUpdate = true
		arg.AvatarUrl = *input.Avatar
	}

	if input.DisplayName != nil {
		arg.DisplayNameToUpdate = true
		arg.DisplayName = *input.DisplayName
	}

	if input.ShortBio != nil {
		arg.ShortBioToUpdate = true
		arg.ShortBio = *input.ShortBio
	}

	host, err := store.UpdateHost(ctx, arg)
	if err != nil {
		return nil, err
	}

	return &Host{
		ID:          host.ID,
		UserID:      int(host.UserID),
		DisplayName: &host.DisplayName.String,
		Avatar:      &host.AvatarUrl.String,
		ShortBio:    &host.ShortBio.String,
	}, nil
}

func (r *mutationResolver) UpdateArtist(ctx context.Context, input UpdateArtist) (*Artist, error) {

	arg := db.UpdateArtistParams{}
	arg.ID = int32(input.ID)
	if input.Avatar != nil {
		arg.AvatarUrlToUpdate = true
		arg.AvatarUrl = *input.Avatar
	}

	if input.DisplayName != nil {
		arg.DisplayNameToUpdate = true
		arg.DisplayName = *input.DisplayName
	}

	if input.ShortBio != nil {
		arg.ShortBioToUpdate = true
		arg.ShortBio = *input.ShortBio
	}

	artist, err := store.UpdateArtist(ctx, arg)
	if err != nil {
		return nil, err
	}

	return &Artist{
		ID:          artist.ID,
		UserID:      int(artist.UserID),
		DisplayName: &artist.DisplayName.String,
		Avatar:      &artist.AvatarUrl.String,
		ShortBio:    &artist.ShortBio.String,
	}, nil
}

func (r *mutationResolver) UpdateSponsor(ctx context.Context, input UpdateSponsor) (*Sponsor, error) {

	arg := db.UpdateSponsorParams{}
	arg.ID = int32(input.ID)
	if input.Avatar != nil {
		arg.AvatarUrlToUpdate = true
		arg.AvatarUrl = *input.Avatar
	}

	if input.DisplayName != nil {
		arg.DisplayNameToUpdate = true
		arg.DisplayName = *input.DisplayName
	}

	if input.ShortBio != nil {
		arg.ShortBioToUpdate = true
		arg.ShortBio = *input.ShortBio
	}

	sponsor, err := store.UpdateSponsor(ctx, arg)
	if err != nil {
		return nil, err
	}

	return &Sponsor{
		ID:          sponsor.ID,
		UserID:      int(sponsor.UserID),
		DisplayName: &sponsor.DisplayName.String,
		Avatar:      &sponsor.AvatarUrl.String,
		ShortBio:    &sponsor.ShortBio.String,
	}, nil
}

func (q *queryResolver) GetCategory(ctx context.Context, id int32) (*Category, error) {

	category, err := store.GetCategory(ctx, id)

	if err != nil {
		return nil, err
	}
	return &Category{
		ID:          category.ID,
		Description: category.Description,
		Status:      int(category.Status),
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
			v.Description,
			v.Image.String,
			int(v.Status),
		})
	}

	return result, nil
}

func (r *mutationResolver) CreateEventFavorite(ctx context.Context, input NewEventFavorite) (*EventFavorite, error) {

	userinfo := middleware.CtxValue(ctx)

	arg := db.CreateFavoriteEventParams{
		EventID: int32(input.EventID),
		UserID:  int32(userinfo.UserID),
	}

	eventFavorite, err := store.CreateFavoriteEvent(ctx, arg)
	if err != nil {
		return nil, err
	}

	return &EventFavorite{
		ID:      eventFavorite.ID,
		EventID: int(eventFavorite.EventID),
		UserID:  int(eventFavorite.UserID),
	}, nil
}

func (r *mutationResolver) CreateVenueFavorite(ctx context.Context, input NewVenueFavorite) (*VenueFavorite, error) {

	arg := db.CreateVenueFavoriteParams{
		VenueID: int32(input.VenueID),
		UserID:  int32(input.UserID),
	}

	venueFavorite, err := store.CreateVenueFavorite(ctx, arg)
	if err != nil {
		return nil, err
	}

	return &VenueFavorite{
		ID:      venueFavorite.ID,
		VenueID: int(venueFavorite.VenueID),
		UserID:  int(venueFavorite.UserID),
	}, nil
}

func GetEventHelper(ctx context.Context, input int32) (*Event, error) {

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
			Name:    &i.Name.String,
			URL:     i.Url,
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
			Name:    &i.Name.String,
			URL:     i.Url,
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
			Price:    v.Price,
			EventID:  int(v.EventID),
			Currency: v.Currency,
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
		Sponsors:    sponsors,
		Status:      int(event.Status),
		Images:      images,
		Videos:      videos,
		Tickets:     tickets,
	}, nil
}

func GetVenueHelper(ctx context.Context, input int32) (*Venue, error) {

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

func (r *mutationResolver) UpdateVenue(ctx context.Context, input UpdateVenue) (*Venue, error) {
	arg := db.UpdateVenueParams{
		ID: int32(input.ID),
	}
	if input.Name != nil {
		arg.Name = *input.Name
		arg.NameToUpdate = true
	}

	if input.Address != nil {
		arg.Address = *input.Address
		arg.AddressToUpdate = true
	}

	if input.PostalCode != nil {
		arg.PostalCode = *input.PostalCode
		arg.PostalCodeToUpdate = true
	}

	if input.City != nil {
		arg.City = *input.City
		arg.CityToUpdate = true
	}

	if input.CountryCode != nil {
		arg.CountryCode = *input.CountryCode
		arg.CountryToUpdate = true
	}

	if input.Longitude != nil {
		arg.Longitude = *input.Longitude
		arg.LongitudeToUpdate = true
	}

	if input.Latitude != nil {
		arg.Latitude = *input.Latitude
		arg.LatitudeToUpdate = true
	}

	if input.Rating != nil {
		arg.Rating = int32(*input.Rating)
		arg.RatingToUpdate = true
	}

	if input.Status != nil {
		arg.Status = int32(*input.Status)
		arg.StatusToUpdate = true
	}

	venue, err := store.UpdateVenue(ctx, arg)
	if err != nil {
		return nil, err
	}
	return &Venue{
		ID:          venue.ID,
		Name:        venue.Name,
		Address:     &venue.Address.String,
		PostalCode:  &venue.PostalCode.String,
		City:        &venue.City.String,
		Province:    &venue.Province.String,
		CountryCode: &venue.CountryCode.String,
		Longitude:   &venue.Longitude.Float64,
		Latitude:    &venue.Latitude.Float64,
	}, nil
}

func GetDateHelper(dateString string) (*time.Time, error) {
	date, err := util.ProcessDateTime("rfc", dateString)
	if err != nil {
		return nil, err
	}
	return date, nil
}
func (r *mutationResolver) CreatePlan(ctx context.Context, input NewPlan) (*Plan, error) {

	arg := db.CreatePlanParams{
		Name:        input.Name,
		Description: input.Description,
		Price:       input.Price,
		NoOfDays:    int32(input.NoOfDays),
	}

	plan, err := store.CreatePlan(ctx, arg)
	if err != nil {
		return nil, err
	}
	return &Plan{
		ID:          plan.ID,
		Name:        plan.Name,
		Description: plan.Description,
		Price:       plan.Price,
		NoOfDays:    int(plan.NoOfDays),
	}, nil
}

func (r *mutationResolver) CreatePromotion(ctx context.Context, input NewPromotion) (*Promotion, error) {

	d1, _ := GetDateHelper(input.StartDate)

	d2, _ := GetDateHelper(input.EndDate)

	arg := db.CreatePromotionParams{
		EventID:   input.EventID,
		UserID:    input.UserID,
		PlanID:    input.PlanID,
		StartDate: *d1,
		EndDate:   *d2,
	}

	promotion, err := store.CreatePromotion(ctx, arg)
	if err != nil {
		return nil, err
	}
	return &Promotion{
		ID:        promotion.ID,
		EventID:   promotion.EventID,
		UserID:    promotion.UserID,
		StartDate: promotion.StartDate.String(),
		EndDate:   promotion.EndDate.String(),
	}, nil
}

func (r *mutationResolver) UpdatePlan(ctx context.Context, input UpdatePlan) (*Plan, error) {
	arg := db.UpdatePlanParams{
		ID: input.ID,
	}
	if input.Name != nil {
		arg.NameToUpdate = true
		arg.Name = *input.Name
	}
	if input.Description != nil {
		arg.Description = *input.Description
		arg.DescriptionToUpdate = true
	}
	if input.Price != nil {
		arg.Price = *input.Price
		arg.PriceToUpdate = true
	}
	if input.NoOfDays != nil {
		arg.NoOfDays = int32(*input.NoOfDays)
		arg.NoOfDaysToUpdate = true
	}

	plan, err := store.UpdatePlan(ctx, arg)
	if err != nil {
		return nil, err
	}
	return &Plan{
		ID:          plan.ID,
		Name:        plan.Name,
		Description: plan.Description,
		Price:       plan.Price,
		NoOfDays:    int(plan.NoOfDays),
	}, nil
}

func (r *mutationResolver) UpdatePromotion(ctx context.Context, input UpdatePromotion) (*Promotion, error) {
	arg := db.UpdatePromotionParams{
		ID: input.ID,
	}

	if input.StartDate != nil {
		d1, _ := GetDateHelper(*input.StartDate)
		arg.StartDate = *d1
		arg.StartDateToUpdate = true
	}

	if input.EndDate != nil {
		d1, _ := GetDateHelper(*input.EndDate)
		arg.EndDate = *d1
		arg.EndDateToUpdate = true
	}

	promotion, err := store.UpdatePromotion(ctx, arg)
	if err != nil {
		return nil, err
	}
	return &Promotion{
		ID:        promotion.ID,
		EventID:   promotion.EventID,
		UserID:    promotion.UserID,
		StartDate: promotion.StartDate.String(),
		EndDate:   promotion.EndDate.String(),
	}, nil
}

func (r *mutationResolver) CreateNews(ctx context.Context, input NewNews) (*News, error) {

	d1, _ := GetDateHelper(input.PublishDate)

	arg := db.CreateNewsParams{
		Title:       input.Title,
		Description: input.Description,
		FeatureImage: sql.NullString{
			String: input.FeatureImage,
			Valid:  true,
		},
		Body:        input.Body,
		PublishDate: *d1,
		UserID:      input.UserID,
		Tags: sql.NullString{
			String: *input.Tags,
			Valid:  true,
		},
		Status: int32(input.Status),
	}

	news, err := store.CreateNews(ctx, arg)
	if err != nil {
		return nil, err
	}
	return &News{
		ID:           news.ID,
		Title:        news.Title,
		Description:  news.Description,
		FeatureImage: news.FeatureImage.String,
		Body:         news.Body,
		PublishDate:  news.PublishDate.String(),
		Tags:         &news.Tags.String,
		Status:       int(news.Status),
	}, nil
}

func (r *mutationResolver) UpdateNews(ctx context.Context, input UpdateNews) (*News, error) {
	arg := db.UpdateNewsParams{
		ID: input.ID,
	}
	d1, _ := GetDateHelper(*input.PublishDate)

	if input.Title != nil {
		arg.TitleIDToUpdate = true
		arg.Title = *input.Title
	}
	if input.Description != nil {
		arg.Description = *input.Description
		arg.DescriptionToUpdate = true
	}
	if input.FeatureImage != nil {
		arg.FeatureImage = *input.FeatureImage
		arg.FeatureImageToUpdate = true
	}
	if input.Body != nil {
		arg.Body = *input.Body
		arg.BodyToUpdate = true
	}

	if input.PublishDate != nil {
		arg.PublishDate = *d1
		arg.PublishDateToUpdate = true
	}

	if input.Tags != nil {
		arg.Tags = *input.Tags
		arg.TagsDateToUpdate = true
	}

	if input.Tags != nil {
		arg.Tags = *input.Tags
		arg.TagsDateToUpdate = true
	}

	if input.Status != nil {
		arg.Status = int32(*input.Status)
		arg.StatusToUpdate = true
	}

	news, err := store.UpdateNews(ctx, arg)
	if err != nil {
		return nil, err
	}
	return &News{
		ID:           news.ID,
		Title:        news.Title,
		Description:  news.Description,
		FeatureImage: news.FeatureImage.String,
		Body:         news.Body,
		PublishDate:  news.PublishDate.String(),
		Tags:         &news.Tags.String,
		Status:       int(news.Status),
	}, nil
}

func (r *mutationResolver) DeleteNews(ctx context.Context, input int32) (bool, error) {

	_, err := store.GetNews(ctx, input)
	if err != nil {
		return false, err
	}

	err = store.DeleteNews(ctx, input)
	if err != nil {
		return false, err
	}
	return true, nil
}

func (r *mutationResolver) DeletePlan(ctx context.Context, input int32) (bool, error) {

	_, err := store.GetNews(ctx, input)
	if err != nil {
		return false, err
	}

	err = store.DeleteNews(ctx, input)
	if err != nil {
		return false, err
	}
	return true, nil
}

func (r *mutationResolver) DeletePromotion(ctx context.Context, input int32) (bool, error) {

	_, err := store.GetPromotion(ctx, input)
	if err != nil {
		return false, err
	}

	err = store.DeletePromotion(ctx, input)
	if err != nil {
		return false, err
	}
	return true, nil
}

func (r *mutationResolver) DeleteTicket(ctx context.Context, input int32) (bool, error) {

	_, err := store.GetTicket(ctx, input)
	if err != nil {
		return false, err
	}

	err = store.DeleteTicket(ctx, input)
	if err != nil {
		return false, err
	}
	return true, nil
}

func (r *queryResolver) GetPlan(ctx context.Context, input int32) (*Plan, error) {

	plan, err := store.GetPlan(ctx, input)

	if err != nil {
		return nil, err
	}
	return &Plan{
		ID:          plan.ID,
		Name:        plan.Name,
		Description: plan.Description,
		Price:       plan.Price,
		NoOfDays:    int(plan.NoOfDays),
	}, nil

}

func (r *queryResolver) GetPromotion(ctx context.Context, input int32) (*Promotion, error) {

	promotion, err := store.GetPromotion(ctx, input)

	if err != nil {
		return nil, err
	}
	return &Promotion{
		ID:        promotion.ID,
		UserID:    promotion.EventID,
		EventID:   promotion.EventID,
		StartDate: promotion.StartDate.String(),
		EndDate:   promotion.EndDate.String(),
	}, nil

}

func (r *queryResolver) GetNews(ctx context.Context, input int32) (*News, error) {

	news, err := store.GetNews(ctx, input)

	if err != nil {
		return nil, err
	}
	return &News{
		ID:           news.ID,
		Title:        news.Title,
		Description:  news.Description,
		FeatureImage: *&news.FeatureImage.String,
		Body:         news.Body,
		PublishDate:  news.PublishDate.String(),
		Tags:         &news.Tags.String,
		Status:       int(news.Status),
	}, nil

}

func (r *queryResolver) GetAllPlans(ctx context.Context) ([]Plan, error) {

	var plans []Plan
	allPlans, err := store.GetAllPlans(ctx)

	if err != nil {
		return nil, err
	}

	for _, plan := range allPlans {
		plans = append(plans, Plan{
			ID:          plan.ID,
			Name:        plan.Name,
			Description: plan.Description,
			Price:       plan.Price,
			NoOfDays:    int(plan.NoOfDays),
		})
	}
	return plans, nil

}

func (r *queryResolver) GetAllPromotions(ctx context.Context) ([]Promotion, error) {
	var promotions []Promotion
	AllPromotions, err := store.GetAllPromotions(ctx)

	if err != nil {
		return nil, err
	}

	for _, promotion := range AllPromotions {

		promotions = append(promotions, Promotion{
			ID:        promotion.ID,
			UserID:    promotion.EventID,
			EventID:   promotion.EventID,
			StartDate: promotion.StartDate.String(),
			EndDate:   promotion.EndDate.String(),
		})
	}
	return promotions, nil

}

func (r *queryResolver) GetAllNews(ctx context.Context) ([]News, error) {
	var allNews []News

	dbNews, err := store.GetAllNews(ctx)

	if err != nil {
		return nil, err
	}
	for _, news := range dbNews {
		allNews = append(allNews, News{
			ID:           news.ID,
			Title:        news.Title,
			Description:  news.Description,
			FeatureImage: *&news.FeatureImage.String,
			Body:         news.Body,
			PublishDate:  news.PublishDate.String(),
			Tags:         &news.Tags.String,
		})
	}
	return allNews, nil

}

func (r *mutationResolver) CreateEventView(ctx context.Context, input NewEventView) (*EventView, error) {

	arg := db.CreateViewEventParams{
		EventID: int32(input.EventID),
		UserID:  int32(input.UserID),
	}

	eventView, err := store.CreateViewEvent(ctx, arg)
	if err != nil {
		return nil, err
	}

	return &EventView{
		ID:      eventView.ID,
		EventID: int(eventView.EventID),
		UserID:  int(eventView.UserID),
	}, nil
}

func (r *mutationResolver) CreateSchedule(ctx context.Context, input NewSchedule) (*Schedule, error) {

	d1, err := util.ProcessDateTime(util.LayoutISODOB, input.Date)
	if err != nil {
		return nil, err
	}

	arg := db.CreateScheduleParams{
		EventID:   int32(input.EventID),
		Date:      *d1,
		StartTime: input.StartTime,
		EndTime:   input.EndTime,
	}

	schedule, err := store.CreateSchedule(ctx, arg)
	if err != nil {
		return nil, err
	}

	return &Schedule{
		ID:        schedule.ID,
		EventID:   schedule.EventID,
		Date:      schedule.Date.String(),
		StartTime: schedule.StartTime,
		EndTime:   schedule.EndTime,
	}, nil
}

func (r *mutationResolver) UpdateSchedule(ctx context.Context, input UpdateSchedule) (*Schedule, error) {

	arg := db.UpdateScheduleParams{
		ID: input.ID,
	}

	if input.Date != nil {
		d1, err := util.ProcessDateTime(util.LayoutISODOB, *input.Date)
		if err != nil {
			return nil, err
		}
		arg.Date = *d1
		arg.DateIDToUpdate = true
	}

	if input.StartTime != nil {
		arg.StartTime = *input.StartTime
		arg.StartTimeToUpdate = true
	}

	if input.EndTime != nil {
		arg.EndTime = *input.EndTime
		arg.EndTimeToUpdate = true
	}

	schedule, err := store.UpdateSchedule(ctx, arg)
	if err != nil {
		return nil, err
	}

	return &Schedule{
		ID:        schedule.ID,
		EventID:   schedule.EventID,
		Date:      schedule.Date.String(),
		StartTime: schedule.StartTime,
		EndTime:   schedule.EndTime,
	}, nil
}
func (r *mutationResolver) CreateDayPlan(ctx context.Context, input NewDayPlan) (*DayPlan, error) {

	arg := db.CreateDayplanParams{
		StartTime:  input.StartTime,
		EndTime:    input.EndTime,
		ScheduleID: input.ScheduleID,
		Title: sql.NullString{
			Valid:  true,
			String: *input.Title,
		},
		Description: sql.NullString{
			Valid:  true,
			String: *input.Description,
		},
		PerformerName: sql.NullString{
			Valid:  true,
			String: *input.PerformerName,
		},
	}

	dayplan, err := store.CreateDayplan(ctx, arg)
	if err != nil {
		return nil, err
	}

	return &DayPlan{
		ID:            dayplan.ID,
		StartTime:     dayplan.StartTime,
		EndTime:       dayplan.EndTime,
		ScheduleID:    dayplan.ScheduleID,
		Title:         &dayplan.Title.String,
		Description:   &dayplan.Description.String,
		PerformerName: &dayplan.PerformerName.String,
	}, nil
}

func (r *mutationResolver) UpdateDayPlan(ctx context.Context, input UpdateDayPlan) (*DayPlan, error) {

	arg := db.UpdateDayPlanParams{
		ID: input.ID,
	}

	if input.StartTime != nil {
		arg.StartTime = *input.StartTime
		arg.StartTimeToUpdate = true
	}

	if input.EndTime != nil {
		arg.EndTime = *input.EndTime
		arg.EndTimeToUpdate = true
	}

	if input.Description != nil {
		arg.Description = *input.Description
		arg.DescriptionToUpdate = true
	}

	if input.PerformerName != nil {
		arg.PerformerName = *input.PerformerName
		arg.PerformerNameToUpdate = true
	}

	dayplan, err := store.UpdateDayPlan(ctx, arg)
	if err != nil {
		return nil, err
	}

	return &DayPlan{
		ID:            dayplan.ID,
		StartTime:     dayplan.StartTime,
		EndTime:       dayplan.EndTime,
		ScheduleID:    dayplan.ScheduleID,
		Title:         &dayplan.Title.String,
		Description:   &dayplan.Description.String,
		PerformerName: &dayplan.PerformerName.String,
	}, nil
}

func (r *mutationResolver) DeleteSchedule(ctx context.Context, input int32) (bool, error) {

	_, err := store.GetSchedule(ctx, input)
	if err != nil {
		return false, err
	}

	err = store.DeleteSchedule(ctx, input)
	if err != nil {
		return false, err
	}
	return true, nil
}

func (r *mutationResolver) DeleteDayPlan(ctx context.Context, input int32) (bool, error) {

	_, err := store.GetDayplan(ctx, input)
	if err != nil {
		return false, err
	}

	err = store.DeleteDayPlan(ctx, input)
	if err != nil {
		return false, err
	}
	return true, nil
}

// Mutation returns MutationResolver implementation.
func (r *Resolver) Mutation() MutationResolver { return &mutationResolver{r} }

// Query returns QueryResolver implementation.
func (r *Resolver) Query() QueryResolver { return &queryResolver{r} }

type mutationResolver struct{ *Resolver }
type queryResolver struct{ *Resolver }
