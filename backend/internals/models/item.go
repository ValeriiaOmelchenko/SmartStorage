package models

type Item struct {
	Id         int
	Name       string
	Categories []string
	Quantity   int
	Info       string
	Custom     any
}
