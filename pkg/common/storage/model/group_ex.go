package model

type WhitelistItem struct {
	Type       string   `json:"type"`
	UserIdList []string `json:"userIdList"`
}

type GroupEx struct {
	Capacity     int             `json:"capacity"`
	RemainingCap int             `json:"remainingCap"`
	Blacklist    []string        `json:"blacklist"`
	MaxAdminNum  int             `json:"maxAdminNum"`
	Whitelist    []WhitelistItem `json:"whitelist"`
}

func NewGroupEx() *GroupEx {
	return &GroupEx{
		Capacity:     200,
		RemainingCap: 220,
		Blacklist:    []string{},
		MaxAdminNum:  2,
		Whitelist:    []WhitelistItem{},
	}
}
