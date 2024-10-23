package delivery

import (
	"go-enigma-laundry/config"
	"go-enigma-laundry/delivery/controller"
	"go-enigma-laundry/middleware"
	"go-enigma-laundry/repository"
	"go-enigma-laundry/usecase"
	"go-enigma-laundry/utils/common"
	"log"

	"github.com/gin-gonic/gin"
)

type Server struct {
	authMiddlware middleware.AuthMiddleware
	customerUc usecase.CustomerUsecase
	serviceUc usecase.ServiceUsecase
	transactionUc usecase.TransactionUsecase
	userUc usecase.UserUsecase
	engine     *gin.Engine
	host       string
}

func (s *Server) setupControllers() {
	controller.NewCustomerController(s.customerUc,s.engine,s.authMiddlware).Route()
	controller.NewServiceController(s.serviceUc,s.engine).Route()
	controller.NewTransactionController(s.transactionUc,s.engine).Route()
	controller.NewUserController(s.userUc,s.engine).Route()
}

func (s *Server) Run(){
	s.setupControllers()
	if err := s.engine.Run(s.host); err != nil {
		log.Fatal("Server Error : ", err.Error())
	}
}

func NewServer() *Server {
	cfg,err := config.NewConfig()

	if err != nil {
		log.Fatal("Config : ",err.Error())
	}

	db,err := config.NewDbConnection(cfg)
	
	if err != nil {
		log.Fatal("DB Connect : ",err.Error())
	}

	
	// Token Service
	jwtToken := common.NewJwtToken(cfg.TokenConfig)

	//Middleware
	authMiddleware := middleware.NewAuthMiddleware(jwtToken)

	// Customer
	customerRepo := repository.NewCustomerRepository(db.Conn())
	customerUc := usecase.NewCustomerUsecase(customerRepo)

	// Service
	serviceRepo := repository.NewServiceRepository(db.Conn())
	serivceUc := usecase.NewServiceUsecase(serviceRepo)

	// Transaction
	trxRepo := repository.NewTransactionRepository(db.Conn())
	trxUc := usecase.NewTransactionUsecase(trxRepo,customerUc,serivceUc)

	// User
	userRepo := repository.NewUserRepository(db.Conn())
	userUc := usecase.NewUserUsecase(userRepo,jwtToken)

	// Gin Engine
	engine := gin.Default()

	return &Server{
		authMiddlware: authMiddleware,
		customerUc: customerUc,
		serviceUc: serivceUc,
		transactionUc: trxUc,
		userUc: userUc,
		engine: engine,
		host: ":8085",
	}
}