package main

import (
	"net/http"
)

func main() {
	http.HandleFunc("/test", Handler)
	http.ListenAndServe(":8080", nil)
}

func Handler(w http.ResponseWriter, r *http.Request) {}

func getData() (*nominatimData, bool, error) {
	res, err := http.Get("https://nominatim.openstreetmap.org/search?q")
	if err != nil {
		//TODO:
	}
	return nil, false, nil

}

type ResponseData struct {
	IsUsingCache bool            `json:"is-using-cache"`
	Data         []nominatimData `json:"data"`
}
type nominatimData struct {
	Address struct {
		ISO31662Lvl4  string `json:"ISO3166-2-lvl4"`
		Borough       string `json:"borough"`
		City          string `json:"city"`
		Country       string `json:"country"`
		CountryCode   string `json:"country_code"`
		Historic      string `json:"historic"`
		HouseNumber   string `json:"house_number"`
		Neighbourhood string `json:"neighbourhood"`
		Postcode      string `json:"postcode"`
		Road          string `json:"road"`
		Suburb        string `json:"suburb"`
	} `json:"address"`
	Boundingbox []string `json:"boundingbox"`
	Class       string   `json:"class"`
	DisplayName string   `json:"display_name"`
	Importance  float64  `json:"importance"`
	Lat         string   `json:"lat"`
	Licence     string   `json:"licence"`
	Lon         string   `json:"lon"`
	OsmID       int      `json:"osm_id"`
	OsmType     string   `json:"osm_type"`
	PlaceID     int      `json:"place_id"`
	Svg         string   `json:"svg"`
	Type        string   `json:"type"`
}
