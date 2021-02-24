package controllers

import (
	"blog-api/api/database"
	"blog-api/api/models"
	"blog-api/api/repository"
	"blog-api/api/repository/crud"
	"strconv"
	"github.com/gin-gonic/gin"
)


func CreateTag(c *gin.Context) {
	tag := models.Tag{}
	err := c.ShouldBind(&tag)
	if err != nil {
		c.JSON(200, gin.H{"code": 201, "msg": "创建失败,表单解析出错！", "data": tag})
		return
	}
	db := database.NewDb()
	repo := crud.NewRepositoryTagsCRUD(db)
	func(tagRepository repository.TagRepository) {
		tag, err := tagRepository.Save(tag)
		if err != nil {
			c.JSON(200, gin.H{"code": 201, "msg": "创建失败！", "data": tag})
			return
		}
		c.JSON(200, gin.H{"code": 200, "msg": "创建成功！", "data": tag})
	}(repo)
}


func DeleteTagById(c *gin.Context) {
	id := c.Query("id")
	db := database.NewDb()
	repo := crud.NewRepositoryTagsCRUD(db)
	func(tagRepository repository.TagRepository) {
		tag, err := tagRepository.DeleteByID(id)
		if err != nil {
			c.JSON(200, gin.H{"code": 201, "msg": "没有查询到资源！", "data": tag})
			return
		}
		c.JSON(200, gin.H{"code": 200, "msg": "删除资源成功！", "data": tag})
	}(repo)
}

func UpdateTagById(c *gin.Context) {
	id := c.Query("id")
	tag := models.Tag{}
	err := c.ShouldBind(&tag)
	if err != nil {
		c.JSON(200, gin.H{"code": 201, "msg": "创建失败,表单解析出错！", "data": tag})
		return
	}
	db := database.NewDb()
	repo := crud.NewRepositoryTagsCRUD(db)
	func(tagRepository repository.TagRepository) {
		tag, err := tagRepository.UpdateByID(id, tag)
		if err != nil {
			c.JSON(200, gin.H{"code": 201, "msg": "没有查询到资源！", "data": tag})
			return
		}
		c.JSON(200, gin.H{"code": 200, "msg": "更新资源成功！", "data": tag})
	}(repo)
}


func GetTagById(c *gin.Context) {
	id := c.Query("id")
	db := database.NewDb()
	repo := crud.NewRepositoryTagsCRUD(db)
	func(tagRepository repository.TagRepository) {
		tag, err := tagRepository.FindByID(id)
		if err != nil {
			c.JSON(200, gin.H{"code": 201, "msg": "没有查询到资源！", "data": tag})
			return
		}
		c.JSON(200, gin.H{"code": 200, "msg": "查询资源成功！", "data": tag})
	}(repo)
}


func GetTagList(c *gin.Context) {
	page, errPage := strconv.Atoi(c.Query("page"))
	size, errSize := strconv.Atoi(c.Query("size"))
	if errPage != nil {
		page = 1
	}
	if errSize != nil {
		size = 8
	}
	db := database.NewDb()
	repo := crud.NewRepositoryTagsCRUD(db)
	func(tagRepository repository.TagRepository) {
		tags, err := tagRepository.FindAll(page,size)
		if err != nil {
			c.JSON(200, gin.H{"code": 201, "msg": "没有查询到资源！", "data": tags})
			return
		}
		c.JSON(200, gin.H{"code": 200, "msg": "查询资源成功！", "data": tags})
	}(repo)

}

func SearchTag(c *gin.Context) {
	q:=c.Query("q")
	if len(q)==0{
		c.JSON(200, gin.H{"code": 201, "msg": "参数错误！", "data":""})
		return
	}
	db := database.NewDb()
	repo := crud.NewRepositoryTagsCRUD(db)
	func(tagRepository repository.TagRepository) {
		tags, err := tagRepository.Search(q)
		if err != nil {
			c.JSON(200, gin.H{"code": 201, "msg": "没有查询到资源！", "data": tags})
			return
		}
		c.JSON(200, gin.H{"code": 200, "msg": "查询资源成功！", "data": tags})
	}(repo)
}