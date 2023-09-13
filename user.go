package anki

type User struct {
	Username       string `json: "Username"`
	Email          string `json: "Email"`
	Password       string `json: "Password"`
	RepeatPassword string `json: "RepeatPassword"`
}

type FinalUser struct {
	Username string `json: "username"`
	Email    string `json: "email"`
	Password string `json: "password"`
}
