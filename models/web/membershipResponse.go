package web

import uuid "github.com/satori/go.uuid"

// Membership Response
type MembershipResponse struct {
	ID           uint      `json:"id"`
	CashierID    uint      `json:"CashierID"`
	Name         string    `json:"name"`
	Code_Member  uuid.UUID `json:"Code_Member"`
	Point        uint      `json:"point"`
	Phone_Number string    `json:"phone_number"`
	Barcode      string    `json:"barcode"`
}
