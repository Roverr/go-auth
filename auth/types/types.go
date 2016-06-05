package authTypes

// RegisterRequest describes how should
// the request body look like for the
// registration endpoint
type RegisterRequest struct {
	RealName string `json:"realName"`
	UserName string `json:"userName"`
	Password string `json:"password"`
}

// ParsedToken is used to describe the JWT token
// information which is parsed before accessing endpoints
// which need authentication first
type ParsedToken struct {
	ID       uint
	UserName string
}
