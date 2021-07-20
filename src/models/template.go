package models

// User is user models property
type Template struct {
  ID     uint   `json:"id" binding:"required"`
  Title  string `json:"title" binding:"required"`
  User   User   `json:"-" binding:"required"`
  UserID uint   `gorm:"not null" json:"user_id"`
}
