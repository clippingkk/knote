package userdb

import "github.com/clippingkk/knote/model"

func (db *DB) FindUser(id int64) (*model.User, error) {
	user := &model.User{}
	has, err := db.Id(id).Get(user)

	if err != nil {
		return nil, err
	}

	if !has {
		return nil, model.ErrNotFound
	}

	return user, nil
}