package crud

import (
	"errors"
	"log"
	"blog-api/api/models"
	"blog-api/api/utils/channels"
	"github.com/jinzhu/gorm"
	"strconv"
)

// RepositoryTagsCRUD is the struct for the Tag CRUD
type RepositoryTagsCRUD struct {
	db *gorm.DB
}

// NewRepositoryTagsCRUD returns a new repository with DB connection
func NewRepositoryTagsCRUD(db *gorm.DB) *RepositoryTagsCRUD {
	return &RepositoryTagsCRUD{db}
}

// Save returns a new tag created or an error
func (r *RepositoryTagsCRUD) Save(tag models.Tag) (models.Tag, error) {
	var err error
	log.Println(tag)
	done := make(chan bool)
	go func(ch chan<- bool) {
		defer close(ch)
		err = r.db.Debug().Model(&models.Tag{}).Create(&tag).Error
		if err != nil {
			ch <- false
			return
		}
		ch <- true
	}(done)
	if channels.OK(done) {
		return tag, nil
	}
	return models.Tag{}, err
}

// FindAll returns all the tags from the DB
func (r *RepositoryTagsCRUD) FindAll(page int,size int) ([]models.Tag, error) {
	var err error
	tags := []models.Tag{}
	done := make(chan bool)
	go func(ch chan<- bool) {
		defer close(ch)
		err = r.db.Debug().Model(&models.Tag{}).Find(&tags).Limit(strconv.Itoa(size)).Offset(strconv.Itoa((page - 1) * size)).Order("-ID").Scan(&tags).Error
		if err != nil {
			ch <- false
			return
		}
		ch <- true
	}(done)
	if channels.OK(done) {
		return tags, nil
	}
	return nil, err
}

// FindByID returns an tag from the DB
func (r *RepositoryTagsCRUD) FindByID(id string) (models.Tag, error) {
	var err error
	tag := models.Tag{}
	done := make(chan bool)
	go func(ch chan<- bool) {
		defer close(ch)
		err = r.db.Debug().Model(&models.Tag{}).Where("id = ?", id).Take(&tag).Error
		if err != nil {
			ch <- false
			return
		}
		ch <- true
	}(done)
	if channels.OK(done) {
		return tag, nil
	}

	if gorm.IsRecordNotFoundError(err) {
		return models.Tag{}, errors.New("Tag Not Found")
	}
	return models.Tag{}, err
}

// Update updates an tag from the DB
func (r *RepositoryTagsCRUD) UpdateByID(id string, tag models.Tag) (int64, error) {
	var rs *gorm.DB
	done := make(chan bool)
	go func(ch chan<- bool) {
		defer close(ch)
		rs = r.db.Debug().Model(&models.Tag{}).Where("id = ?", id).Take(&models.Tag{})
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

// Delete removes an tag from the DB
func (r *RepositoryTagsCRUD) DeleteByID(id string) (int64, error) {
	var rs *gorm.DB
	done := make(chan bool)
	go func(ch chan<- bool) {
		defer close(ch)
		rs = r.db.Debug().Model(&models.Tag{}).Where("id = ?", id).Take(&models.Tag{}).Delete(&models.Tag{})
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

// Search removes an tag from the DB
func (r *RepositoryTagsCRUD) Search(q string) ([]models.Tag, error) {
	var err error
	tags := []models.Tag{}
	done := make(chan bool)
	go func(ch chan<- bool) {
		defer close(ch)
		err = r.db.Debug().Model(&models.Tag{}).Where("key LIKE ?","%"+q+"%").Scan(&tags).Error
		if err != nil {
			ch <- false
			return
		}
		ch <- true
	}(done)
	if channels.OK(done) {
		return tags, nil
	}
	if gorm.IsRecordNotFoundError(err) {
		return []models.Tag{}, errors.New("没有找到！")
	}
	return []models.Tag{}, err
}
