<template>
  <div class="warp">
    <Header @updateNum='getNum' :nume="nume"></Header>
    <div class="content">
      <!-- 首页 -->
      <div v-if="num == 1" class="content-index">
        <div class="content-index-center">
          病害识别系统
        </div>
        <el-carousel height="100vh">
          <el-carousel-item>
            <img style="height: 100%;width: 100%;" src="../assets/img/12.jpg" alt="">
          </el-carousel-item>
          <el-carousel-item>
            <img style="height: 100%;width: 100%;" src="../assets/img/22.jpg" alt="">
          </el-carousel-item>
          <el-carousel-item>
            <img style="height: 100%;width: 100%;" src="../assets/img/34.jpg" alt="">
          </el-carousel-item>
        </el-carousel>
      </div>
      <!-- 病害百科 -->
      <div v-else-if="num == 2" class="content-record">
        <div class="content-index-center-1">
          <div class="main" style="height: 400px;width: 280px;">
            <div class="main-li">
              <div :class="[disname == '矮化叶病' ? 'text active' : 'text']" @click="link('矮化叶病')">矮化叶病:</div>
              <div style="background: #5672c4;" class="main-li-div"></div>
            </div>
            <div class="main-li">
              <div :class="[disname == '锈病' ? 'text active' : 'text']" @click="link('锈病')">锈病:</div>
              <div style="background: #e96366;" class="main-li-div"></div>
            </div>
            <div class="main-li">
              <div :class="[disname == '叶斑病' ? 'text active' : 'text']" @click="link('叶斑病')">叶斑病:</div>
              <div style="background: #7ac1dd;" class="main-li-div"></div>
            </div>
            <div class="main-li">
              <div :class="[disname == '灰斑病' ? 'text active' : 'text']" @click="link('灰斑病')">灰斑病:</div>
              <div style="background: #f7c75f;" class="main-li-div"></div>
            </div>
            <div class="img">
              <!-- 矮化叶病 -->
              <div v-if="disname == '矮化叶病'">
                <div style="margin-bottom: 10px;">示例：<span
                    style="margin-left: 5px;font-size: 16px;color: red;">矮化叶病</span></div>
                <div class="img-warp">
                  <div class="img-li">
                    <img src="../assets/img/aihuayebing107.jpg" alt="">
                  </div>
                  <div class="img-li">
                    <img src="../assets/img/aihuayebing127.jpg" alt="">
                  </div>
                </div>
              </div>
              <!-- 锈病 -->
              <div v-else-if="disname == '锈病'">
                <div style="margin-bottom: 10px;">示例：<span
                    style="margin-left: 5px;font-size: 16px;color: red;">锈病</span></div>
                <div class="img-warp">
                  <div class="img-li">
                    <img src="../assets/img/xiubing172.jpg" alt="">
                  </div>
                  <div class="img-li">
                    <img src="../assets/img/xiubing237.jpg" alt="">
                  </div>
                </div>
              </div>
              <!-- 叶斑病 -->
              <div v-else-if="disname == '叶斑病'">
                <div style="margin-bottom: 10px;">示例：<span
                    style="margin-left: 5px;font-size: 16px;color: red;">叶斑病</span></div>
                <div class="img-warp">
                  <div class="img-li">
                    <img src="../assets/img/yebanbing79.jpg" alt="">
                  </div>
                  <div class="img-li">
                    <img src="../assets/img/yebanbing78.jpg" alt="">
                  </div>
                </div>
              </div>
              <!-- 灰斑病 -->
              <div v-else-if="disname == '灰斑病'">
                <div style="margin-bottom: 10px;">示例：<span
                    style="margin-left: 5px;font-size: 16px;color: red;">灰斑病</span></div>
                <div class="img-warp">
                  <div class="img-li">
                    <img src="../assets/img/huibanbing35.jpg" alt="">
                  </div>
                  <div class="img-li">
                    <img src="../assets/img/huibanbing44.jpg" alt="">
                  </div>
                </div>
              </div>

            </div>
          </div>
          <div class="introduce">
            <div class="introduce-title">
              病虫害介绍
            </div>
            <div class="introduce-text">
              <div>名称: <span style="margin-left: 15px;font-size: 12px;color: #999;">{{
                  dataObj['value']['DiseaseName']
                }}</span></div>
              <div>别名:<span style="margin-left: 15px;font-size: 12px;color: #999;">{{
                  dataObj['value']['Alias']
                }}</span></div>
              <div>简介:<span style="margin-left: 15px;font-size: 12px;color: #999;">{{
                  dataObj['value']['Introduction']
                }}</span></div>
              <div>症状:<span
                  style="margin-left: 15px;font-size: 12px;color: #999;">{{ dataObj['value']['SymptomsOfHarm'] }}</span>
              </div>
              <div>病害条件:<span style="margin-left: 15px;font-size: 12px;color: #999;">{{
                  dataObj['value']['OnsetConditions']
                }}</span></div>
              <div>如何防治:<span
                  style="margin-left: 15px;font-size: 12px;color: #999;">{{ dataObj['value']['Treatment'] }}</span>
              </div>
            </div>
          </div>
        </div>
      </div>
      <!-- 病害检测 -->
      <div v-else-if="num == 3" class="content-detection">
        <div class="content-detection-center">
          <template v-if="is_o">
            <div class="content-detection-center-top">
              <img v-if="url" style="width: 100%;" :src="url" alt="">
              <div class="opload">
                <input type="file" @change="filechange">
                <!-- <img src="../assets/img/上传加号.png" alt=""> -->
              </div>
            </div>
            <div @click="changeIs" class="content-detection-center-btn">
              检测
            </div>
          </template>
          <template v-else>
            <div class="content-detection-center-btn-center">
              <div class="content-detection-center-btn-left">
                <div class="back">
                  <el-button type="info" plain @click="back">返回</el-button>
                </div>
                <div class="top">
                  <img :src="urlTo" alt="">
                </div>
                <div class="bottom">
                  <div style="cursor: pointer;" :class="[activeIndex == index ? 'text active' : 'text']"
                       v-if="arr.length" v-for="(item,index) in arr" @click="diseaseChange(item,index)"><span
                      style="color: #000;margin-right: 5px;">{{ '病害' + (index + 1) + '：' }}</span>{{
                      item['DiseaseName']
                    }}
                  </div>
                  <div v-else>未检测到病害</div>
                </div>
              </div>
              <div class="content-detection-center-btn-right">
                <template v-if="arr.length">
                  <div style="display: flex;align-items: center;">
                    <div>病害信息:<span
                        style="margin-left: 15px;font-size: 12px;color: red;">{{ obj.value['Introduction'] }}</span>
                    </div>
                  </div>
                  <div style="margin-top: 50px;display: flex;align-items: center;">
                    <div @click="console.log(obj)">治疗信息: <span
                        style="margin-left: 15px;font-size: 12px;color: red;">{{
                        obj.value['Treatment']
                      }}</span></div>

                  </div>
                </template>
              </div>
            </div>
          </template>
        </div>
      </div>
      <!-- 检测记录 -->
      <div v-else-if="num == 4" class="content-record">
        <div class="content-detection-center">
          <el-table :data="tableData">
            <el-table-column prop="id" label="ID" width="120"/>
            <el-table-column prop="DiseaseName" label="病害名称" width="120"/>
            <el-table-column prop="CreateTime" label="创建时间" width="180"/>
            <!--            <el-table-column prop="Introduction" label="病害条件" width="600" />-->
            <!--            <el-table-column prop="Treatment" label="治疗方法" width="120" />-->
            <el-table-column label="操作" width="120">
              <template #default="scope">
                <el-button link type="primary" size="small" @click="handleClick(scope.row)">详情</el-button>
              </template>
            </el-table-column>
          </el-table>
          <el-pagination style="display: flex;justify-content: flex-end;margin-top: 25px;" background
                         layout="prev, pager, next" @change="changePage" :total="total"/>
        </div>
      </div>

    </div>

  </div>

</template>

<script setup>
import * as echarts from "echarts";
import axios from 'axios'
import {onMounted, ref, nextTick, watch, reactive} from "vue"
import {ElMessage} from 'element-plus'
import {toRaw} from '@vue/reactivity'
import Header from '@/components/common/Header.vue'

const num = ref('1')
const total = ref(0)
const is_o = ref(true)
const getNum = val => {
  num.value = val
  nume.value = val
}

const handleClick = row => {
  num.value = 2
  disname.value = row['DiseaseName']
  nume.value = '2'
  getDiseases()
}

const back = () => {
  is_o.value = true
}

const tableData = ref([])

const nume = ref('1')

watch(num, (val) => {

})

const page = ref(1)

const link = data => {
  disname.value = data
  getDiseases()
}

const dataObj = reactive({
  ID: '',
  Alias: '',
  Introduction: '',
  DiseaseName: '',
  OnsetConditions: '',
  PathogenicFeatures: '',
  SymptomsOfHarm: '',
  Treatment: '',
})
const disname = ref('矮化叶病')

const url = ref('')

const urlTo = ref('')

const arr = ref([])

// http://47.120.79.124:8081
// /api
const newUrl = ref('/api')

const diseaseName = ref([])

const obj = reactive({
  value: {
    Introduction: '',
    Treatment: ''
  }
})

const activeIndex = ref(0)

const diseaseChange = async (item, index) => {
  console.log({item, index})
  activeIndex.value = index
  const res = await axios({
    url: newUrl.value + `/search?diseaseName=${item['DiseaseName']}`,
    method: 'get',
    headers: {
      'Authorization': "Bearer" + " " + (localStorage.getItem("loginResult") ? JSON.parse(localStorage.getItem("loginResult"))['token'] : '')
    },
  });
  console.log(res)
  if (res.data.code == 1000) {
      obj.value = res.data.data
  }
}

const filechange = e => {
  const resultFile = e.target.files;
  const aBlob = new Blob([resultFile[0]], {type: 'image'})
  const reader = new FileReader();
  const formData = new FormData();
  formData.append("image", e.target.files[0])
  reader.onload = (ev) => {
    const base64Url = ev.target.result; // 5 base64内容
    axios({
      url: newUrl.value + `/upImg`,
      method: 'post',
      headers: {
        'Authorization': "Bearer" + " " + (localStorage.getItem("loginResult") ? JSON.parse(localStorage.getItem("loginResult"))['token'] : '')
      },
      data: formData
    }).then((res) => {
      if (res.data.code == 1000) {
        url.value = base64Url
        urlTo.value = res.data.data['result_image'] ? ("data:image/png;base64," + res.data.data['result_image']) : base64Url
        diseaseName.value = res.data.data['result']
      } else {
        ElMessage({message: res.data.msg, type: 'error',})
      }
    })
  }
  reader.readAsDataURL(aBlob) // 4 将文件转换成指定类型的数据
}

const getDisease = (item, index) => {
  axios({
    url: newUrl.value + `/search?diseaseName=${item}`,
    method: 'get',
    headers: {
      'Authorization': "Bearer" + " " + (localStorage.getItem("loginResult") ? JSON.parse(localStorage.getItem("loginResult"))['token'] : '')
    },
  }).then((res) => {
    if (res.data.code == 1000) {
      if (index == 0) {
        obj.value = res.data.data
      }
      arr.value.push(res.data.data)
    }
  })
}

const getNumPage = () => {
  axios({
    url: newUrl.value + `/count`,
    method: 'get',
    headers: {
      'Authorization': "Bearer" + " " + (localStorage.getItem("loginResult") ? JSON.parse(localStorage.getItem("loginResult"))['token'] : '')
    },
  }).then((res) => {
    if (res.data.code == 1000) {
      console.log(res.data, 'ddd')
      total.value = res.data.data
    }
  })
}

const changePage = e => {
  console.log(e, 'eee')
  page.value = e
  getList()
}

const changeIs = () => {
  arr.value = []
  is_o.value = false
  // diseaseName.value = ref(['锈病','叶斑病'])
  if (diseaseName.value.length) {
    diseaseName.value.forEach((item, index) => {
      getDisease(item, index)
    })
  }
}


const getDiseases = () => {
  axios({
    url: newUrl.value + `/search?diseaseName=${disname.value}`,
    method: 'get',
    headers: {
      'Authorization': "Bearer" + " " + (localStorage.getItem("loginResult") ? JSON.parse(localStorage.getItem("loginResult"))['token'] : '')
    },
  }).then((res) => {
    if (res.data.code == 1000) {
      dataObj.value = res.data.data
    } else {
      dataObj.value = {
        ID: '',
        Alias: '',
        Introduction: '',
        DiseaseName: '',
        OnsetConditions: '',
        PathogenicFeatures: '',
        SymptomsOfHarm: '',
        Treatment: '',
      }
    }
  })
}

// 时间戳：1637244864707
/* 时间戳转换为时间 */
const timestampToTime = (timestamp) => {
  timestamp = timestamp ? timestamp : null;
  let date = new Date(timestamp);//时间戳为10位需*1000，时间戳为13位的话不需乘1000
  let Y = date.getFullYear() + '-';
  let M = (date.getMonth() + 1 < 10 ? '0' + (date.getMonth() + 1) : date.getMonth() + 1) + '-';
  let D = (date.getDate() < 10 ? '0' + date.getDate() : date.getDate()) + ' ';
  let h = (date.getHours() < 10 ? '0' + date.getHours() : date.getHours())
  return Y + M + D
}


const getList = () => {
  // http://47.120.79.124:8081
  axios({
    url: newUrl.value + `/record?page=${page.value}&size=10`,
    method: 'post',
    headers: {
      'Authorization': "Bearer" + " " + (localStorage.getItem("loginResult") ? JSON.parse(localStorage.getItem("loginResult"))['token'] : '')
    },
    data: JSON.stringify({
      username: localStorage.getItem("loginResult") ? JSON.parse(localStorage.getItem("loginResult"))['user_name'] : '',
      password: localStorage.getItem("loginResult") ? JSON.parse(localStorage.getItem("loginResult"))['user_id'] : '',
    })
  }).then((res) => {
    if (res.data.code == 1000) {
      res.data.data.forEach(k => {
        k['CreateTime'] = timestampToTime(k['CreateTime'])
      })
      tableData.value = res.data.data
    }
  })
}

onMounted(() => {
  getDiseases()
  getList()
  getNumPage()
});

</script>


<style scoped>
.img {
  margin-top: 30px;
}

.img-warp {
  display: flex;
  align-items: center;
}

.img-warp img {
  width: 150px;
  height: 200px;
}

.main-li {
  display: flex;
  align-items: center;
  justify-content: space-between;
  cursor: pointer;
}

.img-li + .img-li {
  margin-left: 30px;
}

.active {
  color: red;
}

.text:hover {
  color: red;
}

.main-li + .main-li {
  margin-top: 20px;
}

.main-li-div {
  margin-left: 15px;
  width: 180px;
  height: 30px;
}

.warp {
  display: flex;
  flex-direction: column;
  width: 100vw;
  height: 100vh;
  overflow: hidden;
}

.content {
  height: 80%;
}

/* 首页 */
.content-index {
  height: 100%;
  overflow: hidden;
  /* background-image: url('../assets/img/检测.png'); */
  background-size: 100% 100%;
}

.content-index-center {
  position: absolute;
  z-index: 9;
  top: 55%;
  left: 50%;
  transform: translate(-50%, -50%);
  width: 300px;
  line-height: 70px;
  text-align: center;
  background: rgba(255, 255, 255, .5);
  border-radius: 5px;
  color: #000;
  font-size: 32px;
  font-weight: 300;
  box-shadow: 0 0 2px 1px #fff;
}

.content-index-center-1 {
  display: flex;
  align-items: center;
  position: absolute;
  top: 50%;
  left: 50%;
  transform: translate(-50%, -50%);
}

.introduce {
  margin-left: 150px;
}

.introduce-title {
  font-size: 18px;
  font-weight: bold;
}

.introduce-text {
  margin: 20px 0 0 20px;
  font-size: 14px;
  font-weight: bold;
  height: 400px;
  width: 800px;
}

.introduce-text div {
}

.introduce-text div + div {
  margin-top: 15px;

}

.content-detection {

}

.content-detection-center {
  height: 300px;
  position: absolute;
  top: 45%;
  left: 50%;
  transform: translate(-50%, -50%);
  line-height: 50px;
  text-align: center;
}

.content-detection-center-top {
  position: relative;
  width: 800px;
  height: 400px;
  background: burlywood;
  box-shadow: 0 0 2px 1px #fff;
}

.content-detection-center-top img {
  width: 600px;
  height: 100%
}

.opload {
  position: absolute;
  display: flex;
  align-items: center;
  justify-content: center;
  left: 0;
  right: 0;
  top: 0;
  bottom: 0;
  cursor: pointer;

}

.opload img {
  width: 35px;
  object-fit: contain;
  margin-bottom: 50px;
  border: 1px solid #999;
}

.opload input[type="file"] {
  color: transparent;
  width: 100px;
  cursor: pointer;
}

.content-detection-center-btn {
  z-index: 9;
  margin: 0 auto;
  margin-top: 30px;
  width: 300px;
  line-height: 50px;
  text-align: center;
  background: #007BE5;
  border-radius: 5px;
  color: #fff;
  box-shadow: 0 0 2px 1px #fff;
  cursor: pointer;
}

.content-detection-center-btn-center {
  display: flex;
}

.bottom {
  display: flex;
  align-items: center;
}

.bottom div {
  text-align: left;
}

.bottom div + div {
  margin-left: 50px;
}

.top img {
  width: 800px;
  height: 400px;
}

.content-detection-center-btn-right {
  margin-left: 30px
}

.content-detection-center-btn-right div {
  width: 500px;
  height: 200px;
  overflow-y: scroll;
  padding-right: 20px;
  /* width: 200px; */
}

.content-detection-center-btn-right div span {
  text-align: left;
}

/*滚动条里面轨道*/
.content-detection-center-btn-right div::-webkit-scrollbar-track {
}

/*定义滚动条整体的样式*/
.content-detection-center-btn-right div::-webkit-scrollbar {
  width: 6px;
  border-radius: 15px;
}

/*滚动条的样式*/
.content-detection-center-btn-right div::-webkit-scrollbar-thumb {
  height: 20px;
  background-image: -webkit-gradient(linear,
  left bottom,
  left top,
  color-stop(0.2, rgb(125, 126, 128)),
  color-stop(0.4, rgb(97, 98, 99)),
  color-stop(0.8, rgb(125, 126, 128)));
}

.back {
  display: flex;
  justify-content: flex-start;
  margin-bottom: 10px;
}
</style>