<template>
  <div>
    <div class="gva-form-box">
      <el-form :model="formData" ref="elFormRef" label-position="right" :rules="rule" label-width="80px">
        <el-form-item label="指标名:" prop="metric_name">
          <el-input v-model="formData.metric_name" :clearable="true" placeholder="请输入" />
       </el-form-item>
        <el-form-item label="机器ip:" prop="machine_ip">
          <el-input v-model="formData.machine_ip" :clearable="true" placeholder="请输入" />
       </el-form-item>
        <el-form-item label="创建时间:" prop="sample_time">
          <el-date-picker v-model="formData.sample_time" type="date" placeholder="选择日期" :clearable="true"></el-date-picker>
       </el-form-item>
        <el-form-item label="次数:" prop="count">
          <el-input v-model.number="formData.count" :clearable="true" placeholder="请输入" />
       </el-form-item>
        <el-form-item label="交易所:" prop="exchange">
          <el-input v-model="formData.exchange" :clearable="true" placeholder="请输入" />
       </el-form-item>
        <el-form-item label="耗时(平均):" prop="mean">
          <el-input-number v-model="formData.mean" :precision="2" :clearable="true"></el-input-number>
       </el-form-item>
        <el-form-item label="耗时(平均):" prop="mid">
          <el-input-number v-model="formData.mid" :precision="2" :clearable="true"></el-input-number>
       </el-form-item>
        <el-form-item label="标准差:" prop="std">
          <el-input-number v-model="formData.std" :precision="2" :clearable="true"></el-input-number>
       </el-form-item>
        <el-form-item label="Symbol:" prop="symbol">
          <el-input v-model="formData.symbol" :clearable="true" placeholder="请输入" />
       </el-form-item>
        <el-form-item>
          <el-button type="primary" @click="save">保存</el-button>
          <el-button type="primary" @click="back">返回</el-button>
        </el-form-item>
      </el-form>
    </div>
  </div>
</template>

<script setup>
import {
  createSpeed,
  updateSpeed,
  findSpeed
} from '@/api/speed'

defineOptions({
    name: 'SpeedForm'
})

// 自动获取字典
import { getDictFunc } from '@/utils/format'
import { useRoute, useRouter } from "vue-router"
import { ElMessage } from 'element-plus'
import { ref, reactive } from 'vue'

const route = useRoute()
const router = useRouter()

const type = ref('')
const formData = ref({
            metric_name: '',
            machine_ip: '',
            sample_time: new Date(),
            count: 0,
            exchange: '',
            mean: 0,
            mid: 0,
            std: 0,
            symbol: '',
        })
// 验证规则
const rule = reactive({
})

const elFormRef = ref()

// 初始化方法
const init = async () => {
 // 建议通过url传参获取目标数据ID 调用 find方法进行查询数据操作 从而决定本页面是create还是update 以下为id作为url参数示例
    if (route.query.id) {
      const res = await findSpeed({ ID: route.query.id })
      if (res.code === 0) {
        formData.value = res.data.remetric
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
               res = await createSpeed(formData.value)
               break
             case 'update':
               res = await updateSpeed(formData.value)
               break
             default:
               res = await createSpeed(formData.value)
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
