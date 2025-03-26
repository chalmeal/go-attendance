package domain

type Seat struct {
	// 男性の席
	MaleSeat []MaleSeat `json:"male_seat"`
	// 女性の席
	FemaleSeat []FemaleSeat `json:"female_seat"`
}

type MaleSeat struct {
	// 男性の席番号
	MaleSeatNumber int `json:"male_seat_number"`
	// 男性の名前
	MaleName string `json:"male_name"`
}

type FemaleSeat struct {
	// 女性の席番号
	FemaleSeatNumber int `json:"female_seat_number"`
	// 女性の名前
	FemaleName string `json:"female_name"`
}

// Seatが空かどうかを判定する
func (s Seat) IsEmpty() bool {
	return len(s.MaleSeat) == 0 && len(s.FemaleSeat) == 0
}
