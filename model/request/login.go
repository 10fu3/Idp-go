package request

type LoginRequest struct {
	Mail string `form:"mail"`
	Password string `form:"password"`
}
