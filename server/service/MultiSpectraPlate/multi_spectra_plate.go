package MultiSpectraPlate

import (
	"github.com/flipped-aurora/gin-vue-admin/server/global"
	"github.com/flipped-aurora/gin-vue-admin/server/model/MultiSpectraPlate"
	MultiSpectraPlateReq "github.com/flipped-aurora/gin-vue-admin/server/model/MultiSpectraPlate/request"
	"github.com/flipped-aurora/gin-vue-admin/server/model/common/request"
)

type MultiSpectraPlateService struct {
}

// CreateMultiSpectraPlate 创建MultiSpectraPlate记录
// Author [piexlmax](https://github.com/piexlmax)
func (MtpectraPlateService *MultiSpectraPlateService) CreateMultiSpectraPlate(MtpectraPlate *MultiSpectraPlate.MultiSpectraPlate) (err error) {
	err = global.GVA_DB.Create(MtpectraPlate).Error
	return err
}

// DeleteMultiSpectraPlate 删除MultiSpectraPlate记录
// Author [piexlmax](https://github.com/piexlmax)
func (MtpectraPlateService *MultiSpectraPlateService) DeleteMultiSpectraPlate(MtpectraPlate MultiSpectraPlate.MultiSpectraPlate) (err error) {
	err = global.GVA_DB.Delete(&MtpectraPlate).Error
	return err
}

// DeleteMultiSpectraPlateByIds 批量删除MultiSpectraPlate记录
// Author [piexlmax](https://github.com/piexlmax)
func (MtpectraPlateService *MultiSpectraPlateService) DeleteMultiSpectraPlateByIds(ids request.IdsReq) (err error) {
	err = global.GVA_DB.Delete(&[]MultiSpectraPlate.MultiSpectraPlate{}, "id in ?", ids.Ids).Error
	return err
}

// UpdateMultiSpectraPlate 更新MultiSpectraPlate记录
// Author [piexlmax](https://github.com/piexlmax)
func (MtpectraPlateService *MultiSpectraPlateService) UpdateMultiSpectraPlate(MtpectraPlate MultiSpectraPlate.MultiSpectraPlate) (err error) {
	err = global.GVA_DB.Save(&MtpectraPlate).Error
	return err
}

// GetMultiSpectraPlate 根据id获取MultiSpectraPlate记录
// Author [piexlmax](https://github.com/piexlmax)
func (MtpectraPlateService *MultiSpectraPlateService) GetMultiSpectraPlate(id uint) (MtpectraPlate MultiSpectraPlate.MultiSpectraPlate, err error) {
	err = global.GVA_DB.Where("id = ?", id).First(&MtpectraPlate).Error
	return
}

// GetMultiSpectraPlateInfoList 分页获取MultiSpectraPlate记录
// Author [piexlmax](https://github.com/piexlmax)
func (MtpectraPlateService *MultiSpectraPlateService) GetMultiSpectraPlateInfoList(info MultiSpectraPlateReq.MultiSpectraPlateSearch) (list []MultiSpectraPlate.MultiSpectraPlate, total int64, err error) {
	limit := info.PageSize
	offset := info.PageSize * (info.Page - 1)
	// 创建db
	db := global.GVA_DB.Model(&MultiSpectraPlate.MultiSpectraPlate{})
	var MtpectraPlates []MultiSpectraPlate.MultiSpectraPlate
	// 如果有条件搜索 下方会自动创建搜索语句
	if info.StartCreatedAt != nil && info.EndCreatedAt != nil {
		db = db.Where("created_at BETWEEN ? AND ?", info.StartCreatedAt, info.EndCreatedAt)
	}
	err = db.Count(&total).Error
	if err != nil {
		return
	}

	err = db.Limit(limit).Offset(offset).Find(&MtpectraPlates).Error
	return MtpectraPlates, total, err
}
