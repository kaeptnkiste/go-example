package main

import (
	"encoding/json"
	"html/template"
	"log"
	"net/http"
	"time"
)

type Ubahn []struct {
	TripID string `json:"tripId"`
	Stop   struct {
		Type     string `json:"type"`
		ID       string `json:"id"`
		Name     string `json:"name"`
		Location struct {
			Type      string  `json:"type"`
			ID        string  `json:"id"`
			Latitude  float64 `json:"latitude"`
			Longitude float64 `json:"longitude"`
		} `json:"location"`
		Products struct {
			Suburban bool `json:"suburban"`
			Subway   bool `json:"subway"`
			Tram     bool `json:"tram"`
			Bus      bool `json:"bus"`
			Ferry    bool `json:"ferry"`
			Express  bool `json:"express"`
			Regional bool `json:"regional"`
		} `json:"products"`
		StationDHID string `json:"stationDHID"`
	} `json:"stop"`
	When            time.Time   `json:"when"`
	PlannedWhen     time.Time   `json:"plannedWhen"`
	Delay           int         `json:"delay"`
	Platform        interface{} `json:"platform"`
	PlannedPlatform interface{} `json:"plannedPlatform"`
	PrognosisType   string      `json:"prognosisType"`
	Direction       string      `json:"direction"`
	Provenance      interface{} `json:"provenance"`
	Line            struct {
		Type        string `json:"type"`
		ID          string `json:"id"`
		FahrtNr     string `json:"fahrtNr"`
		Name        string `json:"name"`
		Public      bool   `json:"public"`
		AdminCode   string `json:"adminCode"`
		ProductName string `json:"productName"`
		Mode        string `json:"mode"`
		Product     string `json:"product"`
		Operator    struct {
			Type string `json:"type"`
			ID   string `json:"id"`
			Name string `json:"name"`
		} `json:"operator"`
		Symbol  interface{} `json:"symbol"`
		Nr      int         `json:"nr"`
		Metro   bool        `json:"metro"`
		Express bool        `json:"express"`
		Night   bool        `json:"night"`
	} `json:"line"`
	Remarks     []interface{} `json:"remarks"`
	Origin      interface{}   `json:"origin"`
	Destination struct {
		Type     string `json:"type"`
		ID       string `json:"id"`
		Name     string `json:"name"`
		Location struct {
			Type      string  `json:"type"`
			ID        string  `json:"id"`
			Latitude  float64 `json:"latitude"`
			Longitude float64 `json:"longitude"`
		} `json:"location"`
		Products struct {
			Suburban bool `json:"suburban"`
			Subway   bool `json:"subway"`
			Tram     bool `json:"tram"`
			Bus      bool `json:"bus"`
			Ferry    bool `json:"ferry"`
			Express  bool `json:"express"`
			Regional bool `json:"regional"`
		} `json:"products"`
		StationDHID string `json:"stationDHID"`
	} `json:"destination"`
	CurrentTripPosition struct {
		Type      string  `json:"type"`
		Latitude  float64 `json:"latitude"`
		Longitude float64 `json:"longitude"`
	} `json:"currentTripPosition,omitempty"`
	Occupancy string `json:"occupancy"`
}

func main() {
	http.HandleFunc("/", departureHandler)
	log.Fatal(http.ListenAndServe(":8080", nil))
}

func departureHandler(w http.ResponseWriter, r *http.Request) {
	// BVG API URL
	apiURL := "https://v5.bvg.transport.rest/stops/900000043101/departures?duration=20&linesOfStops=false&remarks=true&language=en"

	// HTTP-Anfrage an die BVG API senden
	resp, err := http.Get(apiURL)
	if err != nil {
		log.Fatal("Fehler bei der Anfrage an die BVG API:", err)
	}
	defer resp.Body.Close()

	// JSON-Daten aus der API-Antwort lesen
	var departures Ubahn
	err = json.NewDecoder(resp.Body).Decode(&departures)
	if err != nil {
		log.Fatal("Fehler beim Parsen der JSON-Daten:", err)
	}

	// HTML-Template erstellen
	tmpl := template.Must(template.ParseFiles("template.html"))

	// Daten an das Template übergeben und in den HTTP-Response schreiben
	err = tmpl.Execute(w, departures)
	if err != nil {
		log.Fatal("Fehler beim Ausführen des Templates:", err)
	}
}
