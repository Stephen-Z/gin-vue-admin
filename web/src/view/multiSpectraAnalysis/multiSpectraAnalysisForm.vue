<template>
  <div>
    <div class="gva-form-box">
      <el-form :model="formData" ref="elFormRef" label-position="right" :rules="rule" label-width="80px">
        <el-form-item label="光谱名:" prop="plateId">
          <el-input v-model.number="formData.multiSpectraName" :clearable="true" placeholder="请输入" />
       </el-form-item>
        <el-form-item label="板块Id:" prop="plateId">
          <el-input v-model.number="formData.plateId" :clearable="true" placeholder="请输入" />
       </el-form-item>
        <el-form-item label="光谱类型id:" prop="multiSpectraTypeId">
          <el-input v-model.number="formData.multiSpectraTypeId" :clearable="true" placeholder="请输入" />
       </el-form-item>
        <el-form-item label="航摄成功id集:" prop="aerialPhotographyIds">
          <el-input v-model="formData.aerialPhotographyIds" :clearable="true" placeholder="请输入" />
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
  name: 'MultiSpectraAnalysis'
}
</script>

<script setup>
import {
  createMultiSpectraAnalysis,
  updateMultiSpectraAnalysis,
  findMultiSpectraAnalysis
} from '@/api/multiSpectraAnalysis'

// 自动获取字典
import { getDictFunc } from '@/utils/format'
import { useRoute, useRouter } from "vue-router"
import { ElMessage } from 'element-plus'
import { ref, reactive } from 'vue'
const route = useRoute()
const router = useRouter()

const type = ref('')
const formData = ref({
            multiSpectraName: '',
            plateId: 0,
            multiSpectraTypeId: 0,
            aerialPhotographyIds: '',
        })
// 验证规则
const rule = reactive({
})

const elFormRef = ref()

// 初始化方法
const init = async () => {
 // 建议通过url传参获取目标数据ID 调用 find方法进行查询数据操作 从而决定本页面是create还是update 以下为id作为url参数示例
    if (route.query.id) {
      const res = await findMultiSpectraAnalysis({ ID: route.query.id })
      if (res.code === 0) {
        formData.value = res.data.reMtSpectraAly
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
               res = await createMultiSpectraAnalysis(formData.value)
               break
             case 'update':
               res = await updateMultiSpectraAnalysis(formData.value)
               break
             default:
               res = await createMultiSpectraAnalysis(formData.value)
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
