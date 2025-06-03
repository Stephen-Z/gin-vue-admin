package MultiSpectraAnalysisPkg

import (
	"errors"
	"github.com/flipped-aurora/gin-vue-admin/server/global"
	"github.com/flipped-aurora/gin-vue-admin/server/model/MultiSpectraAnalysisPkg"
	MultiSpectraAnalysisPkgReq "github.com/flipped-aurora/gin-vue-admin/server/model/MultiSpectraAnalysisPkg/request"
	"github.com/flipped-aurora/gin-vue-admin/server/model/common/request"
	"github.com/flipped-aurora/gin-vue-admin/server/service/NestInfo"
	"github.com/gin-gonic/gin"
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
func (MtSpectraAlyService *MultiSpectraAnalysisService) GetMultiSpectraAnalysisInfoList(info MultiSpectraAnalysisPkgReq.MultiSpectraAnalysisSearch, c *gin.Context) (list []MultiSpectraAnalysisPkg.MultiSpectraAnalysis, total int64, err error) {
	limit := info.PageSize
	offset := info.PageSize * (info.Page - 1)
	// 创建db
	db := global.GVA_DB.Model(&MultiSpectraAnalysisPkg.MultiSpectraAnalysis{})
	var MtSpectraAlys []MultiSpectraAnalysisPkg.MultiSpectraAnalysis
	var RtMtSpectraAlys []MultiSpectraAnalysisPkg.MultiSpectraAnalysis

	//permission 查出所有登录用户下所属机巢下所有航摄成果的光谱分析id集
	ids, err := MtSpectraAlyService.QueryAerialPhotographyTypeIds(c)
	db.Where(" deleted_at is null")
	if ids != "" {
		db.Where("id in (" + ids + ")")
	}

	// 如果有条件搜索 下方会自动创建搜索语句
	if info.StartCreatedAt != nil && info.EndCreatedAt != nil {
		db = db.Where("created_at BETWEEN ? AND ?", info.StartCreatedAt, info.EndCreatedAt)
	}

	db = db.Order("created_at desc")
	err = db.Count(&total).Error
	if err != nil {
		return
	}

	if info.PageSize > 0 && info.Page > 0 {
		db.Limit(limit).Offset(offset)
	}
	err = db.Find(&MtSpectraAlys).Error

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

// QueryAerialPhotographyTypeIds 根据机巢id查询该用户所属机巢下的光谱分析id集
func (MtSpectraAlyService *MultiSpectraAnalysisService) QueryAerialPhotographyTypeIds(c *gin.Context) (Ids string, err error) {
	db := global.GVA_DB.Model(&MultiSpectraAnalysisPkg.MultiSpectraAnalysis{})
	nestInfoService := new(NestInfo.NestInfoService)
	nestIDList, err := nestInfoService.GetNestIDListByUser(c)
	if err != nil {
		return "", errors.New("query nest is nil")
	}
	querySql := "select group_concat(distinct id) aerial_photography_ids from (select id, SUBSTRING_INDEX(SUBSTRING_INDEX(aerial_photography_ids, ',', 1), ',', -1) as aerial_photography_id from multi_spectra_analysis where aerial_photography_ids <> '' union all select id, SUBSTRING_INDEX(SUBSTRING_INDEX(aerial_photography_ids, ',', 2), ',', -1) as aerial_photography_id from multi_spectra_analysis where aerial_photography_ids like '%,%' and aerial_photography_ids <> '' union all select id, SUBSTRING_INDEX(SUBSTRING_INDEX(aerial_photography_ids, ',', 3), ',', -1) as aerial_photography_id from multi_spectra_analysis where aerial_photography_ids like '%,%,%' and aerial_photography_ids <> '') tab where 1 = 1 and aerial_photography_id in (select id from aerial_photography_result where 1 = 1 and status = 2 and load_or_not = 0 and deleted_by = 0  "
	//querySql := "select group_concat(distinct id) aerial_photography_ids from aerial_photography_result where 1 = 1 and status = 2 and load_or_not = 0 and deleted_by = 0    "
	if len(nestIDList) > 0 {
		sqlWhere := " "
		for i, str := range nestIDList {
			if i == 0 {
				//db.Where("nest_ids like ?", str)
				sqlWhere += "nest_ids like  '%" + str + "%'"
			} else {
				//db.Or("nest_ids like ?", str)
				sqlWhere += "or nest_ids like  '%" + str + "%'"
			}
		}
		sqlWhere = "(" + sqlWhere + "))"
		querySql += " and " + sqlWhere
	}
	queryErr := db.Raw(querySql).First(&Ids).Error
	return Ids, queryErr
}
