package domain

type BannerStruct struct {
	MerchantId 		string `json:"merchant_id,omitempty"`
	Scope 			string `json:"scope"`
	Placement 		string `json:"placement"`
	BannerUid 		string `json:"banner_uid"`
	AltDesc 		string `json:"alt_descr,omitempty"`
	BannerStatus 	*bool  `json:"banner_status,omitempty"`
	BannerType 		string `json:"banner_type,omitempty"`
	Caption 		string `json:"caption,omitempty"`
	Title 			string `json:"title,omitempty"`
	HeadTitle 		string `json:"head_title,omitempty"`
	SourceUrl 		string `json:"source_url,omitempty"`
	DirectUrl 		string `json:"direct_url,omitempty"`
	Sort 			*int   `json:"sort,omitempty"`
}

type filterBannerResponse struct {
	TotalHits int64          `json:"total_hits,omitempty"`
	Banners   []BannerStruct `json:"banners,omitempty"`
}