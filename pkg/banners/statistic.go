package banners

import (
	"fmt"
	"github.com/go-pg/pg/v9"
)

type Statistic struct {
	BannerId   int64 `pg:",pk"`
	Banner     *Banner
	GroupId    int64
	Group      *Group
	ViewCount  int64
	ClickCount int64
}

func ClickBanner(bannerId int64, groupId int64) (*Statistic, error) {
	db := pg.Connect(&pg.Options{
		User:     "postgres",
		Database: "banners",
	})
	defer db.Close()
	stat := new(Statistic)
	count, err := db.Model(stat).
		Where("banner_id = ?", bannerId).
		Where("group_id = ?", groupId).
		Count()
	if err != nil {
		return nil, err
	}
	if count == 0 {
		newStatistic := &Statistic{
			BannerId:   bannerId,
			GroupId:    groupId,
			ViewCount:  0,
			ClickCount: 1,
		}
		err := db.Insert(newStatistic)
		if err != nil {
			return nil, err
		}
		return newStatistic, nil
	} else {
		err = db.Model(stat).
			Where("banner_id = ?", bannerId).
			Where("group_id = ?", groupId).
			Select()
		if err != nil {
			return nil, err
		}
		stat.ClickCount++
		err = db.Update(stat)
		banner := Banner{
			Id: bannerId,
		}
		if err != nil {
			return nil, err
		}
		err = db.Select(&banner)
		if err != nil {
			fmt.Println(err)
		}
		SendStatistic("click", banner.SlotId, bannerId, groupId)
		return stat, nil
	}
}
