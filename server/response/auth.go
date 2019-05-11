package response

type Auth struct {
	Token string `json:"token"`
}

type User struct {
	ID interface{} `json:"id"`
}
