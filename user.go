package anki

type User struct {
	// Id 		int    `json: "-"`
	Username       string `json: "username"`
	Email          string `json: "email"`
	Password       string `json: "password"`
	RepeatPassword string `json: "repeat-password"`
}
