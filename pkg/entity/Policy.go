package entity

type Policy struct {
	Id             int    `json:"id,omitempty"`
	CheckinPolicy  string `json:"checkinPolicy,omitempty"`
	CheckoutPolicy string `json:"checkoutPolicy,omitempty"`
	ChildrenPolicy string `json:"childrenPolicy,omitempty"`
	BedPolicy      string `json:"bedPolicy,omitempty"`
	AgePolicy      string `json:"agePolicy,omitempty"`
	Payment        string `json:"payment,omitempty"`
	CheckinStamp   string `json:"checkinStamp,omitempty"`
	CheckoutStamp  string `json:"checkoutStamp,omitempty"`
}
