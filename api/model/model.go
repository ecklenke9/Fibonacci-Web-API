package model

// Fibonacci schema of the fibonaccis table
type Fibonacci struct {
	Ordinal int `json:"ordinal" gorm:"unique;not null"`
	FibNum  int `json:"fibNum"`
}
