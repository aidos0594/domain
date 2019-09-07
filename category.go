package domain

import (
	"time"
)

type MarketingCategory struct {
	Uid                string   `json:"uid"`
	IncludedCategories []string `json:"included_categories"`
	ExcludedCategories []string `json:"excluded_categories"`
	Type               string   `json:"type"`

	//pair of prices
	//[]Pair<int, int>
	PriceRanges []PriceRange `json:"price_range"` // ["0-500", "10000-50000"]
	Rules       [][]string   `json:"rules"`

	Products           []string    `json:"products"`     //ids array of products to include
	ContentType        string      `json:"content_type"` //dynamic & static & directory
	VersionId          string      `json:"version_id"`
	Names              []Languages `json:"names,omitempty"`
	IsActive           bool        `json:"is_active"`
	DisplayBanner      bool        `json:"display_banner"`
	DisplayDescription bool        `json:"display_description"`
	Description        string      `json:"description"`
	IconUrl            string      `json:"icon_url"`
	ProductDisplayType string      `json:"product_display_type"`
	DefaultDisplayType string      `json:"default_display_type"`
	SortingTypes       string      `json:"sorting_types"`
	DefaultSorting     string      `json:"default_sorting"`
	DefaultPaging      int64       `json:"default_paging"`
	ParentId           string      `json:"parent_id"`

	IsCopied    bool   `json:"is_copied"`
	Path        string `json:"path,omitempty"`
	SeoTitle    string `json:"seo_title"`
	SeoKeywords string `json:"seo_keywords"`
	SeoH1       string `json:"seo_h_1"`
	SeoDscr     string `json:"seo_dscr"`
	Sort        int    `json:"sort"`
	Depth       int    `json:"depth"`

	MenuKey string `json:"menu_key"` //mobile && desktop

	//Marketing category
	Childs       []*MarketingCategory `json:"childs,omitempty"`
}

type PriceRange struct {
	From int `json:"from"`
	To   int `json:"to"`
}

//todo rename
type CatalogVersion struct {
	Uid           string          `json:"uid"`
	Name          string          `json:"name"`
	Status        string          `json:"status"`
	Whitelabels   []string        `json:"whitelabels"`
	CategoryCount int64           `json:"category_count"`
	CreatedOn     time.Time       `json:"created_on"`
	UpdatedOn     *time.Time      `json:"updated_on,omitempty"`
	ChangeHistory []ChangeHistory `json:"change_history"`
	IsActive      bool            `json:"is_active"`
}

type ChangeHistory struct {
	FIO        string    `json:"fio"`
	ProfileId  string    `json:"profile_id"`
	Action     string    `json:"action"`
	ActionTime time.Time `json:"action_time,omitempty"`
}

func (c *MarketingCategory) GetParentId() string {
	return (*c).ParentId
}

func (c *MarketingCategory) GetId() string {
	return (*c).Uid
}

func (c *MarketingCategory) SaveChilds(category TreeInterface) {
	c.Childs = append(c.Childs, category.(*MarketingCategory))
}