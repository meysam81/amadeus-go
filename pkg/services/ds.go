package services

// ======================= Data Structures =======================
type FlightLowFareSearchRequest struct {
	Origin        string
	Destination   string
	DepartureDate string
	ReturnDate    string
}

type FlightLowFareSearchResponse struct {
	Data         []*Data       `json:"data"`
	Dictionaries *Dictionaries `json:"dictionaries"`
	Meta         *Meta         `json:"meta"`
}

type FlightInspirationSearchRequest struct {
	Origin   string
	MaxPrice int32
}

type FlightInspirationSearchResponse struct {
	InspirationData []*InspirationData `json:"data"`
	Dictionaries    *Dictionaries      `json:"dictionaries"`
	Meta            *Meta              `json:"meta"`
}

type FlightMostTraveledDestinationsRequest struct {
	OriginCityCode string
	Period         string
}

type FlightMostTraveledDestinationsResponse struct {
	MostTraveledData []*MostTraveledData `json:"data"`
	Meta             *Meta               `json:"meta"`
}

// ===============================================================
type Data struct {
	Type       string       `json:"type"`
	Id         string       `json:"id"`
	OfferItems []*OfferItem `json:"offerItems"`
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

// ===============================================================
type Dictionaries struct {
	Carriers   map[string]string            `json:"carriers"`
	Currencies map[string]string            `json:"currencies"`
	Aircrafts  map[string]string            `json:"aircraft"`
	Locations  map[string]map[string]string `json:"locations"`
}

// ===============================================================
type Meta struct {
	Links    *Links    `json:"links"`
	Currency string    `json:"currency"`
	Defaults *Defaults `json:"defaults"`
	Count    int32     `json:"count"`
}

type Links struct {
	Self string `json:"self"`
}

type Defaults struct {
	NonStop       bool   `json:"nonStop"`
	Adults        int32  `json:"adults"`
	DepartureDate string `json:"departureDate"`
	OneWay        bool   `json:"oneWay"`
	Duration      string `json:"duration"`
	ViewBy        string `json:"viewBy"`
}

// ===============================================================
type InspirationData struct {
	Type          string           `json:"type"`
	Origin        string           `json:"origin"`
	Destination   string           `json:"destination"`
	DepartureDate string           `json:"departureDate"`
	ReturnDate    string           `json:"returnDate"`
	Price         Price            `json:"price"`
	Links         InspirationLinks `json:"links"`
}

type InspirationLinks struct {
	FlightDates  string `json:"flightDates"`
	FlightOffers string `json:"flightOffers"`
}

// ===============================================================
type MostTraveledData struct {
	Type        string    `json:"type"`
	Destination string    `json:"destination"`
	SubType     string    `json:"subType"`
	Analytics   Analytics `json:"analytics"`
}

type Analytics struct {
	Flights   Score `json:"flights"`
	Travelers Score `json:"travelers"`
}

type Score struct {
	Score int32 `json:"score"`
}
