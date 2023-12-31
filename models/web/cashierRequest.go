package web

type CashierCreateRequest struct {
	AdminID  uint   `json:"adminId"`
	Fullname string `json:"fullname" validate:"required,min=1,max=255"`
	Username string `json:"username" validate:"required,alphanum,min=1"`
	Password string `json:"password" validate:"required,min=8"`
}

type CashierLoginRequest struct {
	Username string `json:"username" validate:"required,min=1"`
	Password string `json:"password" validate:"required,min=8"`
}

type CashierUpdateRequest struct {
	AdminID  uint   `json:"adminId"`
	Fullname string `json:"fullname" validate:"required,min=1,max=255"`
	Username string `json:"username" validate:"required,alphanum,min=1"`
	Password string `json:"password" validate:"required,min=8"`
}