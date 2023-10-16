package Theme

import (
	"github.com/flipped-aurora/gin-vue-admin/server/global"
	"github.com/flipped-aurora/gin-vue-admin/server/model/Theme"
	ThemeReq "github.com/flipped-aurora/gin-vue-admin/server/model/Theme/request"
	"github.com/flipped-aurora/gin-vue-admin/server/model/common/request"
)

type ThemeService struct {
}

// CreateTheme 创建Theme记录
// Author [piexlmax](https://github.com/piexlmax)
func (themeService *ThemeService) CreateTheme(theme *Theme.Theme) (err error) {
	if *theme.IsOrNoDefaultTheme == true {
		global.GVA_DB.Exec("update theme  set is_or_no_default_theme = 0")
	}
	err = global.GVA_DB.Create(theme).Error
	return err
}

// DeleteTheme 删除Theme记录
// Author [piexlmax](https://github.com/piexlmax)
func (themeService *ThemeService) DeleteTheme(theme Theme.Theme) (err error) {
	err = global.GVA_DB.Delete(&theme).Error
	return err
}

// DeleteThemeByIds 批量删除Theme记录
// Author [piexlmax](https://github.com/piexlmax)
func (themeService *ThemeService) DeleteThemeByIds(ids request.IdsReq) (err error) {
	err = global.GVA_DB.Delete(&[]Theme.Theme{}, "id in ?", ids.Ids).Error
	return err
}

// UpdateTheme 更新Theme记录
// Author [piexlmax](https://github.com/piexlmax)
func (themeService *ThemeService) UpdateTheme(theme Theme.Theme) (err error) {
	if *theme.IsOrNoDefaultTheme == true {
		global.GVA_DB.Exec("update theme  set is_or_no_default_theme = 0")
	}
	err = global.GVA_DB.Save(&theme).Error
	return err
}

// GetTheme 根据id获取Theme记录
// Author [piexlmax](https://github.com/piexlmax)
func (themeService *ThemeService) GetTheme(id uint) (theme Theme.Theme, err error) {
	err = global.GVA_DB.Where("id = ?", id).First(&theme).Error
	return
}

// GetThemeInfoList 分页获取Theme记录
// Author [piexlmax](https://github.com/piexlmax)
func (themeService *ThemeService) GetThemeInfoList(info ThemeReq.ThemeSearch) (list []Theme.Theme, total int64, err error) {
	limit := info.PageSize
	offset := info.PageSize * (info.Page - 1)
	// 创建db
	db := global.GVA_DB.Model(&Theme.Theme{})
	var themes []Theme.Theme
	// 如果有条件搜索 下方会自动创建搜索语句
	if info.StartCreatedAt != nil && info.EndCreatedAt != nil {
		db = db.Where("created_at BETWEEN ? AND ?", info.StartCreatedAt, info.EndCreatedAt)
	}
	err = db.Count(&total).Error
	if err != nil {
		return
	}

	err = db.Limit(limit).Offset(offset).Find(&themes).Error
	return themes, total, err
}

//智航巡飞登陆页个性化定制功能：
//1、gva后台新建表单：主题名称，用户角色(多选，重复校验)，登陆页公司logo，系统logo，系统名称，背景轮播图(最多10张)，是否默认主题(保存的时候，如果已经存在其他默认主题，其他默认变为非默认)。
//2、接口1: 根据主题ID获取主题配置。接口2:根据角色ID获取主题ID

// GetThemeByRoleId 根据角色id获取Theme记录
// Author [piexlmax](https://github.com/piexlmax)
func (themeService *ThemeService) GetThemeByRoleId(roleId string) (theme Theme.Theme, err error) {
	err = global.GVA_DB.Where("user_roles = ?", roleId).Limit(1).First(&theme).Error
	return
}
