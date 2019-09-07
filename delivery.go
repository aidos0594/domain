package domain

import (
	"fmt"
	"strconv"
	"time"
)

type Delivery struct {
	JsonClass      string       `json:"jsonClass,omitempty"`
	Uid            string       `json:"uid,omitempty"`
	City           string       `json:"city"`
	CityName       string       `json:"city_name"`
	MerchantId     string       `json:"merchant_id"`
	DeliveryOption DeliveryOpt  `json:"delivery_option"`
	DeliveryPrice  []DelPrice   `json:"delivery_price"`
	BusinessDays   []Bdays      `json:"business_days"`
	BankHolidays   []*time.Time `json:"bank_holidays"`
}

type GetCalculatedDeliveryResponse struct {
	JsonClass       string                    `json:"jsonClass,omitempty"`
	DeliveryOptions map[string]DeliveryOption `json:"delivery_options"`
	Error           string                    `json:"error,omitempty"`
}

type DeliveryOption struct {
	StartWorkTime         string  `json:"start_work_time"`
	EndWorkTime           string  `json:"end_work_time"`
	DeliveryPrice         []DelPrice `json:"delivery_price"`
	EstimatedDeliveryDate time.Time  `json:"estimated_delivery_date"`
}


type DeliveryOpt struct {
	OrderBefore  int    `json:"order_before,omitempty"`
	TermTitle    string `json:"term_title,omitempty"`
	Term         string `json:"term"`
	PeriodInDays int    `json:"period_in_days,omitempty"`
}

type DelPrice struct {
	StartPrice float64 `json:"start_price"`
	EndPrice   float64 `json:"end_price"`
	Price      float64 `json:"price"`
}

type Bdays struct {
	StartTime     string `json:"start_time"`
	EndTime       string `json:"end_time"`
	WeekDayName   string `json:"week_day_name"`
	OrdinalNumber int    `json:"ordinal_number"`
}

func (d *Delivery) IsWorkingDay(t time.Time) bool {

	for _, val := range d.BusinessDays {
		if val.OrdinalNumber == int(t.Weekday()) {
			return true
		}
	}
	return false
}

func (d *Delivery) GetWorkingHours(t time.Time) (startTime string, endTime string) {
	for _, val := range d.BusinessDays {
		if val.OrdinalNumber == int(t.Weekday()) {
			return val.StartTime, val.EndTime
		}
	}
	return
}

func (d *Delivery) IsWorkingHours(t time.Time) bool {
	layOut := "2006-01-02 15:04:05"
	timeStamp, _ := time.Parse(layOut, t.Format(layOut))
	hr, _, _ := timeStamp.Clock()

	for _, val := range d.BusinessDays {
		if val.OrdinalNumber == int(t.Weekday()) {
			startTime, _ := strconv.Atoi(val.StartTime)
			endTime, _ := strconv.Atoi(val.EndTime)
			if startTime < hr && hr < endTime {
				return true
			}
		}
	}
	return false
}

func (d *Delivery) DeliveryTerms() time.Time {
	//loc, _ := time.LoadLocation("Asia/Almaty")
	Hours := 6
	now := time.Now().Local().Add(time.Hour * time.Duration(Hours))

	if d.DeliveryOption.Term == "from_two_days" {
		x := d.DeliveryOption.PeriodInDays
		after := now.AddDate(0, 0, 1)
		for x != 0 {

			if d.IsWorkingDay(after) == true {
				x--
			}
			if x > 0 {
				after = after.AddDate(0, 0, 1)
			}
		}
		return after
	}

	if d.DeliveryOption.Term == "next_day" {
		x := 1
		after := now.AddDate(0, 0, 1)
		for x != 0 {

			if d.IsWorkingDay(after) == true {
				x--
			}
			if x > 0 {
				after = after.AddDate(0, 0, 1)
			}
		}
		return after
	}

	if d.DeliveryOption.Term == "same_day" {
		after := now
		layOut := "2006-01-02 15:04:05"
		timeStamp, _ := time.Parse(layOut, after.Format(layOut))
		hr, _, _ := timeStamp.Clock()
		fmt.Println("current Time", after)
		fmt.Println("d.IsWorkingDay(after)", d.IsWorkingDay(after))
		fmt.Println("current hours", hr)
		fmt.Println("int(d.DeliveryOption.OrderBefore)", int(d.DeliveryOption.OrderBefore))
		if d.IsWorkingDay(after) == true && hr > int(d.DeliveryOption.OrderBefore) {
			return after
		} else {
			x := 1
			after := now.AddDate(0, 0, 1)
			for x != 0 {

				if d.IsWorkingDay(after) == true {
					x--
				}
				if x > 0 {
					after = after.AddDate(0, 0, 1)
				}
			}
			return after
		}
	}
	return now
}


func (d *Delivery) GetDeliveryPrice(productPrice float64) []DelPrice {
	return d.DeliveryPrice
}
