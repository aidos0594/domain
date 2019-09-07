package domain

type Catalog struct {
	Scope         []string `json:"scope,omitempty"`
	Path          string   `json:"path,omitempty"`
	ParentId      string   `json:"parentId,omitempty"`
	Title         string   `json:"title,omitempty"`
	SeoTitle      string   `json:"seoTitle,omitempty"`
	SeoKeywords   string   `json:"seoKeywords,omitempty"`
	SeoH1         string   `json:"seoH1,omitempty"`
	SeoDscr       string   `json:"seoDscr,omitempty"`
	Sort          int      `json:"sort,omitempty"`
	Dscr          string   `json:"dscr,omitempty"`
	Depth         int      `json:"depth,omitempty"`
	Token         string   `json:"token,omitempty"`
	OnMain        bool     `json:"onMain,omitempty"`
	IsVisible     bool     `json:"isVisible"`
	IsMarketing   bool     `json:"isMarketing,omitempty"`
	Medias        []Media  `json:"media,omitempty"`
	BizCategories []string `json:"bizCategories,omitempty"`
}

type CategoryCached struct {
	Id     string            `json:"id"`
	Title  string            `json:"nm"`
	Slug   string            `json:"slug,omitempty"`
	Icon   []Media           `json:"media,omitempty"`
	Childs []*CategoryCached `json:"ch,omitempty"`
}
