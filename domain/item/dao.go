package item

type Item struct {
	ID           int         `json:"id"`
	Seller       int         `json:"seller"`
	Title        string      `json:"title"`
	Description  Description `json:"description"`
	Price        float32     `json:"price"`
	Images       []Image     `json:"images"`
	Video        string      `json:"video"`
	AvailableQty int         `json:"available_qty"`
	SoldQty      int         `json:"sold_qty"`
	Status       string      `json:"status"`
}

type Description struct {
	PlainText string `json:"plain_text"`
	HTML      string `json:"html"`
}

type Image struct {
	ID  int    `json:"id"`
	URL string `json:"url"`
}
