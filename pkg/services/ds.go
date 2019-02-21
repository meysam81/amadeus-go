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
	Aircraft   map[string]string            `json:"aircraft"`
	Locations  map[string]map[string]string `json:"locations"`
}

// ===============================================================
type Meta struct {
	Links    *Links    `json:"links"`
	Currency string    `json:"currency"`
	Defaults *Defaults `json:"defaults"`
}

type Links struct {
	Self string `json:"self"`
}

type Defaults struct {
	NonStop bool  `json:"nonStop"`
	Adults  int32 `json:"adults"`
}
