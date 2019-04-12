package toggl

// User contailn all the information of a user
type User struct {
	ID      int  `json:"id"`
	Pid     int  `json:"pid"`
	UID     int  `json:"uid"`
	Wid     int  `json:"wid"`
	Manager bool `json:"manager"`
	Rate    int  `json:"rate"`
}
