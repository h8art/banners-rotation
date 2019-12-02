package handlers

import (
	"banner-otus/pkg/banners"
	"github.com/gin-gonic/gin"
	"strconv"
)

func AddBanner(c *gin.Context) {
	type BannerReq struct {
		SlotId      int64
		Description string
	}
	newBanner := BannerReq{}
	err := c.BindJSON(&newBanner)
	if err != nil {
		c.JSON(500, err)
	}
	banner, err := banners.AddBanner(newBanner.SlotId, newBanner.Description)
	if err != nil {
		c.JSON(500, err)
	}
	c.JSON(200, banner)
}
func DeleteBanner(c *gin.Context) {
	bannerIdString := c.Param("id")
	bannerId, err := strconv.Atoi(bannerIdString)
	if err != nil {
		c.JSON(500, err)
	}
	err = banners.DeleteBanner(int64(bannerId))
	if err != nil {
		c.JSON(500, err)
	}
}
func VisitBanner(c *gin.Context) {
	bannerIdString := c.Param("bannerId")
	groupIdString := c.Param("groupId")
	bannerId, err := strconv.Atoi(bannerIdString)
	if err != nil {
		c.JSON(500, err.Error())
	}
	groupId, err := strconv.Atoi(groupIdString)
	if err != nil {
		c.JSON(500, err.Error())
	}
	stat, err := banners.ClickBanner(int64(bannerId), int64(groupId))
	if err != nil {
		c.JSON(500, err.Error())
	}
	c.JSON(200, stat)
}
func GetBanner(c *gin.Context) {
	slotIdString := c.Param("slotId")
	groupIdString := c.Param("groupId")
	slotId, err := strconv.Atoi(slotIdString)
	if err != nil {
		c.JSON(500, err.Error())
	}
	groupId, err := strconv.Atoi(groupIdString)
	if err != nil {
		c.JSON(500, err.Error())
	}
	bannerId, err := banners.GetBanner(int64(slotId), int64(groupId))
	if err != nil {
		c.JSON(500, err.Error())
	}
	c.JSON(200, bannerId)
}
func AddGroup(c *gin.Context) {
	type group struct {
		Name string
	}
	groupReq := group{}
	err := c.BindJSON(&groupReq)
	if err != nil {
		c.JSON(500, err)
	}
	newGroup, err := banners.AddGroup(groupReq.Name)
	if err != nil {
		c.JSON(500, err)
	}
	c.JSON(200, newGroup)
}
func GetGroups(c *gin.Context) {
	groups, err := banners.GetGroups()
	if err != nil {
		c.JSON(500, err)
	}
	c.JSON(200, groups)
}
func GetBanners(c *gin.Context) {
	bannersArr, err := banners.GetBanners()
	if err != nil {
		c.JSON(500, err)
	}
	c.JSON(200, bannersArr)
}
func GetSlots(c *gin.Context) {
	slotsArr, err := banners.GetSlots()
	if err != nil {
		c.JSON(500, err)
	}
	c.JSON(200, slotsArr)
}
func AddSlot(c *gin.Context) {
	slot, err := banners.AddSlot()
	if err != nil {
		c.JSON(500, err)
	}
	c.JSON(200, slot)
}
