package out

type UserOut struct {
	Total int64       `json:"total"`
	Data  interface{} `json:"data"`
}

type Users struct {
	Id        int64  `json:"id"`
	Username  string `json:"username"`
	Nickname  string `json:"nickname"`
	Phone     string `json:"phone"`
	Email     string `json:"email"`
	Gender    int8   `json:"gender"`
	RoleId    string `json:"role_id"`
	Enable    int8   `json:"enable"`
	CreatedAt string `json:"created_at"`
	UpdatedAt string `json:"updated_at"`
}
