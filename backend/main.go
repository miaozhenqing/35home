package main

import (
	"fmt"
	"log"

	"35home/config"
	"35home/handler"
	"35home/repository"
	"35home/service"

	"github.com/gin-gonic/gin"
)

func main() {
	// 加载配置
	cfg := config.LoadConfig()

	// 初始化数据库
	err := config.InitDatabase(cfg)
	if err != nil {
		log.Fatalf("Failed to initialize database: %v", err)
	}
	defer config.CloseDatabase()

	// 初始化各层
	userRepo := repository.NewUserRepository(config.DB)
	userService := service.NewUserService(userRepo)
	userHandler := handler.NewUserHandler(userService)

	voteRepo := repository.NewVoteRepository(config.DB)
	voteService := service.NewVoteService(voteRepo)
	voteHandler := handler.NewVoteHandler(voteService)

	// 创建Gin引擎
	r := gin.Default()

	// 跨域中间件
	r.Use(func(c *gin.Context) {
		c.Writer.Header().Set("Access-Control-Allow-Origin", "*")
		c.Writer.Header().Set("Access-Control-Allow-Methods", "GET, POST, PUT, DELETE, OPTIONS")
		c.Writer.Header().Set("Access-Control-Allow-Headers", "Content-Type, Authorization")

		if c.Request.Method == "OPTIONS" {
			c.AbortWithStatus(204)
			return
		}

		c.Next()
	})

	// 健康检查
	r.GET("/health", func(c *gin.Context) {
		c.JSON(200, gin.H{
			"status": "ok",
		})
	})

	// API路由组
	api := r.Group("/api")
	{
		// 用户相关路由
		user := api.Group("/user")
		{
			user.POST("/register", userHandler.Register)
			user.POST("/login", userHandler.Login)
		}

		// 投票相关路由
		vote := api.Group("/vote")
		{
			vote.POST("/submit", voteHandler.SubmitVote)
			vote.GET("/stats", voteHandler.GetVoteStats)
		}

		// 压力测评相关路由
		test := api.Group("/test")
		{
			test.GET("/questions", getTestQuestions)
			test.POST("/submit", submitTest)
		}

		// 内容相关路由
		content := api.Group("/content")
		{
			content.GET("/articles", getArticles)
			content.GET("/article/:id", getArticleById)
		}
	}

	// 启动服务器
	port := cfg.ServerPort
	fmt.Printf("Server running on http://localhost:%s\n", port)
	if err := r.Run(fmt.Sprintf(":%s", port)); err != nil {
		log.Fatalf("Failed to start server: %v", err)
	}
}

// 临时处理函数
func getTestQuestions(c *gin.Context) {
	c.JSON(200, gin.H{
		"message": "Get test questions endpoint",
	})
}

func submitTest(c *gin.Context) {
	c.JSON(200, gin.H{
		"message": "Submit test endpoint",
	})
}

func getArticles(c *gin.Context) {
	c.JSON(200, gin.H{
		"message": "Get articles endpoint",
	})
}

func getArticleById(c *gin.Context) {
	c.JSON(200, gin.H{
		"message": "Get article by id endpoint",
	})
}
