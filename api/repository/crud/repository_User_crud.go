package crud

import (
	"errors"
	"log"
	"blog-api/api/models"
	"blog-api/api/utils/channels"
	"github.com/jinzhu/gorm"
	"strconv"
)

// RepositoryUsersCRUD is the struct for the User CRUD
type RepositoryUsersCRUD struct {
	db *gorm.DB
}

// NewRepositoryUsersCRUD returns a new repository with DB connection
func NewRepositoryUsersCRUD(db *gorm.DB) *RepositoryUsersCRUD {
	return &RepositoryUsersCRUD{db}
}

// Save returns a new user created or an error
func (r *RepositoryUsersCRUD) Save(user models.User) (models.User, error) {
	var err error
	log.Println(user)
	done := make(chan bool)
	go func(ch chan<- bool) {
		defer close(ch)
		err = r.db.Debug().Model(&models.User{}).Create(&user).Error
		if err != nil {
			ch <- false
			return
		}
		ch <- true
	}(done)
	if channels.OK(done) {
		return user, nil
	}
	return models.User{}, err
}

// FindAll returns all the users from the DB
func (r *RepositoryUsersCRUD) FindAll(page int,size int) ([]models.User, error) {
	var err error
	users := []models.User{}
	done := make(chan bool)
	go func(ch chan<- bool) {
		defer close(ch)
		err = r.db.Debug().Model(&models.User{}).Find(&users).Limit(strconv.Itoa(size)).Offset(strconv.Itoa((page - 1) * size)).Order("-ID").Scan(&users).Error
		if err != nil {
			ch <- false
			return
		}
		ch <- true
	}(done)
	if channels.OK(done) {
		return users, nil
	}
	return nil, err
}

// FindByID returns an user from the DB
func (r *RepositoryUsersCRUD) FindByID(id string) (models.User, error) {
	var err error
	user := models.User{}
	done := make(chan bool)
	go func(ch chan<- bool) {
		defer close(ch)
		err = r.db.Debug().Model(&models.User{}).Where("id = ?", id).Take(&user).Error
		if err != nil {
			ch <- false
			return
		}
		ch <- true
	}(done)
	if channels.OK(done) {
		return user, nil
	}

	if gorm.IsRecordNotFoundError(err) {
		return models.User{}, errors.New("User Not Found")
	}
	return models.User{}, err
}

// Update updates an user from the DB
func (r *RepositoryUsersCRUD) UpdateByID(id string, user models.User) (int64, error) {
	var rs *gorm.DB
	done := make(chan bool)
	go func(ch chan<- bool) {
		defer close(ch)
		rs = r.db.Debug().Model(&models.User{}).Where("id = ?", id).Take(&models.User{})
		ch <- true
	}(done)

	if channels.OK(done) {
		if rs.Error != nil {
			return 0, rs.Error
		}

		return rs.RowsAffected, nil
	}
	return 0, rs.Error
}

// Delete removes an user from the DB
func (r *RepositoryUsersCRUD) DeleteByID(id string) (int64, error) {
	var rs *gorm.DB
	done := make(chan bool)
	go func(ch chan<- bool) {
		defer close(ch)
		rs = r.db.Debug().Model(&models.User{}).Where("id = ?", id).Take(&models.User{}).Delete(&models.User{})
		ch <- true
	}(done)

	if channels.OK(done) {
		if rs.Error != nil {
			return 0, rs.Error
		}

		return rs.RowsAffected, nil
	}
	return 0, rs.Error
}

// Search removes an user from the DB
func (r *RepositoryUsersCRUD) Search(q string) ([]models.User, error) {
	var err error
	users := []models.User{}
	done := make(chan bool)
	go func(ch chan<- bool) {
		defer close(ch)
		err = r.db.Debug().Model(&models.User{}).Where("key LIKE ?","%"+q+"%").Scan(&users).Error
		if err != nil {
			ch <- false
			return
		}
		ch <- true
	}(done)
	if channels.OK(done) {
		return users, nil
	}
	if gorm.IsRecordNotFoundError(err) {
		return []models.User{}, errors.New("没有找到！")
	}
	return []models.User{}, err
}
