package model

import (
	"encoding/json"

	"gorm.io/gorm"
)

type Menu struct {
	MenuName string
	gorm.Model
}

func (Menu) TableName() string {
	// ini akan membuat sebuah nama tabel (custominasi nama tabel)
	return "m_menu"
}

func (m *Menu) ToString() string {
	menu, err := json.MarshalIndent(m, "", "  ")
	if err != nil {
		return ""
	}
	return string(menu)
}
