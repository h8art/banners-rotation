package banners

import (
	"github.com/go-pg/pg/v9"
	"github.com/go-pg/pg/v9/orm"
)

func InitDB() error {
	db := pg.Connect(&pg.Options{
		User:     "postgres",
		Database: "banners",
	})
	defer db.Close()
	err := createSchema(db)
	if err != nil {
		return err
	}
	return nil
}
func createSchema(db *pg.DB) error {

	for _, model := range []interface{}{(*Banner)(nil), (*Group)(nil), (*Slot)(nil), (*Statistic)(nil)} {
		err := db.CreateTable(model, &orm.CreateTableOptions{})
		if err != nil {
			return err
		}
	}
	return nil
}
