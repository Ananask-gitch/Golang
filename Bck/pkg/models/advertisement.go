package models

type Advertisement struct {
	ID      uint    `gorm:"primaryKey"`
	Name    string  `gorm:"type:varchar(200); not null"`
	Comment string  `gorm:"varchar(1000);not null"`
	Photos  []Photo `gorm:"foreignKey:"AdvertisementID; constraint:OnDelete:CASCADE;not null"`
	Price   uint    `gorm"type:uint;not null"`
}

type Photo struct {
	AdvertisementID uint   `gorm:"primaryKey"`
	PhotoMain       string `gorm:"not null"`
	PhotoSecond     string `gorm:""`
	PhotoSecond2    string `gorm:""`
}
