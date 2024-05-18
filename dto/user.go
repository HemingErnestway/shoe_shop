package dto

type User struct {
	Id        uint32 `json:"id" gorm:"primaryKey"`
	Email     string `json:"email"`
	Password  string `json:"password"`
	BirthDate string `json:"birth_date"`
	Access    string `json:"access"`
}
