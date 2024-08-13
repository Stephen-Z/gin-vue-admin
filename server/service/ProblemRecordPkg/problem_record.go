package ProblemRecordPkg

import (
	"github.com/flipped-aurora/gin-vue-admin/server/global"
	"github.com/flipped-aurora/gin-vue-admin/server/model/ProblemRecordPkg"
	ProblemRecordPkgReq "github.com/flipped-aurora/gin-vue-admin/server/model/ProblemRecordPkg/request"
	"github.com/flipped-aurora/gin-vue-admin/server/model/common/request"
	"github.com/gin-gonic/gin"
	"path/filepath"
	"strings"
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

// GetProblemRecordInfoListByUser 根据用户权限分页获取ProblemRecord记录
// Author [piexlmax](https://github.com/piexlmax)
func (pbRecordService *ProblemRecordService) GetProblemRecordInfoListByUser(info ProblemRecordPkgReq.ProblemRecordSearch, nestIdArr []string) (list []ProblemRecordPkg.ProblemRecord, total int64, err error) {
	limit := info.PageSize
	offset := info.PageSize * (info.Page - 1)
	// 创建db
	db := global.GVA_DB.Model(&ProblemRecordPkg.ProblemRecord{})
	var pbRecords []ProblemRecordPkg.ProblemRecord
	if len(nestIdArr) > 0 {
		for _, nestId := range nestIdArr {
			db = db.Or("locate(?,nestid) > 0", nestId)
		}
	}
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
		err = db.Debug().Limit(limit).Offset(offset).Find(&pbRecords).Error
	}

	return pbRecords, total, err
}

// GetProblemRecordInfoList 分页获取ProblemRecord记录
// Author [piexlmax](https://github.com/piexlmax)
func (pbRecordService *ProblemRecordService) GetProblemRecordInfoList(info ProblemRecordPkgReq.ProblemRecordSearch, c *gin.Context) (list []ProblemRecordPkg.ProblemRecord, total int64, err error) {
	limit := info.PageSize
	offset := info.PageSize * (info.Page - 1)
	// 创建db
	db := global.GVA_DB.Model(&ProblemRecordPkg.ProblemRecord{})
	var pbRecords []ProblemRecordPkg.ProblemRecord
	var rtPbRecords []ProblemRecordPkg.ProblemRecord
	// 如果有条件搜索 下方会自动创建搜索语句
	if info.ExecRecordId != "" {
		db = db.Where(" exec_record_id = ?", info.ExecRecordId)
	}
	if info.MultiSpectraId != "" {
		db = db.Where(" multi_spectra_id = ?", info.MultiSpectraId)
	}
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

	for _, record := range pbRecords {
		if record.MultiTypeId > 0 {
			mtType := make(map[string]interface{})
			first := global.GVA_DB.Raw("select id ID, created_at CreatedAt, aerial_photography_id AerialPhotographyId, aerial_server_address aerialServerAddress, spectra_type spectraType, (select position from aerial_photography_result where id = mst.aerial_photography_id) position , (select JSON_EXTRACT(aerial_photography_file, '$[0].url') from aerial_photography_result where id = mst.aerial_photography_id) FileUrl from multi_spectra_type mst where id = ? limit 1", record.MultiTypeId).First(&mtType)
			if first.Error != nil {
			} else {
				if mtType["FileUrl"] != nil {
					url := mtType["FileUrl"].(string)[1:strings.LastIndex(mtType["FileUrl"].(string), ".")]
					filepaths := filepath.Join(url, mtType["spectraType"].(string), "{z}", "{x}", "{y}"+".png")
					mtType["FileUrl"] = &filepaths
				}
				record.MultiSpectraType = mtType
			}

		}

		rtPbRecords = append(rtPbRecords, record)
	}
	return rtPbRecords, total, err
}
