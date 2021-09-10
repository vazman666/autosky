package models

type Tiss struct {
	Firm       string
	PartNumber string
	Name       string //наименование
	Price      string //цена
	Qty        string //количество
	Amount     string //сумма
	Remark     string
}

var Price []Tiss
var Order []Tiss
var Nakl []Tiss

const Kurs = 75
