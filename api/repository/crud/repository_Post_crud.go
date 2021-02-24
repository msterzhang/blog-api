package crud

import (
	"errors"
	"log"
	"blog-api/api/models"
	"blog-api/api/utils/channels"
	"github.com/jinzhu/gorm"
	"strconv"
)

// RepositoryPostsCRUD is the struct for the Post CRUD
type RepositoryPostsCRUD struct {
	db *gorm.DB
}

// NewRepositoryPostsCRUD returns a new repository with DB connection
func NewRepositoryPostsCRUD(db *gorm.DB) *RepositoryPostsCRUD {
	return &RepositoryPostsCRUD{db}
}

// Save returns a new post created or an error
func (r *RepositoryPostsCRUD) Save(post models.Post) (models.Post, error) {
	var err error
	log.Println(post)
	done := make(chan bool)
	go func(ch chan<- bool) {
		defer close(ch)
		err = r.db.Debug().Model(&models.Post{}).Create(&post).Error
		if err != nil {
			ch <- false
			return
		}
		ch <- true
	}(done)
	if channels.OK(done) {
		return post, nil
	}
	return models.Post{}, err
}

// FindAll returns all the posts from the DB
func (r *RepositoryPostsCRUD) FindAll(page int,size int) ([]models.Post, error) {
	var err error
	posts := []models.Post{}
	done := make(chan bool)
	go func(ch chan<- bool) {
		defer close(ch)
		err = r.db.Debug().Model(&models.Post{}).Find(&posts).Limit(strconv.Itoa(size)).Offset(strconv.Itoa((page - 1) * size)).Order("-ID").Scan(&posts).Error
		if err != nil {
			ch <- false
			return
		}
		ch <- true
	}(done)
	if channels.OK(done) {
		return posts, nil
	}
	return nil, err
}

// FindByID returns an post from the DB
func (r *RepositoryPostsCRUD) FindByID(id string) (models.Post, error) {
	var err error
	post := models.Post{}
	done := make(chan bool)
	go func(ch chan<- bool) {
		defer close(ch)
		err = r.db.Debug().Model(&models.Post{}).Where("id = ?", id).Take(&post).Error
		if err != nil {
			ch <- false
			return
		}
		ch <- true
	}(done)
	if channels.OK(done) {
		return post, nil
	}

	if gorm.IsRecordNotFoundError(err) {
		return models.Post{}, errors.New("Post Not Found")
	}
	return models.Post{}, err
}

// Update updates an post from the DB
func (r *RepositoryPostsCRUD) UpdateByID(id string, post models.Post) (int64, error) {
	var rs *gorm.DB
	done := make(chan bool)
	go func(ch chan<- bool) {
		defer close(ch)
		rs = r.db.Debug().Model(&models.Post{}).Where("id = ?", id).Take(&models.Post{})
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

// Delete removes an post from the DB
func (r *RepositoryPostsCRUD) DeleteByID(id string) (int64, error) {
	var rs *gorm.DB
	done := make(chan bool)
	go func(ch chan<- bool) {
		defer close(ch)
		rs = r.db.Debug().Model(&models.Post{}).Where("id = ?", id).Take(&models.Post{}).Delete(&models.Post{})
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

// Search removes an post from the DB
func (r *RepositoryPostsCRUD) Search(q string) ([]models.Post, error) {
	var err error
	posts := []models.Post{}
	done := make(chan bool)
	go func(ch chan<- bool) {
		defer close(ch)
		err = r.db.Debug().Model(&models.Post{}).Where("key LIKE ?","%"+q+"%").Scan(&posts).Error
		if err != nil {
			ch <- false
			return
		}
		ch <- true
	}(done)
	if channels.OK(done) {
		return posts, nil
	}
	if gorm.IsRecordNotFoundError(err) {
		return []models.Post{}, errors.New("没有找到！")
	}
	return []models.Post{}, err
}
