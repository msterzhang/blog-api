package crud

import (
	"errors"
	"log"
	"blog-api/api/models"
	"blog-api/api/utils/channels"
	"github.com/jinzhu/gorm"
	"strconv"
)

// RepositoryCommentsCRUD is the struct for the Comment CRUD
type RepositoryCommentsCRUD struct {
	db *gorm.DB
}

// NewRepositoryCommentsCRUD returns a new repository with DB connection
func NewRepositoryCommentsCRUD(db *gorm.DB) *RepositoryCommentsCRUD {
	return &RepositoryCommentsCRUD{db}
}

// Save returns a new comment created or an error
func (r *RepositoryCommentsCRUD) Save(comment models.Comment) (models.Comment, error) {
	var err error
	log.Println(comment)
	done := make(chan bool)
	go func(ch chan<- bool) {
		defer close(ch)
		err = r.db.Debug().Model(&models.Comment{}).Create(&comment).Error
		if err != nil {
			ch <- false
			return
		}
		ch <- true
	}(done)
	if channels.OK(done) {
		return comment, nil
	}
	return models.Comment{}, err
}

// FindAll returns all the comments from the DB
func (r *RepositoryCommentsCRUD) FindAll(page int,size int) ([]models.Comment, error) {
	var err error
	comments := []models.Comment{}
	done := make(chan bool)
	go func(ch chan<- bool) {
		defer close(ch)
		err = r.db.Debug().Model(&models.Comment{}).Find(&comments).Limit(strconv.Itoa(size)).Offset(strconv.Itoa((page - 1) * size)).Order("-ID").Scan(&comments).Error
		if err != nil {
			ch <- false
			return
		}
		ch <- true
	}(done)
	if channels.OK(done) {
		return comments, nil
	}
	return nil, err
}

// FindByID returns an comment from the DB
func (r *RepositoryCommentsCRUD) FindByID(id string) (models.Comment, error) {
	var err error
	comment := models.Comment{}
	done := make(chan bool)
	go func(ch chan<- bool) {
		defer close(ch)
		err = r.db.Debug().Model(&models.Comment{}).Where("id = ?", id).Take(&comment).Error
		if err != nil {
			ch <- false
			return
		}
		ch <- true
	}(done)
	if channels.OK(done) {
		return comment, nil
	}

	if gorm.IsRecordNotFoundError(err) {
		return models.Comment{}, errors.New("Comment Not Found")
	}
	return models.Comment{}, err
}

// Update updates an comment from the DB
func (r *RepositoryCommentsCRUD) UpdateByID(id string, comment models.Comment) (int64, error) {
	var rs *gorm.DB
	done := make(chan bool)
	go func(ch chan<- bool) {
		defer close(ch)
		rs = r.db.Debug().Model(&models.Comment{}).Where("id = ?", id).Take(&models.Comment{})
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

// Delete removes an comment from the DB
func (r *RepositoryCommentsCRUD) DeleteByID(id string) (int64, error) {
	var rs *gorm.DB
	done := make(chan bool)
	go func(ch chan<- bool) {
		defer close(ch)
		rs = r.db.Debug().Model(&models.Comment{}).Where("id = ?", id).Take(&models.Comment{}).Delete(&models.Comment{})
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

// Search removes an comment from the DB
func (r *RepositoryCommentsCRUD) Search(q string) ([]models.Comment, error) {
	var err error
	comments := []models.Comment{}
	done := make(chan bool)
	go func(ch chan<- bool) {
		defer close(ch)
		err = r.db.Debug().Model(&models.Comment{}).Where("key LIKE ?","%"+q+"%").Scan(&comments).Error
		if err != nil {
			ch <- false
			return
		}
		ch <- true
	}(done)
	if channels.OK(done) {
		return comments, nil
	}
	if gorm.IsRecordNotFoundError(err) {
		return []models.Comment{}, errors.New("没有找到！")
	}
	return []models.Comment{}, err
}
