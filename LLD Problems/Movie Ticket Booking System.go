// City -> Cinema -> Hall -> Movie{ Showtime, Title, Langauge, genre, ReleaseDate}

// // SearchByTitle()ListofMovie
// // SearchByLanguage() ListofMovie
// // SearchByGenre() ListofMovies
// // SearchByReleaseDate() ListofMovies
// // SearchByCity() ListofCinema
// // Search By

package main

import (
	"time"
)

type BMSService struct {
	Cinemas []CinemaHall
}

func (b *BMSService) GetMovies(date time.Time, city string) []Movie {
	// Implementation
	return nil
}

func (b *BMSService) GetCinemaHalls(city string) []CinemaHall {
	// Implementation
	return nil
}

type CinemaHall struct {
	CinemaHallID   int
	CinemaHallName string
	Address        Address
	AudiList       []Audi
}

func (c *CinemaHall) GetMovies(dateList []time.Time) map[time.Time][]Movie {
	// Implementation
	return nil
}

func (c *CinemaHall) GetShows(dateList []time.Time) map[time.Time][]Show {
	// Implementation
	return nil
}

type Address struct {
	PinCode int
	Street  string
	City    string
	State   string
	Country string
}

type Audi struct {
	AudiID     int
	AudiName   string
	TotalSeats int
	Shows      []Show
}

type Show struct {
	ShowID         int
	Movie          Movie
	StartTime      time.Time
	EndTime        time.Time
	CinemaPlayedAt CinemaHall
	Seats          []Seat
}

type Seat struct {
	SeatID     int
	SeatType   SeatType
	SeatStatus SeatStatus
	Price      float64
}

type SeatType int

const (
	Delux SeatType = iota
	VIP
	Economy
	Other
)

type SeatStatus int

const (
	Booked SeatStatus = iota
	Available
	Reserved
	NotAvailable
)

type Movie struct {
	MovieName      string
	MovieID        int
	DurationInMins int
	Language       string
	Genre          Genre
	ReleaseDate    time.Time
	CityShowMap    map[string][]Show
}

type Genre int

const (
	SciFi Genre = iota
	Drama
	RomCom
	Fantasy
)

type User struct {
	UserID int
	Search Search
}

type SystemMember struct {
	User
	Account Account
	Name    string
	Email   string
	Address Address
}

type Member struct {
	SystemMember
}

func (m *Member) MakeBooking(booking Booking) Booking {
	// Implementation
	return Booking{}
}

func (m *Member) GetBooking() []Booking {
	// Implementation
	return nil
}

type Admin struct {
	SystemMember
}

func (a *Admin) AddMovie(movie Movie) bool {
	// Implementation
	return false
}

func (a *Admin) AddShow(show Show) bool {
	// Implementation
	return false
}

type Account struct {
	UserName string
	Password string
}

type Search struct{}

func (s *Search) SearchMoviesByNames(name string) []Movie {
	// Implementation
	return nil
}

func (s *Search) SearchMoviesByGenre(genre Genre) []Movie {
	// Implementation
	return nil
}

func (s *Search) SearchMoviesByLanguage(language string) []Movie {
	// Implementation
	return nil
}

func (s *Search) SearchMoviesByDate(releaseDate time.Time) []Movie {
	// Implementation
	return nil
}

type Booking struct {
	BookingID     string
	BookingDate   time.Time
	Member        Member
	Audi          Audi
	Show          Show
	BookingStatus BookingStatus
	TotalAmount   float64
	Seats         []Seat
	Payment       Payment
}

func (b *Booking) MakePayment(payment Payment) bool {
	// Implementation
	return false
}

type Payment struct {
	Amount        float64
	PaymentDate   time.Time
	TransactionID int
	PaymentStatus PaymentStatus
}

type BookingStatus int

const (
	Requested BookingStatus = iota
	Pending
	Confirmed
	Cancelled
)

type PaymentStatus int

const (
	Unpaid PaymentStatus = iota
	Pending
	Completed
	Declined
	Cancelled
	Refunded
)

func main() {
	// Main function
}
