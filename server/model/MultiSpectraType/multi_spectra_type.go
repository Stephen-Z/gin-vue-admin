// 自动生成模板MultiSpectraType
package MultiSpectraType

import (
	"github.com/flipped-aurora/gin-vue-admin/server/global"
)

// MultiSpectraType 结构体
type MultiSpectraType struct {
	global.GVA_MODEL
	SpectraType         string `json:"spectraType" form:"spectraType" gorm:"index:idx_aerial_type,unique;column:spectra_type;comment:光谱类型;"`
	AerialPhotographyId uint   `json:"aerialPhotographyId" form:"aerialPhotographyId" gorm:"index:idx_aerial_type,unique;column:aerial_photography_id;comment:航摄成果Id;"`
	AerialServerAddress string `json:"aerialServerAddress" form:"aerialServerAddress" gorm:"column:aerial_server_address;comment:航摄成果地址;"`

	FileUrl  *string `gorm:"-"`
	Position *string `gorm:"-"`
}

// TableName MultiSpectraType 表名
func (MultiSpectraType) TableName() string {
	return "multi_spectra_type"
}
