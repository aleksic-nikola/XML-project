package data

type PaymentOption uint

type Payment struct {
	ID 			  uint          `json:"id" gorm:"primaryKey"`
	PaymentOption PaymentOption `json:"payment_option" gorm:"type:int"`
}

const (
	ONPICKUP PaymentOption = iota
)