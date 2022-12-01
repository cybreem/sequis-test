package repositories

import (
	"github.com/cybreem/sequis-test/app"
	"github.com/cybreem/sequis-test/app/domains"
	"github.com/cybreem/sequis-test/app/domains/models"
	"github.com/cybreem/sequis-test/app/domains/types"
	sdkUtils "github.com/h4lim/go-sdk/utils"
)

type FriendRequestsRepositoryContext struct {
}

func NewFriendRequestsRepository() domains.FriendRequestsRepository {
	return FriendRequestsRepositoryContext{}
}

func (r FriendRequestsRepositoryContext) GetFriendRequests(param types.FriendRequestList) *[]models.FriendRequests {
	db, err := sdkUtils.DBModel.DBOpen()
	defer db.Close()
	if err != nil {
		app.Log.Errorf(app.LogID, "Error occurred %s ", err)
		return nil
	}
	var data []models.FriendRequests
	q := db.Where(models.FriendRequests{
		IsValid: 1,
		Email:   param.Email,
	})
	q.Debug().Find(&data)

	if len(data) == 0 {
		return nil
	}

	return &data
}

func (r FriendRequestsRepositoryContext) SearchFriendRequests(model models.FriendRequests) *[]models.FriendRequests {
	db, err := sdkUtils.DBModel.DBOpen()
	defer db.Close()
	if err != nil {
		app.Log.Errorf(app.LogID, "Error occurred %s ", err)
		return nil
	}
	var data []models.FriendRequests
	model.Status = "pending"
	q := db.Where(model)
	model.Status = "accepted"
	q = q.Or(model)
	model.Status = "accepted"
	Requestor := model.Requestor
	Email := model.Email
	model.Email = Requestor
	model.Requestor = Email
	q = q.Or(model)
	model.Status = "pending"
	q = q.Or(model)
	q.Debug().Find(&data)

	if len(data) == 0 {
		return nil
	}

	return &data
}

func (r FriendRequestsRepositoryContext) Save(data models.FriendRequests) *models.FriendRequests {
	db, err := sdkUtils.DBModel.DBOpen()
	defer db.Close()
	if err != nil {
		app.Log.Errorf(app.LogID, "Error occurred %s ", err)
		return nil
	}
	db.Create(&data)
	return &data
}

func (r FriendRequestsRepositoryContext) Unvalidate(data models.FriendRequests) *models.FriendRequests {
	db, err := sdkUtils.DBModel.DBOpen()
	defer db.Close()
	if err != nil {
		app.Log.Errorf(app.LogID, "Error occurred %s ", err)
		return nil
	}
	data.IsValid = 1
	update := db.Debug().Model(&data).Where(data).Update("is_valid", 0)
	if update.Error == nil && update.RowsAffected > 0 {
		return &data
	} else {
		return nil
	}
}

// func Paginate(page int, pageSize int) func(db *gorm.DB) *gorm.DB {
// 	return func(db *gorm.DB) *gorm.DB {
// 		if page == 0 {
// 			page = 1
// 		}
// 		switch {
// 		case pageSize > 100:
// 			pageSize = 100
// 		case pageSize <= 0:
// 			pageSize = 10
// 		}

// 		offset := (page - 1) * pageSize
// 		return db.Offset(offset).Limit(pageSize)
// 	}
// }
