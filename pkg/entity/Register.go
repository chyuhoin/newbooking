package entity

type Register struct {
	Id             string `xorm:"pk id" json:"id"`
	RoomId         string `xorm:"room_id" json:"roomId"`
	StartTime      string `xorm:"start_time" json:"startTime"`
	EndTime        string `xorm:"end_time" json:"endTime"`
	UserId         string `xorm:"user_id" json:"userId"`
	BookerName     string `xorm:"b_name" json:"bookerName"`
	BookerEmail    string `xorm:"b_email" json:"bookerEmail"`
	BookerState    string `xorm:"b_state" json:"bookerState"`
	BookerPhone    string `xorm:"b_phone" json:"bookerPhone"`
	BookerBirthday string `xorm:"b_birthday" json:"bookerBirthday"`
	RoomerName     string `xorm:"r_name" json:"roomerName"`
	RoomerEmail    string `xorm:"r_email" json:"roomerEmail"`
	Plan           string `xorm:"plan" json:"plan"`
	IsDeleted      bool   `xorm:"is_deleted" json:"isDeleted"`
}
