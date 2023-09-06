package NestInfo

import (
	"encoding/json"

	"github.com/flipped-aurora/gin-vue-admin/server/global"
	"github.com/flipped-aurora/gin-vue-admin/server/model/NestInfo"
	NestInfoReq "github.com/flipped-aurora/gin-vue-admin/server/model/NestInfo/request"
	NestRoleReq "github.com/flipped-aurora/gin-vue-admin/server/model/Nestrolepkg/request"
	"github.com/flipped-aurora/gin-vue-admin/server/model/common/request"
	"github.com/flipped-aurora/gin-vue-admin/server/model/system"
	"github.com/flipped-aurora/gin-vue-admin/server/service/Nestrolepkg"
	serviceSystem "github.com/flipped-aurora/gin-vue-admin/server/service/system"
	"github.com/flipped-aurora/gin-vue-admin/server/utils"
	"github.com/gin-gonic/gin"
	"gorm.io/gorm"
)

type NestInfoService struct {
}

// CreateNestInfo 创建NestInfo记录
// Author [piexlmax](https://github.com/piexlmax)
func (nestinfoService *NestInfoService) CreateNestInfo(nestinfo *NestInfo.NestInfo) (err error) {
	err = global.GVA_DB.Create(nestinfo).Error
	return err
}

// DeleteNestInfo 删除NestInfo记录
// Author [piexlmax](https://github.com/piexlmax)
func (nestinfoService *NestInfoService) DeleteNestInfo(nestinfo NestInfo.NestInfo) (err error) {
	err = global.GVA_DB.Transaction(func(tx *gorm.DB) error {
		if err := tx.Model(&NestInfo.NestInfo{}).Where("id = ?", nestinfo.ID).Update("deleted_by", nestinfo.DeletedBy).Error; err != nil {
			return err
		}
		if err = tx.Delete(&nestinfo).Error; err != nil {
			return err
		}
		return nil
	})
	return err
}

// DeleteNestInfoByIds 批量删除NestInfo记录
// Author [piexlmax](https://github.com/piexlmax)
func (nestinfoService *NestInfoService) DeleteNestInfoByIds(ids request.IdsReq, deleted_by uint) (err error) {
	err = global.GVA_DB.Transaction(func(tx *gorm.DB) error {
		if err := tx.Model(&NestInfo.NestInfo{}).Where("id in ?", ids.Ids).Update("deleted_by", deleted_by).Error; err != nil {
			return err
		}
		if err := tx.Where("id in ?", ids.Ids).Delete(&NestInfo.NestInfo{}).Error; err != nil {
			return err
		}
		return nil
	})
	return err
}

// UpdateNestInfo 更新NestInfo记录
// Author [piexlmax](https://github.com/piexlmax)
func (nestinfoService *NestInfoService) UpdateNestInfo(nestinfo NestInfo.NestInfo) (err error) {
	err = global.GVA_DB.Save(&nestinfo).Error
	return err
}

// GetNestInfo 根据id获取NestInfo记录
// Author [piexlmax](https://github.com/piexlmax)
func (nestinfoService *NestInfoService) GetNestInfo(id uint) (nestinfo NestInfo.NestInfo, err error) {
	err = global.GVA_DB.Where("id = ?", id).First(&nestinfo).Error
	return
}

// GetNestInfoByIds 根据id获取NestInfo记录
// Author [piexlmax](https://github.com/piexlmax)
func (nestinfoService *NestInfoService) GetNestInfoByNestIds(c *gin.Context, nestIds []string) (nestinfoArr []NestInfo.NestInfo, err error) {
	nestIDList, err := nestinfoService.GetNestIDListByUser(c)
	if err != nil {
		return
	}
	db := global.GVA_DB.Model(&NestInfo.NestInfo{})
	db.Where("deleted_by = 0")
	db.Where("nestid IN ? ", nestIDList)
	if len(nestIds) > 0 {
		err = db.Where("nestid in ? ", nestIds).Find(&nestinfoArr).Error
	} else {
		err = db.Find(&nestinfoArr).Error
	}
	return
}

// GetNestInfoInfoList 分页获取NestInfo记录
// Author [piexlmax](https://github.com/piexlmax)
func (nestinfoService *NestInfoService) GetNestInfoInfoList(info NestInfoReq.NestInfoSearch) (list []NestInfo.NestInfo, total int64, err error) {
	limit := info.PageSize
	offset := info.PageSize * (info.Page - 1)
	// 创建db
	db := global.GVA_DB.Model(&NestInfo.NestInfo{})
	var nestinfos []NestInfo.NestInfo
	// 如果有条件搜索 下方会自动创建搜索语句
	if info.StartCreatedAt != nil && info.EndCreatedAt != nil {
		db = db.Where("created_at BETWEEN ? AND ?", info.StartCreatedAt, info.EndCreatedAt)
	}
	err = db.Count(&total).Error
	if err != nil {
		return
	}

	err = db.Limit(limit).Offset(offset).Find(&nestinfos).Error
	return nestinfos, total, err
}

func (nestinfoService *NestInfoService) GetAllUserList() (list []system.SysUser, total int64, err error) {
	db := global.GVA_DB.Model(&system.SysUser{})
	var userList []system.SysUser
	err = db.Count(&total).Error
	if err != nil {
		return
	}
	err = db.Preload("Authorities").Preload("Authority").Find(&userList).Error
	return userList, total, err
}

func (nestinfoService *NestInfoService) GetNestInfoInfoListWithUser(c *gin.Context) (list []NestInfo.NestInfo, total int64, err error) {
	nestIDList, err := nestinfoService.GetNestIDListByUser(c)
	if err != nil {
		return nil, 0, err
	}
	db := global.GVA_DB.Model(&NestInfo.NestInfo{})
	db.Where("nestid IN ?", nestIDList)
	var nestinfos []NestInfo.NestInfo
	err = db.Find(&nestinfos).Error
	return nestinfos, total, err
}

func (nestinfoService *NestInfoService) GetNestIDListByUser(c *gin.Context) (list []string, err error) {
	user, err := utils.GetClaims(c)
	if err != nil {
		return
	}
	us := new(serviceSystem.UserService)
	nestroleservice := new(Nestrolepkg.NestRoleService)
	nestroleSearch := NestRoleReq.NestRoleSearch{}
	nestroleSearch.Page = 1
	nestroleSearch.PageSize = 9999
	nestroleSearch.AuthID, err = us.GetUserAuthorities(user.BaseClaims.ID)
	if err != nil {
		return
	}
	var nestIDList []string
	if !utils.UintArrContains(nestroleSearch.AuthID, 888) {
		nestRoleList, _, err := nestroleservice.GetNestRoleInfoList(nestroleSearch)
		if err != nil {
			return nil, err
		}
		for _, item := range nestRoleList {
			var nestIDListStr []string
			json.Unmarshal([]byte(item.Nestid), &nestIDListStr)
			// nestIDListStr := strings.Split(item.Nestid, ",")
			for _, id := range nestIDListStr {
				if !utils.StringArrContains(nestIDList, id) {
					nestIDList = append(nestIDList, id)
				}
			}
		}
	} else {
		db := global.GVA_DB.Model(&NestInfo.NestInfo{})
		err := db.Select("nestid").Find(&nestIDList).Error
		if err != nil {
			return nil, err
		}
	}
	return nestIDList, nil
}
