package model

import "time"

type Book struct {
	ID            uint      `json:"id" gorm:"primary_key"`
	Title         string    `json:"title" gorm:"not null"`
	Author        string    `json:"author"`
	ISBN          string    `json:"isbn" gorm:"not null;unique"`
	Publisher     string    `json:"publisher"`
	PublishedDate time.Time `json:"published_date"`
	Description   string    `json:"description"`
	Genre         string    `json:"genre"`
	CoverImageURL string    `json:"cover_image_url"`
	CreatedAt     time.Time `json:"created_at"`
	UpdatedAt     time.Time `json:"updated_at"`
}