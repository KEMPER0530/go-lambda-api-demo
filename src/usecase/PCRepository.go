package usecase

import (
    "github.com/guregu/dynamo"
    "go-lambda-api-demo/src/domain"
)

type PCRepository interface {
    Find(table *dynamo.Table) (demo domain.ProductCatalog, err error)
}
