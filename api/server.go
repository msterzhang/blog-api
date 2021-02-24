package api

import (
	"fmt"
	"github.com/gin-contrib/cors"
	"github.com/gin-gonic/gin"
	"blog-api/api/controllers"
	"blog-api/auto"
	"blog-api/config"
	"time"
)

func init() {
	auto.Load()
}


func Run() {
	// Disable Console Color, you don't need console color when writing the logs to file.
	gin.DisableConsoleColor()

	// Logging to a file.
	//f, _ := os.Create("gin.log")
	//gin.DefaultWriter = io.MultiWriter(f)
	//gin.SetMode(gin.ReleaseMode)

	//系统初始化
	r := gin.Default()
	r.Use(cors.New(cors.Config{
		AllowOriginFunc:  func(origin string) bool { return true },
		AllowOrigins:     []string{"http://localhost:8002", "http://syrme.top"},
		AllowMethods:     []string{"GET", "POST", "PUT", "DELETE", "PATCH"},
		AllowHeaders:     []string{"Origin", "Content-Length", "Content-Type"},
		AllowCredentials: true,
		MaxAge:           12 * time.Hour,
	}))

	//网络测试
	r.GET("/ping", func(c *gin.Context) {
		c.JSON(200, gin.H{
			"message": "pong",
		})
	})
	
    user := r.Group("/v1/api/user")
    user.POST("/create", controllers.CreateUser)
    user.POST("/update", controllers.UpdateUserById)
    user.POST("/delete", controllers.DeleteUserById)
    user.POST("/id", controllers.GetUserById)
    user.POST("/list", controllers.GetUserList)
    
    post := r.Group("/v1/api/post")
    post.POST("/create", controllers.CreatePost)
    post.POST("/update", controllers.UpdatePostById)
    post.POST("/delete", controllers.DeletePostById)
    post.POST("/id", controllers.GetPostById)
    post.POST("/list", controllers.GetPostList)
    
    tag := r.Group("/v1/api/tag")
    tag.POST("/create", controllers.CreateTag)
    tag.POST("/update", controllers.UpdateTagById)
    tag.POST("/delete", controllers.DeleteTagById)
    tag.POST("/id", controllers.GetTagById)
    tag.POST("/list", controllers.GetTagList)
    
    comment := r.Group("/v1/api/comment")
    comment.POST("/create", controllers.CreateComment)
    comment.POST("/update", controllers.UpdateCommentById)
    comment.POST("/delete", controllers.DeleteCommentById)
    comment.POST("/id", controllers.GetCommentById)
    comment.POST("/list", controllers.GetCommentList)
    
    link := r.Group("/v1/api/link")
    link.POST("/create", controllers.CreateLink)
    link.POST("/update", controllers.UpdateLinkById)
    link.POST("/delete", controllers.DeleteLinkById)
    link.POST("/id", controllers.GetLinkById)
    link.POST("/list", controllers.GetLinkList)
    
	r.Run(fmt.Sprintf(":%d", config.PORT))

}
