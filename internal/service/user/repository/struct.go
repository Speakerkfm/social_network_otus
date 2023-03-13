package repository

type SocialUser struct {
	ID             string
	FirstName      string
	SecondName     string
	Age            int
	Sex            int
	City           string
	Biography      string
	HashedPassword string
}

type UserSession struct {
	ID     string
	UserID string
	Token  string
}
