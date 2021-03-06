package services

// ==================================== RPC ====================================
type Response struct {
	Data         []*Data         `json:"data"`
	Dictionaries *Dictionaries   `json:"dictionaries"`
	Meta         *Meta           `json:"meta"`
	Warnings     []*ErrorWarning `json:"warnings"`
	Errors       []*ErrorWarning `json:"errors"`
}

type FlightLowFareSearchRequest struct {
	Origin        string
	Destination   string
	DepartureDate string

	ReturnDate      string
	ArrivalBy       string
	ReturnBy        string
	Adults          int32
	Children        int32
	Infants         int32
	Seniors         int32
	TravelClass     *TravelClass
	IncludeAirlines string
	ExcludeAirlines string
	NonStop         bool
	Currency        string
	MaxPrice        int64
	Max             int32
}

type FlightInspirationSearchRequest struct {
	Origin string

	DepartureDate string
	OneWay        bool
	Duration      string
	NonStop       bool
	MaxPrice      int64
	Currency      string
}

type FlightCheapestDateSearchRequest struct {
	Origin      string
	Destination string

	DepartureDate string
	OneWay        bool
	Duration      string
	NonStop       bool
	MaxPrice      int64
	Currency      string
	ViewBy        *ViewBy
}

type FlightMostSearchedDestinationsRequest struct {
	OriginCityCode    string
	SearchPeriod      string
	MarketCountryCode string

	Max        int32
	Fields     string
	PageLimit  int32
	PageOffset int32
}

type FlightMostSearchedByDestinationRequest struct {
	OriginCityCode      string
	DestinationCityCode string
	SearchPeriod        string
	MarketCountryCode   string

	Fields string
}

type FlightCheckInLinksRequest struct {
	AirlineCode string
	Language    string
}

type FlightMostTraveledDestinationsRequest struct {
	OriginCityCode string
	Period         string

	Max        int32
	Fields     string
	PageLimit  int32
	PageOffset int32
	Sort       *Sort
}

type FlightMostBookedDestinationsRequest struct {
	OriginCityCode string
	Period         string

	Max        int32
	Fields     string
	PageLimit  int32
	PageOffset int32
	Sort       *Sort
}

type FlightBusiestTravelingPeriodRequest struct {
	CityCode  string
	Period    string
	Direction *Direction
}

type AirportNearestRelevantRequest struct {
	Latitude   float32
	Longitude  float32

	Radius     int32
	PageLimit  int32
	PageOffset int32
	Sort       *RelevantSort
}

type AirportAndCitySearchRequest struct {
	SubType     string
	Keyword     string

	CountryCode string
	PageLimit int32
	PageOffset int32
	View *View
}

type AirlineCodeLookupRequest struct {
	AirlineCodes string
}

// ============================== Data Structures ==============================
type Data struct {
	Type           string                  `json:"type"`
	Id             string                  `json:"id"`
	OfferItems     []*OfferItem            `json:"offerItems"`
	Destination    string                  `json:"destination"`
	SubType        string                  `json:"subType"`
	Analytics      *Analytics              `json:"analytics"`
	Period         string                  `json:"period"`
	Name           string                  `json:"name"`
	DetailedName   string                  `json:"detailedName"`
	TimeZoneOffset string                  `json:"timeZoneOffset"`
	IataCode       string                  `json:"iataCode"`
	GeoCode        *GeoCode                `json:"geoCode"`
	Address        *Address                `json:"address"`
	Distance       *Distance               `json:"distance"`
	Relevance      float32                 `json:"relevance"`
	Origin         string                  `json:"origin"`
	DepartureDate  string                  `json:"departureDate"`
	ReturnDate     string                  `json:"returnDate"`
	Price          *Price                  `json:"price"`
	Links          *Links                  `json:"links"`
	Self           *Self                   `json:"links"`
	Href           string                  `json:"href"`
	Channel        string                  `json:"channel"`
	Parameters     map[string]*ParamDetail `json:"parameters"`
	IcaoCode       string                  `json:"icaoCode"`
	BusinessName   string                  `json:"businessName"`
	CommonName     string                  `json:"commonName"`
}

type OfferItem struct {
	Services      []*Service `json:"services"`
	Price         *Price     `json:"price"`
	PricePerAdult *Price     `json:"pricePerAdult"`
}

type Service struct {
	Segments []*Segment `json:"segments"`
}

type Segment struct {
	FlightSegment         *FlightSegment         `json:"flightSegment"`
	PricingDetailPerAdult *PricingDetailPerAdult `json:"pricingDetailPerAdult"`
}

type FlightSegment struct {
	Departure   *DepartureArrival `json:"departure"`
	Arrival     *DepartureArrival `json:"arrival"`
	CarrierCode string            `json:"carrierCode"`
	Number      string            `json:"number"`
	Aircraft    *Aircraft         `json:"aircraft"`
	Operating   *Operating        `json:"operating"`
	Duration    string            `json:"duration"`
}

type DepartureArrival struct {
	IataCode string `json:"iatacode"`
	Terminal string `json:"terminal"`
	At       string `json:"at"`
}

type Aircraft struct {
	Code string `json:"code"`
}

type Operating struct {
	CarrierCode string `json:"carrierCode"`
	Number      string `json:"number"`
}

type PricingDetailPerAdult struct {
	TravelClass  string `json:"travelClass"`
	FareClass    string `json:"fareClass"`
	Availability int32  `json:"availability"`
	FareBasis    string `json:"fareBasis"`
}

type Price struct {
	Total      string `json:"total"`
	TotalTaxes string `json:"totalTaxes"`
}

type Analytics struct {
	Flights   *Score `json:"flights"`
	Travelers *Score `json:"travelers"`
	Searches  *Score `json:"searches,omitemtpy"`
}

type Score struct {
	Score            int32             `json:"score"`
	NumberOfSearches *NumberOfSearches `json:"numberOfSearches"`
}

type GeoCode struct {
	Latitude  float32 `json:"Latitude"`
	Longitude float32 `json:"Longitude"`
}

type Address struct {
	CityName    string `json:"cityName"`
	CityCode    string `json:"cityCode"`
	CountryName string `json:"countryName"`
	CountryCode string `json:"countryCode"`
	StateCode   string `json:"stateCode"`
	RegionCode  string `json:"regionCode"`
}

type Distance struct {
	Value int32  `json:"value"`
	Unit  string `json:"unit"`
}

type Self struct {
	Href    string   `json:"href"`
	Methods []string `json:"methods"`
}

type Dictionaries struct {
	Carriers   map[string]string            `json:"carriers"`
	Currencies map[string]string            `json:"currencies"`
	Aircrafts  map[string]string            `json:"aircraft"`
	Locations  map[string]map[string]string `json:"locations"`
}

type Meta struct {
	Links    *Links    `json:"links"`
	Currency string    `json:"currency"`
	Defaults *Defaults `json:"defaults"`
	Count    int32     `json:"count"`
}

type Links struct {
	Self               string `json:"self"`
	Next               string `json:"next"`
	Last               string `json:"last"`
	FlightDates        string `json:"flightDates"`
	FlightOffers       string `json:"flightOffers"`
	FlightDestinations string `json:"flightDestinations"`
}

type Defaults struct {
	NonStop       bool   `json:"nonStop"`
	Adults        int32  `json:"adults"`
	DepartureDate string `json:"departureDate"`
	OneWay        bool   `json:"oneWay"`
	Duration      string `json:"duration"`
	ViewBy        string `json:"viewBy"`
}

type ErrorWarning struct {
	Status int32   `json:"status"`
	Code   int32   `json:"code"`
	Title  string  `json:"title"`
	Detail string  `json:"title"`
	Source *Source `json:"source"`
}

type Source struct {
	Pointer   string `json:"pointer"`
	Parameter string `json:"parameter"`
	Example   string `json:"example"`
}

type NumberOfSearches struct {
	PerTripDuration  map[string]string `json:"perTripDuration"`
	PerDaysInAdvance map[string]string `json:"perDaysInAdvance"`
}

type ParamDetail struct {
	Description string `json:"description"`
	Type        string `json:"type"`
	Format      string `json:"format"`
}

type TravelClass int

const (
	TravelClass_ECONOMY TravelClass = iota
	TravelClass_PREMIUM_ECONOMY
	TravelClass_BUSINESS
	TravelClass_FIRST
)

func (t TravelClass) String() string {
	return [...]string{"ECONOMY", "PREMIUM_ECONOMY", "BUSINESS", "FIRST"}[t]
}

type ViewBy int

const (
	ViewBy_DATE ViewBy = iota
	ViewBy_DURATION
	ViewBy_WEEK
)

func (v ViewBy) String() string {
	return [...]string{"DATE", "DURATION", "WEEK"}[v]
}

type Sort int

const (
	Sort_FLIGHTS Sort = iota
	Sort_TRAVELERS
)

func (v Sort) String() string {
	return [...]string{"analytics.flights.score", "analytics.travelers.score"}[v]
}

type Direction int

const (
	Direction_ARRIVING Direction = iota
	Direction_DEPARTING
)

func (v Direction) String() string {
	return [...]string{"ARRIVING", "DEPARTING"}[v]
}

type RelevantSort int

const (
	RelevantSort_RELEVANCE RelevantSort = iota
	RelevantSort_DISTANCE
	RelevantSort_FLIGHTS
	RelevantSort_TRAVELERS
)

func (v RelevantSort) String() string {
	return [...]string{
		"relevance",
		"distance",
		"analytics.flights.score",
		"analytics.travelers.score",
	}[v]
}

type View int

const (
	View_LIGHT View = iota
	View_FULL
)

func (v View) String() string {
	return [...]string{
		"LIGHT",
		"FULL",
	}[v]
}

