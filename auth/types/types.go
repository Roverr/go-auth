package authTypes

// RegisterRequest describes how should
// the request body look like for the
// registration endpoint
type RegisterRequest struct {
	RealName string `json:"realName"`
	UserName string `json:"userName"`
	Password string `json:"password"`
}
