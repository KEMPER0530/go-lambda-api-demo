package database

import (
    "errors"
    "github.com/jinzhu/gorm"
    "go-lambda-api-demo/src/domain"
)

type DemoRepository struct {}

func (repo *DemoRepository) Find(db *gorm.DB) (demo domain.Demo, err error) {
    demo = domain.Demo{}
    db.Order("id asc").Find(&demo)
    if demo.Id <= 0 {
        return domain.Demo{}, errors.New("demo is not found")
    }
    return demo, nil
}
