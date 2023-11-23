package model

type User struct {
	Id         int64  `json:"id"`
	First_name string `json:"first_name"`
	Last_name  string `json:"last_name"`
	Email      string `json:"email"`
	Password   string `json:"password"`
	Is_admin   string `json:"is_admin"`
	Created_at string `json:"created_at"`
	Updated_at string `json:"updated_at"`
}
