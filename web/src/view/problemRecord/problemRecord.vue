<template>
  <div>
    <div class="gva-search-box">
      <el-form ref="elSearchFormRef" :inline="true" :model="searchInfo" class="demo-form-inline" :rules="searchRule" @keyup.enter="onSubmit">
      <el-form-item label="创建日期" prop="createdAt">
      <template #label>
        <span>
          创建日期
          <el-tooltip content="搜索范围是开始日期（包含）至结束日期（不包含）">
            <el-icon><QuestionFilled /></el-icon>
          </el-tooltip>
        </span>
      </template>
      <el-date-picker v-model="searchInfo.startCreatedAt" type="datetime" placeholder="开始日期" :disabled-date="time=> searchInfo.endCreatedAt ? time.getTime() > searchInfo.endCreatedAt.getTime() : false"></el-date-picker>
       —
      <el-date-picker v-model="searchInfo.endCreatedAt" type="datetime" placeholder="结束日期" :disabled-date="time=> searchInfo.startCreatedAt ? time.getTime() < searchInfo.startCreatedAt.getTime() : false"></el-date-picker>
      </el-form-item>
        <el-form-item>
          <el-button type="primary" icon="search" @click="onSubmit">查询</el-button>
          <el-button icon="refresh" @click="onReset">重置</el-button>
        </el-form-item>
      </el-form>
    </div>
    <div class="gva-table-box">
        <div class="gva-btn-list">
            <el-button type="primary" icon="plus" @click="openDialog">新增</el-button>
            <el-popover v-model:visible="deleteVisible" placement="top" width="160">
            <p>确定要删除吗？</p>
            <div style="text-align: right; margin-top: 8px;">
                <el-button type="primary" link @click="deleteVisible = false">取消</el-button>
                <el-button type="primary" @click="onDelete">确定</el-button>
            </div>
            <template #reference>
                <el-button icon="delete" style="margin-left: 10px;" :disabled="!multipleSelection.length" @click="deleteVisible = true">删除</el-button>
            </template>
            </el-popover>
        </div>
        <el-table
        ref="multipleTable"
        style="width: 100%"
        tooltip-effect="dark"
        :data="tableData"
        row-key="ID"
        @selection-change="handleSelectionChange"
        >
        <el-table-column type="selection" width="55" />
        <el-table-column align="left" label="日期" width="180">
            <template #default="scope">{{ formatDate(scope.row.CreatedAt) }}</template>
        </el-table-column>
        <el-table-column align="left" label="序号" prop="orderNum" width="120" />
         <el-table-column align="left" label="登记日期" width="180">
            <template #default="scope">{{ formatDate(scope.row.registerDate) }}</template>
         </el-table-column>
        <el-table-column align="left" label="问题名称" prop="problemName" width="120" />
        <el-table-column align="left" label="问题类型" prop="problemType" width="120" />
        <el-table-column align="left" label="问题来源" prop="problemSource" width="120" />
        <el-table-column align="left" label="问题描述" prop="problemDesc" width="120" />
        <el-table-column align="left" label="负责人" prop="personCharge" width="120" />
        <el-table-column align="left" label="作业记录Id" prop="execRecordId" width="120" />
          <el-table-column label="问题图片" width="200">
              <template #default="scope">
                <el-image style="width: 100px; height: 100px" :src="getUrl(scope.row.problemImage)" fit="cover"/>
              </template>
          </el-table-column>
        <el-table-column align="left" label="处理措施" prop="handMeasurce" width="120" />
        <el-table-column align="left" label="当前状态" prop="currState" width="120" />
        <el-table-column align="left" label="经度" prop="lng" width="120" />
        <el-table-column align="left" label="纬度" prop="lat" width="120" />
        <el-table-column align="left" label="机巢列表" prop="nestid" width="120" />
        <el-table-column align="left" label="光谱id" prop="multiSpectraId" width="120" />
         <el-table-column align="left" label="光谱类型id" prop="multiTypeId" width="120" />
        <el-table-column align="left" label="航摄成果id" prop="AerialPhotographyId" width="120" />
        <el-table-column align="left" label="航摄成果地址" prop="aerialServerAddress" width="120" />
        <el-table-column align="left" label="操作">
            <template #default="scope">
            <el-button type="primary" link icon="edit" class="table-button" @click="updateProblemRecordFunc(scope.row)">变更</el-button>
            <el-button type="primary" link icon="delete" @click="deleteRow(scope.row)">删除</el-button>
            </template>
        </el-table-column>
        </el-table>
        <div class="gva-pagination">
            <el-pagination
            layout="total, sizes, prev, pager, next, jumper"
            :current-page="page"
            :page-size="pageSize"
            :page-sizes="[10, 30, 50, 100]"
            :total="total"
            @current-change="handleCurrentChange"
            @size-change="handleSizeChange"
            />
        </div>
    </div>
    <el-dialog v-model="dialogFormVisible" :before-close="closeDialog" :title="type==='create'?'添加':'修改'" destroy-on-close>
      <el-form :model="formData" label-position="right" ref="elFormRefRecord" :rules="rule" label-width="80px">
        <el-form-item label="序号:"  prop="orderNum" >
          <el-input v-model.number="formData.orderNum" :clearable="true" placeholder="请输入" />
        </el-form-item>
         <el-form-item label="作业记录Id:"  prop="problemName" >
          <el-input v-model="formData.execRecordId" :clearable="true"  placeholder="请输入" />
        </el-form-item>
        <el-form-item label="登记日期:"  prop="registerDate" >
          <el-date-picker v-model="formData.registerDate" type="date" style="width:100%" placeholder="选择日期" :clearable="true"  />
        </el-form-item>
        <el-form-item label="问题名称:"  prop="problemName" >
          <el-input v-model="formData.problemName" :clearable="true"  placeholder="请输入" />
        </el-form-item>
        <el-form-item label="问题类型:"  prop="problemType" >
          <el-input v-model="formData.problemType" :clearable="true"  placeholder="请输入" />
        </el-form-item>
        <el-form-item label="问题来源:"  prop="problemSource" >
          <el-input v-model="formData.problemSource" :clearable="true"  placeholder="请输入" />
        </el-form-item>
        <el-form-item label="问题描述:"  prop="problemDesc" >
          <el-input v-model="formData.problemDesc" :clearable="true"  placeholder="请输入" />
        </el-form-item>
        <el-form-item label="负责人:"  prop="personCharge" >
          <el-input v-model="formData.personCharge" :clearable="true"  placeholder="请输入" />
        </el-form-item>
         <el-form-item label="光谱类型Id:"  prop="multiTypeId" >
          <el-input v-model.number="formData.multiTypeId" :clearable="true"  placeholder="请输入" />
        </el-form-item>
        <el-form-item label="问题图片:"  prop="problemImage" >
            <SelectImage v-model="formData.problemImage" :is-multiple="true" />
        </el-form-item>
        <el-form-item label="处理措施:"  prop="handMeasurce" >
          <el-input v-model="formData.handMeasurce" :clearable="true"  placeholder="请输入" />
        </el-form-item>
        <el-form-item label="当前状态:"  prop="currState" >
            <el-select v-model="formData.currState" placeholder="请选择" style="width:100%" :clearable="true" >
               <el-option v-for="item in ['已处理','未处理']" :key="item" :label="item" :value="item" />
            </el-select>
        </el-form-item>
        <el-form-item label="经度:"  prop="lng" >
          <el-input-number v-model="formData.lng"  style="width:100%" :precision="6" :clearable="true"  />
        </el-form-item>
        <el-form-item label="纬度:"  prop="lat" >
          <el-input-number v-model="formData.lat"  style="width:100%" :precision="6" :clearable="true"  />
        </el-form-item>
        <el-row>
          <el-col :span="24" class="grid-cell">
            <el-form-item label="机巢" prop="nestid">
              <el-select
                v-model="formData.nestid"
                class="full-width-input"
                clearable
                filterable
                multiple
                remote
                :remote-method="nestSearch"
                :loading="nestLoading"
              >
              <el-option
                v-for="(item, index) in nestidOptions"
                :key="index"
                :label="item.label"
                :value="item.value"
                :disabled="item.disabled"
              ></el-option>
            </el-select>
          </el-form-item>
        </el-col>
      </el-row>
      <el-row>
          <el-col :span="24" class="grid-cell">
            <el-form-item label="航摄成果" prop="AerialPhotographyId">
              <el-select
                v-model.number="formData.AerialPhotographyId"
                class="full-width-input"
                clearable
                filterable
                remote
                :remote-method="aerialPhotographySearch"
                :loading="aerialPhotographyloading"
              >
              <el-option
                v-for="(item, index) in allaerialPhotographyOptions"
                :key="index"
                :label="item.label"
                :value="item.value"
                :disabled="item.disabled"
              ></el-option>
            </el-select>
          </el-form-item>
        </el-col>
      </el-row>
      <el-row>
          <el-col :span="24" class="grid-cell">
            <el-form-item label="光谱" prop="multiSpectraId">
              <el-select
                v-mode="formData.multiSpectraId"
                class="full-width-input"
                clearable
                filterable
                remote
                :remote-method="multiSpectraSearch"
                :loading="multiSpectraloading"
              >
              <el-option
                v-for="(item, index) in allMultiSpectraOptions"
                :key="index"
                :label="item.label"
                :value="item.value"
                :disabled="item.disabled"
              ></el-option>
            </el-select>
          </el-form-item>
        </el-col>
      </el-row>
      <el-form-item label="航摄成果地址:"  prop="aerialServerAddress" >
          <el-input v-model="formData.aerialServerAddress" :clearable="true"  placeholder="请输入" />
        </el-form-item>
      </el-form>
      <template #footer>
        <div class="dialog-footer">
          <el-button @click="closeDialog">取 消</el-button>
          <el-button type="primary" @click="enterDialog">确 定</el-button>
        </div>
      </template>
    </el-dialog>
  </div>
</template>

<script>
export default {
  name: 'ProblemRecord'
}
</script>

<script setup>
import {
  createProblemRecord,
  deleteProblemRecord,
  deleteProblemRecordByIds,
  updateProblemRecord,
  findProblemRecord,
  getProblemRecordList
} from '@/api/problemRecord'
import { getUrl } from '@/utils/image'
// 图片选择组件
import SelectImage from '@/components/selectImage/selectImage.vue'
import { getNestInfoList } from "@/api/nestInfo";
import { getAerialPhotographyResultList } from "@/api/aerialPhotographyResult";
import { getMultiSpectraAnalysisList } from "@/api/multiSpectraAnalysis";


// 全量引入格式化工具 请按需保留
import { getDictFunc, formatDate, formatBoolean, filterDict } from '@/utils/format'
import { ElMessage, ElMessageBox } from 'element-plus'
import { ref, reactive,nextTick } from 'vue'

// 自动化生成的字典（可能为空）以及字段
const formData = ref({
        orderNum: 0,
        registerDate: new Date(),
        problemName: '',
        problemType: '',
        problemSource: '',
        problemDesc: '',
        personCharge: '',
        handMeasurce: '',
        lng: 0,
        lat: 0,
        nestid: '',
        multiSpectraId: '',
        AerialPhotographyId: 0,
        aerialServerAddress: '',
        execRecordId: '',
        multiTypeId: 0,
        })

// 验证规则
const rule = reactive({
  nestid: [
    {
      required: true,
      message: "",
      trigger: ["input", "blur"],
    },
    {
      whitespace: true,
      message: "不能只输入空格",
      trigger: ["input", "blur"],
    },
  ],
});
const loading = ref(false);
//机巢
const nestLoading = ref(false);
const nestidOptions = ref([]);
const allNestidOptions = ref([]);
const nestSearch = (query) => {
  loading.value = true;
  const tmpArr = nestidOptions.value.filter((item) => {
    return item.label.indexOf(query) !== -1;
  });
  loading.value = false;
};
//航摄成果
const aerialPhotographyloading = ref(false);
const aerialPhotographyOptions = ref([]);
const allaerialPhotographyOptions = ref([]);
//光谱
const multiSpectraloading = ref(false);
const multiSpectraOptions = ref([]);
const allMultiSpectraOptions = ref([]);

// 初始化方法
const init = async () => {
  //获取所有机巢并初始化
  getNestInfoList({ page: 1, pageSize: 9999 }).then((res) => {
    const data = res.data;
    for (const item of data.list) {
      allNestidOptions.value.push({
        value: `${item.nestid}`,
        label: item.nestName,
      });
    }
    nestidOptions.value = allNestidOptions.value;
  });
  //获取所有航摄成果并初始化数据数组
  getAerialPhotographyResultList({ page: 1, pageSize: 99999 }).then((res) => {
    const data = res.data;
    for (const item of data.list) {
      allaerialPhotographyOptions.value.push({
        value: `${item.ID}`,
        label: item.name,
      });
    }
    aerialPhotographyOptions.value = allaerialPhotographyOptions.value;
  });
  //获取所有光谱数据并初始化数据数组
  getMultiSpectraAnalysisList({ page: 1, pageSize: 99999 }).then((res) => {
    const data = res.data;
    for (const item of data.list) {
      allMultiSpectraOptions.value.push({
        value: `${item.ID}`,
        label: item.multiSpectraName,
      });
    }
    multiSpectraOptions.value = allMultiSpectraOptions.value;
  });
}
init();
const searchRule = reactive({
  createdAt: [
    { validator: (rule, value, callback) => {
      if (searchInfo.value.startCreatedAt && !searchInfo.value.endCreatedAt) {
        callback(new Error('请填写结束日期'))
      } else if (!searchInfo.value.startCreatedAt && searchInfo.value.endCreatedAt) {
        callback(new Error('请填写开始日期'))
      } else if (searchInfo.value.startCreatedAt && searchInfo.value.endCreatedAt && (searchInfo.value.startCreatedAt.getTime() === searchInfo.value.endCreatedAt.getTime() || searchInfo.value.startCreatedAt.getTime() > searchInfo.value.endCreatedAt.getTime())) {
        callback(new Error('开始日期应当早于结束日期'))
      } else {
        callback()
      }
    }, trigger: 'change' }
  ],
})

const elFormRefRecord = ref()
const elSearchFormRef = ref()

// =========== 表格控制部分 ===========
const page = ref(1)
const total = ref(0)
const pageSize = ref(10)
const tableData = ref([])
const searchInfo = ref({})

// 重置
const onReset = () => {
  searchInfo.value = {}
  getTableData()
}

// 搜索
const onSubmit = () => {
  elSearchFormRef.value?.validate(async(valid) => {
    if (!valid) return
    page.value = 1
    pageSize.value = 10
    getTableData()
  })
}

// 分页
const handleSizeChange = (val) => {
  pageSize.value = val
  getTableData()
}

// 修改页面容量
const handleCurrentChange = (val) => {
  page.value = val
  getTableData()
}

// 查询
const getTableData = async() => {
  const table = await getProblemRecordList({ page: page.value, pageSize: pageSize.value, ...searchInfo.value })
  if (table.code === 0) {
    tableData.value = table.data.list
    total.value = table.data.total
    page.value = table.data.page
    pageSize.value = table.data.pageSize
  }
}

getTableData()

// ============== 表格控制部分结束 ===============

// 获取需要的字典 可能为空 按需保留
const setOptions = async () =>{
}

// 获取需要的字典 可能为空 按需保留
setOptions()


// 多选数据
const multipleSelection = ref([])
// 多选
const handleSelectionChange = (val) => {
    multipleSelection.value = val
}

// 删除行
const deleteRow = (row) => {
    ElMessageBox.confirm('确定要删除吗?', '提示', {
        confirmButtonText: '确定',
        cancelButtonText: '取消',
        type: 'warning'
    }).then(() => {
            deleteProblemRecordFunc(row)
        })
    }


// 批量删除控制标记
const deleteVisible = ref(false)

// 多选删除
const onDelete = async() => {
      const ids = []
      if (multipleSelection.value.length === 0) {
        ElMessage({
          type: 'warning',
          message: '请选择要删除的数据'
        })
        return
      }
      multipleSelection.value &&
        multipleSelection.value.map(item => {
          ids.push(item.ID)
        })
      const res = await deleteProblemRecordByIds({ ids })
      if (res.code === 0) {
        ElMessage({
          type: 'success',
          message: '删除成功'
        })
        if (tableData.value.length === ids.length && page.value > 1) {
          page.value--
        }
        deleteVisible.value = false
        getTableData()
      }
    }

// 行为控制标记（弹窗内部需要增还是改）
const type = ref('')

// 更新行
const updateProblemRecordFunc = async(row) => {
    const res = await findProblemRecord({ ID: row.ID })
    type.value = 'update'
    console.log(res)
    if (res.code === 0) {
        formData.value = res.data.repbRecord
        try {
          const nest = JSON.parse(row.nestid)
          // row.data
          formData.value.nestid = nest
        } catch (error) {
          
        }
        console.log(res.data.repbRecord, row)
        dialogFormVisible.value = true
    }
}

// 删除行
const deleteProblemRecordFunc = async (row) => {
    const res = await deleteProblemRecord({ ID: row.ID })
    if (res.code === 0) {
        ElMessage({
                type: 'success',
                message: '删除成功'
            })
            if (tableData.value.length === 1 && page.value > 1) {
            page.value--
        }
        getTableData()
    }
}

// 弹窗控制标记
const dialogFormVisible = ref(false)



// 打开弹窗
const openDialog = () => {
    type.value = 'create'
    dialogFormVisible.value = true
    
}

// 关闭弹窗
const closeDialog = () => {
    dialogFormVisible.value = false
    formData.value = {
        orderNum: 0,
        registerDate: new Date(),
        problemName: '',
        problemType: '',
        problemSource: '',
        problemDesc: '',
        personCharge: '',
        handMeasurce: '',
        lng: 0,
        lat: 0,
        nestid: '',
        }
        emit('cal')
}
// 弹窗确定
const enterDialog = async () => {
  // debugger;
  console.log(elFormRefRecord.value)
     elFormRefRecord.value?.validate( async (valid) => {
      console.log("valid", valid)
            //  if (!valid) return
              let res
              formData.value.nestid = JSON.stringify(formData.value.nestid)
              console.log('val', formData.value)
              switch (type.value) {
                case 'create':
                  res = await createProblemRecord(formData.value)
                  break
                case 'update':
                  res = await updateProblemRecord(formData.value)
                  break
                default:
                  res = await createProblemRecord(formData.value)
                  break
              }
              if (res.code === 0) {
                ElMessage({
                  type: 'success',
                  message: '创建/更改成功'
                })
                closeDialog()
                getTableData()
              }
      })
}

</script>

<style>

</style>
