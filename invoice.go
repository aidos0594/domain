package domain

import (
	stdjwt "github.com/dgrijalva/jwt-go"
	"time"
)

type Invoice struct {
	InvoiceId          string         `json:"invoiceId,omitempty"`
	Items              []InvoiceItems `json:"items,omitempty"`
	OrderRef           string         `json:"orderRef,omitempty"`
	State              string         `json:"state,omitempty"`
	PayAmount          float64        `json:"payAmount,omitempty"`
	PaymentOption      PaymentOption  `json:"paymentOption,omitempty"`
	InvoiceCreatedDate string         `json:"invoiceCreatedDate,omitempty"`
	InvoicePaidDate    string         `json:"invoicePaidDate,omitempty"`
	Jwt                string         `json:"jwt,omitempty"`
	Amount             float64        `json:"amount,omitempty"`
	TerminalId         string         `json:"terminalId,omitempty"`
	InstallmentPeriod  int32          `json:"installmentPeriod,omitempty"`
}

type CompleteInvoice struct {
	JsonClass  string  `json:"jsonClass,omitempty"`
	InvoiceId  string  `json:"invoiceId,omitempty"`
	TerminalId string  `json:"terminalId,omitempty"`
	Command    string  `json:"cammand,omitempty"`
	Amount     float64 `json:"amount,omitempty"`
	ReplyTo    string  `json:"replyTo,omitempty"`
}

type PayInvoice struct {
	JsonClass       string        `json:"jsonClass,omitempty"`
	InvoiceId       string        `json:"invoiceId,omitempty"`
	MerchantId      string        `json:"merchantId,omitempty"`
	TerminalId      string        `json:"terminalId,omitempty"`
	PaymentOption   PaymentOption `json:"paymentOption,omitempty"`
	Amount          float64       `json:"amount,omitempty"`
	InvoicePaidDate *time.Time    `json:"invoicePaidDate,omitempty"`
}

type PaymentOption struct {
	JsonClass      string          `json:"jsonClass,omitempty"`
	PaymentType    string          `json:"paymentType,omitempty"`
	PaidAmount     float64         `json:"paidAmount,omitempty"`
	Change         float64         `json:"change"`
	InstallmentId  string          `json:"installmentId,omitempty"`
	Period		   int32		   `json:"period,omitempty"`
	PaymentOptions []PaymentOption `json:"paymentOptions,omitempty"`
}

type InvoiceItems struct {
	Code     string  `json:"code,omitempty"`
	Desc     string  `json:"desc,omitempty"`
	ImgURL   string  `json:"imgURL,omitempty"`
	Quantity float64 `json:"quantity,omitempty"`
	Cost     float64 `json:"cost,omitempty"`
	Unit     string  `json:"unit,omitempty"`
	Cashback float64 `json:"cashBack,omitempty"`
}

type MyClaim struct {
	Amount       float64 `json:"amount"`
	MerchantId   string  `json:"merchant_id"`
	OrderRef     string  `json:"order_ref"`
	MerchantName string  `json:"merchant_name"`
	Receiver     string  `json:"receiver"`
	Reference    string  `json:"reference"`
}

type MyCustomClaims struct {
	MyClaim
	stdjwt.StandardClaims
}

type PaymentCheque struct {
	JsonClass          string        `json:"jsonClass,omitempty"`
	InvoiceId          string        `json:"invoiceId,omitempty"`
	PaymentSystemName  string        `json:"paymentSystemName"`
	PaymentSystem      string        `json:"paymentSystem,omitempty"`
	Amount             float64       `json:"amount,omitempty"`
	TransactionAccount PaymentOption `json:"transactionAccount,omitempty"`
	TransactionId      string        `json:"transactionId,omitempty"`
	Currency           string        `json:"currency,omitempty"`
}
