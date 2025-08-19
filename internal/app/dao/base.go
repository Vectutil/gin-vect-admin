package dao

import (
	"context"
	"gin-vect-admin/internal/app/model/common"
	"gorm.io/gorm"
)

type BaseDao struct {
	db *gorm.DB
}

func NewBaseDao(db *gorm.DB) *BaseDao {
	return &BaseDao{db: db}
}

func (d *BaseDao) Create(ctx context.Context, model common.IModel) error {
	return d.db.WithContext(ctx).Create(model).Error
}

func (d *BaseDao) Update(ctx context.Context, model common.IModel) error {
	return d.db.WithContext(ctx).
		Model(model).Where("id = ?", model.GetID()).Updates(model).Error
}

func (d *BaseDao) Delete(ctx context.Context, model common.IModel) error {
	return d.db.WithContext(ctx).
		Model(model).Where("id = ?", model.GetID()).Update("deleted_at", gorm.Expr("NOW()")).Error
}

func (d *BaseDao) GetById(ctx context.Context, id int64, model common.IModel) error {
	err := d.db.WithContext(ctx).First(&model, id).Error
	if err != nil {
		return err
	}
	return nil
}
