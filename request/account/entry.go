package account

type EntryRequest struct {
	Mail     string `form:"mail"`
	Password string `form:"password"`
	Name     string `form:"name"`
}
