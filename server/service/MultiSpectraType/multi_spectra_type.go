package MultiSpectraType

import (
	"github.com/flipped-aurora/gin-vue-admin/server/global"
	"github.com/flipped-aurora/gin-vue-admin/server/model/MultiSpectraType"
	MultiSpectraTypeReq "github.com/flipped-aurora/gin-vue-admin/server/model/MultiSpectraType/request"
	"github.com/flipped-aurora/gin-vue-admin/server/model/common/request"
)

type MultiSpectraTypeService struct {
}

// CreateMultiSpectraType 创建MultiSpectraType记录
// Author [piexlmax](https://github.com/piexlmax)
func (MtSpectraTypeService *MultiSpectraTypeService) CreateMultiSpectraType(MtSpectraType *MultiSpectraType.MultiSpectraType) (err error) {
	err = global.GVA_DB.Create(MtSpectraType).Error
	return err
}

// DeleteMultiSpectraType 删除MultiSpectraType记录
// Author [piexlmax](https://github.com/piexlmax)
func (MtSpectraTypeService *MultiSpectraTypeService) DeleteMultiSpectraType(MtSpectraType MultiSpectraType.MultiSpectraType) (err error) {
	err = global.GVA_DB.Delete(&MtSpectraType).Error
	return err
}

// DeleteMultiSpectraTypeByIds 批量删除MultiSpectraType记录
// Author [piexlmax](https://github.com/piexlmax)
func (MtSpectraTypeService *MultiSpectraTypeService) DeleteMultiSpectraTypeByIds(ids request.IdsReq) (err error) {
	err = global.GVA_DB.Delete(&[]MultiSpectraType.MultiSpectraType{}, "id in ?", ids.Ids).Error
	return err
}

// UpdateMultiSpectraType 更新MultiSpectraType记录
// Author [piexlmax](https://github.com/piexlmax)
func (MtSpectraTypeService *MultiSpectraTypeService) UpdateMultiSpectraType(MtSpectraType MultiSpectraType.MultiSpectraType) (err error) {
	err = global.GVA_DB.Save(&MtSpectraType).Error
	return err
}

// GetMultiSpectraType 根据id获取MultiSpectraType记录
// Author [piexlmax](https://github.com/piexlmax)
func (MtSpectraTypeService *MultiSpectraTypeService) GetMultiSpectraType(id uint) (MtSpectraType MultiSpectraType.MultiSpectraType, err error) {
	err = global.GVA_DB.Where("id = ?", id).First(&MtSpectraType).Error
	return
}

// GetMultiSpectraTypeInfoList 分页获取MultiSpectraType记录
// Author [piexlmax](https://github.com/piexlmax)
func (MtSpectraTypeService *MultiSpectraTypeService) GetMultiSpectraTypeInfoList(info MultiSpectraTypeReq.MultiSpectraTypeSearch) (list []MultiSpectraType.MultiSpectraType, total int64, err error) {
	limit := info.PageSize
	offset := info.PageSize * (info.Page - 1)
	// 创建db
	db := global.GVA_DB.Model(&MultiSpectraType.MultiSpectraType{})
	var MtSpectraTypes []MultiSpectraType.MultiSpectraType
	if info.AerialPhotographyId > 0 {
		db = db.Where(" aerial_photography_id = ?", info.AerialPhotographyId)
	}
	if info.ID > 0 {
		db = db.Where(" id = ?", info.ID)
	}
	// 如果有条件搜索 下方会自动创建搜索语句
	if info.StartCreatedAt != nil && info.EndCreatedAt != nil {
		db = db.Where("created_at BETWEEN ? AND ?", info.StartCreatedAt, info.EndCreatedAt)
	}
	err = db.Count(&total).Error
	if err != nil {
		return
	}

	err = db.Limit(limit).Offset(offset).Find(&MtSpectraTypes).Error
	return MtSpectraTypes, total, err
}
