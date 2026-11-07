package pokeapi

type LocationAreasResp struct {
	Count    int     `json:"count"`
	Next     *string `json:"next"`     //	The best way to represent a nullable string is defining it as a pointer to a string, because when it's null, the pointer will be nil
	Previous *string `json:"previous"` //	The best way to represent a nullable string is defining it as a pointer to a string, because when it's null, the pointer will be nil
	Results  []struct {
		Name string `json:"name"`
		URL  string `json:"url"`
	} `json:"results"`
}
