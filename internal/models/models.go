package models

type RegisterRequest struct {
	Email    string
	Password string
}

type Login struct {
	Apikey   string
	Email    string
	Password string
}

type OutputAuth struct {
	Username        string
	Email           string
	ProfilePhotoUrl string
	Role            string
}
