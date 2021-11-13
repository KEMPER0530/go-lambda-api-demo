package database

import (
    "errors"
    "github.com/guregu/dynamo"
    "go-lambda-api-demo/src/domain"
    "log"
    "fmt"
)

type PCRepository struct {}

func (repo *PCRepository) Find(table *dynamo.Table) (pc domain.ProductCatalog, err error) {
    var results []domain.ProductCatalog
    // use placeholders in filter expressions (see Expressions section below)
    err = table.Scan().All(&results)
    if (err != nil)  {
        log.Println(err)
        return pc, errors.New(fmt.Sprintf("%T\n", err))
    }
    // dynamoDBの項目表示
    log.Println(results)
    pc = results[0]
    return pc, nil
}
