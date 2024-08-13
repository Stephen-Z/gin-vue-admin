package MultiSpectraAnalysisPkg

import (
	"github.com/flipped-aurora/gin-vue-admin/server/global"
	"github.com/flipped-aurora/gin-vue-admin/server/model/MultiSpectraAnalysisPkg"
	MultiSpectraAnalysisPkgReq "github.com/flipped-aurora/gin-vue-admin/server/model/MultiSpectraAnalysisPkg/request"
	"github.com/flipped-aurora/gin-vue-admin/server/model/common/request"
)

type MultiSpectraAnalysisService struct {
}

// CreateMultiSpectraAnalysis 创建MultiSpectraAnalysis记录
// Author [piexlmax](https://github.com/piexlmax)
func (MtSpectraAlyService *MultiSpectraAnalysisService) CreateMultiSpectraAnalysis(MtSpectraAly *MultiSpectraAnalysisPkg.MultiSpectraAnalysis) (err error) {
	err = global.GVA_DB.Create(MtSpectraAly).Error
	return err
}

// DeleteMultiSpectraAnalysis 删除MultiSpectraAnalysis记录
// Author [piexlmax](https://github.com/piexlmax)
func (MtSpectraAlyService *MultiSpectraAnalysisService) DeleteMultiSpectraAnalysis(MtSpectraAly MultiSpectraAnalysisPkg.MultiSpectraAnalysis) (err error) {
	err = global.GVA_DB.Delete(&MtSpectraAly).Error
	return err
}

// DeleteMultiSpectraAnalysisByIds 批量删除MultiSpectraAnalysis记录
// Author [piexlmax](https://github.com/piexlmax)
func (MtSpectraAlyService *MultiSpectraAnalysisService) DeleteMultiSpectraAnalysisByIds(ids request.IdsReq) (err error) {
	err = global.GVA_DB.Delete(&[]MultiSpectraAnalysisPkg.MultiSpectraAnalysis{}, "id in ?", ids.Ids).Error
	return err
}

// UpdateMultiSpectraAnalysis 更新MultiSpectraAnalysis记录
// Author [piexlmax](https://github.com/piexlmax)
func (MtSpectraAlyService *MultiSpectraAnalysisService) UpdateMultiSpectraAnalysis(MtSpectraAly MultiSpectraAnalysisPkg.MultiSpectraAnalysis) (err error) {
	err = global.GVA_DB.Save(&MtSpectraAly).Error
	return err
}

// GetMultiSpectraAnalysis 根据id获取MultiSpectraAnalysis记录
// Author [piexlmax](https://github.com/piexlmax)
func (MtSpectraAlyService *MultiSpectraAnalysisService) GetMultiSpectraAnalysis(id uint) (MtSpectraAly MultiSpectraAnalysisPkg.MultiSpectraAnalysis, err error) {
	err = global.GVA_DB.Where("id = ?", id).First(&MtSpectraAly).Error
	return
}

// GetMultiSpectraAnalysisInfoList 分页获取MultiSpectraAnalysis记录
// Author [piexlmax](https://github.com/piexlmax)
func (MtSpectraAlyService *MultiSpectraAnalysisService) GetMultiSpectraAnalysisInfoList(info MultiSpectraAnalysisPkgReq.MultiSpectraAnalysisSearch) (list []MultiSpectraAnalysisPkg.MultiSpectraAnalysis, total int64, err error) {
	limit := info.PageSize
	offset := info.PageSize * (info.Page - 1)
	// 创建db
	db := global.GVA_DB.Model(&MultiSpectraAnalysisPkg.MultiSpectraAnalysis{})
	var MtSpectraAlys []MultiSpectraAnalysisPkg.MultiSpectraAnalysis
	var RtMtSpectraAlys []MultiSpectraAnalysisPkg.MultiSpectraAnalysis

	// 如果有条件搜索 下方会自动创建搜索语句
	if info.StartCreatedAt != nil && info.EndCreatedAt != nil {
		db = db.Where("created_at BETWEEN ? AND ?", info.StartCreatedAt, info.EndCreatedAt)
	}

	db = db.Order("created_at desc")
	err = db.Count(&total).Error
	if err != nil {
		return
	}

	err = db.Limit(limit).Offset(offset).Find(&MtSpectraAlys).Error

	type TmpPbParam struct {
		Id           uint `json:"id" gorm:"column:id;comment:;"`
		ProblemCount int  `json:"problem_count" gorm:"column:problem_count;comment:;"`
	}
	var tmpPbParamArr []TmpPbParam
	queryPbRecordCountSql := " select msa.id, (select count(1) from problem_record where multi_spectra_id = msa.id and deleted_at is null) problem_count from multi_spectra_analysis msa where deleted_at is null "
	global.GVA_DB.Raw(queryPbRecordCountSql).Find(&tmpPbParamArr)
	for _, aly := range MtSpectraAlys {
		for _, pb := range tmpPbParamArr {
			if aly.ID == pb.Id {
				aly.ProblemCount = pb.ProblemCount
				RtMtSpectraAlys = append(RtMtSpectraAlys, aly)
			}
		}
	}
	return RtMtSpectraAlys, total, err
}

// GetMultiSpectraAnalysisInfoList 分页获取MultiSpectraAnalysis记录并统计问题记录数
// Author [piexlmax](https://github.com/piexlmax)
func (MtSpectraAlyService *MultiSpectraAnalysisService) GetMultiSpectraAnalysisInfoWithPbRecordCountList(info MultiSpectraAnalysisPkgReq.MultiSpectraAnalysisSearch) (list []MultiSpectraAnalysisPkg.MultiSpectraAnalysis, total int64, err error) {
	limit := info.PageSize
	offset := info.PageSize * (info.Page - 1)
	querySql := "select msa.*, (select count(1) from problem_record where multi_spectra_id = msa.id) problem_count from multi_spectra_analysis msa where deleted_at is null order by created_at "
	// 创建db
	db := global.GVA_DB.Model(&MultiSpectraAnalysisPkg.MultiSpectraAnalysis{})
	var MtSpectraAlys []MultiSpectraAnalysisPkg.MultiSpectraAnalysis
	err = db.Count(&total).Error
	if err != nil {
		return
	}
	err = db.Limit(limit).Offset(offset).Raw(querySql).Find(&MtSpectraAlys).Error
	return MtSpectraAlys, total, err
}
