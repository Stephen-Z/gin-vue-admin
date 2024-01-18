<template>
  <div>
    <div class="gva-form-box">
      <el-form :model="formData" ref="elFormRef" label-position="right" :rules="rule" label-width="80px">
        <el-form-item label="序号:" prop="orderNum">
          <el-input v-model.number="formData.orderNum" :clearable="true" placeholder="请输入" />
       </el-form-item>
        <el-form-item label="登记日期:" prop="registerDate">
          <el-date-picker v-model="formData.registerDate" type="date" placeholder="选择日期" :clearable="true"></el-date-picker>
       </el-form-item>
        <el-form-item label="问题名称:" prop="problemName">
          <el-input v-model="formData.problemName" :clearable="true" placeholder="请输入" />
       </el-form-item>
        <el-form-item label="问题类型:" prop="problemType">
          <el-input v-model="formData.problemType" :clearable="true" placeholder="请输入" />
       </el-form-item>
        <el-form-item label="问题来源:" prop="problemSource">
          <el-input v-model="formData.problemSource" :clearable="true" placeholder="请输入" />
       </el-form-item>
        <el-form-item label="问题描述:" prop="problemDesc">
          <el-input v-model="formData.problemDesc" :clearable="true" placeholder="请输入" />
       </el-form-item>
        <el-form-item label="负责人:" prop="personCharge">
          <el-input v-model="formData.personCharge" :clearable="true" placeholder="请输入" />
       </el-form-item>
        <el-form-item label="问题图片:" prop="problemImage">
          <SelectImage v-model="formData.problemImage" />
       </el-form-item>
        <el-form-item label="处理措施:" prop="handMeasurce">
          <el-input v-model="formData.handMeasurce" :clearable="true" placeholder="请输入" />
       </el-form-item>
        <el-form-item label="当前状态:" prop="currState">
        <el-select v-model="formData.currState" placeholder="请选择" style="width:100%" :clearable="true">
          <el-option v-for="item in ['已处理','未处理']" :key="item" :label="item" :value="item" />
        </el-select>
       </el-form-item>
        <el-form-item label="经度:" prop="lng">
          <el-input-number v-model="formData.lng" :precision="2" :clearable="true"></el-input-number>
       </el-form-item>
        <el-form-item label="纬度:" prop="lat">
          <el-input-number v-model="formData.lat" :precision="2" :clearable="true"></el-input-number>
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
  name: 'ProblemRecord'
}
</script>

<script setup>
import {
  createProblemRecord,
  updateProblemRecord,
  findProblemRecord
} from '@/api/problemRecord'

// 自动获取字典
import { getDictFunc } from '@/utils/format'
import { useRoute, useRouter } from "vue-router"
import { ElMessage } from 'element-plus'
import { ref, reactive } from 'vue'
import SelectImage from '@/components/selectImage/selectImage.vue'
const route = useRoute()
const router = useRouter()

const type = ref('')
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
            nestid: "",
        })
// 验证规则
const rule = reactive({
})

const elFormRef = ref()

// 初始化方法
const init = async () => {
 // 建议通过url传参获取目标数据ID 调用 find方法进行查询数据操作 从而决定本页面是create还是update 以下为id作为url参数示例
    if (route.query.id) {
      const res = await findProblemRecord({ ID: route.query.id })
      if (res.code === 0) {
        formData.value = res.data.repbRecord
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
