package domain

import (
	"errors"
	"github.com/sony/sonyflake"
	"strconv"
	"strings"
	"time"
)

type (
	OrderStatus string
	PayType string
	DeliveryType string
)

const (
	//order statuses
	New  = "new"  //initial
	PendingApprove = "pending_approve" // merchant must confirm existence of goods
	PrePendingApprove = "pre_pending_approve" // merchant confirm before bank scoring
	PendingApproveByBank = "pending_approve_by_bank" // pending bank scoring
	PendingPickup = "pending_pickup"
	Taken = "taken"
	OnDelivery = "on_delivery"
	Delivered = "delivered"
	Cancelled = "cancelled"
	Returned = "returned"

	//pay types
	CARD  = "CARD"
	COD = "COD"
	ACQUIRING = "ACQUIRING"
	FORTE_EXPRESS = "FORTE_EXPRESS"
	BONUS = "BONUS"
	CREDIT = "CREDIT"

	//delivery types
	DELIVERY = "delivery" // доставка
	Pickup = "pickup" // самовывоз
)

type Order struct {
	JsonClass             string       `json:"jsonClass,omitempty"`
	BasketId              string       `json:"basket_id,omitempty"`  //elastic
	Uid                   string       `json:"uid,omitempty"`        //elastic
	Scope                 string       `json:"scope,omitempty"`      //elastic
	CreatedOn             time.Time    `json:"created_on,omitempty"` //elastic
	UpdatedOn             time.Time    `json:"updated_on,omitempty"` //elastic
	OwnType               string       `json:"own_type,omitempty"`   //dar_profile_id или merchant_id(darbiz) //elastic
	Owner                 string       `json:"owner,omitempty"`      //elastic
	Comment               string       `json:"comment,omitempty"`
	PackageType           string       `json:"package_type,omitempty"` // тип групировки по городам или по магазинам
	CommonPrice           float64      `json:"common_price,omitempty"` //elastic
	CommonOldPrice        float64      `json:"common_old_price,omitempty"`
	CanceledComments      string       `json:"canceled_comments,omitempty"`
	CanceledType          string       `json:"canceled_type,omitempty"`
	CommonQuantity        int          `json:"common_quantity,omitempty"`
	Packages              []Package    `json:"packages,omitempty"` //elastic sku товара
	Items                 []BasketItem `json:"items,omitempty"`
	ReceiverContacts      *Contacts    `json:"receiver_contacts,omitempty"` //elastic Name ,Surname,Mobile,email
	MerchantsId           []string     `json:"merchants_id,omitempty"`      //elastic
	PayTypes              string       `json:"pay_types,omitempty"`         //elastic
	PayTitle              string       `json:"pay_title,omitempty"`
	PaidTitle             string       `json:"paid_type,omitempty"`
	PaidAmount            float64      `json:"paid_amount,omitempty"`
	Paid                  bool         `json:"paid"`
	PaymentAdditional	  string 	   `json:"payment_additional,omitempty"`
	InvoiceRef            string       `json:"invoice_ref,omitempty"`
	Discount              string       `json:"discount,omitempty"`
	DiscountName          string       `json:"discount_name,omitempty"`
	Status                string       `json:"status,omitempty"`       //elastic
	StatusTitle           string       `json:"status_title,omitempty"` //elastic
	DeliveryPrice         float64      `json:"delivery_price,omitempty"`
	DeliveryTypes         string       `json:"delivery_types,omitempty"`
	DeliveryTitle         string       `json:"delivery_title,omitempty"`
	EstimatedDeliveryDate *time.Time   `json:"estimated_delivery_date,omitempty"`
	DeliveryCommission    float64      `json:"delivery_commission,omitempty"`
	RefundCommission      float64      `json:"refund_commission,omitempty"`
	AditionalInfo         string       `json:"aditional_info,omitempty"`
	InvoiceIsCreated      bool         `json:"invoice_is_created"`
	ChangeFrom            float64      `json:"change_from,omitempty"`
	PaidDate              time.Time    `json:"paidDate,omitempty"`
	LocalId               string       `json:"localId,omitempty"`
	OwnerTitle            string       `json:"ownerTitle,omitempty"`
	MerchantId            string       `json:"merchantId,omitempty"`
	TerminalId            string       `json:"terminalId,omitempty"`
	Token                 string       `json:"token,omitempty"`
	StockId               string       `json:"stock_id,omitempty"`
	Otp                   int          `json:"otp,omitempty"`
	CityId                string       `json:"city_id,omitempty"`
	StatusHistory         []Statuses   `json:"status_history,omitempty"`
	UpdatedByProfileId    string       `json:"updated_by_profile_id,omitempty"`
	UpdatedByName         string       `json:"updated_by_name,omitempty"`
	UpdatedByMobilePhone  string       `json:"updated_by_mobile_phone,omitempty"`
	PickupDate            *time.Time   `json:"pickup_date,omitempty"`
	CalculateByDelivery   bool         `json:"calculate_by_delivery"`
}

type Statuses struct {
	Status               string    `json:"status,omitempty"`
	StatusWeight         float64   `json:"status_weight,omitempty"`
	StatusTitle          string    `json:"status_title,omitempty"`
	CreatedOn            time.Time `json:"created_on,omitempty"`
	UpdatedByProfileId   string    `json:"updated_by_profile_id,omitempty"`
	UpdatedByName        string    `json:"updated_by_name,omitempty"`
	UpdatedByMobilePhone string    `json:"updated_by_mobile_phone,omitempty"`
}

type ForteKassaOrder struct {
	ParentInvoiceRef string                 `json:"parent_invoice_ref,omitempty"`
	JsonClass        string                 `json:"jsonClass,omitempty"`
	Uid              string                 `json:"id"`
	LocalId          string                 `json:"localId,omitempty"`
	OwnType          string                 `json:"ownType,omitempty"`
	Owner            string                 `json:"owner,omitempty"`
	PayTypes         string                 `json:"payTypes,omitempty"`
	PayTitle         string                 `json:"payTitle,omitempty"`
	PaidTitle        string                 `json:"paid_type,omitempty"`
	PaidAmount       float64                `json:"paid_amount,omitempty"`
	PaidDate         string                 `json:"paidDate,omitempty"`
	Paid             bool                   `json:"paid"`
	MerchantId       string                 `json:"merchantId,omitempty"`
	MerchantTitle    string                 `json:"merchantTitle,omitempty"`
	TerminalId       string                 `json:"terminalId,omitempty"`
	Token            string                 `json:"token,omitempty"`
	Price            float64                `json:"price"`
	Quantity         float64                `json:"quantity"`
	OrderItems       []ForteKassaOrderItems `json:"orderItems,omitempty"`
	CreatedOn        time.Time              `json:"created_on"`
	UpdatedOn        time.Time              `json:"updated_on"`
	InvoiceRef       string                 `json:"invoice_ref"`
	OrderStatus      string                 `json:"order_status"`
	OrderStatusTitle string                 `json:"order_status_title"`

	ParentOrderId   string  `json:"parentOrderId,omitempty"`
	OrderType       string  `json:"orderType,omitempty"`
	ReturnReason    string  `json:"returnReason,omitempty"`
	ReturnPayAmount float64 `json:"returnPayAmount,omitempty"`
}

type ForteKassaOrderItems struct {
	Nomenclature
	//ItemTitle          string  `json:"itemTitle"`
	//Quantity           float64 `json:"quantity"`
	//Price              float64 `json:"price"`
	//NomenclatureId     string  `json:"nomenclatureId"`
	//StockId            string  `json:"stockId"`
	CategoryGroupId    string `json:"categoryGroupId"`
	CategoryGroupTitle string `json:"categoryGroupTitle"`
	Status             string `json:"status"`
	StatusTitle        string `json:"statusTitle"`
}

type OfflineOrderItems struct {
	ProductId string  `json:"productId"`
	SkuId     string  `json:"skuId"`
	Quantity  float64 `json:"quantity"`
}

type OrderInElastic struct {
	Uid              string    `json:"uid,omitempty"`
	LocalId          string    `json:"localId,omitempty"`
	InvoiceRef       string    `json:"invoice_ref,omitempty"`
	Scope            string    `json:"scope,omitempty"`
	FIO              string    `json:"fio,omitempty"`
	Mobile           string    `json:"mobile,omitempty"`
	WalletMobile     string    `json:"wallet_mobile,omitempty"`
	OrderStatus      string    `json:"order_status,omitempty"`
	OrderStatusTitle string    `json:"order_status_title,omitempty"`
	CreatedOn        time.Time `json:"created_on,omitempty"`
	Paid             bool      `json:"paid"`
	UpdatedOn        time.Time `json:"updated_on,omitempty"`
	CommonPrice      float64   `json:"common_price,omitempty"`
	DeliveryPrice    float64   `json:"delivery_price,omitempty"`
	DeliveryTypes    string    `json:"delivery_types"`
	DeliveryTitle    string    `json:"delivery_title,omitempty"`
	PayTitle         string    `json:"pay_title,omitempty"`
	PayTypes         string    `json:"pay_types,omitempty"`
	MerchantIds      []string  `json:"merchant_ids,omitempty"`
	StockIds         []string  `json:"stock_ids,omitempty"`
	StoreIds         []string  `json:"store_ids,omitempty"`
	OwnType          string    `json:"own_type,omitempty"` //dar account или merchant(darbiz) //elastic
	Owner            string    `json:"owner,omitempty"`
	Address          string    `json:"address,omitempty"`
	PaidTitle        string    `json:"paid_title,omitempty"`
	ProductIds       []string  `json:"product_ids,omitempty"`
	MerchantId       string    `json:"merchantId,omitempty"`
	TerminalId       string    `json:"terminalId,omitempty"`
	CityId           string    `json:"city_id,omitempty"`
	PickupPoints     []string  `json:"pickup_points,omitempty"`
	OrderDump        string    `json:"orderdump,omitempty"`
}

func (o *Order) SetStatusCancelled(comments string) error {
	o.CanceledComments = comments
	return o.SetStatusToPackages(Cancelled, "Отменен", 2500)
}

func (o *Order) SetStatusOnDelivery() error {
	return o.SetStatusToPackages(OnDelivery, "На доставке", 2300)
}


func (o *Order) SetInitialStatus() error {  //SetInitialStatus()
	var err error
	flake := sonyflake.NewSonyflake(sonyflake.Settings{})
	id, err := flake.NextID()
	if err != nil {
		return err
	}
	s := strconv.FormatUint(id, 10)
	o.Uid = s[0:len(s)-5] + "-" + s[len(s)-5:]
	switch strings.ToUpper(o.PayTypes) {
	case COD:
		err = o.SetStatusToPackages(PendingApprove, "Ожидает подтверждения продавцом", 400)
	default:
		err = o.SetStatusToPackages(New, "Новый заказ", 400)
	}
	return err
}

func (o *Order) SetStatusCanceled() error {
	var err error
	if o.Status == Cancelled {
		return errors.New("Заказ уже отменен ")
	}
	if o.Status != Taken && o.Status != Delivered && o.Status != Returned {
		err = o.SetStatusToPackages(Cancelled, "Заказ отменен", 2500)
	} else {
		return errors.New("Заказ не может быть отменен ")
	}
	return err
}

func (o *Order) SetStatusMerchantView() error {

	if o.Status == "merchant_view" || o.Status != "new" {

		return errors.New("Заказ не может быть merchant_view")
	}
	return o.SetStatusToPackages("merchant_view", "Предварительно подтвержден", 400)
}

func (o *Order) SetStatusGiveAway(otp int) error {

	if o.Status != PendingPickup && o.Status != OnDelivery {
		return errors.New("Заказ не может быть giveAway, статус = " + o.Status)
	}

	if o.Otp != otp {
		return errors.New("Неверный код otp ")
	}

	if o.PayTypes != COD && !o.Paid { // each payType except COD should be paid
		return errors.New("Заказ не оплачен ")
	}

	var err error
	if o.DeliveryTypes != DELIVERY {
		err = o.SetStatusToPackages(Taken, "Выдан", 2300)
	} else {
		err = o.SetStatusToPackages(Delivered, "Доставлен", 2300)
	}

	return err
}

func (o *Order) SetStatusReturned(otp int) error {

	if o.Status != Taken && o.Status != Delivered {
		return errors.New("Заказ не может быть returned, статус = " + o.Status)
	}
	if o.Otp != otp {
		return errors.New("Неверный код otp ")
	}
	if !o.Paid {
		return errors.New("Заказ не оплачен ")
	}

	return o.SetStatusToPackages(Returned, "Возвращен", 2300)
}

func (o *Order) SetStatusAccept() error {

	if o.Status == PendingPickup || o.Status == OnDelivery {
		return errors.New("Заказ уже принят, статус = " + o.Status)
	}

	if o.Status != PendingApprove && o.Status != PrePendingApprove {
		return errors.New("Заказ не может быть принят ")
	}

	if o.Status == PrePendingApprove {
		return o.SetStatusToPackages(PendingApproveByBank, "Ожидает подтверждения банком", 2300)
	}

	if o.PayTypes != COD && !o.Paid { // each payType except COD should be paid
		return errors.New("Заказ не может быть принят, заказ не оплачен ")
	}

	var err error
	if o.DeliveryTypes != DELIVERY {
		err = o.SetStatusToPackages(PendingPickup, "Ожидает выдачи", 2300)
	} else {
		err = o.SetStatusToPackages(OnDelivery, "На доставке", 2300)
	}

	return err
}

func (o *Order) SetStatusPendingApprove() error {

	if o.Status != New && o.Status != "" && o.Status != PendingApproveByBank {
		return errors.New("Заказ не может быть pending_approve")
	}

	return o.SetStatusToPackages(PendingApprove, "Ожидает подтверждения", 400)
}

func (o *Order) SetStatusPrePendingApprove() error {

	if o.Status != New && o.Status != "" {
		return errors.New("Заказ не может быть pre_pending_approve")
	}

	return o.SetStatusToPackages(PrePendingApprove, "Ожидает подтверждения продавцом до банка", 400)
}

func (o *Order) SetStatusToPackages(status string, statusTitle string, weight float64) error {
	if o.Status == status {
		return errors.New("Заказ уже " + status)
	}

	o.Status = status
	o.StatusTitle = statusTitle

	createTime := time.Now().UTC()
	o.UpdatedOn = createTime

	var statusHistory Statuses
	statusHistory.Status = o.Status
	statusHistory.StatusTitle = o.StatusTitle
	statusHistory.CreatedOn = createTime
	statusHistory.StatusWeight = weight
	statusHistory.UpdatedByMobilePhone = o.UpdatedByMobilePhone
	statusHistory.UpdatedByProfileId = o.UpdatedByProfileId
	statusHistory.UpdatedByName = o.UpdatedByName

	o.StatusHistory = append(o.StatusHistory, statusHistory)

	for i, _ := range o.Items {
		o.Items[i].Status = status
	}

	return nil
}

func (o *ForteKassaOrder) SetForteKassaOrderStatusDone() error {
	o.SetForteKassaOrdersStatusToPackages("done", "Завершен", 2000)
	return nil
}

func (o *ForteKassaOrder) SetForteKassaOrdersStatusToPackages(status string, statusTitle string, weight float64) error {

	o.OrderStatus = status
	o.OrderStatusTitle = statusTitle
	o.UpdatedOn = time.Now().UTC()

	for i, _ := range o.OrderItems {
		o.OrderItems[i].Status = status
		o.OrderItems[i].StatusTitle = statusTitle
	}

	return nil
}

// unused
func (o *Order) ApproveAll() error {
	o.Status = "approved"
	o.StatusTitle = "Подтвержден"
	createTime := time.Now()

	for i, _ := range o.Packages {
		var pkgstatus Statuses
		pkgstatus.Status = o.Status
		pkgstatus.StatusTitle = o.StatusTitle
		pkgstatus.CreatedOn = createTime
		pkgstatus.StatusWeight = 1000

		o.Packages[i].Status = o.Status
		o.Packages[i].StatusTitle = o.StatusTitle
		o.Packages[i].PackageStatus = append(o.Packages[i].PackageStatus, pkgstatus)
		for j, _ := range o.Packages[i].Products {
			o.Packages[i].Products[j].Product.Status = "approved"
			o.Packages[i].Products[j].Product.UpdatedOn = createTime
		}
	}
	o.UpdatedOn = createTime

	return nil
}

func (o *Order) SetStatusDeliveredNotPaid() error {
	return o.SetStatusToPackages("delivered_not_paid", "Доставлен", 1200)
}

func (o *Order) SetStatusPendingClietApprove() error {
	return o.SetStatusToPackages("pending_client_approve", "Подтверждение клиентом", 500)
}

func (o *Order) SetStatusApproved() error {
	return o.SetStatusToPackages("approved", "Подтвержден", 600)
}

func (o *Order) SetStatusPrepare() error {
	return o.SetStatusToPackages("prepare", "В обработке", 2300)
}

func (o *Order) SetStatusDone() error {
	return 	o.SetStatusToPackages("done", "Завершен", 2000)
}

func (o *Order) CalculateCommon() {
	o.CommonPrice = 0
	o.CommonOldPrice = 0
	o.CommonQuantity = 0
	for i, _ := range o.Packages {
		o.Packages[i].CommonPrice = 0
		o.Packages[i].CommonOldPrice = 0
		o.Packages[i].CommonQuantity = 0

		for _, prod := range o.Packages[i].Products {
			if prod.Product.Status != "not_approved" {
				o.Packages[i].CommonPrice += prod.Product.Skus[0].Price * float64(prod.Product.Skus[0].Amount)
				if prod.Product.Skus[0].Price > prod.Product.Skus[0].OldPrice {
					o.Packages[i].CommonOldPrice += prod.Product.Skus[0].Price * float64(prod.Product.Skus[0].Amount)
				} else {
					o.Packages[i].CommonOldPrice += prod.Product.Skus[0].OldPrice * float64(prod.Product.Skus[0].Amount)
				}
				o.Packages[i].CommonQuantity += prod.Product.Skus[0].Amount
			}

		}
		o.CommonPrice += o.Packages[i].CommonPrice
		o.CommonOldPrice += o.Packages[i].CommonOldPrice
		o.CommonQuantity += o.Packages[i].CommonQuantity

	}

}

func (o *Order) CheckApprove() error {
	CommonCount := 0
	ApprovedCount := 0
	NotApprovedCount := 0

	for i, _ := range o.Packages {
		for _, prod := range o.Packages[i].Products {
			if prod.Product.Status == "approved" {
				ApprovedCount++
			} else {
				NotApprovedCount++
			}
			CommonCount++
		}

	}
	if CommonCount == ApprovedCount {
		o.SetStatusApproved()
		//o.SetStatusOnDelivery()
	} else if CommonCount == NotApprovedCount {
		o.SetStatusCancelled("not_approved")
		o.CalculateCommon()

	} else {
		o.SetStatusPendingClietApprove()
		o.CalculateCommon()
	}
	return nil
}

func (o *Order) ChangeOrderAddress(contacts Contacts) error {

	if contacts.Name != "" {
		o.ReceiverContacts.Name = contacts.Name
	}
	if contacts.Surname != "" {
		o.ReceiverContacts.Surname = contacts.Surname
	}
	if contacts.Title != "" {
		o.ReceiverContacts.Title = contacts.Title
	}
	if contacts.City != "" {
		o.ReceiverContacts.City = contacts.City
	}
	if contacts.Address != "" {
		o.ReceiverContacts.Address = contacts.Address
	}
	if contacts.Email != "" {
		o.ReceiverContacts.Email = contacts.Email
	}
	if contacts.HouseNumber != "" {
		o.ReceiverContacts.HouseNumber = contacts.HouseNumber
	}
	if contacts.FlatNumber != "" {
		o.ReceiverContacts.FlatNumber = contacts.FlatNumber
	}
	if contacts.Mobile != "" {
		o.ReceiverContacts.Mobile = contacts.Mobile
	}

	return nil
}
