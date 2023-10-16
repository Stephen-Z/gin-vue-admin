// 自动生成模板Theme
package Theme

import (
	"github.com/flipped-aurora/gin-vue-admin/server/global"
)

// Theme 结构体
type Theme struct {
	global.GVA_MODEL
	ThemeName          string `json:"themeName" form:"themeName" gorm:"column:theme_name;comment:主题名;"`
	UserRoles          string `json:"userRoles" form:"userRoles" gorm:"type:longtext;column:user_roles;comment:用户角色;"`
	LoginViewLogo      string `json:"loginViewLogo" form:"loginViewLogo" gorm:"type:longtext;column:login_view_logo;comment:登录页logo;"`
	SystemLogo         string `json:"systemLogo" form:"systemLogo" gorm:"type:longtext;column:system_logo;comment:系统logo;"`
	SystemName         string `json:"systemName" form:"systemName" gorm:"column:system_name;comment:系统名称;"`
	BackgroundPhoto    string `json:"backgroundPhoto" form:"backgroundPhoto" gorm:"type:longtext;column:background_photo;comment:背景轮播图;"`
	IsOrNoDefaultTheme *bool  `json:"isOrNoDefaultTheme" form:"isOrNoDefaultTheme" gorm:"column:is_or_no_default_theme;comment:是否默认主题;"`
}

// TableName Theme 表名
func (Theme) TableName() string {
	return "theme"
}
