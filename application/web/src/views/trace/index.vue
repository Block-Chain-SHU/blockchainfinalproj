<template>
  <div class="trace-container">
    <el-input v-model="input" placeholder="请输入学籍追踪码查询" style="width: 300px;margin-right: 15px;" />
    <el-button type="primary" plain @click="queryTraceInfo"> 查询 </el-button>
    <el-button type="success" plain @click="queryAllTraceInfo"> 获取所有培养信息 </el-button>
    <el-table
      :data="tracedata"
      style="width: 100%"
    >
      <el-table-column type="expand">
        <template slot-scope="props">
          <el-form label-position="left" inline class="demo-table-expand">
            <div><span class="trace-text" style="color: #67C23A;">学生基础信息</span></div>
            <el-form-item label="学生姓名：">
              <span>{{ props.row.farmer_input.fa_studentName }}</span>
            </el-form-item>
            <el-form-item label="专业方向：">
              <span>{{ props.row.farmer_input.fa_major }}</span>
            </el-form-item>
            <el-form-item label="入学时间：">
              <span>{{ props.row.farmer_input.fa_enrollTime }}</span>
            </el-form-item>
            <el-form-item label="毕业时间：">
              <span>{{ props.row.farmer_input.fa_graduationTime }}</span>
            </el-form-item>
            <el-form-item label="培养导师：">
              <span>{{ props.row.farmer_input.fa_mentorName }}</span>
            </el-form-item>
            <el-form-item v-if="props.row.farmer_input.fa_img" label="相关图片（点击下载）：" class="image-item">
              <a :href="`${baseApi}getImg/${props.row.farmer_input.fa_img}`" target="_blank">
                <el-image
                  style="width: 100px; height: 100px;"
                  :src="`${baseApi}getImg/${props.row.farmer_input.fa_img}`"
                  fit="cover"
                />
              </a>
            </el-form-item>
            <el-form-item label="区块链交易ID：">
              <span>{{ props.row.farmer_input.fa_txid }}</span>
            </el-form-item>
            <el-form-item label="区块链交易时间：">
              <span>{{ props.row.farmer_input.fa_timestamp }}</span>
            </el-form-item>
            <div><span class="trace-text" style="color: #409EFF;">课程/实验室信息</span></div>
            <el-form-item label="课程/实验室名称：">
              <span>{{ props.row.factory_input.fac_courseOrLabName }}</span>
            </el-form-item>
            <el-form-item label="课程批次：">
              <span>{{ props.row.factory_input.fac_courseBatch }}</span>
            </el-form-item>
            <el-form-item label="开课/实验时间：">
              <span>{{ props.row.factory_input.fac_courseTime }}</span>
            </el-form-item>
            <el-form-item label="实验室/课程平台：">
              <span>{{ props.row.factory_input.fac_platformName }}</span>
            </el-form-item>
            <el-form-item label="负责人联系方式：">
              <span>{{ props.row.factory_input.fac_contactNumber }}</span>
            </el-form-item>
            <el-form-item v-if="props.row.factory_input.fac_img" label="相关图片（点击下载）：" class="image-item">
              <a :href="`${baseApi}getImg/${props.row.factory_input.fac_img}`" target="_blank">
                <el-image
                  style="width: 100px; height: 100px;"
                  :src="`${baseApi}getImg/${props.row.factory_input.fac_img}`"
                  fit="cover"
                />
              </a>
            </el-form-item>
            <el-form-item label="区块链交易ID：">
              <span>{{ props.row.factory_input.fac_txid }}</span>
            </el-form-item>
            <el-form-item label="区块链交易时间：">
              <span>{{ props.row.factory_input.fac_timestamp }}</span>
            </el-form-item>
            <div><span class="trace-text" style="color: #E6A23C;">项目实践/竞赛信息</span></div>
            <el-form-item label="导师姓名：">
              <span>{{ props.row.driver_input.dr_mentorName }}</span>
            </el-form-item>
            <el-form-item label="指导年限：">
              <span>{{ props.row.driver_input.dr_experienceYears }}</span>
            </el-form-item>
            <el-form-item label="联系电话：">
              <span>{{ props.row.driver_input.dr_phone }}</span>
            </el-form-item>
            <el-form-item label="项目/竞赛编号：">
              <span>{{ props.row.driver_input.dr_projectCode }}</span>
            </el-form-item>
            <el-form-item label="项目贡献说明：">
              <span>{{ props.row.driver_input.dr_contribution }}</span>
            </el-form-item>
            <el-form-item v-if="props.row.driver_input.dr_img" label="相关图片（点击下载）：" class="image-item">
              <a :href="`${baseApi}getImg/${props.row.driver_input.dr_img}`" target="_blank">
                <el-image
                  style="width: 100px; height: 100px;"
                  :src="`${baseApi}getImg/${props.row.driver_input.dr_img}`"
                  fit="cover"
                />
              </a>
            </el-form-item>
            <el-form-item label="区块链交易ID：">
              <span>{{ props.row.driver_input.dr_txid }}</span>
            </el-form-item>
            <el-form-item label="区块链交易时间：">
              <span>{{ props.row.driver_input.dr_timestamp }}</span>
            </el-form-item>
            <div><span class="trace-text" style="color: #909399;">就业/升学去向信息</span></div>
            <el-form-item label="签约/录取时间：">
              <span>{{ props.row.shop_input.sh_offerTime }}</span>
            </el-form-item>
            <el-form-item label="入职/入学时间：">
              <span>{{ props.row.shop_input.sh_onboardTime }}</span>
            </el-form-item>
            <el-form-item label="单位/目标院校：">
              <span>{{ props.row.shop_input.sh_orgName }}</span>
            </el-form-item>
            <el-form-item label="单位/院校地址：">
              <span>{{ props.row.shop_input.sh_orgAddress }}</span>
            </el-form-item>
            <el-form-item label="联系人电话：">
              <span>{{ props.row.shop_input.sh_contactPhone }}</span>
            </el-form-item>
            <el-form-item v-if="props.row.shop_input.sh_img" label="相关图片(点击下载）：" class="image-item">
              <a :href="`${baseApi}getImg/${props.row.shop_input.sh_img}`" target="_blank">
                <el-image
                  style="width: 100px; height: 100px;"
                  :src="`${baseApi}getImg/${props.row.shop_input.sh_img}`"
                  fit="cover"
                />
              </a>
            </el-form-item>
            <el-form-item label="区块链交易ID：">
              <span>{{ props.row.shop_input.sh_txid }}</span>
            </el-form-item>
            <el-form-item label="区块链交易时间：">
              <span>{{ props.row.shop_input.sh_timestamp }}</span>
            </el-form-item>
          </el-form>
        </template>
      </el-table-column>
      <el-table-column
        label="学籍追踪码"
        prop="student_trace_id"
      />
      <el-table-column
        label="学生姓名"
        prop="farmer_input.fa_studentName"
      />
      <el-table-column
        label="毕业时间"
        prop="farmer_input.fa_graduationTime"
      />
    </el-table>
  </div>
</template>

<script>
import { mapGetters } from 'vuex'
import { getTraceInfo, getTraceList, getAllTraceInfo, getTraceHistory } from '@/api/trace'

export default {
  name: 'Trace',
  data() {
    return {
      tracedata: [],
      loading: false,
      input: '',
      baseApi: process.env.VUE_APP_BASE_API
    }
  },
  computed: {
    ...mapGetters([
      'name',
      'userType'
    ])
  },
  created() {
    const code = this.$route.params.student_trace_id
    if (code) {
      this.input = code
      this.queryTraceInfo()
    } else {
      getTraceList().then(res => {
        this.tracedata = JSON.parse(res.data)
      })
    }
  },
  methods: {
    queryAllTraceInfo() {
      getAllTraceInfo().then(res => {
        this.tracedata = JSON.parse(res.data)
      })
    },
    queryTraceHistory() {
      getTraceHistory().then(res => {
        // console.log(res)
      })
    },
    queryTraceInfo() {
      var formData = new FormData()
      formData.append('student_trace_id', this.input)
      getTraceInfo(formData).then(res => {
        if (res.code === 200) {
          // console.log(res)
          this.tracedata = []
          this.tracedata[0] = JSON.parse(res.data)
          return
        } else {
          this.$message.error(res.message)
        }
      })
    }
  }
}
</script>

<style lang="scss" scoped>

.demo-table-expand {
    font-size: 0;
  }
  .demo-table-expand label {
    width: 90px;
    color: #99a9bf;
  }
  .demo-table-expand .el-form-item {
    margin-right: 0;
    margin-bottom: 0;
    width: 50%;
  }
.trace {
  &-container {
    margin: 30px;
  }
  &-text {
    font-size: 30px;
    line-height: 46px;
  }
}

.demo-table-expand {
  font-size: 0;
}

.demo-table-expand label {
  width: 90px;
  color: #99a9bf;
}

.demo-table-expand .el-form-item {
  margin-right: 0;
  margin-bottom: 0;
  width: 50%;
  display: inline-block;
  vertical-align: top;
}

.demo-table-expand .image-item {
  width: 100%;
  margin-top: 10px;
  margin-bottom: 10px;
}

.demo-table-expand .image-item .el-form-item__content {
  display: flex;
  align-items: center;
  gap: 10px;
}

</style>
