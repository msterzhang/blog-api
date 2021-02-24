package controllers

import (
	"blog-api/api/database"
	"blog-api/api/models"
	"blog-api/api/repository"
	"blog-api/api/repository/crud"
	"strconv"
	"github.com/gin-gonic/gin"
)


func CreatePost(c *gin.Context) {
	post := models.Post{}
	err := c.ShouldBind(&post)
	if err != nil {
		c.JSON(200, gin.H{"code": 201, "msg": "创建失败,表单解析出错！", "data": post})
		return
	}
	db := database.NewDb()
	repo := crud.NewRepositoryPostsCRUD(db)
	func(postRepository repository.PostRepository) {
		post, err := postRepository.Save(post)
		if err != nil {
			c.JSON(200, gin.H{"code": 201, "msg": "创建失败！", "data": post})
			return
		}
		c.JSON(200, gin.H{"code": 200, "msg": "创建成功！", "data": post})
	}(repo)
}


func DeletePostById(c *gin.Context) {
	id := c.Query("id")
	db := database.NewDb()
	repo := crud.NewRepositoryPostsCRUD(db)
	func(postRepository repository.PostRepository) {
		post, err := postRepository.DeleteByID(id)
		if err != nil {
			c.JSON(200, gin.H{"code": 201, "msg": "没有查询到资源！", "data": post})
			return
		}
		c.JSON(200, gin.H{"code": 200, "msg": "删除资源成功！", "data": post})
	}(repo)
}

func UpdatePostById(c *gin.Context) {
	id := c.Query("id")
	post := models.Post{}
	err := c.ShouldBind(&post)
	if err != nil {
		c.JSON(200, gin.H{"code": 201, "msg": "创建失败,表单解析出错！", "data": post})
		return
	}
	db := database.NewDb()
	repo := crud.NewRepositoryPostsCRUD(db)
	func(postRepository repository.PostRepository) {
		post, err := postRepository.UpdateByID(id, post)
		if err != nil {
			c.JSON(200, gin.H{"code": 201, "msg": "没有查询到资源！", "data": post})
			return
		}
		c.JSON(200, gin.H{"code": 200, "msg": "更新资源成功！", "data": post})
	}(repo)
}


func GetPostById(c *gin.Context) {
	id := c.Query("id")
	db := database.NewDb()
	repo := crud.NewRepositoryPostsCRUD(db)
	func(postRepository repository.PostRepository) {
		post, err := postRepository.FindByID(id)
		if err != nil {
			c.JSON(200, gin.H{"code": 201, "msg": "没有查询到资源！", "data": post})
			return
		}
		c.JSON(200, gin.H{"code": 200, "msg": "查询资源成功！", "data": post})
	}(repo)
}


func GetPostList(c *gin.Context) {
	page, errPage := strconv.Atoi(c.Query("page"))
	size, errSize := strconv.Atoi(c.Query("size"))
	if errPage != nil {
		page = 1
	}
	if errSize != nil {
		size = 8
	}
	db := database.NewDb()
	repo := crud.NewRepositoryPostsCRUD(db)
	func(postRepository repository.PostRepository) {
		posts, err := postRepository.FindAll(page,size)
		if err != nil {
			c.JSON(200, gin.H{"code": 201, "msg": "没有查询到资源！", "data": posts})
			return
		}
		c.JSON(200, gin.H{"code": 200, "msg": "查询资源成功！", "data": posts})
	}(repo)

}

func SearchPost(c *gin.Context) {
	q:=c.Query("q")
	if len(q)==0{
		c.JSON(200, gin.H{"code": 201, "msg": "参数错误！", "data":""})
		return
	}
	db := database.NewDb()
	repo := crud.NewRepositoryPostsCRUD(db)
	func(postRepository repository.PostRepository) {
		posts, err := postRepository.Search(q)
		if err != nil {
			c.JSON(200, gin.H{"code": 201, "msg": "没有查询到资源！", "data": posts})
			return
		}
		c.JSON(200, gin.H{"code": 200, "msg": "查询资源成功！", "data": posts})
	}(repo)

}