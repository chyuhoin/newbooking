package entity

type Register struct {
	Id             string `xorm:"id" json:"id,omitempty"`
	RoomId         string `xorm:"room_id" json:"roomId,omitempty"`
	StartTime      string `xorm:"start_time" json:"startTime,omitempty"`
	EndTime        string `xorm:"end_time" json:"endTime,omitempty"`
	UserId         string `xorm:"user_id" json:"userId,omitempty"`
	BookerName     string `xorm:"b_name" json:"bookerName,omitempty"`
	BookerEmail    string `xorm:"b_email" json:"bookerEmail,omitempty"`
	BookerState    string `xorm:"b_state" json:"bookerState,omitempty"`
	BookerPhone    string `xorm:"b_phone" json:"bookerPhone,omitempty"`
	BookerBirthday string `xorm:"b_birthday" json:"bookerBirthday,omitempty"`
	RoomerName     string `xorm:"r_name" json:"roomerName,omitempty"`
	RoomerEmail    string `xorm:"r_email" json:"roomerEmail,omitempty"`
	Plan           string `xorm:"plan" json:"plan,omitempty"`
	IsDeleted      string `xorm:"is_deleted" json:"isDeleted,omitempty"`
}
