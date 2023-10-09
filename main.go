package main

import (
	"encoding/json"
	"fmt"
	"log"
	"net/http"
	"net/url"
)

func main() {
	http.HandleFunc("/api", Handler)
	http.ListenAndServe(":8080", nil)
}

func Handler(w http.ResponseWriter, r *http.Request) {
	fmt.Println("hit the handler")

	q := r.URL.Query().Get("q")

	data, err := getData(q)
	if err != nil {
		log.Printf("error calling data source : %v\n", err)
		return
	}

	resp := apiResponse{
		Cache: false,
		Data:  data,
	}
	err = json.NewEncoder(w).Encode(resp)
	if err != nil {
		log.Printf("error encoding json : %v\n", err)
		w.WriteHeader(http.StatusInternalServerError)
		return
	}

}
func getData(q string) ([]NominatimResponse, error) {

	// is query cached

	// San Francisco
	// to emit space we use %20 like i,e: San%20Francisco
	escapedQuery := url.PathEscape(q)
	address := fmt.Sprintf("https://nominatim.openstreetmap.org/search?q=%s&format=json", escapedQuery)

	res, err := http.Get(address)
	if err != nil {
		return nil, err
	}

	data := make([]NominatimResponse, 0)

	err = json.NewDecoder(res.Body).Decode(&data)
	if err != nil {
		return nil, err
	}

	return data, nil
}

type apiResponse struct {
	Cache bool                `json:"cache"`
	Data  []NominatimResponse `json:"data"`
}

type NominatimResponse struct {
	PlaceID     int      `json:"place_id"`
	Licence     string   `json:"licence"`
	OsmType     string   `json:"osm_type"`
	OsmID       int      `json:"osm_id"`
	Boundingbox []string `json:"boundingbox"`
	Latitude    string   `json:"lat"`
	Longitude   string   `json:"lon"`
	DisplayName string   `json:"display_name"`
	Class       string   `json:"class"`
	Type        string   `json:"type"`
	Importance  float64  `json:"importance"`
	Icon        string   `json:"icon"`
}
