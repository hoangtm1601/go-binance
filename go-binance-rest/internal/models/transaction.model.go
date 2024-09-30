package models

import (
	"gorm.io/gorm"
	"time"
)

type TransactionStatus string

const (
	Failed    TransactionStatus = `Failed`
	Succeeded TransactionStatus = `Succeeded`
)

type TransactionType string

const (
	Charge TransactionType = "charge"
	Refund TransactionType = "refund"
)

type Transaction struct {
	gorm.Model
	UserId          uint              `gorm:"not null"`
	Amount          float64           `gorm:"not null"`
	TransactionDate time.Time         `gorm:"not null"`
	Status          TransactionStatus `gorm:"not null"`
	Currency        string            `gorm:"not null" default:"USD"`
	Type            TransactionType   `gorm:"not null"`
}
