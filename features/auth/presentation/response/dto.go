package response

type user struct {
	ID    int    `json:"id" form:"id"`
	Token string `json:"token" form:"token"`
}

func ToResponse(id int, token string) user {
	return user{
		ID:    id,
		Token: token,
	}
}
