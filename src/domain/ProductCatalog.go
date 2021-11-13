package domain

type ProductCatalog struct {
    Id int
    Price int
    Title string
}

type ProductCatalogForGet struct {
	Id       int    `dynamo:"Id"     json:"id"`
	Price    int    `dynamo:"Price"  json:"price"`
	Title    string `dynamo:"Title"   json:"title"`
}

func (p *ProductCatalog) BuildForGet() ProductCatalogForGet {
		ProductCatalog := ProductCatalogForGet{}
    ProductCatalog.Id = p.Id
    ProductCatalog.Price = p.Price
    ProductCatalog.Title = p.Title
    return ProductCatalog
}
