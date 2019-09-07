package domain

import (
	"errors"
	"fmt"
	satori "github.com/satori/go.uuid"
	"strconv"
	"time"
)

type Product struct {
	ProductWeightNumber float64         `json:"product_weight_number,omitempty"`
	VendorCode          string          `json:"vendor_code,omitempty"`
	Brand               string          `json:"brand,omitempty"`
	TnVedCod            string          `json:"tn_ved_cod,omitempty"`
	Uid                 string          `json:"uid,omitempty"`
	Scope               []string        `json:"scope,omitempty"`       //обязательное
	MerchantID          string          `json:"merchant_id,omitempty"` //обязательное
	Name                string          `json:"name,omitempty"`        //обязательное
	Description         string          `json:"description,omitempty"`
	ShortDescription    string          `json:"short_description,omitempty"`
	Medias              []Media         `json:"media,omitempty"`
	GroupId             []string        `json:"group_id,omitempty"`
	CategoriesArray     []string        `json:"categories_array,omitempty"` ///нужен ли
	SeoTitle            string          `json:"seoTitle,omitempty"`
	SeoKeywords         string          `json:"seoKeywords,omitempty"`
	SeoH1               string          `json:"seoH1,omitempty"`
	SeoDscr             string          `json:"seoDscr,omitempty"`
	CreatedOn           time.Time       `json:"created_on,omitempty"`
	UpdatedOn           time.Time       `json:"updated_on,omitempty"`
	Cashe               string          `json:"cashe,omitempty"`
	Status              string          `json:"status,omitempty"`
	ProductPrice        float64         `json:"product_price,omitempty"`
	ProductOldPrice     float64         `json:"product_old_price,omitempty"`
	Weight              float64         `json:"weight,omitempty"`
	Dimensions          Dimensions      `json:"dimensions,omitempty"`
	IsVisible           bool            `json:"is_visible"`
	P_Type              string          `json:"p_type,omitempty"` //goods, services
	DefaultCategory     string          `json:"default_category,omitempty"`
	ProductParam        []ProductParams `json:"product_param,omitempty"`
	MPN                 string          `json:"mpn,omitempty"`
	Label               []string        `json:"label,omitempty"` //new top
	Skus                []ProductSku    `json:"skus"`
	AggsRating          float64         `json:"aggs_rating,omitempty"`
	AggsMerchantRating  float64         `json:"aggs_merchant_rating,omitempty"`
	Tag                 map[string]int  `json:"tag,omitempty"`
	PurchaseCount       int64           `json:"purchase_count,omitempty"`
	GlobalBarcode       []string        `json:"global_barcode,omitempty"`
	Popularity			float64			`json:"popularity,omitempty"`
}

type CharacteristicBody struct {
	CharacteristicUid string `json:"characs_uid"`
	CharacteristicValueId string `json:"characs_value_id"`
}
type Dimensions struct {
	Height float64 `json:"height,omitempty"`
	Length float64 `json:"length,omitempty"`
	Width  float64 `json:"width,omitempty"`
}

type ProductSku struct {
	Uid               string          `json:"uid,omitempty"`
	Medias            []Media         `json:"media,omitempty"`
	Price             float64         `json:"price"`
	OldPrice          float64         `json:"old_price,omitempty"`
	SavedPrice        float64         `json:"saved_price,omitempty"`
	Amount            int             `json:"amount"`
	IsVisible         bool            `json:"is_visible"`
	Metatags          []string        `json:"metatags,omitempty"`
	SKU               string          `json:"sku,omitempty"`
	ProductParam      []ProductParams `json:"product_param,omitempty"`
	CreatedOn         time.Time       `json:"created_on,omitempty"`
	UpdatedOn         time.Time       `json:"updated_on,omitempty"`
	Label             []string        `json:"label,omitempty"`
	Status            string          `json:"status,omitempty"`
	Nomenclature      string          `json:"nomenclature,omitempty"`
	PromoPriceExpDate *time.Time      `json:"promo_price_exp_date,omitempty"`
	ProductPromoPrice float64         `json:"product_promo_price,omitempty"`
	GSKU              string          `json:"gsku,omitempty"`
	SkuNumber		  int64           `json:"sku_number,omitempty"`
	Characteristics   []CharacteristicBody `json:"characs"`
	CharacteristicSet []string `json:"characteristic_set"`
}

type ProductParams struct {
	ParamsType       string  `json:"params_type,omitempty"`
	ParamsTitle      string  `json:"params_title,omitempty"`
	CodeValue        string  `json:"code_value,omitempty"`
	ParamsValueID    string  `json:"params_value_id,omitempty"`
	ParamsValueTitle string  `json:"params_value_title,omitempty"`
	Value            float64 `json:"value,omitempty"`
}

func (p *Product) UpdateSku(sku *ProductSku) (founded bool, err error) {

	addskuparams := make(map[string]bool)
	var found = false
	var updateSkuNumber int
	for i := range p.Skus {
		if p.Skus[i].Uid == sku.Uid {
			updateSkuNumber = i
			for _, skuParam := range p.Skus[i].ProductParam {
				addskuparams[skuParam.ParamsType+skuParam.ParamsValueID+strconv.FormatFloat(skuParam.Value, 'f', -1, 64)] = true
			}

			fmt.Println("in UpdateSku, sku found")
			if sku.Amount < 1 {
				fmt.Println("changing visibility to false")
				sku.IsVisible = false
			}
			found = true
			break
		}
	}

	if !found {
		return false, errors.New("Ресурс не найден")
	}

	//check params
	if p.Status != "ontagging" {
		skulen := len(addskuparams)
		for _, psku := range p.Skus {

			if skulen != len(psku.ProductParam) || psku.Uid == sku.Uid {
				continue
			}

			d := make(map[string]bool)
			for key, value := range addskuparams {
				d[key] = value
			}

			for _, skuParam := range psku.ProductParam {
				d[skuParam.ParamsType+skuParam.ParamsValueID+strconv.FormatFloat(skuParam.Value, 'f', -1, 64)] = true

			}
			if skulen == len(d) {
				return true, errors.New("SKU с такими параметрами существует uid=" + psku.Uid)
			}
		}
	}

	sku.CreatedOn = p.Skus[updateSkuNumber].CreatedOn
	sku.UpdatedOn = time.Now()
	p.UpdatedOn = sku.UpdatedOn
	p.Skus[updateSkuNumber] = *sku

	return false, nil
}

func (p *Product) AddSku(sku *ProductSku) (err error) {

	//check params
	addskuparams := make(map[string]bool)

	for _, skuParam := range sku.ProductParam {
		addskuparams[skuParam.ParamsType+skuParam.ParamsValueID+strconv.FormatFloat(skuParam.Value, 'f', -1, 64)] = true
	}

	skulen := len(addskuparams)

	for _, psku := range p.Skus {

		if skulen != len(psku.ProductParam) {
			continue
		}
		d := make(map[string]bool)
		for key, value := range addskuparams {
			d[key] = value
		}
		for _, skuParam := range psku.ProductParam {
			d[skuParam.ParamsType+skuParam.ParamsValueID+strconv.FormatFloat(skuParam.Value, 'f', -1, 64)] = true
		}
		if skulen == len(d) {
			return errors.New("SKU с такими параметрами существует uid=" + psku.Uid)
		}
	}

	sku.Uid = satori.NewV1().String()
	sku.CreatedOn = time.Now()
	sku.UpdatedOn = sku.CreatedOn
	sku.Status = "created"
	p.UpdatedOn = sku.CreatedOn

	p.Skus = append(p.Skus, *sku)
	return nil
}

func (p *Product) IsStatusTagging() (ontagging bool) {
	if p.Status == "ontagging" {
		return true
	}
	return false

}
