package model

type UserData struct {
	Token             string `json:"token"`
	ID                uint `json:"id"`
	FName             string `json:"first_name"`
	LName             string `json:"last_name"`
	FullName          string `json:"full_name"`
	Email             string `json:"email"`
	Password          string `json:"password"`
	ProfilePictureUrl string `json:"profile_picture_url"`
	Phone             string `json:"phone"`
}

type LoginData struct {
	Email    string `json:"email"`
	Password string `json:"password"`
}

type UpdateProfilePicture struct {
	Url string `json:"url"`
}
