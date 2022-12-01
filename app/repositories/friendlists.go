package repositories

import (
	"github.com/cybreem/sequis-test/app"
	"github.com/cybreem/sequis-test/app/domains"
	"github.com/cybreem/sequis-test/app/domains/models"
	"github.com/cybreem/sequis-test/app/domains/types"
	sdkUtils "github.com/h4lim/go-sdk/utils"
)

type FriendlistRepositoryContext struct {
}

func NewFriendlistsRepository() domains.FriendlistsRepository {
	return FriendlistRepositoryContext{}
}

func (r FriendlistRepositoryContext) GetFriendlists(param types.FriendlistRequest) *[]models.Friendlists {
	db, err := sdkUtils.DBModel.DBOpen()
	defer db.Close()
	if err != nil {
		app.Log.Errorf(app.LogID, "Error occurred %s ", err)
		return nil
	}
	var data []models.Friendlists
	q := db.Where(models.Friendlists{
		IsValid: 1,
		User:    param.Email,
	}).Where("block=?", 0)
	q.Debug().Find(&data)

	if len(data) == 0 {
		return nil
	}

	return &data
}

func (r FriendlistRepositoryContext) GetMutuallists(param types.MutualRequest) *[]models.Friendlists {
	db, err := sdkUtils.DBModel.DBOpen()
	defer db.Close()
	if err != nil {
		app.Log.Errorf(app.LogID, "Error occurred %s ", err)
		return nil
	}
	var data []models.Friendlists
	q := db.Raw("select f1.user as friend from friendlists as f1 join friendlists as f2 using (user) where f1.friend = ? and f2.friend = ? and f1.is_valid=1 and f2.is_valid=1", param.Friends[0], param.Friends[1])
	q.Debug().Scan(&data)

	if len(data) == 0 {
		return nil
	}

	return &data
}

func (r FriendlistRepositoryContext) Save(data models.Friendlists) *models.Friendlists {
	db, err := sdkUtils.DBModel.DBOpen()
	defer db.Close()
	if err != nil {
		app.Log.Errorf(app.LogID, "Error occurred %s ", err)
		return nil
	}
	db.Create(&data)
	return &data
}

func (r FriendlistRepositoryContext) Block(data models.Friendlists) *models.Friendlists {
	db, err := sdkUtils.DBModel.DBOpen()
	defer db.Close()
	if err != nil {
		app.Log.Errorf(app.LogID, "Error occurred %s ", err)
		return nil
	}
	db.Create(&data)
	return &data
}

func (r FriendlistRepositoryContext) GetBlockedUser(model models.FriendRequests) *[]models.Friendlists {
	db, err := sdkUtils.DBModel.DBOpen()
	defer db.Close()
	if err != nil {
		app.Log.Errorf(app.LogID, "Error occurred %s ", err)
		return nil
	}
	var data []models.Friendlists
	q := db.Where(models.Friendlists{
		Block:   1,
		IsValid: 1,
		User:    model.Email,
		Friend:  model.Requestor,
	}).Or(models.Friendlists{
		Block:   1,
		IsValid: 1,
		User:    model.Requestor,
		Friend:  model.Email,
	})
	q.Debug().Find(&data)

	if len(data) == 0 {
		return nil
	}

	return &data
}

func (r FriendlistRepositoryContext) GetBlockedUserStrict(model models.FriendRequests) *[]models.Friendlists {
	db, err := sdkUtils.DBModel.DBOpen()
	defer db.Close()
	if err != nil {
		app.Log.Errorf(app.LogID, "Error occurred %s ", err)
		return nil
	}
	var data []models.Friendlists
	q := db.Where(models.Friendlists{
		Block:   1,
		IsValid: 1,
		User:    model.Requestor,
		Friend:  model.Email,
	})
	q.Debug().Find(&data)

	if len(data) == 0 {
		return nil
	}

	return &data
}

func (r FriendlistRepositoryContext) Unblock(data models.Friendlists) *models.Friendlists {
	db, err := sdkUtils.DBModel.DBOpen()
	defer db.Close()
	if err != nil {
		app.Log.Errorf(app.LogID, "Error occurred %s ", err)
		return nil
	}
	update := db.Debug().Model(&data).Where("user = ? and friend = ? and is_valid = 1 and block = 1", data.User, data.Friend).Update("is_valid", 0)
	if update.Error == nil && update.RowsAffected > 0 {
		return &data
	} else {
		return nil
	}
}

func (r FriendlistRepositoryContext) Unvalidate(data models.Friendlists) *models.Friendlists {
	db, err := sdkUtils.DBModel.DBOpen()
	defer db.Close()
	if err != nil {
		app.Log.Errorf(app.LogID, "Error occurred %s ", err)
		return nil
	}
	update := db.Debug().Model(&data).Where("(user = ? and friend = ? and is_valid = 1) or (user = ? and friend = ? and is_valid = 1)", data.User, data.Friend, data.Friend, data.User).Update("is_valid", 0)
	if update.Error == nil && update.RowsAffected > 0 {
		return &data
	} else {
		return nil
	}
}
