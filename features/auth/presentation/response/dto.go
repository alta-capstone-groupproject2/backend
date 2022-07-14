package response

type user struct {
	ID    int    `json:"id"`
	Role  string `json:"role"`
	Token string `json:"token"`
}

func ToResponse(id int, role string, token string) user {
	return user{
		ID:    id,
		Role:  role,
		Token: token,
	}
}
