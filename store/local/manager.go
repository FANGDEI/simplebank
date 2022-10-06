package local

import (
	"fmt"

	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

type Manager struct {
	handler *gorm.DB
}

func New() (*Manager, error) {
	db, err := gorm.Open(
		postgres.Open(
			fmt.Sprintf("host=%s port=%d user=%s dbname=%s password=%s TimeZone=%s",
				C.Host,
				C.Port,
				C.User,
				C.DB,
				C.Password,
				C.TimeZone,
			),
		),
		&gorm.Config{},
	)
	return &Manager{
		handler: db,
	}, err
}
