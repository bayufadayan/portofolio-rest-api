package model

type PersonalInformation struct {
	ID           int    `json:"id" gorm:"primaryKey"`
	Name         string `json:"name"`
	Email        string `json:"email"`
	Phone        string `json:"phone"`
	Address      string `json:"address"`
	Description  string `json:"description"`
	ProfileImage string `json:"profile_image"`
	SocialMedia  int    `json:"social_media"`
	JobTitle     int    `json:"job_title"`
}
