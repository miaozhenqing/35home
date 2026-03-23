package main

import (
	"fmt"
	"log"
	"net/http"
	"time"

	"35home/config"
	"35home/models"

	"github.com/gin-gonic/gin"
	"golang.org/x/crypto/bcrypt"
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
			user.POST("/register", register)
			user.POST("/login", login)
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

// 注册请求结构体
type RegisterRequest struct {
	Username   string `json:"username" binding:"required,min=2,max=50"`
	Email      string `json:"email" binding:"required,email"`
	Password   string `json:"password" binding:"required,min=6"`
	ConfirmPassword string `json:"confirm_password" binding:"required"`
	Gender     string `json:"gender"`
	Age        int    `json:"age"`
	Occupation string `json:"occupation"`
	Hobbies    string `json:"hobbies"`
}

// 登录请求结构体
type LoginRequest struct {
	Email    string `json:"email" binding:"required,email"`
	Password string `json:"password" binding:"required"`
}

// 注册处理函数
func register(c *gin.Context) {
	var req RegisterRequest
	
	// 绑定请求参数
	if err := c.ShouldBindJSON(&req); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"error": "Invalid request parameters",
			"details": err.Error(),
		})
		return
	}
	
	// 验证密码是否一致
	if req.Password != req.ConfirmPassword {
		c.JSON(http.StatusBadRequest, gin.H{
			"error": "Passwords do not match",
		})
		return
	}
	
	// 检查邮箱是否已存在
	var existingUser models.User
	if err := config.DB.Where("email = ?", req.Email).First(&existingUser).Error; err == nil {
		c.JSON(http.StatusConflict, gin.H{
			"error": "Email already registered",
		})
		return
	}
	
	// 加密密码
	hashedPassword, err := bcrypt.GenerateFromPassword([]byte(req.Password), bcrypt.DefaultCost)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{
			"error": "Failed to hash password",
		})
		return
	}
	
	// 创建用户
	user := models.User{
		Username:   req.Username,
		Email:      req.Email,
		Password:   string(hashedPassword),
		Gender:     req.Gender,
		Age:        req.Age,
		Occupation: req.Occupation,
		Hobbies:    req.Hobbies,
	}
	
	if err := config.DB.Create(&user).Error; err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{
			"error": "Failed to create user",
			"details": err.Error(),
		})
		return
	}
	
	// 返回成功响应
	c.JSON(http.StatusCreated, gin.H{
		"message": "User registered successfully",
		"user": gin.H{
			"id":         user.ID,
			"username":   user.Username,
			"email":      user.Email,
			"gender":     user.Gender,
			"age":        user.Age,
			"occupation": user.Occupation,
			"hobbies":    user.Hobbies,
			"created_at": user.CreatedAt,
		},
	})
}

// 登录处理函数
func login(c *gin.Context) {
	var req LoginRequest
	
	// 绑定请求参数
	if err := c.ShouldBindJSON(&req); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"error": "Invalid request parameters",
			"details": err.Error(),
		})
		return
	}
	
	// 查找用户
	var user models.User
	if err := config.DB.Where("email = ?", req.Email).First(&user).Error; err != nil {
		c.JSON(http.StatusUnauthorized, gin.H{
			"error": "Invalid email or password",
		})
		return
	}
	
	// 验证密码
	err := bcrypt.CompareHashAndPassword([]byte(user.Password), []byte(req.Password))
	if err != nil {
		c.JSON(http.StatusUnauthorized, gin.H{
			"error": "Invalid email or password",
		})
		return
	}
	
	// 更新最后登录时间
	user.LastLoginAt = time.Now()
	config.DB.Save(&user)
	
	// 返回成功响应
	c.JSON(http.StatusOK, gin.H{
		"message": "Login successful",
		"user": gin.H{
			"id":         user.ID,
			"username":   user.Username,
			"email":      user.Email,
			"gender":     user.Gender,
			"age":        user.Age,
			"occupation": user.Occupation,
			"hobbies":    user.Hobbies,
			"last_login": user.LastLoginAt,
		},
		"token": "dummy-token", // 实际项目中应该生成JWT token
	})
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
