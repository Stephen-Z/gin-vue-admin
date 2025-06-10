package FlyResultPkg

import (
	"archive/zip"
	"bytes"
	"crypto/md5"
	"fmt"
	"github.com/flipped-aurora/gin-vue-admin/server/global"
	"github.com/flipped-aurora/gin-vue-admin/server/model/FlyResultPkg"
	FlyResultPkgReq "github.com/flipped-aurora/gin-vue-admin/server/model/FlyResultPkg/request"
	NestAirlinePkgReq "github.com/flipped-aurora/gin-vue-admin/server/model/NestAirlinePkg/request"
	"github.com/flipped-aurora/gin-vue-admin/server/model/common/request"
	"github.com/flipped-aurora/gin-vue-admin/server/model/common/response"
	"github.com/flipped-aurora/gin-vue-admin/server/service"
	"github.com/flipped-aurora/gin-vue-admin/server/utils"
	"github.com/gin-gonic/gin"
	"go.uber.org/zap"
	"io"
	"log"
	"os"
	"path"
	"path/filepath"
	"strings"
	"time"
)

func init() {
	go TimeToClearCompressFile()
}

type FlyResultApi struct {
}

var NestAirlineService = service.ServiceGroupApp.NestAirlinePkgServiceGroup.NestAirlineService
var NestExecRecordService = service.ServiceGroupApp.NestExecRecordPkgServiceGroup.NestExecRecordService
var FlyRtService = service.ServiceGroupApp.FlyResultPkgServiceGroup.FlyResultService

// CreateFlyResult 创建FlyResult
// @Tags FlyResult
// @Summary 创建FlyResult
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Param data body FlyResultPkg.FlyResult true "创建FlyResult"
// @Success 200 {string} string "{"success":true,"data":{},"msg":"获取成功"}"
// @Router /FlyRt/createFlyResult [post]
func (FlyRtApi *FlyResultApi) CreateFlyResult(c *gin.Context) {
	var FlyRt FlyResultPkg.FlyResult
	err := c.ShouldBindJSON(&FlyRt)
	if err != nil {
		response.FailWithMessage(err.Error(), c)
		return
	}
	FlyRt.CreatedBy = utils.GetUserID(c)
	if err := FlyRtService.CreateFlyResult(&FlyRt); err != nil {
		global.GVA_LOG.Error("创建失败!", zap.Error(err))
		response.FailWithMessage("创建失败", c)
	} else {
		response.OkWithMessage("创建成功", c)
	}
}

// DeleteFlyResult 删除FlyResult
// @Tags FlyResult
// @Summary 删除FlyResult
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Param data body FlyResultPkg.FlyResult true "删除FlyResult"
// @Success 200 {string} string "{"success":true,"data":{},"msg":"删除成功"}"
// @Router /FlyRt/deleteFlyResult [delete]
func (FlyRtApi *FlyResultApi) DeleteFlyResult(c *gin.Context) {
	var FlyRt FlyResultPkg.FlyResult
	err := c.ShouldBindJSON(&FlyRt)
	if err != nil {
		response.FailWithMessage(err.Error(), c)
		return
	}
	FlyRt.DeletedBy = utils.GetUserID(c)
	if err := FlyRtService.DeleteFlyResult(FlyRt); err != nil {
		global.GVA_LOG.Error("删除失败!", zap.Error(err))
		response.FailWithMessage("删除失败", c)
	} else {
		response.OkWithMessage("删除成功", c)
	}
}

// DeleteFlyResultByIds 批量删除FlyResult
// @Tags FlyResult
// @Summary 批量删除FlyResult
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Param data body request.IdsReq true "批量删除FlyResult"
// @Success 200 {string} string "{"success":true,"data":{},"msg":"批量删除成功"}"
// @Router /FlyRt/deleteFlyResultByIds [delete]
func (FlyRtApi *FlyResultApi) DeleteFlyResultByIds(c *gin.Context) {
	var IDS request.IdsReq
	err := c.ShouldBindJSON(&IDS)
	if err != nil {
		response.FailWithMessage(err.Error(), c)
		return
	}
	deletedBy := utils.GetUserID(c)
	//清空作业记录的全景链接
	if len(IDS.RecordIds) > 0 {
		for _, rId := range IDS.RecordIds {
			record, GetNErr := NestExecRecordService.GetNestExecRecord(uint(rId))
			if GetNErr == nil {
				record.PanoramaLink = ""
				record.ID = uint(rId)
				NestExecRecordService.UpdateNestExecRecord(record)
			}
		}
	}
	if err := FlyRtService.DeleteFlyResultByIds(IDS, deletedBy); err != nil {
		global.GVA_LOG.Error("批量删除失败!", zap.Error(err))
		response.FailWithMessage("批量删除失败", c)
	} else {
		response.OkWithMessage("批量删除成功", c)
	}
}

// UpdateFlyResult 更新FlyResult
// @Tags FlyResult
// @Summary 更新FlyResult
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Param data body FlyResultPkg.FlyResult true "更新FlyResult"
// @Success 200 {string} string "{"success":true,"data":{},"msg":"更新成功"}"
// @Router /FlyRt/updateFlyResult [put]
func (FlyRtApi *FlyResultApi) UpdateFlyResult(c *gin.Context) {
	var FlyRt FlyResultPkg.FlyResult
	err := c.ShouldBindJSON(&FlyRt)
	if err != nil {
		response.FailWithMessage(err.Error(), c)
		return
	}
	FlyRt.UpdatedBy = utils.GetUserID(c)
	if err := FlyRtService.UpdateFlyResult(FlyRt); err != nil {
		global.GVA_LOG.Error("更新失败!", zap.Error(err))
		response.FailWithMessage("更新失败", c)
	} else {
		response.OkWithMessage("更新成功", c)
	}
}

// FindFlyResult 用id查询FlyResult
// @Tags FlyResult
// @Summary 用id查询FlyResult
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Param data query FlyResultPkg.FlyResult true "用id查询FlyResult"
// @Success 200 {string} string "{"success":true,"data":{},"msg":"查询成功"}"
// @Router /FlyRt/findFlyResult [get]
func (FlyRtApi *FlyResultApi) FindFlyResult(c *gin.Context) {
	var FlyRt FlyResultPkg.FlyResult
	err := c.ShouldBindQuery(&FlyRt)
	if err != nil {
		response.FailWithMessage(err.Error(), c)
		return
	}
	if reFlyRt, err := FlyRtService.GetFlyResult(FlyRt.ID); err != nil {
		global.GVA_LOG.Error("查询失败!", zap.Error(err))
		response.FailWithMessage("查询失败", c)
	} else {
		response.OkWithData(gin.H{"reFlyRt": reFlyRt}, c)
	}
}

// GetFlyResultList 分页获取FlyResult列表
// @Tags FlyResult
// @Summary 分页获取FlyResult列表
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Param data query FlyResultPkgReq.FlyResultSearch true "分页获取FlyResult列表"
// @Success 200 {string} string "{"success":true,"data":{},"msg":"获取成功"}"
// @Router /FlyRt/getFlyResultList [get]
func (FlyRtApi *FlyResultApi) GetFlyResultList(c *gin.Context) {
	var pageInfo FlyResultPkgReq.FlyResultSearch
	err := c.ShouldBindQuery(&pageInfo)
	if err != nil {
		response.FailWithMessage(err.Error(), c)
		return
	}
	if list, total, err := FlyRtService.GetFlyResultInfoList(pageInfo, c); err != nil {
		global.GVA_LOG.Error("获取失败!", zap.Error(err))
		response.FailWithMessage("获取失败", c)
	} else {
		response.OkWithDetailed(response.PageResult{
			List:     list,
			Total:    total,
			Page:     pageInfo.Page,
			PageSize: pageInfo.PageSize,
		}, "获取成功", c)
	}
}

// QueryAirlineRecordFlyResult 查询航线下所有作业记录及成果
// @Tags FlyResult
// @Summary 查询航线下所有作业记录及成果
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Param data query FlyResultPkgReq.FlyResultSearch true "查询航线下所有作业记录及成果"
// @Success 200 {string} string "{"success":true,"data":{},"msg":"获取成功"}"
// @Router /FlyRt/getFlyResultList [get]
func (FlyRtApi *FlyResultApi) QueryAirlineRecordFlyResult(c *gin.Context) {
	//var NtAirline NestAirlinePkg.NestAirline
	var pageInfo NestAirlinePkgReq.NestAirlineSearch
	err := c.ShouldBindQuery(&pageInfo)
	if err != nil {
		response.FailWithMessage(err.Error(), c)
		return
	}

	//获取航线列表
	airlineList, count, err := NestAirlineService.NoPageGetNestAirlineInfoList(pageInfo, c)
	if err != nil {
		global.GVA_LOG.Error("获取失败!", zap.Error(err))
		response.FailWithMessage("获取失败", c)
	}
	//获取记录列表
	recordList, err := NestExecRecordService.NoPageGetNestExecRecordInfoList(c)
	if err != nil {
		global.GVA_LOG.Error("获取失败!", zap.Error(err))
		response.FailWithMessage("获取失败", c)
	}
	//获取成果列表
	resultList, err := FlyRtService.NoPageGetFlyResultInfoList(c)
	if err != nil {
		global.GVA_LOG.Error("获取失败!", zap.Error(err))
		response.FailWithMessage("获取失败", c)
	}
	for _, record := range recordList {
		imgCount := 0
		videoCount := 0
		for _, result := range resultList {
			if result["execute_id"] != nil && result["file_name"] != nil {
				fileName := result["file_name"].(string)
				if result["type"] == 1 {
					result["file_link"] = global.GVA_CONFIG.FileServer.MainLink + "\\" + path.Join("", result["execute_id"].(string), fileName)
					result["file_thumbnails_link"] = global.GVA_CONFIG.FileServer.MainLink + "\\" + path.Join(result["execute_id"].(string), global.GVA_CONFIG.FileServer.ThumbnailsImgLink, fileName)
				} else {
					result["file_link"] = global.GVA_CONFIG.FileServer.MainLink + "\\" + path.Join("", result["execute_id"].(string), fileName)
					result["file_thumbnails_link"] = global.GVA_CONFIG.FileServer.MainLink + "\\" + path.Join(result["execute_id"].(string), global.GVA_CONFIG.FileServer.ThumbnailsImgLink, strings.Replace(fileName, ".MP4", ".jpg", -1))
				}

			} else {
				result["file_link"] = ""
				result["file_thumbnails_link"] = ""
			}
			if _, exist := record["fly_result_arr"]; !exist {
				resArr := make([]map[string]interface{}, 0, 0)
				record["fly_result_arr"] = resArr
			}
			if record["execute_id"] == result["execute_id"] {
				if _, exist := record["fly_result_arr"]; exist {
					record["fly_result_arr"] = append(record["fly_result_arr"].([]map[string]interface{}), result)
				}

				if result["type"] != nil {
					if *result["type"].(*int) == 1 {
						//照片
						imgCount += 1
					} else if *result["type"].(*int) == 2 {
						//视频
						videoCount += 1
					}
				}
			}
		}
		record["img_count"] = imgCount
		record["video_count"] = videoCount
	}

	for _, airline := range airlineList {
		panoramaCount := 0
		imgCount := 0
		videoCount := 0
		//executeId := ""
		for _, record := range recordList {
			if _, exist := airline["exec_record_arr"]; !exist {
				resArr := make([]map[string]interface{}, 0, 0)
				airline["exec_record_arr"] = resArr
			}
			if airline["missionid"] == record["missionid"] {
				if _, exist := airline["exec_record_arr"]; exist {
					if _, ok := airline["task_max_new_time"]; !ok {
						airline["task_max_new_time"] = record["execute_at"]
					}
					airline["exec_record_arr"] = append(airline["exec_record_arr"].([]map[string]interface{}), record)
				}

				if record["panorama_link"].(string) != "" {
					panoramaCount += 1
				}
				if record["img_count"] != nil {
					imgCount = imgCount + record["img_count"].(int)
				}
				if record["video_count"] != nil {
					videoCount = videoCount + record["video_count"].(int)
				}
				//if record["execute_id"] != nil {
				//	executeId = record["execute_id"].(string)
				//}
			}
		}
		airline["panorama_count"] = panoramaCount
		airline["img_count"] = imgCount
		airline["video_count"] = videoCount
		//airline["execute_id"] = executeId
	}
	response.OkWithData(gin.H{"airlineList": airlineList, "total": count, "page": pageInfo.Page, "pagesize": pageInfo.PageSize}, c)
}

// DataresultDownload 数据成果打包下载
// @Tags FlyResult
// @Summary 数据成果打包下载
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Param data query FlyResultPkgReq.FlyResultSearch true "数据成果打包下载"
// @Success 200 {string} string "{"success":true,"data":{},"msg":"获取成功"}"
// @Router /FlyRt/getFlyResultList [get]
func (FlyRtApi *FlyResultApi) DataResultDownload(c *gin.Context) {
	var Reqs request.IdsReq
	err := c.ShouldBindJSON(&Reqs)
	if err != nil {
		response.FailWithMessage(err.Error(), c)
		return
	}
	if len(Reqs.DataPaths) > 0 {
		fileList := make([]string, 0, 0)
		finalPath := ""
		for _, rPath := range Reqs.DataPaths {
			if rPath != "" && global.GVA_CONFIG.FileServer.FlyRecordPath != "" {
				completePath := path.Join(global.GVA_CONFIG.FileServer.FlyRecordPath, rPath)
				fileList = append(fileList, completePath)
				finalPath += rPath
			}
		}
		encodeFileName := md5.Sum([]byte(finalPath))
		zipFileName := fmt.Sprintf("%x.zip", encodeFileName)
		//zipFileName := fmt.Sprintf("%s_%v-result.zip", uuid.NewString(), time.Now().Format("2006-01-02-15-04-05"))
		_, exist := os.Stat(path.Join(global.GVA_CONFIG.FileServer.FlyRecordCompressDownload, zipFileName))
		if exist == nil {
			//已存在同文件的压缩包
			response.OkWithData(gin.H{"filePath": path.Join(global.GVA_CONFIG.FileServer.FlyRecordCompressDownload, zipFileName)}, c)
		} else {
			//不存在
			if zipErr := utils.ZipFiles(path.Join(global.GVA_CONFIG.FileServer.FlyRecordCompressDownload, zipFileName), fileList, global.GVA_CONFIG.FileServer.FlyRecordPath, ""); err != nil {
				log.Println(zipErr.Error())
				response.FailWithMessage("打包失败", c)
			} else {
				//c.Header("Content-Type", "application/octet-stream")
				//c.Header("Content-Disposition", "attachment; filename="+zipFileName) //添加此header触发http下载动作
				//c.Header("Cache-Control", "no-cache")
				//c.File("./" + zipFileName)
				response.OkWithData(gin.H{"filePath": path.Join(global.GVA_CONFIG.FileServer.FlyRecordCompressDownload, zipFileName)}, c)
			}
		}
	} else {
		response.FailWithMessage("请勾选文件", c)
	}
}

// DataResultDownload 根据 表 nest_exec_record 的 execute_id 返回所有 关于正射的文件
func (FlyRtApi *FlyResultApi) DataResultDownloadSource(c *gin.Context) {

	type RequestObj struct {
		ExecuteId string `json:"execute_id" binding:"required"`
	}
	var requestObj RequestObj
	errShouldBindJSON := c.ShouldBindJSON(&requestObj)
	if errShouldBindJSON != nil {
		response.FailWithMessage(errShouldBindJSON.Error(), c)
	}
	execRecord, errGetNestExecRecordByExecuteId := NestExecRecordService.GetNestExecRecordByExecuteId(requestObj.ExecuteId)
	if errGetNestExecRecordByExecuteId != nil {
		c.JSON(200, map[string]interface{}{"code": -1, "msg": "该条记录不存在"})
		return
	}
	if execRecord.Status == nil || (*execRecord.Status != int(6) && (*execRecord.Status != int(7))) {
		c.JSON(200, map[string]interface{}{"code": -1, "msg": "该条记录格式不正确"})
		return
	}

	filePathString := path.Join(global.GVA_CONFIG.FileServer.FlyRecordPath, requestObj.ExecuteId) // 线上的文件夹
	//filePathString := path.Join("D:", "图片", "ZHXF-DATA", "flyrecords", requestObj.ExecuteId) // 本地测试的文件夹
	binary, errCompressFolderToBinary := compressFolderToBinary(filePathString)
	if errCompressFolderToBinary != nil {
		c.JSON(200, map[string]interface{}{"code": -1, "msg": errCompressFolderToBinary.Error()})
		return
	}

	c.Header("Content-Disposition", "attachment; filename="+requestObj.ExecuteId+".zip")
	c.Data(200, "application/octet-stream", binary)
	return
}

// 每日凌晨定时清除压缩完的作业成果
func TimeToClearCompressFile() {
	t := time.NewTimer(SetTime(0, 1, 0))
	defer t.Stop()
	for {
		select {
		case <-t.C:
			t.Reset(time.Hour * 24)
			// 定时任务函数
			log.Printf("凌晨删除打包的压缩文件: %v\n", time.Now())
			//以下为定时执行的操作
			err := os.RemoveAll(global.GVA_CONFIG.FileServer.FlyRecordCompressDownload)
			if err != nil {
				global.GVA_LOG.Error("删除作业成果压缩包失败!", zap.Error(err))
			} else {
				global.GVA_LOG.Info("删除作业成果压缩包成功!", zap.Error(err))
				os.MkdirAll(global.GVA_CONFIG.FileServer.FlyRecordCompressDownload, os.ModePerm)
			}
		}
	}
}
func SetTime(hour, min, second int) (d time.Duration) {
	now := time.Now()
	setTime := time.Date(now.Year(), now.Month(), now.Day(), hour, min, second, 0, now.Location())
	d = setTime.Sub(now)
	if d > 0 {
		return
	}
	return d + time.Hour*24
}

// 压缩文件夹 并返回[]byte 格式为.zip
func compressFolderToBinary(srcDir string) ([]byte, error) {
	// 使用 bytes.Buffer 来保存 zip 数据
	buf := new(bytes.Buffer)
	zipWriter := zip.NewWriter(buf)

	// 遍历源文件夹中的所有文件并添加到 zip
	err := filepath.Walk(srcDir, func(filePath string, info os.FileInfo, err error) error {
		if err != nil {
			return err
		}

		// 构建 ZIP 中的相对路径
		relPath, err := filepath.Rel(srcDir, filePath)
		if err != nil {
			return err
		}

		// 创建 header
		header, err := zip.FileInfoHeader(info)
		if err != nil {
			return err
		}

		header.Name = relPath
		// 如果是目录，则确保路径以斜杠结尾
		if info.IsDir() {
			header.Name += "/"
		} else {
			header.Method = zip.Deflate
		}

		// 创建 writer
		writer, err := zipWriter.CreateHeader(header)
		if err != nil {
			return err
		}

		// 如果不是目录，则写入文件内容
		if !info.IsDir() {
			file, errOpen := os.Open(filePath)
			if errOpen != nil {
				return errOpen
			}
			defer file.Close()
			_, errCopy := io.Copy(writer, file)
			if errCopy != nil {
				return errCopy
			}
		}
		return nil
	})

	// 确保关闭 zip.Writer 以完成写入操作
	if closeErr := zipWriter.Close(); closeErr != nil && err == nil {
		err = closeErr
	}

	if err != nil {
		return nil, err
	}

	return buf.Bytes(), nil
}
