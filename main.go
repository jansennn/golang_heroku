package main

import (
	"github.com/gin-contrib/cors"
	"github.com/gin-gonic/gin"
	"github.com/jansennn/golang_heroku/config"
	v1 "github.com/jansennn/golang_heroku/handler/v1"
	"github.com/jansennn/golang_heroku/middleware"
	"github.com/jansennn/golang_heroku/repo"
	"github.com/jansennn/golang_heroku/service"
	"gorm.io/gorm"
)

var (
	db             		*gorm.DB               		= config.SetupDatabaseConnection()
	userRepo       		repo.UserRepository    		= repo.NewUserRepo(db)
	productRepo    		repo.ProductRepository 		= repo.NewProductRepo(db)
	descriptionRepo  	repo.DescriptionRepository  = repo.NewDescriptionRepo(db)
	projectRepo 		repo.ProjectRepository		= repo.NewProjectRepo(db)
	careerRepo 			repo.CareerRepository		= repo.NewCareerRepo(db)
	authService    		service.AuthService    		= service.NewAuthService(userRepo)
	jwtService     		service.JWTService     		= service.NewJWTService()
	userService    		service.UserService    		= service.NewUserService(userRepo)
	productService 		service.ProductService 		= service.NewProductService(productRepo)
	descriptionService  service.DescriptionService  = service.NewDescriptionService(descriptionRepo)
	projectService 		service.ProjectService		= service.NewProjectService(projectRepo)
	careerService 		service.CareerService		= service.NewCareerService(careerRepo)
	authHandler    		v1.AuthHandler         		= v1.NewAuthHandler(authService, jwtService, userService)
	userHandler    		v1.UserHandler         		= v1.NewUserHandler(userService, jwtService)
	productHandler 		v1.ProductHandler      		= v1.NewProductHandler(productService, jwtService)
	descriptionHandler  v1.DescriptionHandler  		= v1.NewDescriptionHandler(descriptionService, jwtService)
	projectHandler		v1.ProjectHandler	 		= v1.NewProjectHandler(projectService)
	careerHandler		v1.CareerHandler			= v1.NewCareerHandler(careerService)
)

func main() {
	defer config.CloseDatabaseConnection(db)
	server := gin.Default()
	config := cors.DefaultConfig()
	config.AllowAllOrigins = true
	server.Use(cors.New(config))

	server.GET("/api/description-public/:id", descriptionHandler.FindOneDescriptionById)
	server.GET("/api/career-public/", careerHandler.All)
	authRoutes := server.Group("api/auth")
	{
		authRoutes.POST("/login", authHandler.Login)
		authRoutes.POST("/register", authHandler.Register)
	}

	userRoutes := server.Group("api/user", middleware.AuthorizeJWT(jwtService))
	{
		userRoutes.GET("/profile", userHandler.Profile)
		userRoutes.PUT("/profile", userHandler.Update)
	}

	productRoutes := server.Group("api/product", middleware.AuthorizeJWT(jwtService))
	{
		productRoutes.GET("/", productHandler.All)
		productRoutes.POST("/", productHandler.CreateProduct)
		productRoutes.GET("/:id", productHandler.FindOneProductByID)
		productRoutes.PUT("/:id", productHandler.UpdateProduct)
		productRoutes.DELETE("/:id", productHandler.DeleteProduct)
	}

	descriptionRoutes := server.Group("api/description")
	{
		//descriptionRoutes.GET("/:id", descriptionHandler.FindOneDescriptionById)
		descriptionRoutes.POST("/", descriptionHandler.CreateDescription)
		descriptionRoutes.PUT("/:id", descriptionHandler.UpdateDescription)
	}

	projectRoutes := server.Group("api/project")
	{
		projectRoutes.GET("/", projectHandler.All)
		projectRoutes.POST("/", projectHandler.CreateProject)
		projectRoutes.GET("/:id", projectHandler.FindOneProjectById)
		projectRoutes.PUT("/:id", projectHandler.UpdateProject)
	}

	careerRoutes := server.Group("api/career")
	{
		careerRoutes.GET("/", careerHandler.All)
		careerRoutes.POST("/", careerHandler.CreateCareer)
		careerRoutes.GET("/:id", careerHandler.FindOneCareerById)
		careerRoutes.PUT("/:id", careerHandler.UpdateCareer)
	}

	checkRoutes := server.Group("api/check")
	{
		checkRoutes.GET("health", v1.Health)
	}

	server.Run()
}
