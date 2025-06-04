// 自动生成模板MultiSpectraAnalysis
package MultiSpectraAnalysisPkg

import (
	"github.com/flipped-aurora/gin-vue-admin/server/global"
)

// MultiSpectraAnalysis 结构体
type MultiSpectraAnalysis struct {
	global.GVA_MODEL
	PlateId              *int   `json:"plateId" form:"plateId" gorm:"column:plate_lpogid;comment:板块Id;"`
	MultiSpectraName     string `json:"multiSpectraName" form:"multiSpectraName" gorm:"column:multi_spectra_name;comment:光谱名称;"`
	MultiSpectraTypeIds  string `json:"multiSpectraTypeIds" form:"multiSpectraTypeIds" gorm:"column:multi_spectra_type_ids;comment:光谱类型id;"`
	AerialPhotographyIds string `json:"aerialPhotographyIds" form:"aerialPhotographyIds" gorm:"column:aerial_photography_ids;comment:航摄成果id集;"`
	Desc                 string `json:"desc" form:"desc" gorm:"type:longtext;column:desc;comment:描述;"`
	LandId               string `json:"landId" form:"landId" gorm:"type:longtext;column:landId;comment:田块id;"`

	ProblemCount int `gorm:"-" json:"problem_count"` //问题记录数
}

// TableName MultiSpectraAnalysis 表名
func (MultiSpectraAnalysis) TableName() string {
	return "multi_spectra_analysis"
}
