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
          />
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
          />
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
          label="主题名"
          prop="themeName"
          width="120"
        />
        <el-table-column align="left" label="用户角色" width="120">
          <template #default="scope">
            {{
              authorityList.find((item) => item.value === scope.row.userRoles)
                ?.label
            }}
          </template>
        </el-table-column>
        <el-table-column
          align="left"
          label="登录页logo"
          prop="loginViewLogo"
          width="150"
        >
          <template #default="scope">
            <el-image
              v-if="typeof scope.row.loginViewLogo === 'object'"
              style="width: 100px; height: 100px"
              :src="'../api/' + scope.row.loginViewLogo[0].url"
              :zoom-rate="1.2"
              :max-scale="7"
              :min-scale="0.2"
              :preview-src-list="
                scope.row.loginViewLogo?.map((item) => '../api/' + item.url)
              "
              :preview-teleported="true"
              fit="contain"
            />
          </template>
        </el-table-column>
        <el-table-column
          align="left"
          label="系统logo"
          prop="systemLogo"
          width="150"
        >
          <template #default="scope">
            <el-image
              v-if="typeof scope.row.systemLogo === 'object'"
              style="width: 100px; height: 100px"
              :src="'../api/' + scope.row.systemLogo[0].url"
              :zoom-rate="1.2"
              :max-scale="7"
              :min-scale="0.2"
              :preview-src-list="
                scope.row.systemLogo.map((item) => '../api/' + item.url)
              "
              fit="contain"
              :preview-teleported="true"
            />
          </template>
        </el-table-column>
        <el-table-column
          align="left"
          label="系统名称"
          prop="systemName"
          width="120"
        />
        <el-table-column
          align="left"
          label="背景轮播图"
          prop="backgroundPhoto"
          width="150"
        >
          <template #default="scope">
            <el-image
              v-if="typeof scope.row.backgroundPhoto === 'object'"
              style="width: 100px; height: 100px"
              :src="'../api/' + scope.row.backgroundPhoto[0].url"
              :zoom-rate="1.2"
              :max-scale="7"
              :min-scale="0.2"
              :preview-src-list="
                scope.row.backgroundPhoto.map((item) => '../api/' + item.url)
              "
              fit="contain"
              :preview-teleported="true"
            />
          </template>
        </el-table-column>
        <el-table-column
          align="left"
          label="是否默认主题"
          prop="isOrNoDefaultTheme"
          width="120"
        >
          <template #default="scope">{{
            formatBoolean(scope.row.isOrNoDefaultTheme)
          }}</template>
        </el-table-column>
        <el-table-column align="left" label="操作">
          <template #default="scope">
            <el-button
              type="primary"
              link
              icon="edit"
              class="table-button"
              @click="updateThemeFunc(scope.row)"
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
    >
      <el-form
        ref="elFormRef"
        :model="formData"
        label-position="right"
        :rules="rule"
        label-width="120px"
      >
        <el-form-item label="主题名:" prop="themeName">
          <el-input
            v-model="formData.themeName"
            :clearable="true"
            placeholder="请输入"
          />
        </el-form-item>
        <el-form-item label="用户角色:" prop="userRoles">
          <el-select
            v-model="formData.userRoles"
            class="m-2"
            placeholder="Select"
            size="large"
          >
            <el-option
              v-for="item in authorityList"
              :key="item.value"
              :label="item.label"
              :value="item.value"
            />
          </el-select>
        </el-form-item>
        <el-form-item label="登录页logo:" prop="loginViewLogo">
          <uploadImg v-model="formData.loginViewLogo" :limit="1" />
        </el-form-item>
        <el-form-item label="系统logo:" prop="systemLogo">
          <uploadImg v-model="formData.systemLogo" :limit="1" />
        </el-form-item>
        <el-form-item label="系统名称:" prop="systemName">
          <el-input
            v-model="formData.systemName"
            :clearable="true"
            placeholder="请输入"
          />
        </el-form-item>
        <el-form-item label="背景轮播图:" prop="backgroundPhoto">
          <uploadImg v-model="formData.backgroundPhoto" />
        </el-form-item>
        <el-form-item label="是否默认主题:" prop="isOrNoDefaultTheme">
          <el-switch
            v-model="formData.isOrNoDefaultTheme"
            active-color="#13ce66"
            inactive-color="#ff4949"
            active-text="是"
            inactive-text="否"
            clearable
          />
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
  name: "Theme",
};
</script>

<script setup>
import uploadImg from "@/components/upload/commonImg.vue";
import {
  createTheme,
  deleteTheme,
  deleteThemeByIds,
  updateTheme,
  findTheme,
  getThemeList,
} from "@/api/theme";
import { getAuthorityList } from "@/api/authority";

// 全量引入格式化工具 请按需保留
import { formatDate, formatBoolean } from "@/utils/format";
import { ElMessage, ElMessageBox } from "element-plus";
import { ref, reactive } from "vue";

const authorityList = ref();
const getAuthorityTableData = async () => {
  const table = await getAuthorityList({ page: 1, pageSize: 100 });
  if (table.code === 0) {
    authorityList.value = table.data.list.map((item) => ({
      value: String(item.authorityId),
      label: item.authorityName,
    }));
  }
};
getAuthorityTableData();
// 自动化生成的字典（可能为空）以及字段
const formData = ref({
  themeName: "",
  userRoles: "",
  loginViewLogo: [],
  systemLogo: [],
  systemName: "",
  backgroundPhoto: [],
  isOrNoDefaultTheme: false,
});

// 验证规则
const rule = reactive({
  themeName: [
    { required: true, message: "主题名称不能为空！", trigger: "blur" },
  ],
  userRoles: [
    { required: true, message: "用户角色不能为空！", trigger: "blur" },
  ],
  loginViewLogo: [
    { required: true, message: "登录页logo不能为空！", trigger: "blur" },
  ],
  systemLogo: [
    { required: true, message: "系统logo不能为空！", trigger: "blur" },
  ],
  systemName: [
    { required: true, message: "系统名称不能为空！", trigger: "blur" },
  ],
  backgroundPhoto: [
    { required: true, message: "背景轮播图不能为空！", trigger: "blur" },
  ],
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
    if (searchInfo.value.isOrNoDefaultTheme === "") {
      searchInfo.value.isOrNoDefaultTheme = null;
    }
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
  const table = await getThemeList({
    page: page.value,
    pageSize: pageSize.value,
    ...searchInfo.value,
  });
  table.data.list.forEach((item) => {
    item.loginViewLogo =
      item.loginViewLogo === "" ? [] : JSON.parse(item.loginViewLogo);
    item.systemLogo = item.systemLogo === "" ? [] : JSON.parse(item.systemLogo);
    item.backgroundPhoto =
      item.backgroundPhoto === "" ? [] : JSON.parse(item.backgroundPhoto);
  });
  console.log(table.data.list);
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
    deleteThemeFunc(row);
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
  const res = await deleteThemeByIds({ ids });
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
const updateThemeFunc = async (row) => {
  const res = await findTheme({ ID: row.ID });
  type.value = "update";
  if (res.code === 0) {
    formData.value = res.data.retheme;
    dialogFormVisible.value = true;
    formData.value.loginViewLogo =
      formData.value.loginViewLogo === ""
        ? []
        : JSON.parse(formData.value.loginViewLogo);
    formData.value.systemLogo =
      formData.value.systemLogo === ""
        ? []
        : JSON.parse(formData.value.systemLogo);
    formData.value.backgroundPhoto =
      formData.value.backgroundPhoto === ""
        ? []
        : JSON.parse(formData.value.backgroundPhoto);
  }
};

// 删除行
const deleteThemeFunc = async (row) => {
  const res = await deleteTheme({ ID: row.ID });
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
    themeName: "",
    userRoles: "",
    loginViewLogo: [],
    systemLogo: [],
    systemName: "",
    backgroundPhoto: [],
    isOrNoDefaultTheme: false,
  };
};
// 弹窗确定
const enterDialog = async () => {
  elFormRef.value?.validate(async (valid) => {
    if (!valid) return;
    let res;

    const formDataPrame = JSON.parse(JSON.stringify(formData.value));
    if (formDataPrame.loginViewLogo.length !== 0) {
      formDataPrame.loginViewLogo.forEach((item) => {
        item.url = item.url.replaceAll("../api/", "");
      });
    }
    if (formDataPrame.systemLogo.length !== 0) {
      formDataPrame.systemLogo.forEach((item) => {
        item.url = item.url.replaceAll("../api/", "");
      });
    }
    if (formDataPrame.backgroundPhoto.length !== 0) {
      formDataPrame.backgroundPhoto.forEach((item) => {
        item.url = item.url.replaceAll("../api/", "");
      });
    }
    formDataPrame.loginViewLogo = JSON.stringify(formDataPrame.loginViewLogo);
    formDataPrame.systemLogo = JSON.stringify(formDataPrame.systemLogo);
    formDataPrame.backgroundPhoto = JSON.stringify(
      formDataPrame.backgroundPhoto
    );
    switch (type.value) {
      case "create":
        res = await createTheme(formDataPrame);
        break;
      case "update":
        res = await updateTheme(formDataPrame);
        break;
      default:
        res = await createTheme(formData.value);
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

<style></style>
