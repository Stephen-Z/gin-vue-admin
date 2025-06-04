// 自动生成模板MultiSpectraPlate
package MultiSpectraPlate

import (
	"github.com/flipped-aurora/gin-vue-admin/server/global"
)

// MultiSpectraPlate 结构体
type MultiSpectraPlate struct {
	global.GVA_MODEL
	PlateName string `json:"plateName" form:"plateName" gorm:"column:plate_name;comment:板块名;"`
}

// TableName MultiSpectraPlate 表名
func (MultiSpectraPlate) TableName() string {
	return "multi_spectra_plate"
}
