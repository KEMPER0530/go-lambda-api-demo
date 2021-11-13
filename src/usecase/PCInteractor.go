package usecase

import (
    "log"
    "go-lambda-api-demo/src/domain"
    "github.com/guregu/dynamo"
)

type PCInteractor struct {
    PC PCRepository
}

func (interactor *PCInteractor) Healthcheck(d *dynamo.DB) (pc domain.ProductCatalogForGet, resultStatus ResultStatus) {
    table := d.Table("ProductCatalog")
    foundPC, err := interactor.PC.Find(&table)
    if err != nil {
        return domain.ProductCatalogForGet{}, NewResultStatus(500)
    }
    log.Println(table)
    pc = foundPC.BuildForGet()

    return pc, NewResultStatus(200)
}
