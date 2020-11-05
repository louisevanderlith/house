package core

import (
	"github.com/louisevanderlith/husk"
	"github.com/louisevanderlith/husk/hsk"
	"github.com/louisevanderlith/husk/op"
	"github.com/louisevanderlith/husk/records"
)

type HouseContext interface {
	GetProperty(key hsk.Key) (Property, error)
	FindLatestProperties(page, size int) (records.Page, error)
	CreateProperty(obj Property) (hsk.Key, error)
	UpdateProperty(key hsk.Key, obj Property) error
}

func CreateContext() HouseContext {
	ctx = context{
		Properties: husk.NewTable(Property{}),
	}

	return ctx
}

func Shutdown() {
	ctx.Properties.Save()
}

func Context() HouseContext {
	return ctx
}

type context struct {
	Properties husk.Table
}

var ctx context

func (c context) GetProperty(key hsk.Key) (Property, error) {
	rec, err := c.Properties.FindByKey(key)

	if err != nil {
		return Property{}, err
	}

	return rec.GetValue().(Property), nil
}

func (c context) FindLatestProperties(page, size int) (records.Page, error) {
	return c.Properties.Find(page, size, op.Everything())
}

func (c context) CreateProperty(obj Property) (hsk.Key, error) {
	return c.Properties.Create(obj)
}

func (c context) UpdateProperty(key hsk.Key, obj Property) error {
	return c.Properties.Update(key, obj)
}
