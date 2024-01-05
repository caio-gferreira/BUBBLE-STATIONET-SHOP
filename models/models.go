package models

type Product struct {
	Sku          int     `json:sku`
	Amount       int     `json:amount`
	Price_unit   float32 `json:price_unit`
	Product_Name string  `json:product_name`
}

type Order struct {
	Client      Client    `json:client_name`
	Products    []Product `json:products`
	Order_Id    int       `json:order_id`
	Total_Value int       `json:total_value`
}

type Client struct {
	Name      string `json:name`
	Phone     string `json:phone`
	Client_Id int    `json:client_id`
}
