package models

// type User struct {
// 	UserId       uint64    `json:"userId"`
// 	UserName     string    `json:"userName"`
// 	UserEmail    string    `json:"userEmail"`
// 	UserPassword string    `json:"userPassword"`
// 	CreatedAt    time.Time `json:"createdAt"`
// }

type User struct {
	UserId       uint64 `json:"userId"`
	UserName     string `json:"userName"`
	UserEmail    string `json:"userEmail"`
	UserPassword string `json:"userPassword"`
	CreatedAt    string `json:"createdAt"`
}

type UserResponse struct {
	UserId    uint64 `json:"userId"`
	UserName  string `json:"userName"`
	UserEmail string `json:"userEmail"`
	CreatedAt string `json:"createdAt"`
}
