package controllers

import (
	"blog-api/api/database"
	"blog-api/api/models"
	"blog-api/api/repository"
	"blog-api/api/repository/crud"
	"strconv"
	"github.com/gin-gonic/gin"
)


func CreateUser(c *gin.Context) {
	user := models.User{}
	err := c.ShouldBind(&user)
	if err != nil {
		c.JSON(200, gin.H{"code": 201, "msg": "创建失败,表单解析出错！", "data": user})
		return
	}
	db := database.NewDb()
	repo := crud.NewRepositoryUsersCRUD(db)
	func(userRepository repository.UserRepository) {
		user, err := userRepository.Save(user)
		if err != nil {
			c.JSON(200, gin.H{"code": 201, "msg": "创建失败！", "data": user})
			return
		}
		c.JSON(200, gin.H{"code": 200, "msg": "创建成功！", "data": user})
	}(repo)
}


func DeleteUserById(c *gin.Context) {
	id := c.Query("id")
	db := database.NewDb()
	repo := crud.NewRepositoryUsersCRUD(db)
	func(userRepository repository.UserRepository) {
		user, err := userRepository.DeleteByID(id)
		if err != nil {
			c.JSON(200, gin.H{"code": 201, "msg": "没有查询到资源！", "data": user})
			return
		}
		c.JSON(200, gin.H{"code": 200, "msg": "删除资源成功！", "data": user})
	}(repo)
}

func UpdateUserById(c *gin.Context) {
	id := c.Query("id")
	user := models.User{}
	err := c.ShouldBind(&user)
	if err != nil {
		c.JSON(200, gin.H{"code": 201, "msg": "创建失败,表单解析出错！", "data": user})
		return
	}
	db := database.NewDb()
	repo := crud.NewRepositoryUsersCRUD(db)
	func(userRepository repository.UserRepository) {
		user, err := userRepository.UpdateByID(id, user)
		if err != nil {
			c.JSON(200, gin.H{"code": 201, "msg": "没有查询到资源！", "data": user})
			return
		}
		c.JSON(200, gin.H{"code": 200, "msg": "更新资源成功！", "data": user})
	}(repo)
}


func GetUserById(c *gin.Context) {
	id := c.Query("id")
	db := database.NewDb()
	repo := crud.NewRepositoryUsersCRUD(db)
	func(userRepository repository.UserRepository) {
		user, err := userRepository.FindByID(id)
		if err != nil {
			c.JSON(200, gin.H{"code": 201, "msg": "没有查询到资源！", "data": user})
			return
		}
		c.JSON(200, gin.H{"code": 200, "msg": "查询资源成功！", "data": user})
	}(repo)
}


func GetUserList(c *gin.Context) {
	page, errPage := strconv.Atoi(c.Query("page"))
	size, errSize := strconv.Atoi(c.Query("size"))
	if errPage != nil {
		page = 1
	}
	if errSize != nil {
		size = 8
	}
	db := database.NewDb()
	repo := crud.NewRepositoryUsersCRUD(db)
	func(userRepository repository.UserRepository) {
		users, err := userRepository.FindAll(page,size)
		if err != nil {
			c.JSON(200, gin.H{"code": 201, "msg": "没有查询到资源！", "data": users})
			return
		}
		c.JSON(200, gin.H{"code": 200, "msg": "查询资源成功！", "data": users})
	}(repo)

}

func SearchUser(c *gin.Context) {
	q:=c.Query("q")
	if len(q)==0{
		c.JSON(200, gin.H{"code": 201, "msg": "参数错误！", "data":""})
		return
	}
	db := database.NewDb()
	repo := crud.NewRepositoryUsersCRUD(db)
	func(userRepository repository.UserRepository) {
		users, err := userRepository.Search(q)
		if err != nil {
			c.JSON(200, gin.H{"code": 201, "msg": "没有查询到资源！", "data": users})
			return
		}
		c.JSON(200, gin.H{"code": 200, "msg": "查询资源成功！", "data": users})
	}(repo)

}