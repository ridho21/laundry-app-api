package controller

import (
	"go-enigma-laundry/model"
	"go-enigma-laundry/usecase"
	"net/http"
	"github.com/gin-gonic/gin"
)


type ServiceController struct {
	uc usecase.ServiceUsecase
	rg *gin.RouterGroup
}

func (cc *ServiceController) registerHandler(ctx *gin.Context) {
	var newService model.Service
	if err := ctx.ShouldBindJSON(&newService); err != nil {
		ctx.JSON(http.StatusBadRequest,gin.H{
			"message" : err.Error(),
		})
		return
	}
	service,err := cc.uc.RegisterService(newService)
	if err != nil {
		ctx.JSON(http.StatusBadRequest,gin.H{
			"message" : err.Error(),
		})
		return
	}
	ctx.JSON(http.StatusCreated,gin.H{
		"message" : "Success Register new Service",
		"data" : service,
	})	
}


func (cc *ServiceController) findAllHandler(ctx *gin.Context){
	customers,err := cc.uc.FindAllService()
	if err != nil {
		ctx.JSON(http.StatusBadRequest,gin.H{
			"message" : err.Error(),
		})
		return
	}
	ctx.JSON(http.StatusCreated,gin.H{
		"message" : "Success Find All Service",
		"data" : customers,
	})	
}

func (cc *ServiceController) Route(){
	router := cc.rg.Group("/services")
	router.POST("",cc.registerHandler)	
	router.GET("",cc.findAllHandler)	
}

func NewServiceController(uc usecase.ServiceUsecase,router *gin.Engine) *ServiceController {
	return &ServiceController{
		uc: uc,
		rg : &router.RouterGroup,
	}
}