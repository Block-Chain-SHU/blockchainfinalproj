<template>
  <div class="uplink-container">
    <div style="color:#909399;margin-bottom: 30px">
      当前用户：{{ name }};
      用户角色: {{ userType }}
    </div>
    <div>
      <el-form ref="form" :model="tracedata" label-width="80px" size="mini" style="">
        <el-form-item v-show="userType!='学生'&userType!='学院/企业/本人查询'" label="学籍追踪码:" style="width: 300px" label-width="120px">
          <el-input v-model="tracedata.student_trace_id" />
        </el-form-item>

        <div v-show="userType=='学生'">
          <el-form-item label="学生姓名:" style="width: 300px" label-width="120px">
            <el-input v-model="tracedata.Farmer_input.Fa_studentName" />
          </el-form-item>
          <el-form-item label="专业方向:" style="width: 300px" label-width="120px">
            <el-input v-model="tracedata.Farmer_input.Fa_major" />
          </el-form-item>
          <el-form-item label="入学时间:" style="width: 300px" label-width="120px">
            <el-input v-model="tracedata.Farmer_input.Fa_enrollTime" />
          </el-form-item>
          <el-form-item label="毕业时间:" style="width: 300px" label-width="120px">
            <el-input v-model="tracedata.Farmer_input.Fa_graduationTime" />
          </el-form-item>
          <el-form-item label="培养导师:" style="width: 300px" label-width="120px">
            <el-input v-model="tracedata.Farmer_input.Fa_mentorName" />
          </el-form-item>
        </div>
        <div v-show="userType=='课程/实验室'">
          <el-form-item label="课程/实验室名称:" style="width: 300px" label-width="120px">
            <el-input v-model="tracedata.Factory_input.Fac_courseOrLabName" />
          </el-form-item>
          <el-form-item label="课程批次:" style="width: 300px" label-width="120px">
            <el-input v-model="tracedata.Factory_input.Fac_courseBatch" />
          </el-form-item>
          <el-form-item label="开课/实验时间:" style="width: 300px" label-width="120px">
            <el-input v-model="tracedata.Factory_input.Fac_courseTime" />
          </el-form-item>
          <el-form-item label="实验室/课程平台:" style="width: 300px" label-width="120px">
            <el-input v-model="tracedata.Factory_input.Fac_platformName" />
          </el-form-item>
          <el-form-item label="负责人联系方式:" style="width: 300px" label-width="120px">
            <el-input v-model="tracedata.Factory_input.Fac_contactNumber" />
          </el-form-item>
        </div>
        <div v-show="userType=='项目实践/竞赛导师'">
          <el-form-item label="导师姓名:" style="width: 300px" label-width="120px">
            <el-input v-model="tracedata.Driver_input.Dr_mentorName" />
          </el-form-item>
          <el-form-item label="指导年限:" style="width: 300px" label-width="120px">
            <el-input v-model="tracedata.Driver_input.Dr_experienceYears" />
          </el-form-item>
          <el-form-item label="联系电话:" style="width: 300px" label-width="120px">
            <el-input v-model="tracedata.Driver_input.Dr_phone" />
          </el-form-item>
          <el-form-item label="项目/竞赛编号:" style="width: 300px" label-width="120px">
            <el-input v-model="tracedata.Driver_input.Dr_projectCode" />
          </el-form-item>
          <el-form-item label="项目贡献说明：" style="width: 300px" label-width="120px">
            <el-input v-model="tracedata.Driver_input.Dr_contribution" />
          </el-form-item>
        </div>
        <div v-show="userType=='就业单位/研究生去向'">
          <el-form-item label="签约/录取时间:" style="width: 300px" label-width="120px">
            <el-input v-model="tracedata.Shop_input.Sh_offerTime" />
          </el-form-item>
          <el-form-item label="入职/入学时间:" style="width: 300px" label-width="120px">
            <el-input v-model="tracedata.Shop_input.Sh_onboardTime" />
          </el-form-item>
          <el-form-item label="单位/目标院校:" style="width: 300px" label-width="120px">
            <el-input v-model="tracedata.Shop_input.Sh_orgName" />
          </el-form-item>
          <el-form-item label="单位/院校地址:" style="width: 300px" label-width="120px">
            <el-input v-model="tracedata.Shop_input.Sh_orgAddress" />
          </el-form-item>
          <el-form-item label="联系人电话:" style="width: 300px" label-width="120px">
            <el-input v-model="tracedata.Shop_input.Sh_contactPhone" />
          </el-form-item>
        </div>
        <el-form-item v-show="userType != '学院/企业/本人查询'" label="上传佐证材料:" style="width: 300px" label-width="120px">
          <el-upload
            action="#"
            class="upload-demo"
            :show-file-list="false"
            :on-change="onImageSelected"
            :before-upload="beforeUpload"
            accept="image/*"
          >
            <el-button type="primary" size="mini">选择图片</el-button>
          </el-upload>

          <div v-if="imagePreview" style="margin-top: 10px;">
            <img :src="imagePreview" alt="预览图" style="max-width: 100%; max-height: 150px; border: 1px solid #dcdfe6;">
          </div>
        </el-form-item>
      </el-form>
      <span slot="footer" style="color: gray;" class="dialog-footer">
        <el-button v-show="userType != '学院/企业/本人查询'" type="primary" plain style="margin-left: 220px;" @click="submittracedata()">提 交</el-button>
      </span>
      <span v-show="userType == '学院/企业/本人查询'" slot="footer" style="color: gray;" class="dialog-footer">
        查询角色没有录入权限，请使用追踪查询功能。
      </span>
    </div>
  </div>
</template>

<script>
import { mapGetters } from 'vuex'
import { uplink } from '@/api/trace'

export default {
  name: 'Uplink',
  data() {
    return {
      tracedata: {
        student_trace_id: '',
        Farmer_input: {
          Fa_studentName: '',
          Fa_major: '',
          Fa_enrollTime: '',
          Fa_graduationTime: '',
          Fa_mentorName: '',
          Fa_img: null
        },
        Factory_input: {
          Fac_courseOrLabName: '',
          Fac_courseBatch: '',
          Fac_courseTime: '',
          Fac_platformName: '',
          Fac_contactNumber: ''
        },
        Driver_input: {
          Dr_mentorName: '',
          Dr_experienceYears: '',
          Dr_phone: '',
          Dr_projectCode: '',
          Dr_contribution: ''
        },
        Shop_input: {
          Sh_offerTime: '',
          Sh_onboardTime: '',
          Sh_orgName: '',
          Sh_orgAddress: '',
          Sh_contactPhone: ''
        }
      },
      imageFile: null,
      imagePreview: null,
      loading: false
    }
  },
  computed: {
    ...mapGetters([
      'name',
      'userType'
    ])
  },
  methods: {
    submittracedata() {
      console.log(this.tracedata)
      const loading = this.$loading({
        lock: true,
        text: '数据上链中...',
        spinner: 'el-icon-loading',
        background: 'rgba(0, 0, 0, 0.7)'
      })
      var formData = new FormData()
      formData.append('student_trace_id', this.tracedata.student_trace_id)
      formData.append('file', this.imageFile)
      // 根据不同的用户给arg1、arg2、arg3..赋值,
      switch (this.userType) {
        case '学生':
          formData.append('arg1', this.tracedata.Farmer_input.Fa_studentName)
          formData.append('arg2', this.tracedata.Farmer_input.Fa_major)
          formData.append('arg3', this.tracedata.Farmer_input.Fa_enrollTime)
          formData.append('arg4', this.tracedata.Farmer_input.Fa_graduationTime)
          formData.append('arg5', this.tracedata.Farmer_input.Fa_mentorName)
          break
        case '课程/实验室':
          formData.append('arg1', this.tracedata.Factory_input.Fac_courseOrLabName)
          formData.append('arg2', this.tracedata.Factory_input.Fac_courseBatch)
          formData.append('arg3', this.tracedata.Factory_input.Fac_courseTime)
          formData.append('arg4', this.tracedata.Factory_input.Fac_platformName)
          formData.append('arg5', this.tracedata.Factory_input.Fac_contactNumber)
          break
        case '项目实践/竞赛导师':
          formData.append('arg1', this.tracedata.Driver_input.Dr_mentorName)
          formData.append('arg2', this.tracedata.Driver_input.Dr_experienceYears)
          formData.append('arg3', this.tracedata.Driver_input.Dr_phone)
          formData.append('arg4', this.tracedata.Driver_input.Dr_projectCode)
          formData.append('arg5', this.tracedata.Driver_input.Dr_contribution)
          break
        case '就业单位/研究生去向':
          formData.append('arg1', this.tracedata.Shop_input.Sh_offerTime)
          formData.append('arg2', this.tracedata.Shop_input.Sh_onboardTime)
          formData.append('arg3', this.tracedata.Shop_input.Sh_orgName)
          formData.append('arg4', this.tracedata.Shop_input.Sh_orgAddress)
          formData.append('arg5', this.tracedata.Shop_input.Sh_contactPhone)
          break
      }
      uplink(formData).then(res => {
        if (res.code === 200) {
          loading.close()
          this.$message({
            message: '上链成功，交易ID：' + res.txid + '\n学籍追踪码：' + res.student_trace_id,
            type: 'success'
          })
        } else {
          loading.close()
          this.$message({
            message: '上链失败',
            type: 'error'
          })
        }
      }).catch(err => {
        loading.close()
        console.log(err)
      })
    },
    onImageSelected(file) {
      // 生成预览图
      this.imagePreview = URL.createObjectURL(file.raw)
      this.imageFile = file.raw
    },
    beforeUpload() {
      // 禁止自动上传
      return false
    }
  }
}

</script>

<style lang="scss" scoped>
.uplink {
  &-container {
    margin: 30px;
  }
  &-text {
    font-size: 30px;
    line-height: 46px;
  }
}
</style>
