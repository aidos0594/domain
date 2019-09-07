package domain

import (
	"encoding/xml"
	"fmt"
	"io/ioutil"
	"net/http"
	"strconv"
	"strings"
	"time"
)

type XMLFile struct {
	Shop XMLShop    `xml:"shop"`
	Date customTime `xml:"date,attr"`
}

func (x *XMLFile) GetFromURL(url, login, pass string) (xmlByteData *[]byte, err error) {
	req, err := http.NewRequest("GET", url, nil)
	if login != "" && pass != "" {
		req.SetBasicAuth(login, pass)
	}

	cli := &http.Client{Timeout:time.Duration(30 * time.Second)}
	resp, err := cli.Do(req)
	if err != nil {
		return nil, err
	}
	defer resp.Body.Close()

	if resp.StatusCode != 200 {
		if resp.StatusCode == 401 {
			return nil, fmt.Errorf("Для проверки файла необходимо ввести логин и пароль")
		}
		return nil, fmt.Errorf("Responce status=%d", resp.StatusCode)
	}
	if resp.ContentLength == 0 {
		return nil, fmt.Errorf("Responce contentLength=%d", resp.ContentLength)
	}

	//var testbytes []byte

	//	if 1==1 {
	//		testbytes = []byte(`
	//<?xml version="1.0" encoding="UTF-8"?>
	//	<fm_catalog date="2017-02-05 17:22">
	//	<shop>
	//	<merchant_id>3423423</merchant_id>
	//	</shop>
	//	</fm_catalog>`)
	//	}

	byteValue, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		return nil, err
	}
	//fmt.Println(string(byteValue))
	err = xml.Unmarshal(byteValue, x)
	if err != nil {
		return nil, err
	}

	return &byteValue, nil
}

type customTime struct {
	time.Time
}

type timeStampToTime struct {
	time.Time
}

func (t *timeStampToTime) UnmarshalXMLAttr(attr xml.Attr) error {
	i, err := strconv.ParseInt(attr.Value, 10, 64)
	if err != nil {
		return err
	}

	tm := time.Unix(i, 0)
	*t = timeStampToTime{tm}
	return nil
}

func (c *customTime) UnmarshalXMLAttr(attr xml.Attr) error {
	const shortFrom = "2006-01-02 15:04"
	parse, err := time.Parse(shortFrom, attr.Value)
	if err != nil {
		return err
	}
	*c = customTime{parse}
	return nil
}

type XMLShop struct {
	MerchantId string     `xml:"merchant-id"`
	Offers     []XMLOffer `xml:"offers>offer"`
}

type XMLOffer struct {
	Sku             string              `xml:"sku,attr"`
	Name            string              `xml:"name"`
	Vendor          string              `xml:"vendor"`
	Barcodes        []string            `xml:"barcodes>barcode"`
	Price           float64             `xml:"price"`
	PickupOptions   []XMLPickupOption   `xml:"pickup-options>pickup-option"`
	DeliveryOptions []XMLDeliveryOption `xml:"delivery-options>delivery-option"`
	CityPrices      []XMLCityPrice      `xml:"city-prices>city-price"`
}

func (offer *XMLOffer) Filter(storeMap map[string]string, cityMap map[string]float64, deliveryMap map[string]bool) string {
	var xmlLogs string
	var xmlPO []XMLPickupOption
	var xmlDO []XMLDeliveryOption
	var xmlCP []XMLCityPrice

	sellsMap := make(map[string]bool)
	hasStoresMap := make(map[string]bool)
	for key, _ := range storeMap {
		hasStoresMap[storeMap[key]] = true
	}
	for _, value := range offer.PickupOptions {
		value.Id = strings.TrimSpace(value.Id)
		if _, ok := storeMap[value.Id]; ok {
			xmlPO = append(xmlPO, value)
			sellsMap[storeMap[value.Id]] = true
		} else {
			xmlLogs += "SKU: " + offer.Sku + ", точка продажи id: " + value.Id + " отсутствует в базе \n"
		}
	}
	offer.PickupOptions = xmlPO

	for _, value := range offer.DeliveryOptions {
		_, hasDeliveryInCity := deliveryMap[value.CityId]
		if _, isExistsCityInOurSystem := cityMap[value.CityId];  isExistsCityInOurSystem {
			if hasDeliveryInCity {
				xmlDO = append(xmlDO, value)
			} else {
				xmlLogs += "SKU: " + offer.Sku + " не заведена доставка с кодом города: " + value.CityId + "\n"
			}
		} else {
			xmlLogs += "SKU: " + offer.Sku + " ,указан не корректный код города в DeliveryOptions: " + value.CityId + "\n"
		}
	}
	offer.DeliveryOptions = xmlDO

	for _, value := range offer.CityPrices {
		if _, ok := cityMap[value.CityId]; ok {
			xmlCP = append(xmlCP, value)
		} else {
			xmlLogs += "SKU: " + offer.Sku + ", код города (CityPrices): " + value.CityId + " не соответствует \n"
		}
	}
	offer.CityPrices = xmlCP

	return xmlLogs
}

func (offer *XMLOffer) IsChanged(nomenclature Nomenclature) bool {

	pickupOptionBool := len(offer.PickupOptions) != len(nomenclature.PickupOptions)
	deliveryOptionBool := len(offer.DeliveryOptions) != len(nomenclature.DeliveryOptions)
	cityOptionBool := len(offer.CityPrices) != len(nomenclature.CityPrices)
	changedPriceBool := offer.Price != nomenclature.Price
	// nomenclature.Available == false needs for availabling nomenclature offer consist
	if pickupOptionBool || deliveryOptionBool || cityOptionBool || changedPriceBool || nomenclature.Available == false {
		return true
	} else {

		pickupOptionsMap := make(map[string]XMLPickupOption)
		for _, body := range offer.PickupOptions {
			pickupOptionsMap[body.Id] = body
		}
		for _, value := range nomenclature.PickupOptions {
			if option, isOk := pickupOptionsMap[value.Id]; isOk {
				if (option.PreOrderDate != nil && value.PreOrderDate == nil) || (option.PreOrderDate == nil && value.PreOrderDate != nil) {
					return true
				} else if option.PreOrderDate != nil && value.PreOrderDate != nil {
					if option.PreOrderDate.Sub(value.PreOrderDate.Time) > 10 || option.PreOrderDate.Sub(value.PreOrderDate.Time) < -10 {
						return true
					}
				}
			} else {
				return true
			}
		}

		deliveryOptionsMap := make(map[string]XMLDeliveryOption)
		for _, body := range offer.DeliveryOptions {
			deliveryOptionsMap[body.CityId] = body
		}
		for _, value := range nomenclature.DeliveryOptions {
			if option, isOk := deliveryOptionsMap[value.CityId]; isOk {
				if option.Cost != value.Cost || option.Days != value.Days {
					return true
				}
			} else {
				return true
			}
		}

		cityOptionsMap := make(map[string]XMLCityPrice)
		for _, body := range offer.CityPrices {
			cityOptionsMap[body.CityId] = body
		}
		for _, value := range nomenclature.CityPrices {
			if option, isOk := cityOptionsMap[value.CityId]; isOk {
				if option.CityPrice != value.CityPrice {
					return true
				}
			} else {
				return true
			}
		}
	}
	return false
}

func (offer *XMLOffer) IsChanged2(nomenclature Nomenclature, citiesmap map[string]float64) bool {
	if !nomenclature.Available ||
		(nomenclature.ProductId != "" && nomenclature.SkuId != "" && !nomenclature.IsVisible && nomenclature.Status == "approved") {
		return true
	}
	if len(nomenclature.CityPriceMap) != len(citiesmap) {
		return true
	}
	for key, _ := range nomenclature.CityPriceMap {
		if _, ok := citiesmap[key]; ok {
			if nomenclature.CityPriceMap[key] != citiesmap[key] {
				return true
			}
		} else {
			return true
		}
	}

	return false
}

type XMLPickupOption struct {
	Id           string           `xml:"id,attr" json:"id,omitempty"`
	PreOrderDate *timeStampToTime `xml:"preorder-date,attr" json:"pre_order_date,omitempty"`
}

type XMLDeliveryOption struct {
	CityId string  `xml:"city-id,attr" json:"city_id,omitempty"`
	Cost   float64 `xml:"cost,attr" json:"cost,omitempty"`
	Days   int64   `xml:"days,attr" json:"days,omitempty"`
}

type XMLCityPrice struct {
	CityId    string  `xml:"city-id,attr" json:"city_id,omitempty"`
	CityPrice float64 `xml:",chardata" json:"city_price,omitempty"`
}

type DataUploadedHistory struct {
	UploadedTime    time.Time `json:"uploaded_time,omitempty"`
	PublishedTime   time.Time `json:"published_time,omitempty"`
	Status          string    `json:"status,omitempty"`
	GoodItems       int64     `json:"good_items,omitempty"`
	XmlFile         string    `json:"xml_file,omitempty"`
	XmlFileId       string    `json:"xml_file_id,omitempty"`
	XmlFileName     string    `json:"xml_file_name,omitempty"`
	PriceListNumber int64     `json:"price_list_number"`

	OfferAmount       int `json:"offer_amount"`
	UnIdentifiedGoods int `json:"un_identified_goods"`
	IdentifiedGoods   int `json:"identified_goods"`
	ProcessedGoods    int `json:"processed_goods"`
	UnProcessedGoods  int `json:"un_processed_goods"`
	WarningGoods      int `json:"warning"`

	UpdatedOffers   int    `json:"updated_offers,omitempty"`
	NewOffers       int    `json:"new_offers,omitempty"`
	UpdatedShowcase int    `json:"updated_showcase,omitempty"`
	NotValid        bool   `json:"not_valid,omitempty"`
	BlobId          string `json:"blob_id,omitempty"`
}

type DataUploadSettings struct {
	MerchantId   string    `json:"merchant_id,omitempty"`
	Url          string    `json:"url,omitempty"`
	Login        string    `json:"login,omitempty"`
	Password     string    `json:"password,omitempty"`
	Scheduler    string    `json:"scheduler,omitempty"`
	IdScheduler  string    `json:"id_scheduler,omitempty"`
	Status       string    `json:"status,omitempty"`
	UpdatedOn    time.Time `json:"updated_on,omitempty"`
	ActionStatus string    `json:"action_status,omitempty"`
}

type DataPriceFile struct {
	Filename string `json:"filename,omitempty"`
	Id       string `json:"id,omitempty"`
	Logs     string `json:"logs,omitempty"`
}
