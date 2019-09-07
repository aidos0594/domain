package domain

import (
	"time"
)

// поиск по городам товары

type ProductShowcase struct {
	Uid             string               `json:"uid,omitempty"`
	CreatedOn       time.Time            `json:"created_on,omitempty"`
	UpdatedOn       time.Time            `json:"updated_on,omitempty"`
	Status          string               `json:"status,omitempty"`
	Name            string               `json:"name,omitempty"` //обязательное
	Description     string               `json:"description,omitempty"`
	Brand           string               `json:"brand,omitempty"`
	Scope           []string             `json:"scope,omitempty"` //обязательное
	AggsRating      float64              `json:"aggs_rating,omitempty"`
	Medias          []Media              `json:"media,omitempty"`
	CategoriesArray []string             `json:"categories_array,omitempty"` ///нужен ли
	ProductPrice    float64              `json:"product_price,omitempty"`
	ProductOldPrice float64              `json:"product_old_price,omitempty"`
	SalesPrice      float64              `json:"sales_price,omitempty"`
	IsVisible       bool                 `json:"is_visible"`
	DefaultCategory string               `json:"default_category,omitempty"`
	SystemCategory  string               `json:"system_category,omitempty"`
	Label           []string             `json:"label,omitempty"` //new top
	Skus            []ProductShowcaseSku `json:"skus"`
	Dimensions      Dimensions           `json:"dimensions,omitempty"`
	Popularity		float64				 `json:"popularity,omitempty"`
}

type Crutch struct {
	Price        float64 `json:"price"`
	ProductId    string  `json:"product_id"`
	SkuId        string  `json:"sku_id"`
	Installment  int32   `json:"installment"`
	Cashback     float64 `json:"cashback,omitempty"`
	CityId       string  `json:"city_id"`
	CommissionFM float64 `json:"commission_fm,omitempty"`
	Brand        string  `json:"brand,omitempty"`
}

type UpdatedShowcases struct {
	JsonClass        string            `json:"json_class,omitempty"`
	UpdatedShowcases []ProductShowcase `json:"updated_showcase,omitempty"`
	CreatedShowcases []ProductShowcase `json:"created_showcase,omitempty"`
	DeletedShowcases []ProductShowcase `json:"deleted_showcase,omitempty"`
}

func (s *ProductShowcase) CreateShowcaseFromProduct(prd *Product) error {
	s.Uid = prd.Uid
	s.CreatedOn = time.Now().UTC()
	s.UpdatedOn = s.CreatedOn
	s.Status = "created"
	s.Name = prd.Name
	s.Description = prd.Description
	s.Brand = prd.Brand
	s.Scope = prd.Scope
	s.AggsRating = prd.AggsRating
	s.Medias = prd.Medias
	s.CategoriesArray = prd.CategoriesArray
	s.ProductPrice = prd.ProductPrice
	s.ProductOldPrice = prd.ProductOldPrice
	s.IsVisible = true
	s.DefaultCategory = prd.DefaultCategory
	s.Label = prd.Label
	s.Dimensions = prd.Dimensions

	for _, value := range prd.Skus {
		var item ProductShowcaseSku
		item.Uid = value.Uid
		item.Medias = value.Medias

		item.GSKU = append(item.GSKU, value.GSKU)
		item.CreatedOn = s.CreatedOn
		item.UpdatedOn = s.UpdatedOn

		item.ProductParam = value.ProductParam
		item.MinPrice = make(map[string]float64)
		item.MinNomenId = make(map[string]string)

		s.Skus = append(s.Skus, item)
	}

	return nil
}

func (s *ProductShowcase) DeleteSku(nomenSkuid string) error {
	for idx, value := range s.Skus {
		if value.Uid == nomenSkuid {
			s.Skus = append(s.Skus[:idx], s.Skus[idx+1:]...)
			break
		}
	}

	return nil
}

func (s *ProductShowcase) AddSku(prd *Product) error {
	for _, value := range prd.Skus {
		var item ProductShowcaseSku
		item.Uid = value.Uid
		item.Medias = value.Medias

		item.GSKU = append(item.GSKU, value.GSKU)
		item.CreatedOn = time.Now().UTC()
		item.UpdatedOn = item.CreatedOn

		item.ProductParam = value.ProductParam
		item.MinPrice = make(map[string]float64)
		item.MinNomenId = make(map[string]string)

		s.Skus = append(s.Skus, item)
	}

	return nil
}

func (s *ProductShowcase) UpdatePrice(nomen *Nomenclature) (bool, error) {
	// if current price will change to upper price, we must recalculate our minimum price for showcase
	edited := false
	tempTime := time.Now().UTC()
	for _, value := range s.Skus {
		if value.Uid == nomen.SkuId {
			for _, city := range nomen.Cities {
				if _, ok := value.MinPrice[city]; ok {
					if value.MinPrice[city] > nomen.Price {
						value.MinPrice[city] = nomen.Price
						value.MinNomenId[city] = nomen.Uid
						edited = true
						value.UpdatedOn = tempTime
					}
				} else {
					value.MinPrice[city] = nomen.Price
					value.MinNomenId[city] = nomen.Uid
					edited = true
					value.UpdatedOn = tempTime
				}
			}
		}
	}

	if edited {
		s.UpdatedOn = tempTime
	}

	return edited, nil
}

type ProductShowcaseSku struct {
	Uid      string             `json:"uid,omitempty"`
	Medias   []Media            `json:"media,omitempty"`
	MinPrice map[string]float64 `json:"min_price"` //минимальная цена по этому городу
	//IsVisible    map[string]bool    `json:"is_visible"` //есть в этом городе
	GSKU            []string             `json:"gsku"` //коды от завода производителя
	ProductParam    []ProductParams      `json:"product_param,omitempty"`
	CreatedOn       time.Time            `json:"created_on,omitempty"`
	UpdatedOn       time.Time            `json:"updated_on,omitempty"`
	Label           []string             `json:"label,omitempty"`
	Status          string               `json:"status,omitempty"`
	MinNomenId      map[string]string    `json:"min_nomen_id"`
	SkuNumber       int64                `json:"sku_number,omitempty"`
	Characteristics []CharacteristicBody `json:"characs"`
}

type Nomenclature struct {
	//обязательные поля
	Uid        string   `json:"uid"`
	MerchantId string   `json:"merchant_id,omitempty"` // ID мерчанта
	Name       string   `json:"name,omitempty"`        // имя товара скопированное из справочника
	Unit       string   `json:"unit,omitempty"`        // шт,кг
	Amount     float64  `json:"amount"`                // аггрегированное количество товара
	Price      float64  `json:"price"`                 // цена продажи
	StockId    []string `json:"stock_id,omitempty"`    // ID склада
	ScaleId    string   `json:"scale_id,omitempty"`    //ID в весах
	//не обязательные поля
	//GSKU         string          `json:"gsku,omitempty"`          // глобальный sku
	Articul                string              `json:"articul,omitempty"`       //артикуль
	Status                 string              `json:"status,omitempty"`        // статус номенклатуры
	GroupId                []string            `json:"group_id,omitempty"`      // группа продукта
	BarCode                []string            `json:"bar_code,omitempty"`      // штрих код для продажи
	SkuId                  string              `json:"sku_id,omitempty"`        // ID sku
	ProductId              string              `json:"product_id,omitempty"`    // ID продукта
	Medias                 []Media             `json:"media,omitempty"`         // ссылка на картинку видео
	CreatedOn              time.Time           `json:"created_on,omitempty"`    // дата создания
	UpdatedOn              time.Time           `json:"updated_on,omitempty"`    // дата обновления
	Scope                  []string            `json:"scope,omitempty"`         // обязательное
	SaleChannels           []string            `json:"sale_channels,omitempty"` // каналы продаж
	IsVisible              bool                `json:"is_visible"`              // активный
	IsBlocked              *bool               `json:"is_blocked,omitempty"`
	Cities                 []string            `json:"cities,omitempty"`         // города в котором будут продаваться
	InStock                map[string]bool     `json:"in_stock,omitempty"`       // в каких городах есть
	Available              bool                `json:"available"`                // если false, мне не показываем
	OldPrice               float64             `json:"old_price,omitempty"`      // старая цена
	SalesPrice             float64             `json:"sales_price,omitempty"`    // скидка
	PurchasePrice          float64             `json:"purchase_price,omitempty"` //цена закупки
	Margin                 float64             `json:"margin,omitempty"`         //Маржа
	PickupOptions          []XMLPickupOption   `json:"pickup_options,omitempty"`
	DeliveryOptions        []XMLDeliveryOption `json:"delivery_options,omitempty"`
	CityPrices             []XMLCityPrice      `json:"city_prices,omitempty"`
	Vendor                 string              `json:"vendor,omitempty"`
	CityPriceMap           map[string]float64  `json:"city_price_map,omitempty"`
	Comment                string              `json:"comment,omitempty"`
	Reason                 string              `json:"reason,omitempty"`
	ApplicationDescription string              `json:"application_description,omitempty"`
	ApplicationLink        string              `json:"application_link,omitempty"`
	CharacteristicInfo     string              `json:"characteristic_info,omitempty"`

	ExtraInfo       string   `json:"extra_info,omitempty"`
	CategoriesArray []string `json:"categories_array,omitempty"`
	NameEBT         string   `json:"name_ebt,omitempty"`
}

func (nomenclature *Nomenclature) CreateFromOffer(xmlOffer XMLOffer, merchantId string, uid string, salesChannel []string) error {
	nomenclature.PickupOptions = xmlOffer.PickupOptions
	nomenclature.DeliveryOptions = xmlOffer.DeliveryOptions
	nomenclature.CityPrices = xmlOffer.CityPrices
	nomenclature.Price = xmlOffer.Price
	nomenclature.Name = xmlOffer.Name
	nomenclature.Vendor = xmlOffer.Vendor
	nomenclature.BarCode = xmlOffer.Barcodes //todo
	nomenclature.UpdatedOn = time.Now()
	nomenclature.CreatedOn = nomenclature.UpdatedOn
	nomenclature.MerchantId = merchantId
	nomenclature.Articul = xmlOffer.Sku
	nomenclature.Status = "created"
	nomenclature.Uid = uid
	nomenclature.Available = true
	nomenclature.SaleChannels = salesChannel
	return nil
}

func (nomenclature *Nomenclature) CreateFromXls(sku, brand, model, merchantId, uid string, price float64, salesChannel []string, xlsPo []XMLPickupOption) error {
	nomenclature.PickupOptions = xlsPo
	nomenclature.Price = price
	nomenclature.Name = model
	nomenclature.Vendor = brand
	nomenclature.UpdatedOn = time.Now()
	nomenclature.CreatedOn = nomenclature.UpdatedOn
	nomenclature.MerchantId = merchantId
	nomenclature.Articul = sku
	nomenclature.Status = "created"
	nomenclature.Uid = uid
	nomenclature.Available = true
	nomenclature.SaleChannels = salesChannel
	return nil
}

type GoodsDirectory struct {
	SKU         string `json:"sku,omitempty"`
	Name        string `json:"name,omitempty"`
	Description string `json:"description,omitempty"`
	Category    string `json:"category,omitempty"`
	Brand       string `json:"brand,omitempty"`
}
