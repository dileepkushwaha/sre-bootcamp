package model

type Student struct {
	ID    uint   `gorm:"primaryKey"`
	Name  string `gorm:"size:255"`
	Age   int
	Email string `gorm:"size:255;unique"`
}

