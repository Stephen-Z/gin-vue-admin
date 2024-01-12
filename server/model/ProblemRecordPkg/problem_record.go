// 自动生成模板ProblemRecord
package ProblemRecordPkg

import (
	"github.com/flipped-aurora/gin-vue-admin/server/global"
	"time"
)

// ProblemRecord 结构体
type ProblemRecord struct {
	global.GVA_MODEL
	OrderNum      *int       `json:"orderNum" form:"orderNum" gorm:"column:order_num;comment:;"`
	RegisterDate  *time.Time `json:"registerDate" form:"registerDate" gorm:"column:register_date;comment:;"`
	ProblemName   string     `json:"problemName" form:"problemName" gorm:"column:problem_name;comment:;"`
	ProblemType   string     `json:"problemType" form:"problemType" gorm:"column:problem_type;comment:;"`
	ProblemSource string     `json:"problemSource" form:"problemSource" gorm:"column:problem_source;comment:;"`
	ProblemDesc   string     `json:"problemDesc" form:"problemDesc" gorm:"column:problem_desc;comment:;"`
	PersonCharge  string     `json:"personCharge" form:"personCharge" gorm:"column:person_charge;comment:;"`
	ProblemImage  string     `json:"problemImage" form:"problemImage" gorm:"column:problem_image;comment:;"`
	HandMeasurce  string     `json:"handMeasurce" form:"handMeasurce" gorm:"column:hand_measurce;comment:;"`
	CurrState     string     `json:"currState" form:"currState" gorm:"column:curr_state;type:enum('已处理','未处理');comment:;"`
	Lng           *float64   `json:"lng" form:"lng" gorm:"column:lng;comment:;"`
	Lat           *float64   `json:"lat" form:"lat" gorm:"column:lat;comment:;"`
}

// TableName ProblemRecord 表名
func (ProblemRecord) TableName() string {
	return "problem_record"
}
