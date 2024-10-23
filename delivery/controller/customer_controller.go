package controller

import (
	"go-enigma-laundry/middleware"
	"go-enigma-laundry/model"
	"go-enigma-laundry/model/dto/response"
	"go-enigma-laundry/usecase"
	"net/http"
	"strconv"

	"github.com/gin-gonic/gin"
	"github.com/google/uuid"
)

type CustomerController struct {
	uc usecase.CustomerUsecase
	rg *gin.RouterGroup
	authMiddleware middleware.AuthMiddleware
}

func (cc *CustomerController) registerHandler(ctx *gin.Context) {
	var newCustomer model.Customer	
	if err := ctx.ShouldBindJSON(&newCustomer); err != nil {
		response.SendSingleResponseError(
			ctx,
			http.StatusBadRequest,
			err.Error(),
		)
		return
	}    
	newCustomer.Id = uuid.NewString();
	err := cc.uc.RegisterCustomer(newCustomer)
	if err != nil {
		response.SendSingleResponseError(
			ctx,
			http.StatusBadRequest,
			err.Error(),
		)
		return
	}
	response.SendSingleResponseCreated(
		ctx,
		newCustomer,
		"Success Register new Customer",
	)
}

func (cc *CustomerController) findAllPageHandler(ctx *gin.Context){
	page,_ := strconv.Atoi(ctx.Query("page"))
	size,_ := strconv.Atoi(ctx.Query("size"))
	customers,paging,err := cc.uc.FindAllPaging(page,size)
	if err != nil {
		response.SendSingleResponseError(
			ctx,
			http.StatusBadRequest,
			err.Error(),
		)
	}
	var data []any
	data = append(data, customers)
	response.SendSinglePageResponse(
		ctx,
		data,
		"Success Get List Customer",
		paging,
	)
}

func (cc *CustomerController) findByIdHandler(ctx *gin.Context){
	id := ctx.Param("id")
	customer,err := cc.uc.FindCustomerById(id)
	if err != nil {
		ctx.JSON(http.StatusBadRequest,gin.H{
			"message" : err.Error(),
		})
		return
	}
	ctx.JSON(http.StatusCreated,gin.H{
		"message" : "Success Get Customer By Id",
		"data" : customer,
	})	
}


func (cc *CustomerController) deleteByIdHandler(ctx *gin.Context){
	id := ctx.Param("id")
	err := cc.uc.DeleteCustomerById(id)
	if err != nil {
		ctx.JSON(http.StatusBadRequest,gin.H{
			"message" : err.Error(),
		})
		return
	}
	ctx.JSON(http.StatusCreated,gin.H{
		"message" : "Success Delete Customer By Id",		
	})	
}

func (cc *CustomerController) updateHandler(ctx *gin.Context){	
	var customer model.Customer
	if err := ctx.ShouldBindJSON(&customer); err != nil {
		ctx.JSON(http.StatusBadRequest,gin.H{
			"message" : err.Error(),
		})
		return
	}
	err := cc.uc.UpdateCustomer(customer)
	if err != nil {
		ctx.JSON(http.StatusBadRequest,gin.H{
			"message" : err.Error(),
		})
		return
	}
	ctx.JSON(http.StatusCreated,gin.H{
		"message" : "Success Update Customer By Id",
		"data" : customer,		
	})	
}

func (cc *CustomerController) Route(){
	router := cc.rg.Group("/customers")
	router.Use(cc.authMiddleware.RequireToken("ADMIN"))
	router.POST("",cc.registerHandler)
	router.GET("",cc.findAllPageHandler)
	router.GET("/:id",cc.findByIdHandler)
	router.DELETE("/:id",cc.deleteByIdHandler)
	router.PUT("",cc.updateHandler)
}

func NewCustomerController(
	uc usecase.CustomerUsecase,
	router *gin.Engine,
	authMiddleware middleware.AuthMiddleware,
	) *CustomerController {
	return &CustomerController{
		uc: uc,
		rg : &router.RouterGroup,
		authMiddleware: authMiddleware,
	}
}