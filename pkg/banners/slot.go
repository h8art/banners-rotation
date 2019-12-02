package banners

import "github.com/go-pg/pg/v9"

type Slot struct {
	Id int64
}

func AddSlot() (*Slot, error) {
	db := pg.Connect(&pg.Options{
		User:     "postgres",
		Database: "banners",
	})
	defer db.Close()
	newSlot := &Slot{}
	err := db.Insert(newSlot)
	if err != nil {
		return nil, err
	}
	return newSlot, nil
}

func GetSlots() ([]Slot, error) {
	db := pg.Connect(&pg.Options{
		User:     "postgres",
		Database: "banners",
	})
	defer db.Close()
	var slots []Slot
	err := db.Model(&slots).Select()
	if err != nil {
		return nil, err
	}
	return slots, nil
}
