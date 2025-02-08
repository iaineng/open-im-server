package model

type GroupEx struct {
	Capacity     int `json:"capacity"`
	RemainingCap int `json:"remaining_cap"`
}

func NewGroupEx() *GroupEx {
	return &GroupEx{
		Capacity:     200,
		RemainingCap: 220,
	}
}
