package banners

import "github.com/go-pg/pg/v9"

type Group struct {
	Id   int64
	Name string
}

func AddGroup(name string) (*Group, error) {
	db := pg.Connect(&pg.Options{
		User:     "postgres",
		Database: "banners",
	})
	defer db.Close()
	newGroup := &Group{
		Name: name,
	}
	err := db.Insert(newGroup)
	if err != nil {
		return nil, err
	}
	return newGroup, nil
}
func GetGroups() ([]Group, error) {
	db := pg.Connect(&pg.Options{
		User:     "postgres",
		Database: "banners",
	})
	defer db.Close()
	var groups []Group
	err := db.Model(&groups).Select()
	if err != nil {
		return nil, err
	}
	return groups, nil
}
