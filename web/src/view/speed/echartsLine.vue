<template>
  
  <div class="dashboard-line-box">
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
    
      <el-form-item label="指标名" prop="metric_name">
        <el-input v-model="searchInfo.metric_name" placeholder="搜索条件" />

      </el-form-item>
      <el-form-item label="机器ip" prop="machine_ip">
        <el-input v-model="searchInfo.machine_ip" placeholder="搜索条件" />

      </el-form-item>
      <el-form-item label="交易所" prop="exchange">
        <el-input v-model="searchInfo.exchange" placeholder="搜索条件" />

      </el-form-item>
      <el-form-item label="Symbol" prop="symbol">
         <el-input v-model="searchInfo.symbol" placeholder="搜索条件" />
      </el-form-item>

      <el-form-item>
        <el-button type="primary" icon="search" @click="onSubmit">查询</el-button>
        <el-button icon="refresh" @click="onReset">重置</el-button>
      </el-form-item>
    </el-form>
    </div>
    <div
      ref="echart"
      class="dashboard-line"
    />
    
  </div>
</template>
<script setup>
import { getDictFunc, formatDate, formatBoolean, filterDict, ReturnArrImg, onDownloadFile } from '@/utils/format'
import {
  createSpeed,
  deleteSpeed,
  deleteSpeedByIds,
  updateSpeed,
  findSpeed,
  getSpeedList
} from '@/api/speed'
import * as echarts from 'echarts'
import { nextTick, onMounted, onUnmounted, ref, reactive, onBeforeMount } from 'vue'
import { useWindowResize } from '@/hooks/use-windows-resize'
var dataAxis = []

const tableData = ref(null)
const page = ref(1)
const total = ref(0)
const pageSize = ref(10)
const searchInfo = ref({})
const maxCost = ref(0)

var midData = []
var meadData = []
var stdData = []

const getTableData = async() => {
  const table = await getSpeedList({ page: page.value, pageSize: pageSize.value, ...searchInfo.value })
  if (table.code === 0) {
    console.log(table.data)
    midData = []
    dataAxis = []
    table.data.list.forEach(element => {
      if (element.mid > maxCost) {
        maxCost = element.mid
      }
      midData.push(element.mid)
      meadData.push(element.mean)
      stdData.push(element.std)
      dataAxis.push(formatDate(element.sample_time))
    });
    console.log(dataAxis)
    console.log(midData)
    tableData.value = table.data.list
    total.value = table.data.total
    page.value = table.data.page
    pageSize.value = table.data.pageSize
    setOptions()
  }
}

let chart = null
const echart = ref(null)

useWindowResize(() => {
  if (!chart) {
    return
  }
  chart.resize()
})

const initChart = () => {
  if (chart) {
    chart = null
  }
  chart = echarts.init(echart.value)
  getTableData()
}

const setOptions = () => {
  console.log("set options")
  chart.setOption({
    title: {
      text: "耗时趋势"
    },
    grid: {
      left: '3%',
      right: '4%',
      bottom: '3%',
      containLabel: true
    },
    tooltip: {
      trigger: 'axis'
    },
    legend: {
      data: ['mid', 'mean', 'std']
    },
    toolbox: {
      feature: {
        saveAsImage: {}
      }
    },
    xAxis: {
      data: dataAxis,
      type: 'category',
      boundaryGap: false,
    },
    yAxis: {
      type: "value"
    },
    dataZoom: [
      {
        type: 'inside',
      },
    ],
    series: [
      {
        name: 'mid',
        type: 'line',
        type: 'line',
        stack: 'Total',
        data: midData,
      },
      {
        name: 'mean',
        type: 'line',
        type: 'line',
        stack: 'Total',
        data: meadData,
      },
      {
        name: 'std',
        type: 'line',
        type: 'line',
        stack: 'Total',
        data: stdData,
      },
    ],
  })
}

onMounted(async() => {
  await nextTick()
  initChart()
})

onUnmounted(() => {
  if (!chart) {
    return
  }
  chart.dispose()
  chart = null
})


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
const elSearchFormRef = ref()

// 搜索
const onSubmit = () => {
  elSearchFormRef.value?.validate(async(valid) => {
    if (!valid) return
    page.value = 1
    pageSize.value = 10
    getTableData()
  })
}

// 重置
const onReset = () => {
  searchInfo.value = {}
  getTableData()
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
</script>
<style lang="scss" scoped>
.dashboard-line-box {
  .dashboard-line {
    background-color: #fff;
    height: 360px;
    width: 100%;
  }
  .dashboard-line-title {
    font-weight: 600;
    margin-bottom: 12px;
  }
}
</style>
