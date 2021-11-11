package usecase

import (
    "go-lambda-api-demo/src/domain"
    "reflect"
)

type DemoInteractor struct {
    DB DBRepository
    Demo DemoRepository
}

func (interactor *DemoInteractor) Healthcheck() (demo domain.DemoForGet, resultStatus ResultStatus) {
    db := interactor.DB.Connect()
    if (db == nil) || reflect.ValueOf(db).IsNil()  {
      return demo, NewResultStatus(500)
    }
    foundDemo, err := interactor.Demo.Find(db)
    if err != nil {
        return domain.DemoForGet{}, NewResultStatus(404)
    }
    demo = foundDemo.BuildForGet()

    return demo, NewResultStatus(200)
}
