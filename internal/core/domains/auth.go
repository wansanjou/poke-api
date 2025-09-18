package domains

type LoginRequest struct {
	Username string `bson:"username"`
	Password string `bson:"password"`
}

type LoginResponse struct {
	Token string `json:"token"`
}

type RegisterRequest struct {
	Username string `json:"username"`
	Password string `json:"password"`
}
