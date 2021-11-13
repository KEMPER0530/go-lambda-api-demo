package controllers

import (
    "go-lambda-api-demo/src/interfaces/database"
    "go-lambda-api-demo/src/usecase"
    "github.com/guregu/dynamo"
)

type PCController struct {
    Interactor usecase.PCInteractor
}

func NewPCController() *PCController {
    return &PCController{
        Interactor: usecase.PCInteractor{
            PC: &database.PCRepository{},
        },
    }
}

func (controller *PCController) Healthcheck(c Context, d *dynamo.DB) {
    PC,res := controller.Interactor.Healthcheck(d)
    _ = PC
    c.JSON(res.Code, res)
}
