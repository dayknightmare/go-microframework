package dummyRepository

import (
	"context"
	"encoding/json"
	"go-skeleton/application/domain/dummy"
	"go-skeleton/pkg/database"
	"time"
)

var ctx = context.Background()

type DummyRepository struct {
	Mysql *database.MySql
	Redis *database.Redis
}

func NewBaseRepository() *DummyRepository {
	return &DummyRepository{}
}

func (repo *DummyRepository) Get(id string) (dummy.Dummy, error) {
	var Data dummy.Dummy
	repo.Mysql.Db.First(&Data, "dummy_id = ?", id)
	return Data, nil
}

func (repo *DummyRepository) Create(d *dummy.Dummy) bool {
	repo.Redis.Cache.Del(ctx, "dummy_list").Result()

	repo.Mysql.Db.Create(d)
	return true
}

func (repo *DummyRepository) List() []dummy.Dummy {
	res, err := repo.Redis.Cache.Get(ctx, "dummy_list").Result()

	if res == "" || err != nil {
		var data []dummy.Dummy

		repo.Mysql.Db.Limit(100).Find(&data)
		out, _ := json.Marshal(data)
		repo.Redis.Cache.Set(ctx, "dummy_list", string(out), time.Minute*10)

		return data
	}

	var result []dummy.Dummy

	json.Unmarshal([]byte(res), &result)

	return result
}

func (repo *DummyRepository) Edit(d *dummy.Dummy) bool {
	repo.Redis.Cache.Del(ctx, "dummy_list").Result()

	repo.Mysql.Db.Updates(d)
	return true
}

func (repo *DummyRepository) Delete(d *dummy.Dummy) bool {
	repo.Redis.Cache.Del(ctx, "dummy_list").Result()

	repo.Mysql.Db.Delete(d)
	return true
}
