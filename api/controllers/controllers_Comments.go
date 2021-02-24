package controllers

import (
	"blog-api/api/database"
	"blog-api/api/models"
	"blog-api/api/repository"
	"blog-api/api/repository/crud"
	"strconv"
	"github.com/gin-gonic/gin"
)


func CreateComment(c *gin.Context) {
	comment := models.Comment{}
	err := c.ShouldBind(&comment)
	if err != nil {
		c.JSON(200, gin.H{"code": 201, "msg": "创建失败,表单解析出错！", "data": comment})
		return
	}
	db := database.NewDb()
	repo := crud.NewRepositoryCommentsCRUD(db)
	func(commentRepository repository.CommentRepository) {
		comment, err := commentRepository.Save(comment)
		if err != nil {
			c.JSON(200, gin.H{"code": 201, "msg": "创建失败！", "data": comment})
			return
		}
		c.JSON(200, gin.H{"code": 200, "msg": "创建成功！", "data": comment})
	}(repo)
}


func DeleteCommentById(c *gin.Context) {
	id := c.Query("id")
	db := database.NewDb()
	repo := crud.NewRepositoryCommentsCRUD(db)
	func(commentRepository repository.CommentRepository) {
		comment, err := commentRepository.DeleteByID(id)
		if err != nil {
			c.JSON(200, gin.H{"code": 201, "msg": "没有查询到资源！", "data": comment})
			return
		}
		c.JSON(200, gin.H{"code": 200, "msg": "删除资源成功！", "data": comment})
	}(repo)
}

func UpdateCommentById(c *gin.Context) {
	id := c.Query("id")
	comment := models.Comment{}
	err := c.ShouldBind(&comment)
	if err != nil {
		c.JSON(200, gin.H{"code": 201, "msg": "创建失败,表单解析出错！", "data": comment})
		return
	}
	db := database.NewDb()
	repo := crud.NewRepositoryCommentsCRUD(db)
	func(commentRepository repository.CommentRepository) {
		comment, err := commentRepository.UpdateByID(id, comment)
		if err != nil {
			c.JSON(200, gin.H{"code": 201, "msg": "没有查询到资源！", "data": comment})
			return
		}
		c.JSON(200, gin.H{"code": 200, "msg": "更新资源成功！", "data": comment})
	}(repo)
}


func GetCommentById(c *gin.Context) {
	id := c.Query("id")
	db := database.NewDb()
	repo := crud.NewRepositoryCommentsCRUD(db)
	func(commentRepository repository.CommentRepository) {
		comment, err := commentRepository.FindByID(id)
		if err != nil {
			c.JSON(200, gin.H{"code": 201, "msg": "没有查询到资源！", "data": comment})
			return
		}
		c.JSON(200, gin.H{"code": 200, "msg": "查询资源成功！", "data": comment})
	}(repo)
}


func GetCommentList(c *gin.Context) {
	page, errPage := strconv.Atoi(c.Query("page"))
	size, errSize := strconv.Atoi(c.Query("size"))
	if errPage != nil {
		page = 1
	}
	if errSize != nil {
		size = 8
	}
	db := database.NewDb()
	repo := crud.NewRepositoryCommentsCRUD(db)
	func(commentRepository repository.CommentRepository) {
		comments, err := commentRepository.FindAll(page,size)
		if err != nil {
			c.JSON(200, gin.H{"code": 201, "msg": "没有查询到资源！", "data": comments})
			return
		}
		c.JSON(200, gin.H{"code": 200, "msg": "查询资源成功！", "data": comments})
	}(repo)

}

func SearchComment(c *gin.Context) {
	q:=c.Query("q")
	if len(q)==0{
		c.JSON(200, gin.H{"code": 201, "msg": "参数错误！", "data":""})
		return
	}
	db := database.NewDb()
	repo := crud.NewRepositoryCommentsCRUD(db)
	func(commentRepository repository.CommentRepository) {
		comments, err := commentRepository.Search(q)
		if err != nil {
			c.JSON(200, gin.H{"code": 201, "msg": "没有查询到资源！", "data": comments})
			return
		}
		c.JSON(200, gin.H{"code": 200, "msg": "查询资源成功！", "data": comments})
	}(repo)

}