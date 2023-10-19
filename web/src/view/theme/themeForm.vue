<template>
  <div>
    <div class="gva-form-box">
      <el-form :model="formData" ref="elFormRef" label-position="right" :rules="rule" label-width="80px">
        <el-form-item label="主题名:" prop="themeName">
          <el-input v-model="formData.themeName" :clearable="true" placeholder="请输入" />
       </el-form-item>
        <el-form-item label="用户角色:" prop="userRoles">
          <el-input v-model="formData.userRoles" :clearable="true" placeholder="请输入" />
       </el-form-item>
        <el-form-item label="登录页logo:" prop="loginViewLogo">
          <el-input v-model="formData.loginViewLogo" :clearable="true" placeholder="请输入" />
       </el-form-item>
        <el-form-item label="系统logo:" prop="systemLogo">
          <el-input v-model="formData.systemLogo" :clearable="true" placeholder="请输入" />
       </el-form-item>
        <el-form-item label="系统名称:" prop="systemName">
          <el-input v-model="formData.systemName" :clearable="true" placeholder="请输入" />
       </el-form-item>
        <el-form-item label="背景轮播图:" prop="backgroundPhoto">
          <el-input v-model="formData.backgroundPhoto" :clearable="true" placeholder="请输入" />
       </el-form-item>
        <el-form-item label="是否默认主题:" prop="isOrNoDefaultTheme">
          <el-switch v-model="formData.isOrNoDefaultTheme" active-color="#13ce66" inactive-color="#ff4949" active-text="是" inactive-text="否" clearable ></el-switch>
       </el-form-item>
        <el-form-item>
          <el-button type="primary" @click="save">保存</el-button>
          <el-button type="primary" @click="back">返回</el-button>
        </el-form-item>
      </el-form>
    </div>
  </div>
</template>

<script>
export default {
  name: 'Theme'
}
</script>

<script setup>
import {
  createTheme,
  updateTheme,
  findTheme
} from '@/api/theme'

// 自动获取字典
import { getDictFunc } from '@/utils/format'
import { useRoute, useRouter } from "vue-router"
import { ElMessage } from 'element-plus'
import { ref, reactive } from 'vue'
const route = useRoute()
const router = useRouter()

const type = ref('')
const formData = ref({
            themeName: '',
            userRoles: '',
            loginViewLogo: '',
            systemLogo: '',
            systemName: '',
            backgroundPhoto: '',
            isOrNoDefaultTheme: false,
        })
// 验证规则
const rule = reactive({
})

const elFormRef = ref()

// 初始化方法
const init = async () => {
 // 建议通过url传参获取目标数据ID 调用 find方法进行查询数据操作 从而决定本页面是create还是update 以下为id作为url参数示例
    if (route.query.id) {
      const res = await findTheme({ ID: route.query.id })
      if (res.code === 0) {
        formData.value = res.data.retheme
        type.value = 'update'
      }
    } else {
      type.value = 'create'
    }
}

init()
// 保存按钮
const save = async() => {
      elFormRef.value?.validate( async (valid) => {
         if (!valid) return
            let res
           switch (type.value) {
             case 'create':
               res = await createTheme(formData.value)
               break
             case 'update':
               res = await updateTheme(formData.value)
               break
             default:
               res = await createTheme(formData.value)
               break
           }
           if (res.code === 0) {
             ElMessage({
               type: 'success',
               message: '创建/更改成功'
             })
           }
       })
}

// 返回按钮
const back = () => {
    router.go(-1)
}

</script>

<style>
</style>
