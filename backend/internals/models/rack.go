package models

type Rack struct {
	Id      int
	Shelves map[int]Shelve
}

type Shelve struct {
	Items []Item
}
