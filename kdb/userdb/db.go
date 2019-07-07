package userdb

import (
	"github.com/clippingkk/knote/model"
	"github.com/go-xorm/xorm"
)

type ReaderWriter interface {
	FindUser(id int64) (*model.User, error)
}

type DB struct {
	*xorm.Engine
}

func NewReaderWriter(db *xorm.Engine) ReaderWriter {
	return &DB{
		Engine: db,
	}
}
