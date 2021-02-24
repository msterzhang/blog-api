package crud

import (
	"errors"
	"log"
	"blog-api/api/models"
	"blog-api/api/utils/channels"
	"github.com/jinzhu/gorm"
	"strconv"
)

// RepositoryLinksCRUD is the struct for the Link CRUD
type RepositoryLinksCRUD struct {
	db *gorm.DB
}

// NewRepositoryLinksCRUD returns a new repository with DB connection
func NewRepositoryLinksCRUD(db *gorm.DB) *RepositoryLinksCRUD {
	return &RepositoryLinksCRUD{db}
}

// Save returns a new link created or an error
func (r *RepositoryLinksCRUD) Save(link models.Link) (models.Link, error) {
	var err error
	log.Println(link)
	done := make(chan bool)
	go func(ch chan<- bool) {
		defer close(ch)
		err = r.db.Debug().Model(&models.Link{}).Create(&link).Error
		if err != nil {
			ch <- false
			return
		}
		ch <- true
	}(done)
	if channels.OK(done) {
		return link, nil
	}
	return models.Link{}, err
}

// FindAll returns all the links from the DB
func (r *RepositoryLinksCRUD) FindAll(page int,size int) ([]models.Link, error) {
	var err error
	links := []models.Link{}
	done := make(chan bool)
	go func(ch chan<- bool) {
		defer close(ch)
		err = r.db.Debug().Model(&models.Link{}).Find(&links).Limit(strconv.Itoa(size)).Offset(strconv.Itoa((page - 1) * size)).Order("-ID").Scan(&links).Error
		if err != nil {
			ch <- false
			return
		}
		ch <- true
	}(done)
	if channels.OK(done) {
		return links, nil
	}
	return nil, err
}

// FindByID returns an link from the DB
func (r *RepositoryLinksCRUD) FindByID(id string) (models.Link, error) {
	var err error
	link := models.Link{}
	done := make(chan bool)
	go func(ch chan<- bool) {
		defer close(ch)
		err = r.db.Debug().Model(&models.Link{}).Where("id = ?", id).Take(&link).Error
		if err != nil {
			ch <- false
			return
		}
		ch <- true
	}(done)
	if channels.OK(done) {
		return link, nil
	}

	if gorm.IsRecordNotFoundError(err) {
		return models.Link{}, errors.New("Link Not Found")
	}
	return models.Link{}, err
}

// Update updates an link from the DB
func (r *RepositoryLinksCRUD) UpdateByID(id string, link models.Link) (int64, error) {
	var rs *gorm.DB
	done := make(chan bool)
	go func(ch chan<- bool) {
		defer close(ch)
		rs = r.db.Debug().Model(&models.Link{}).Where("id = ?", id).Take(&models.Link{})
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

// Delete removes an link from the DB
func (r *RepositoryLinksCRUD) DeleteByID(id string) (int64, error) {
	var rs *gorm.DB
	done := make(chan bool)
	go func(ch chan<- bool) {
		defer close(ch)
		rs = r.db.Debug().Model(&models.Link{}).Where("id = ?", id).Take(&models.Link{}).Delete(&models.Link{})
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

// Search removes an link from the DB
func (r *RepositoryLinksCRUD) Search(q string) ([]models.Link, error) {
	var err error
	links := []models.Link{}
	done := make(chan bool)
	go func(ch chan<- bool) {
		defer close(ch)
		err = r.db.Debug().Model(&models.Link{}).Where("key LIKE ?","%"+q+"%").Scan(&links).Error
		if err != nil {
			ch <- false
			return
		}
		ch <- true
	}(done)
	if channels.OK(done) {
		return links, nil
	}
	if gorm.IsRecordNotFoundError(err) {
		return []models.Link{}, errors.New("没有找到！")
	}
	return []models.Link{}, err
}
