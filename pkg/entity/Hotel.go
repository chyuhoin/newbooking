package entity

type Hotel struct {
	Id         int     `json:"id"`
	Name       string  `json:"name"`
	Province   string  `json:"province"`
	City       string  `json:"city"`
	Location   string  `json:"location"`
	GreenLevel int     `json:"greenLevel"`
	Comment    int     `json:"comment"`
	Rating     float64 `json:"rating"`
	Star       int8    `json:"star"`
	Distance   float64 `json:"distance"`
}
