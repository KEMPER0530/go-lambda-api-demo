package controllers

import (
    "go-lambda-api-demo/src/interfaces/database"
    "go-lambda-api-demo/src/usecase"
)

type DemoController struct {
    Interactor usecase.DemoInteractor
}

func NewDemoController(db database.DB) *DemoController {
    return &DemoController{
        Interactor: usecase.DemoInteractor{
            DB: &database.DBRepository{ DB: db },
            Demo: &database.DemoRepository{},
        },
    }
}

func (controller *DemoController) Healthcheck(c Context) {
    Demo,res := controller.Interactor.Healthcheck()
    _ = Demo
    c.JSON(res.Code, res)
}
