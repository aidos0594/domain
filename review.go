package domain

import (
	"errors"
	"strings"
	"time"
)

type Review struct {
	Uid            string    `json:"uid"`
	MerchantId     string    `json:"merchant_id"`
	OrderId        string    `json:"order_id"`
	WLabelId       string    `json:"w_label_id"`
	CustomerName   string    `json:"customer_name"`
	ProfileId      string    `json:"profile_id"`
	Mobile         string    `json:"mobile"`
	Phone          string    `json:"phone,omitempty"`
	Mark           int       `json:"mark"`
	ReviewText     string    `json:"review_text,omitempty"`
	Status         string    `json:"status"`
	ProblemStatus  string    `json:"problem_status,omitempty"`
	CreatedOn      time.Time `json:"created_on"`
	QueuedOn       time.Time `json:"queued_on"`
	UpdatedOn      time.Time `json:"updated_on"`
	OrderCreatedOn time.Time `json:"order_created_on"`

	MerchantMessage       string    `json:"merchant_message,omitempty"`
	MerchantName          string    `json:"merchant_name,omitempty"`
	MerchantCommentedTime time.Time `json:"merchant_commented_time,omitempty"`

	AdminMessage       string    `json:"admin_message,omitempty"`
	AdminPublishedTime time.Time `json:"published_time,omitempty"`
	AdminId            string    `json:"admin_id,omitempty"`
	AdminName          string    `json:"admin_name,omitempty"`
	OrderSnapshotDump  string    `json:"ord_snap_dump"`
}

type AggDailyRating struct {
	MerchantId              string    `json:"merchant_id"`
	UntimelyProcessedOrders int64     `json:"untimely_processed_orders"`
	CancelledOrders         int64     `json:"cancelled_orders"`
	FinishedOrders          int64     `json:"finished_orders"`
	ReturnedOrders          int64     `json:"returned_orders"`
	IsActive                bool      `json:"is_active"`
	UpdatedOn               time.Time `json:"datetime"`
	ReviewsAmount           int64     `json:"reviews_amount"`
	RatingSum               float32   `json:"rating_sum"`
	Rating                  float32   `json:"rating"`
}

type MerchantRatingInElastic struct {
	MerchantId                  string    `json:"merchant_id"`
	MerchantName                string    `json:"merchant_name"`
	UntimelyProcessedPercentage float32   `json:"untimely_processed_percentage"`
	CancelledPercentage         float32   `json:"cancelled_percentage"`
	ReturnedPercentage          float32   `json:"returned_percentage"`
	IsActive                    bool      `json:"is_active"`
	UpdatedOn                   time.Time `json:"datetime"`
	Rating                      float32   `json:"rating"`
	StatisticsDump              string    `json:"statistics_dump"`
}

type OrderSnapshot struct {
	ProductName           string        `json:"product_name"`
	ImageUrl              string        `json:"product_image"`
	Price                 float64       `json:"price"`
	OrderDate             time.Time     `json:"order_date"`
	DeliveryType          string        `json:"delivery_type"`
	Address               *PickupAddres `json:"address"`
	DeliveryDate          time.Time     `json:"delivery_date"`
	MerchantId            string        `json:"merchant_id"`
	ReceiverMobile        string        `json:"receiver_mobile"`
	ReceiverName          string        `json:"receiver_name"`
	Owner                 string        `json:"owner"`
	WLabelId              string        `json:"w_label_id"`
	MerchantName          string        `json:"merchant_name"`
	CreatedOn             time.Time     `json:"created_on,omitempty"`
	UpdatedOn             time.Time     `json:"updated_on,omitempty"`
	OrderStatus           string        `json:"order_status,omitempty"`
	OrderStatusTitle      string        `json:"order_status_title,omitempty"`
	Articul               string        `json:"articul,omitempty"`
	AcceptedByName        string        `json:"accepted_by_name,omitempty"`
	AcceptedByMobilePhone string        `json:"accepted_by_mobile_phone,omitempty"`
	GivenByName           string        `json:"given_by_name,omitempty"`
	GivenByMobilePhone    string        `json:"given_by_mobile_phone,omitempty"`
}

func (r *Review) SetStatus(status string) error {
	if r.Status == status {
		return errors.New("Отзыв уже " + status)
	}

	r.UpdatedOn = time.Now().UTC()
	r.Status = status

	return nil
}

func (r *Review) SetProblemStatus(problemStatus string) error {
	if r == nil || r.Status != "accepted" {
		return errors.New("отзыв должен быть опубликован")
	}

	if r.ProblemStatus == problemStatus {
		return errors.New("Отзыв уже " + problemStatus)
	}

	switch strings.ToLower(problemStatus) {
	case "resolved", "unresolved":
		if r.ProblemStatus != "pre_resolved" && r.ProblemStatus != "conflict" {
			return errors.New("отзыв не может быть " + problemStatus)
		}
	case "pre_resolved":
		if r.Mark > 3 {
			return errors.New("отзыв не может быть " + problemStatus)
		}
	default:
		return errors.New("нет првильный action = " + problemStatus)
	}

	r.UpdatedOn = time.Now().UTC()
	r.ProblemStatus = problemStatus

	return nil
}
