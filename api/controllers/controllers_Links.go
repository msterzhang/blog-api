package controllers

import (
	"blog-api/api/database"
	"blog-api/api/models"
	"blog-api/api/repository"
	"blog-api/api/repository/crud"
	"strconv"
	"github.com/gin-gonic/gin"
)


func CreateLink(c *gin.Context) {
	link := models.Link{}
	err := c.ShouldBind(&link)
	if err != nil {
		c.JSON(200, gin.H{"code": 201, "msg": "创建失败,表单解析出错！", "data": link})
		return
	}
	db := database.NewDb()
	repo := crud.NewRepositoryLinksCRUD(db)
	func(linkRepository repository.LinkRepository) {
		link, err := linkRepository.Save(link)
		if err != nil {
			c.JSON(200, gin.H{"code": 201, "msg": "创建失败！", "data": link})
			return
		}
		c.JSON(200, gin.H{"code": 200, "msg": "创建成功！", "data": link})
	}(repo)
}


func DeleteLinkById(c *gin.Context) {
	id := c.Query("id")
	db := database.NewDb()
	repo := crud.NewRepositoryLinksCRUD(db)
	func(linkRepository repository.LinkRepository) {
		link, err := linkRepository.DeleteByID(id)
		if err != nil {
			c.JSON(200, gin.H{"code": 201, "msg": "没有查询到资源！", "data": link})
			return
		}
		c.JSON(200, gin.H{"code": 200, "msg": "删除资源成功！", "data": link})
	}(repo)
}

func UpdateLinkById(c *gin.Context) {
	id := c.Query("id")
	link := models.Link{}
	err := c.ShouldBind(&link)
	if err != nil {
		c.JSON(200, gin.H{"code": 201, "msg": "创建失败,表单解析出错！", "data": link})
		return
	}
	db := database.NewDb()
	repo := crud.NewRepositoryLinksCRUD(db)
	func(linkRepository repository.LinkRepository) {
		link, err := linkRepository.UpdateByID(id, link)
		if err != nil {
			c.JSON(200, gin.H{"code": 201, "msg": "没有查询到资源！", "data": link})
			return
		}
		c.JSON(200, gin.H{"code": 200, "msg": "更新资源成功！", "data": link})
	}(repo)
}


func GetLinkById(c *gin.Context) {
	id := c.Query("id")
	db := database.NewDb()
	repo := crud.NewRepositoryLinksCRUD(db)
	func(linkRepository repository.LinkRepository) {
		link, err := linkRepository.FindByID(id)
		if err != nil {
			c.JSON(200, gin.H{"code": 201, "msg": "没有查询到资源！", "data": link})
			return
		}
		c.JSON(200, gin.H{"code": 200, "msg": "查询资源成功！", "data": link})
	}(repo)
}


func GetLinkList(c *gin.Context) {
	page, errPage := strconv.Atoi(c.Query("page"))
	size, errSize := strconv.Atoi(c.Query("size"))
	if errPage != nil {
		page = 1
	}
	if errSize != nil {
		size = 8
	}
	db := database.NewDb()
	repo := crud.NewRepositoryLinksCRUD(db)
	func(linkRepository repository.LinkRepository) {
		links, err := linkRepository.FindAll(page,size)
		if err != nil {
			c.JSON(200, gin.H{"code": 201, "msg": "没有查询到资源！", "data": links})
			return
		}
		c.JSON(200, gin.H{"code": 200, "msg": "查询资源成功！", "data": links})
	}(repo)

}

func SearchLink(c *gin.Context) {
	q:=c.Query("q")
	if len(q)==0{
		c.JSON(200, gin.H{"code": 201, "msg": "参数错误！", "data":""})
		return
	}
	db := database.NewDb()
	repo := crud.NewRepositoryLinksCRUD(db)
	func(linkRepository repository.LinkRepository) {
		links, err := linkRepository.Search(q)
		if err != nil {
			c.JSON(200, gin.H{"code": 201, "msg": "没有查询到资源！", "data": links})
			return
		}
		c.JSON(200, gin.H{"code": 200, "msg": "查询资源成功！", "data": links})
	}(repo)

}