package domain

type Demo struct {
    Id int
    Title string
}

type DemoForGet struct {
	Id       int    `gorm:"primary_key;not null"     json:"id"`
	Title    string `gorm:"type:varchar(2000);null"   json:"title"`
}

func (d *Demo) BuildForGet() DemoForGet {
		demo := DemoForGet{}
    demo.Id = d.Id
    demo.Title = d.Title
    return demo
}
