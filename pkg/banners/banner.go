package banners

import (
	"fmt"
	"github.com/go-pg/pg/v9"
	"math"
	"sort"
	"time"
)

type Banner struct {
	Id          int64
	SlotId      int64
	Slot        *Slot
	Description string
}

func AddBanner(slotId int64, desc string) (*Banner, error) {
	db := pg.Connect(&pg.Options{
		User:     "postgres",
		Database: "banners",
	})
	defer db.Close()
	newBanner := &Banner{
		SlotId:      slotId,
		Description: desc,
	}
	err := db.Insert(newBanner)
	if err != nil {
		return nil, err
	}
	return newBanner, nil
}

func DeleteBanner(id int64) error {
	db := pg.Connect(&pg.Options{
		User:     "postgres",
		Database: "banners",
	})
	defer db.Close()
	_, err := db.Model((*Banner)(nil)).Where("id = (?)", id).Delete()
	if err != nil {
		return err
	}
	return nil
}

func GetBanner(slotId int64, groupId int64) (int64, error) {
	db := pg.Connect(&pg.Options{
		User:     "postgres",
		Database: "banners",
	})
	defer db.Close()
	totalViews, err := getTotalViewsSlot(slotId, groupId)
	if err != nil {
		return 0, nil
	}
	var banners []Banner
	err = db.Model(&banners).Select()
	if err != nil {
		return 0, err
	}
	type bannerBandit struct {
		id  int64
		val float64
	}
	var bannersBandit []bannerBandit
	for _, b := range banners {
		val, err := getBanditVal(db, b.Id, groupId, totalViews)
		if err != nil {
			return 0, err
		}
		bannersBandit = append(bannersBandit, bannerBandit{
			id:  b.Id,
			val: val,
		})
	}
	sort.Slice(bannersBandit, func(i, j int) bool {
		return bannersBandit[i].val > bannersBandit[j].val
	})
	fmt.Println(bannersBandit)
	ViewBanner(bannersBandit[0].id, groupId, slotId)
	return bannersBandit[0].id, nil
}

func ViewBanner(bannerId int64, groupId int64, slotId int64) {
	db := pg.Connect(&pg.Options{
		User:     "postgres",
		Database: "banners",
	})
	defer db.Close()
	stat := new(Statistic)
	err := db.Model(stat).
		Where("banner_id = ?", bannerId).
		Where("group_id = ?", groupId).
		Select()
	if stat.BannerId == 0 {
		newStatistic := &Statistic{
			BannerId:   bannerId,
			GroupId:    groupId,
			ViewCount:  1,
			ClickCount: 0,
		}
		err := db.Insert(newStatistic)
		if err != nil {
			fmt.Println(err)
		}
		return
	}
	if err != nil {
		fmt.Println(err)
	}
	stat.ViewCount++
	SendStatistic("view", slotId, bannerId, groupId)
	err = db.Update(stat)
	if err != nil {
		fmt.Println(err)
	}

}

func SendStatistic(statType string, slotId int64, bannerId int64, groupId int64) {
	fmt.Println(statType, slotId, bannerId, groupId, time.Now())
}

func getBanditVal(db *pg.DB, bannerId int64, groupId int64, totalViews int64) (float64, error) {
	var stat Statistic
	var value float64
	_ = db.Model(&stat).
		Where("banner_id = (?)", bannerId).
		Where("group_id = (?)", groupId).
		Select()
	views := int64(1)
	clicks := int64(1)
	if stat.ViewCount != 0 {
		views = stat.ViewCount
	}
	if stat.ClickCount != 0 {
		clicks = stat.ClickCount
	}
	numerator := 2 * math.Log(float64(totalViews))
	denominator := float64(views)
	if denominator == 0 {
		denominator = 1
	}
	value = float64(clicks) + math.Sqrt(numerator/denominator)
	return value, nil
}

func getTotalViewsSlot(slotId int64, groupId int64) (int64, error) {
	db := pg.Connect(&pg.Options{
		User:     "postgres",
		Database: "banners",
	})
	defer db.Close()
	var totalViews int64
	var banners []Banner
	err := db.Model(&banners).Select()
	if err != nil {
		return 0, err
	}
	var bannersId []int64
	for _, b := range banners {
		bannersId = append(bannersId, b.Id)
	}
	var stats []Statistic
	err = db.Model(&stats).
		Where("banner_id in (?)", pg.In(bannersId)).
		Where("group_id = (?)", groupId).
		Select()
	if err != nil {
		return 0, err
	}
	for _, s := range stats {
		totalViews += s.ViewCount
	}
	return totalViews, nil
}

func getAvgClicks(groupId int64) (int64, error) {
	db := pg.Connect(&pg.Options{
		User:     "postgres",
		Database: "banners",
	})
	defer db.Close()
	var stats []Statistic
	err := db.Model(&stats).Select()
	if err != nil {
		return 0, err
	}
	var total int64
	for _, s := range stats {
		total += s.ClickCount
	}
	return total / int64(len(stats)), nil
}

func GetBanners() ([]Banner, error) {
	db := pg.Connect(&pg.Options{
		User:     "postgres",
		Database: "banners",
	})
	defer db.Close()
	var banners []Banner
	err := db.Model(&banners).Select()
	if err != nil {
		return nil, err
	}
	return banners, nil
}
