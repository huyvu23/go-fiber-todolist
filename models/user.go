package models

// ? => json kia có nghĩa là gì ?
// => Nó nó nghĩa là nó sẽ map các field như Username trả về cho FE là username
type User struct {
	ID       int    `json:"id"`
	Username string `json:"username"`
	Password string `json:"password"`
}
