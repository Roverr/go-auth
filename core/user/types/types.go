package userTypes

// ClientObject describes the structure of
// the public data about the user which can be
// sent back to the clientside
type ClientObject struct {
	RealName string `json:"realName"`
	UserName string `json:"userName"`
	ID       uint   `json:"id"`
}
