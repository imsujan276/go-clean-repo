package todocontrollers

type TodoInput struct {
	ID          uint   `gorm:"primary_key"`
	Title       string `gorm:"not null"`
	Description string `gorm:""`
	Completed   bool   `gorm:"not null"`
	UserId      uint   `gorm:"not null"`
}

type TodoStatusInput struct {
	ID        uint `gorm:"primary_key, not null"`
	Completed bool `gorm:"not null"`
	UserId    uint `gorm:""`
}
