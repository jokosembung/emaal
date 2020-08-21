package models

import (
	"errors"
	"github.com/jinzhu/gorm"
)

type Kartu struct {
	ID       uint32 `gorm:"primary_key;auto_increment" json:"id"`
	No_kartu string `gorm:"size:255;not null;unique" json:"no_kartu"`
	Nama     string `gorm:"size:100;not null;unique" json:"nama"`
	Saldo    string `gorm:"size:100;not null;" json:"saldo"`
}

func (u *Kartu) FindkartuByCustomerId(db *gorm.DB, uid uint32) (*Kartu, error) {
	var err error
	err = db.Debug().Model(Kartu{}).Where("customer_id = ?", uid).Take(&u).Error
	if err != nil {
		return &Kartu{}, err
	}
	if gorm.IsRecordNotFoundError(err) {
		return &Kartu{}, errors.New("Kartu tidak ditemukan silahkan aktivasi ulang")
	}
	return u, err
}

func (u *Kartu) FindkartuByNomor(db *gorm.DB, uid string) (*Kartu, error) {
	var err error
	err = db.Debug().Model(Kartu{}).Where("no_kartu = ?", uid).Take(&u).Error
	if err != nil {
		return &Kartu{}, err
	}
	if gorm.IsRecordNotFoundError(err) {
		return &Kartu{}, errors.New("Kartu tidak ditemukan silahkan aktivasi ulang")
	}
	return u, err
}
