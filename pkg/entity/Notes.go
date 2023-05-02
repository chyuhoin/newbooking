package entity

type Notes struct {
	Id          int    `json:"id,omitempty"`
	NoteContent string `json:"note,omitempty" xorm:"notes"`
}
