<template>
  <div>
    <div class="gva-search-box">
      <el-form
        ref="elSearchFormRef"
        :inline="true"
        :model="searchInfo"
        class="demo-form-inline"
        :rules="searchRule"
        @keyup.enter="onSubmit"
      >
        <el-form-item label="创建日期" prop="createdAt">
          <template #label>
            <span>
              创建日期
              <el-tooltip
                content="搜索范围是开始日期（包含）至结束日期（不包含）"
              >
                <el-icon><QuestionFilled /></el-icon>
              </el-tooltip>
            </span>
          </template>
          <el-date-picker
            v-model="searchInfo.startCreatedAt"
            type="datetime"
            placeholder="开始日期"
            :disabled-date="
              (time) =>
                searchInfo.endCreatedAt
                  ? time.getTime() > searchInfo.endCreatedAt.getTime()
                  : false
            "
          ></el-date-picker>
          —
          <el-date-picker
            v-model="searchInfo.endCreatedAt"
            type="datetime"
            placeholder="结束日期"
            :disabled-date="
              (time) =>
                searchInfo.startCreatedAt
                  ? time.getTime() < searchInfo.startCreatedAt.getTime()
                  : false
            "
          ></el-date-picker>
        </el-form-item>
        <el-form-item>
          <el-button type="primary" icon="search" @click="onSubmit"
            >查询</el-button
          >
          <el-button icon="refresh" @click="onReset">重置</el-button>
        </el-form-item>
      </el-form>
    </div>
    <div class="gva-table-box">
      <div class="gva-btn-list">
        <el-button type="primary" icon="plus" @click="openDialog"
          >新增</el-button
        >
        <el-popover v-model:visible="deleteVisible" placement="top" width="160">
          <p>确定要删除吗？</p>
          <div style="text-align: right; margin-top: 8px">
            <el-button type="primary" link @click="deleteVisible = false"
              >取消</el-button
            >
            <el-button type="primary" @click="onDelete">确定</el-button>
          </div>
          <template #reference>
            <el-button
              icon="delete"
              style="margin-left: 10px"
              :disabled="!multipleSelection.length"
              @click="deleteVisible = true"
              >删除</el-button
            >
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
          <template #default="scope">{{
            formatDate(scope.row.CreatedAt)
          }}</template>
        </el-table-column>
        <el-table-column
          align="left"
          label="航线id"
          prop="missionid"
          width="120"
        />
        <el-table-column
          align="left"
          label="航线名称"
          prop="name"
          width="120"
        />
        <el-table-column align="left" label="航线类型" prop="type" width="120">
          <template #default="scope">{{
            convertTableValue(scope.row.type, typeList)
          }}</template>
        </el-table-column>
        <el-table-column
          align="left"
          label="起飞模式"
          prop="gotoFirstWaypointMode"
          width="120"
        >
          <template #default="scope">{{
            convertTableValue(
              scope.row.gotoFirstWaypointMode,
              gotoFirstWaypointModeList
            )
          }}</template>
        </el-table-column>
        <el-table-column
          align="left"
          label="结束模式"
          prop="finishAction"
          width="120"
        >
          <template #default="scope">{{
            convertTableValue(scope.row.finishAction, finishActionList)
          }}</template>
        </el-table-column>
        <el-table-column
          align="left"
          label="路径模式"
          prop="flightPathMode"
          width="120"
        >
          <template #default="scope">{{
            convertTableValue(scope.row.flightPathMode, flightPathModeList)
          }}</template>
        </el-table-column>
        <el-table-column
          align="left"
          label="朝向模式"
          prop="headingMode"
          width="120"
        >
          <template #default="scope">{{
            convertTableValue(scope.row.headingMode, headingModeList)
          }}</template>
        </el-table-column>
        <el-table-column
          align="left"
          label="自动飞行速度"
          prop="autoFlightSpeed"
          width="120"
        />
        <!-- <el-table-column align="left" label="参数体" prop="param" width="120" /> -->
        <el-table-column align="left" label="安全" prop="safealt" width="120" />
        <el-table-column align="left" label="kml" prop="kml" width="120" />
        <el-table-column align="left" label="gps" prop="gps" width="120" />
        <el-table-column
          align="left"
          label="站点id"
          prop="station"
          width="120"
        />
        <el-table-column
          align="left"
          label="明确定位"
          prop="clearHomeLocation"
          width="120"
        />
        <el-table-column
          align="left"
          label="制作人"
          prop="producer"
          width="120"
        />
        <el-table-column
          align="left"
          label="制作单位"
          prop="productionUnit"
          width="120"
        />
        <el-table-column
          align="left"
          label="isactive"
          prop="isActive"
          width="120"
        />
        <el-table-column
          align="left"
          label="是否多光谱"
          prop="useMultispectrum"
          width="120"
        />
        <el-table-column
          align="left"
          label="固定返航点"
          prop="fixedReturnPoint"
          width="120"
        />
        <el-table-column
          align="left"
          label="机巢id"
          prop="nestId"
          width="120"
        />
        <el-table-column align="left" label="备注" prop="remark" width="120" />
        <el-table-column
          align="left"
          label="返航高度"
          prop="goHomeHeight"
          width="120"
        />
        <el-table-column
          align="left"
          label="预估距离"
          prop="execDistance"
          width="120"
        />
        <el-table-column
          align="left"
          label="预估执行时间"
          prop="execTimeSpend"
          width="120"
        />
        <el-table-column align="left" label="操作">
          <template #default="scope">
            <el-button
              type="primary"
              link
              icon="edit"
              class="table-button"
              @click="updateNestAirlineFunc(scope.row)"
              >变更</el-button
            >
            <el-button
              type="primary"
              link
              icon="delete"
              @click="deleteRow(scope.row)"
              >删除</el-button
            >
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
    <el-dialog
      v-model="dialogFormVisible"
      :before-close="closeDialog"
      :title="type === 'create' ? '添加' : '修改'"
      destroy-on-close
      class="dialog__container"
    >
      <el-form
        :model="formData"
        label-position="right"
        ref="elFormRef"
        :rules="rule"
        label-width="100px"
      >
        <el-form-item label="航线名称:" prop="name">
          <el-input
            v-model="formData.name"
            :clearable="true"
            placeholder="请输入"
          />
        </el-form-item>
        <el-form-item label="航线类型:" prop="type">
          <el-select
            v-model="formData.type"
            class="m-2"
            placeholder="Select"
            size="small"
          >
            <el-option
              v-for="item in typeList"
              :key="item.value"
              :label="item.label"
              :value="item.value"
            />
          </el-select>
        </el-form-item>
        <el-form-item label="起飞模式:" prop="gotoFirstWaypointMode">
          <el-select
            v-model="formData.gotoFirstWaypointMode"
            class="m-2"
            placeholder="Select"
            size="small"
          >
            <el-option
              v-for="item in gotoFirstWaypointModeList"
              :key="item.value"
              :label="item.label"
              :value="item.value"
            />
          </el-select>
        </el-form-item>
        <el-form-item label="结束模式:" prop="finishAction">
          <el-select
            v-model="formData.finishAction"
            class="m-2"
            placeholder="Select"
            size="small"
          >
            <el-option
              v-for="item in finishActionList"
              :key="item.value"
              :label="item.label"
              :value="item.value"
            />
          </el-select>
        </el-form-item>
        <el-form-item label="路径模式:" prop="flightPathMode">
          <el-select
            v-model="formData.flightPathMode"
            class="m-2"
            placeholder="Select"
            size="small"
          >
            <el-option
              v-for="item in flightPathModeList"
              :key="item.value"
              :label="item.label"
              :value="item.value"
            />
          </el-select>
        </el-form-item>
        <el-form-item label="朝向模式:" prop="headingMode">
          <el-select
            v-model="formData.headingMode"
            class="m-2"
            placeholder="Select"
            size="small"
          >
            <el-option
              v-for="item in headingModeList"
              :key="item.value"
              :label="item.label"
              :value="item.value"
            />
          </el-select>
        </el-form-item>
        <el-form-item label="机巢:" prop="nestId">
          <el-select
            v-model="formData.nestId"
            class="m-2"
            placeholder="Select"
            size="small"
          >
            <el-option
              v-for="item in nestinfoAll"
              :key="item.value"
              :label="item.label"
              :value="item.value"
            />
          </el-select>
        </el-form-item>
        <el-form-item label="参数体:" prop="param">
          <el-input
            v-model="formData.param"
            :clearable="true"
            placeholder="请输入"
          />
        </el-form-item>
        <el-form-item label="安全:" prop="safealt">
          <el-input
            v-model="formData.safealt"
            :clearable="true"
            placeholder="请输入"
          />
        </el-form-item>
        <el-form-item label="kml:" prop="kml">
          <el-input
            v-model="formData.kml"
            :clearable="true"
            placeholder="请输入"
          />
        </el-form-item>
        <el-form-item label="gps:" prop="gps">
          <el-input
            v-model="formData.gps"
            :clearable="true"
            placeholder="请输入"
          />
        </el-form-item>
        <el-form-item label="站点id:" prop="station">
          <el-input
            v-model="formData.station"
            :clearable="true"
            placeholder="请输入"
          />
        </el-form-item>
        <el-form-item label="明确定位:" prop="clearHomeLocation">
          <el-input
            v-model="formData.clearHomeLocation"
            :clearable="true"
            placeholder="请输入"
          />
        </el-form-item>
        <el-form-item label="制作人:" prop="producer">
          <el-input
            v-model="formData.producer"
            :clearable="true"
            placeholder="请输入"
          />
        </el-form-item>
        <el-form-item label="制作单位:" prop="productionUnit">
          <el-input
            v-model="formData.productionUnit"
            :clearable="true"
            placeholder="请输入"
          />
        </el-form-item>
        <el-form-item label="isactive:" prop="isActive">
          <el-input
            v-model="formData.isActive"
            :clearable="true"
            placeholder="请输入"
          />
        </el-form-item>
         <el-form-item label="是否多光谱:" prop="useMultispectrum">
          <el-input-number
            v-model="formData.useMultispectrum"
            :clearable="true"
            placeholder="请输入"
          />
        </el-form-item>
        <el-form-item label="固定返航点:" prop="fixedReturnPoint">
          <el-input
            v-model="formData.fixedReturnPoint"
            :clearable="true"
            placeholder="请输入"
          />
        </el-form-item>
        <el-form-item label="备注:" prop="remark">
          <el-input
            v-model="formData.remark"
            :clearable="true"
            placeholder="请输入"
          />
        </el-form-item>
        <el-form-item label="自动飞行速度:" prop="autoFlightSpeed">
          <el-input-number
            v-model="formData.autoFlightSpeed"
            :min="1"
            :max="15"
            style="width: 100%"
            :precision="2"
            :clearable="true"
          />
        </el-form-item>
        <el-form-item label="返航高度:" prop="goHomeHeight">
          <el-input-number
            v-model="formData.goHomeHeight"
            style="width: 100%"
            :precision="2"
            :min="10"
            :clearable="true"
          />
        </el-form-item>
        <el-form-item label="预估距离:" prop="execDistance">
          <el-input-number
            v-model="formData.execDistance"
            disabled
            style="width: 100%"
            :precision="0"
            :min="0"
            :clearable="true"
          ></el-input-number>
        </el-form-item>
        <el-form-item label="预估执行时间:" prop="execTimeSpend">
          <el-input-number
            v-model="formData.execTimeSpend"
            disabled
            style="width: 100%"
            :precision="0"
            :min="0"
            :clearable="true"
          ></el-input-number>
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
  name: "NestAirline",
};
</script>

<script setup>
import { v4 } from "uuid";
import {
  createNestAirline,
  deleteNestAirline,
  deleteNestAirlineByIds,
  updateNestAirline,
  findNestAirline,
  getNestAirlineList,
} from "@/api/nestAirline";
import { findNestInfoByNestIds } from "@/api/aerialPhotographyResult";
// 全量引入格式化工具 请按需保留
import {
  getDictFunc,
  formatDate,
  formatBoolean,
  filterDict,
} from "@/utils/format";
import { ElMessage, ElMessageBox } from "element-plus";
import { ref, reactive } from "vue";

// 自动化生成的字典（可能为空）以及字段
const formData = ref({
  missionid: v4(),
  name: "",
  type: 0,
  autoFlightSpeed: 1,
  gotoFirstWaypointMode: 1,
  finishAction: 0,
  flightPathMode: 1,
  headingMode: 1,
  param: "",
  safealt: "",
  kml: "",
  gps: "",
  station: "",
  clearHomeLocation: "",
  producer: "",
  productionUnit: "",
  isActive: "",
  useMultispectrum: 0,
  fixedReturnPoint: "",
  nestId: "",
  remark: "",
  goHomeHeight: 60,
  execDistance: 0,
  execTimeSpend: 0,
});

// 验证规则
const rule = reactive({
  name: [
    { required: true, message: "航线名称不能为空", trigger: "blur" },
    { min: 2, message: "航线名称不能少于两个字", trigger: "blur" },
  ],
  nestId: [{ required: true, message: "请选择机巢", trigger: "blur" }],
});

const searchRule = reactive({
  createdAt: [
    {
      validator: (rule, value, callback) => {
        if (searchInfo.value.startCreatedAt && !searchInfo.value.endCreatedAt) {
          callback(new Error("请填写结束日期"));
        } else if (
          !searchInfo.value.startCreatedAt &&
          searchInfo.value.endCreatedAt
        ) {
          callback(new Error("请填写开始日期"));
        } else if (
          searchInfo.value.startCreatedAt &&
          searchInfo.value.endCreatedAt &&
          (searchInfo.value.startCreatedAt.getTime() ===
            searchInfo.value.endCreatedAt.getTime() ||
            searchInfo.value.startCreatedAt.getTime() >
              searchInfo.value.endCreatedAt.getTime())
        ) {
          callback(new Error("开始日期应当早于结束日期"));
        } else {
          callback();
        }
      },
      trigger: "change",
    },
  ],
});

const elFormRef = ref();
const elSearchFormRef = ref();

const typeList = [
  {
    value: 0,
    label: "航线",
  },
  {
    value: 1,
    label: "全景",
  },
  {
    value: 2,
    label: "正射",
  },
  {
    value: 3,
    label: "三维",
  },
];
const gotoFirstWaypointModeList = [
  {
    value: 1,
    label: "稳健模式",
  },
  {
    value: 2,
    label: "点到点",
  },
];
const finishActionList = [
  {
    value: 0,
    label: "无操作",
  },
  {
    value: 1,
    label: "返回机巢",
  },
  {
    value: 2,
    label: "降落到结束点",
  },
  {
    value: 3,
    label: "回到起始点",
  },
];
const flightPathModeList = [
  {
    value: 1,
    label: "正常",
  },
  {
    value: 2,
    label: "协调转弯",
  },
];
const headingModeList = [
  {
    value: 1,
    label: "自动",
  },
  {
    value: 2,
    label: "手动控制",
  },
  {
    value: 3,
    label: "导航点控制",
  },
  {
    value: 4,
    label: "平滑旋转",
  },
];
const convertTableValue = (value, list) => {
  return list.find((val) => val.value === value).label;
};
let nestinfoAll = [];
const getNestInfoByNestIds = async () => {
  nestinfoAll = await findNestInfoByNestIds();
  nestinfoAll = nestinfoAll.data.renestinfoList;
  nestinfoAll = nestinfoAll.map((item) => ({
    label: item.nestName,
    value: item.nestid,
  }));
};
getNestInfoByNestIds();
// =========== 表格控制部分 ===========
const page = ref(1);
const total = ref(0);
const pageSize = ref(10);
const tableData = ref([]);
const searchInfo = ref({});

// 重置
const onReset = () => {
  searchInfo.value = {};
  getTableData();
};

// 搜索
const onSubmit = () => {
  elSearchFormRef.value?.validate(async (valid) => {
    if (!valid) return;
    page.value = 1;
    pageSize.value = 10;
    getTableData();
  });
};

// 分页
const handleSizeChange = (val) => {
  pageSize.value = val;
  getTableData();
};

// 修改页面容量
const handleCurrentChange = (val) => {
  page.value = val;
  getTableData();
};

// 查询
const getTableData = async () => {
  const table = await getNestAirlineList({
    page: page.value,
    pageSize: pageSize.value,
    ...searchInfo.value,
  });
  if (table.code === 0) {
    tableData.value = table.data.list;
    total.value = table.data.total;
    page.value = table.data.page;
    pageSize.value = table.data.pageSize;
  }
};

getTableData();

// ============== 表格控制部分结束 ===============

// 获取需要的字典 可能为空 按需保留
const setOptions = async () => {};

// 获取需要的字典 可能为空 按需保留
setOptions();

// 多选数据
const multipleSelection = ref([]);
// 多选
const handleSelectionChange = (val) => {
  multipleSelection.value = val;
};

// 删除行
const deleteRow = (row) => {
  ElMessageBox.confirm("确定要删除吗?", "提示", {
    confirmButtonText: "确定",
    cancelButtonText: "取消",
    type: "warning",
  }).then(() => {
    deleteNestAirlineFunc(row);
  });
};

// 批量删除控制标记
const deleteVisible = ref(false);

// 多选删除
const onDelete = async () => {
  const ids = [];
  if (multipleSelection.value.length === 0) {
    ElMessage({
      type: "warning",
      message: "请选择要删除的数据",
    });
    return;
  }
  multipleSelection.value &&
    multipleSelection.value.map((item) => {
      ids.push(item.ID);
    });
  const res = await deleteNestAirlineByIds({ ids });
  if (res.code === 0) {
    ElMessage({
      type: "success",
      message: "删除成功",
    });
    if (tableData.value.length === ids.length && page.value > 1) {
      page.value--;
    }
    deleteVisible.value = false;
    getTableData();
  }
};

// 行为控制标记（弹窗内部需要增还是改）
const type = ref("");

// 更新行
const updateNestAirlineFunc = async (row) => {
  const res = await findNestAirline({ ID: row.ID });
  type.value = "update";
  if (res.code === 0) {
    formData.value = res.data.reNtAirline;
    dialogFormVisible.value = true;
  }
};

// 删除行
const deleteNestAirlineFunc = async (row) => {
  const res = await deleteNestAirline({ ID: row.ID });
  if (res.code === 0) {
    ElMessage({
      type: "success",
      message: "删除成功",
    });
    if (tableData.value.length === 1 && page.value > 1) {
      page.value--;
    }
    getTableData();
  }
};

// 弹窗控制标记
const dialogFormVisible = ref(false);

// 打开弹窗
const openDialog = () => {
  type.value = "create";
  dialogFormVisible.value = true;
};

// 关闭弹窗
const closeDialog = () => {
  dialogFormVisible.value = false;
  formData.value = {
    missionid: "",
    name: "",
    type: 0,
    autoFlightSpeed: 1,
    gotoFirstWaypointMode: 1,
    finishAction: 0,
    flightPathMode: 1,
    headingMode: 1,
    param: "",
    safealt: "",
    kml: "",
    gps: "",
    station: "",
    clearHomeLocation: "",
    producer: "",
    productionUnit: "",
    isActive: "",
    fixedReturnPoint: "",
    nestId: "",
    remark: "",
    goHomeHeight: 60,
    execDistance: 0,
    execTimeSpend: 0,
  };
};
// 弹窗确定
const enterDialog = async () => {
  elFormRef.value?.validate(async (valid) => {
    if (!valid) return;
    let res;
    switch (type.value) {
      case "create":
        res = await createNestAirline(formData.value);
        break;
      case "update":
        res = await updateNestAirline(formData.value);
        break;
      default:
        res = await createNestAirline(formData.value);
        break;
    }
    if (res.code === 0) {
      ElMessage({
        type: "success",
        message: "创建/更改成功",
      });
      closeDialog();
      getTableData();
    }
  });
};
</script>

<style lang="scss">
.dialog__container {
  height: 70%;
  .el-dialog__body {
    height: 73%;
    overflow: auto;
  }
}
</style>
