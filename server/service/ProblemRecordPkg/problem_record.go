package ProblemRecordPkg

import (
	"github.com/flipped-aurora/gin-vue-admin/server/global"
	"github.com/flipped-aurora/gin-vue-admin/server/model/ProblemRecordPkg"
	ProblemRecordPkgReq "github.com/flipped-aurora/gin-vue-admin/server/model/ProblemRecordPkg/request"
	"github.com/flipped-aurora/gin-vue-admin/server/model/common/request"
)

type ProblemRecordService struct {
}

// CreateProblemRecord 创建ProblemRecord记录
// Author [piexlmax](https://github.com/piexlmax)
func (pbRecordService *ProblemRecordService) CreateProblemRecord(pbRecord *ProblemRecordPkg.ProblemRecord) (err error) {
	err = global.GVA_DB.Create(pbRecord).Error
	return err
}

// DeleteProblemRecord 删除ProblemRecord记录
// Author [piexlmax](https://github.com/piexlmax)
func (pbRecordService *ProblemRecordService) DeleteProblemRecord(pbRecord ProblemRecordPkg.ProblemRecord) (err error) {
	err = global.GVA_DB.Delete(&pbRecord).Error
	return err
}

// DeleteProblemRecordByIds 批量删除ProblemRecord记录
// Author [piexlmax](https://github.com/piexlmax)
func (pbRecordService *ProblemRecordService) DeleteProblemRecordByIds(ids request.IdsReq) (err error) {
	err = global.GVA_DB.Delete(&[]ProblemRecordPkg.ProblemRecord{}, "id in ?", ids.Ids).Error
	return err
}

// UpdateProblemRecord 更新ProblemRecord记录
// Author [piexlmax](https://github.com/piexlmax)
func (pbRecordService *ProblemRecordService) UpdateProblemRecord(pbRecord ProblemRecordPkg.ProblemRecord) (err error) {
	err = global.GVA_DB.Save(&pbRecord).Error
	return err
}

// GetProblemRecord 根据id获取ProblemRecord记录
// Author [piexlmax](https://github.com/piexlmax)
func (pbRecordService *ProblemRecordService) GetProblemRecord(id uint) (pbRecord ProblemRecordPkg.ProblemRecord, err error) {
	err = global.GVA_DB.Where("id = ?", id).First(&pbRecord).Error
	return
}

// GetProblemRecordInfoList 分页获取ProblemRecord记录
// Author [piexlmax](https://github.com/piexlmax)
func (pbRecordService *ProblemRecordService) GetProblemRecordInfoList(info ProblemRecordPkgReq.ProblemRecordSearch) (list []ProblemRecordPkg.ProblemRecord, total int64, err error) {
	limit := info.PageSize
	offset := info.PageSize * (info.Page - 1)
	// 创建db
	db := global.GVA_DB.Model(&ProblemRecordPkg.ProblemRecord{})
	var pbRecords []ProblemRecordPkg.ProblemRecord
	// 如果有条件搜索 下方会自动创建搜索语句
	if info.StartCreatedAt != nil && info.EndCreatedAt != nil {
		db = db.Where("created_at BETWEEN ? AND ?", info.StartCreatedAt, info.EndCreatedAt)
	}
	err = db.Count(&total).Error
	if err != nil {
		return
	}
	if info.Page == 0 || info.PageSize == 0 {
		err = db.Find(&pbRecords).Error
	} else {
		err = db.Limit(limit).Offset(offset).Find(&pbRecords).Error
	}

	return pbRecords, total, err
}
