package controller

import (
	"go-enigma-laundry/model/dto/request"
	"go-enigma-laundry/usecase"
	"net/http"

	"github.com/gin-gonic/gin"
)


type TransactionController struct {
	uc usecase.TransactionUsecase
	rg *gin.RouterGroup
}

func (tc *TransactionController) createTrxHandler(ctx *gin.Context) {
	var newTrx request.TransactionRequest
	if err := ctx.ShouldBindJSON(&newTrx); err != nil {
		ctx.JSON(http.StatusBadRequest,gin.H{
			"message" : err.Error(),
		})
		return
	}
	trxCreated,err := tc.uc.CreateTransaction(newTrx)
	if err != nil {
		ctx.JSON(http.StatusBadRequest,gin.H{
			"message" : err.Error(),
		})
		return
	}
	ctx.JSON(http.StatusCreated,gin.H{
		"message" : "Success Create Transaction",
		"data" : trxCreated,
	})	
}

func (tc *TransactionController) getListTrxHandler(ctx *gin.Context) {
	trxCreated,err := tc.uc.FindAllTrx()
	if err != nil {
		ctx.JSON(http.StatusInternalServerError,gin.H{
			"message" : err.Error(),
		})
		return
	}
	ctx.JSON(http.StatusCreated,gin.H{
		"message" : "Success get list Transaction",
		"data" : trxCreated,
	})	
}



func (tc *TransactionController) Route(){
	router := tc.rg.Group("/transactions")
	router.POST("",tc.createTrxHandler)	
	router.GET("",tc.getListTrxHandler)	
}

func NewTransactionController(uc usecase.TransactionUsecase,router *gin.Engine) *TransactionController {
	return &TransactionController{
		uc: uc,
		rg : &router.RouterGroup,
	}
}