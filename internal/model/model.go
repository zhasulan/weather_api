package model

/*
	{
		"status": "Failed",
		"error": "error message"
	}

	{
		"status": "City successfully saved",
		"id": 1
	}
*/
type Response struct {
	Status string `json:"status"`
	ID     int    `json:"id,omitempty"`
	IDs    []int  `json:"ids,omitempty"`
	Error  string `json:"error,omitempty"`
}

type WeatherStatus struct {
	Code    int    `json:"code"`
	Message string `json:"message"`
}

type WeatherResponse struct {
	CityName       string  `json:"city_name"`
	CelsiusTemp    float64 `json:"temp_c"`
	FahrenheitTemp float64 `json:"temp_f"`
	Humidity       int     `json:"humidity"`
	Condition      string  `json:"condition"`
}
