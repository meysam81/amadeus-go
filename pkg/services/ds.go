package services

// ======================= Data Structures =======================
type FlightLowFareSearchRequest struct {
	Origin        string
	Destination   string
	DepartureDate string
	ReturnDate    string
}

type FlightLowFareSearchResponse struct {
	Data         []*Data       `json:"data,omitempty"`
	Dictionaries *Dictionaries `json:"dictionaries,omitempty"`
	Meta         *Meta         `json:"meta,omitempty"`
}

type FlightInspirationSearchRequest struct {
	Origin   string
	MaxPrice int32
}

type FlightInspirationSearchResponse struct {
	Data         []*Data       `json:"data,omitempty"`
	Dictionaries *Dictionaries `json:"dictionaries,omitempty"`
	Meta         *Meta         `json:"meta,omitempty"`
}

type FlightMostTraveledDestinationsRequest struct {
	OriginCityCode string
	Period         string
}

type FlightMostTraveledDestinationsResponse struct {
	Data []*Data `json:"data,omitempty"`
	Meta *Meta   `json:"meta,omitempty"`
}

type FlightMostBookedDestinationsRequest struct {
	OriginCityCode string
	Period         string
}

type FlightMostBookedDestinationsResponse struct {
	Data []*Data `json:"data,omitempty"`
	Meta *Meta   `json:"meta,omitempty"`
}

type FlightBusiestTravelingPeriodRequest struct {
	CityCode  string
	Period    string
	Direction string
}

type FlightBusiestTravelingPeriodResponse struct {
	Data []*Data `json:"data,omitempty"`
	Meta *Meta   `json:"meta,omitempty"`
}

type AirportNearestRelevantRequest struct {
	Latitude  float32
	Longitude float32
	Sort      string
}

type AirportNearestRelevantResponse struct {
	Data []*Data `json:"data,omitempty"`
	Meta *Meta   `json:"meta,omitempty"`
}

// ===============================================================
type Data struct {
	Type           string       `json:"type,omitempty"`
	Id             string       `json:"id,omitempty"`
	OfferItems     []*OfferItem `json:"offerItems,omitempty"`
	Destination    string       `json:"destination,omitempty"`
	SubType        string       `json:"subType,omitempty"`
	Analytics      Analytics    `json:"analytics,omitempty"`
	Period         string       `json:"period,omitempty"`
	Name           string       `json:"name,omitempty"`
	DetailedName   string       `json:"detailedName,omitempty"`
	TimeZoneOffset string       `json:"timeZoneOffset,omitempty"`
	IataCode       string       `json:"iataCode,omitempty"`
	GeoCode        GeoCode      `json:"geoCode,omitempty"`
	Address        Address      `json:"address,omitempty"`
	Distance       Distance     `json:"distance,omitempty"`
	Relevance      float32      `json:"relevance,omitempty"`
	Origin         string       `json:"origin,omitempty"`
	DepartureDate  string       `json:"departureDate,omitempty"`
	ReturnDate     string       `json:"returnDate,omitempty"`
	Price          Price        `json:"price,omitempty"`
	Links          Links        `json:"links,omitempty"`
}

type OfferItem struct {
	Services      []*Service `json:"services,omitempty"`
	Price         *Price     `json:"price,omitempty"`
	PricePerAdult *Price     `json:"pricePerAdult,omitempty"`
}

type Service struct {
	Segments []*Segment `json:"segments,omitempty"`
}

type Segment struct {
	FlightSegment         *FlightSegment         `json:"flightSegment,omitempty"`
	PricingDetailPerAdult *PricingDetailPerAdult `json:"pricingDetailPerAdult,omitempty"`
}

type FlightSegment struct {
	Departure   *DepartureArrival `json:"departure,omitempty"`
	Arrival     *DepartureArrival `json:"arrival,omitempty"`
	CarrierCode string            `json:"carrierCode,omitempty"`
	Number      string            `json:"number,omitempty"`
	Aircraft    *Aircraft         `json:"aircraft,omitempty"`
	Operating   *Operating        `json:"operating,omitempty"`
	Duration    string            `json:"duration,omitempty"`
}

type DepartureArrival struct {
	IataCode string `json:"iatacode,omitempty"`
	Terminal string `json:"terminal,omitempty"`
	At       string `json:"at,omitempty"`
}

type Aircraft struct {
	Code string `json:"code,omitempty"`
}

type Operating struct {
	CarrierCode string `json:"carrierCode,omitempty"`
	Number      string `json:"number,omitempty"`
}

type PricingDetailPerAdult struct {
	TravelClass  string `json:"travelClass,omitempty"`
	FareClass    string `json:"fareClass,omitempty"`
	Availability int32  `json:"availability,omitempty"`
	FareBasis    string `json:"fareBasis,omitempty"`
}

type Price struct {
	Total      string `json:"total,omitempty"`
	TotalTaxes string `json:"totalTaxes,omitempty"`
}

type Analytics struct {
	Flights   Score `json:"flights,omitempty"`
	Travelers Score `json:"travelers,omitempty"`
}

type Score struct {
	Score int32 `json:"score,omitempty"`
}

type GeoCode struct {
	Latitude  float32 `json:"Latitude,omitempty"`
	Longitude float32 `json:"Longitude,omitempty"`
}

type Address struct {
	CityName    string `json:"cityName,omitempty"`
	CityCode    string `json:"cityCode,omitempty"`
	CountryName string `json:"countryName,omitempty"`
	CountryCode string `json:"countryCode,omitempty"`
	RegionCode  string `json:"regionCode,omitempty"`
}

type Distance struct {
	Value int32  `json:"value,omitempty"`
	Unit  string `json:"unit,omitempty"`
}

type Dictionaries struct {
	Carriers   map[string]string            `json:"carriers,omitempty"`
	Currencies map[string]string            `json:"currencies,omitempty"`
	Aircrafts  map[string]string            `json:"aircraft,omitempty"`
	Locations  map[string]map[string]string `json:"locations,omitempty"`
}

type Meta struct {
	Links    *Links    `json:"links,omitempty"`
	Currency string    `json:"currency,omitempty"`
	Defaults *Defaults `json:"defaults,omitempty"`
	Count    int32     `json:"count,omitempty"`
}

type Links struct {
	Self         string `json:"self,omitempty"`
	Next         string `json:"next,omitempty"`
	Last         string `json:"last,omitempty"`
	FlightDates  string `json:"flightDates,omitempty"`
	FlightOffers string `json:"flightOffers,omitempty"`
}

type Defaults struct {
	NonStop       bool   `json:"nonStop,omitempty"`
	Adults        int32  `json:"adults,omitempty"`
	DepartureDate string `json:"departureDate,omitempty"`
	OneWay        bool   `json:"oneWay,omitempty"`
	Duration      string `json:"duration,omitempty"`
	ViewBy        string `json:"viewBy,omitempty"`
}
