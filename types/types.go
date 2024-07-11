package types

type Station struct {
	ID         string `json:"id"`
	Name       string `json:"name"`
	Score      any    `json:"score"`
	Coordinate struct {
		Type string  `json:"type"`
		X    float64 `json:"x"`
		Y    float64 `json:"y"`
	} `json:"coordinate"`
	Distance any `json:"distance"`
}

type Coordinate struct {
	Type string  `json:"type"`
	X    float64 `json:"x"`
	Y    float64 `json:"y"`
}
type Point struct {
	Station            Station  `json:"station"`
	Arrival            Datetime `json:"arrival"`
	ArrivalTimestamp   any      `json:"arrivalTimestamp"`
	Departure          Datetime `json:"departure"`
	DepartureTimestamp int      `json:"departureTimestamp"`
	Delay              int      `json:"delay"`
	Platform           string   `json:"platform"`
	Prognosis          struct {
		Platform    any    `json:"platform"`
		Arrival     any    `json:"arrival"`
		Departure   string `json:"departure"`
		Capacity1St any    `json:"capacity1st"`
		Capacity2Nd any    `json:"capacity2nd"`
	} `json:"prognosis"`
	RealtimeAvailability any `json:"realtimeAvailability"`
	Location             struct {
		ID         string     `json:"id"`
		Name       string     `json:"name"`
		Score      any        `json:"score"`
		Coordinate Coordinate `json:"coordinate"`
		Distance   any        `json:"distance"`
	} `json:"location"`
}

type Horaires struct {
	Connections []struct {
		From        Point    `json:"from"`
		To          Point    `json:"to"`
		Duration    string   `json:"duration"`
		Transfers   int      `json:"transfers"`
		Service     any      `json:"service"`
		Products    []string `json:"products"`
		Capacity1St any      `json:"capacity1st"`
		Capacity2Nd any      `json:"capacity2nd"`
		Sections    []struct {
			Journey struct {
				Name         string  `json:"name"`
				Category     string  `json:"category"`
				Subcategory  any     `json:"subcategory"`
				CategoryCode any     `json:"categoryCode"`
				Number       string  `json:"number"`
				Operator     string  `json:"operator"`
				To           string  `json:"to"`
				PassList     []Point `json:"passList"`
				Capacity1St  any     `json:"capacity1st"`
				Capacity2Nd  any     `json:"capacity2nd"`
			} `json:"journey"`
			Walk      any   `json:"walk"`
			Departure Point `json:"departure"`
			Arrival   Point `json:"arrival"`
		} `json:"sections"`
	} `json:"connections"`
	From struct {
		ID         string     `json:"id"`
		Name       string     `json:"name"`
		Score      any        `json:"score"`
		Coordinate Coordinate `json:"coordinate"`
		Distance   any        `json:"distance"`
	} `json:"from"`
	To struct {
		ID         string     `json:"id"`
		Name       string     `json:"name"`
		Score      any        `json:"score"`
		Coordinate Coordinate `json:"coordinate"`
		Distance   any        `json:"distance"`
	} `json:"to"`
	Stations struct {
		From []struct {
			ID         string     `json:"id"`
			Name       string     `json:"name"`
			Score      any        `json:"score"`
			Coordinate Coordinate `json:"coordinate"`
			Distance   any        `json:"distance"`
		} `json:"from"`
		To []struct {
			ID         string     `json:"id"`
			Name       string     `json:"name"`
			Score      any        `json:"score"`
			Coordinate Coordinate `json:"coordinate"`
			Distance   any        `json:"distance"`
		} `json:"to"`
	} `json:"stations"`
}

func (station Station) ToString() string {
	return station.Name
}
