package usecase

import (
    "github.com/jinzhu/gorm"

    "go-lambda-api-demo/src/domain"
)

type DemoRepository interface {
    Find(db *gorm.DB) (demo domain.Demo, err error)
}
