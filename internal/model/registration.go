package model

type Registration struct {
	Id      int64  `gorm:"primaryKey;autoIncrement" `
	UserId  int64  `gorm:"not null;uniqueIndex:idx_user_event"`
	User    *User  `gorm:"foreignKey:UserId"`
	EventId int    `gorm:"not null;uniqueIndex:idx_user_event"`
	Event   *Event `gorm:"foreignKey:EventId"`
}
