package models

import (
	"time"

	"gorm.io/gorm"
)

type Reservation struct {
	Id          uint `gorm:"primaryKey"`
	Title       string
	TimeFrom    time.Time
	TimeTo      time.Time
	ResvBy      uint
	User        User `gorm:"foreignKey:ResvBy"`
	Description string
	Location    string
	Color       string
	CreatedAt   time.Time
}

func NewReservation(id uint, title string, timeFrom time.Time, timeTo time.Time, resvBy uint, color, description, location string) *Reservation {
	return &Reservation{
		Id:          id,
		Title:       title,
		TimeFrom:    timeFrom,
		TimeTo:      timeTo,
		ResvBy:      resvBy,
		Color:       color,
		Description: description,
		Location:    location,
		CreatedAt:   time.Now(),
	}
}

func CreateReservation(db *gorm.DB, r *Reservation) error {
	return db.Create(r).Error
}

func GetReservation(db *gorm.DB, id uint) (*Reservation, error) {
	r := &Reservation{}
	if err := db.First(r, id).Error; err != nil {
		return nil, err
	}
	return r, nil
}

func UpdateReservation(db *gorm.DB, r *Reservation) error {
	return db.Save(r).Error
}

func DeleteReservation(db *gorm.DB, id uint) error {
	return db.Delete(&Reservation{}, id).Error
}
