package repository

import (
	"fmt"

	"github.com/jinzhu/gorm"
	"github.com/ronymmoura/codepix/domain/model"
)

type TransactionRepositoryDb struct {
	Db *gorm.DB
}

func (r TransactionRepositoryDb) Register(t *model.Transaction) error {
	err := r.Db.Create(t).Error

	if err != nil {
		return err
	}

	return nil
}

func (r TransactionRepositoryDb) Save(t *model.Transaction) error {
	err := r.Db.Save(t).Error

	if err != nil {
		return err
	}

	return nil
}

func (r TransactionRepositoryDb) Find(id string) (*model.Transaction, error) {
	var transaction model.Transaction

	r.Db.Preload("AccountFrom.Bank").First(&transaction, "id = ?", id)

	if transaction.ID == "" {
		return nil, fmt.Errorf("no transaction was found")
	}

	return &transaction, nil
}
