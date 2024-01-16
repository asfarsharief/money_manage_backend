package payload

type WomenListingResponsePayload struct {
	APIResponse  string         `json:"api_response"`
	APIMsgCode   int            `json:"api_msg_code"`
	ProductMinis []ProductMinis `json:"product_minis"`
}

type ProductMinis struct {
	ProductID          string      `json:"product_id,omitempty"`
	ProductCatalogueID int         `json:"product_catalogue_id,omitempty"`
	ProductCategoryID  int         `json:"product_category_id"`
	ProductName        string      `json:"product_name"`
	Cost               float32     `json:"cost"`
	DiscountedPrize    float32     `json:"discounted_prize"`
	DiscountPercetage  string      `json:"discount_percetage"`
	ImageUrls          []ImageUrls `json:"image_urls"`
	ColorsAvailable    []string    `json:"colors_available,omitempty"`
	ColorsInStock      []string    `json:"colors_in_stock,omitempty"`
	SizeAvailable      []string    `json:"size_available,omitempty"`
	SizeInStock        []string    `json:"size_in_stock,omitempty"`
}

type ImageUrls struct {
	Prefrence int    `json:"prefrence"`
	ImageURL  string `json:"image_url"`
}

// https://i.ibb.co/3spR9Cc/accessories.jpg
// https://i.ibb.co/X7mTrc9/banner-image.jpg
// https://i.ibb.co/xCrPMSf/bottoms.jpg
// https://i.ibb.co/3s2S1vd/category-Left.jpg
// https://i.ibb.co/n73JdTd/category-Right.jpg
// https://i.ibb.co/FWk8xSj/crochet.jpg
// https://i.ibb.co/hMh9VNh/dress.webp
// https://i.ibb.co/qRtSwFm/kaftan.jpg
// https://i.ibb.co/RysvVVV/topImage.jpg
// https://i.ibb.co/x5CkZYL/Nykaa-Logo.png
