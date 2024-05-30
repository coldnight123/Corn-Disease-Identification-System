<template>
  <header class="header">
    <div class="header-left">
      <!-- 左边图标 -->
      <div class='content-logo'>
        <img src="@/assets/img/corn.png">
        <h1> 玉米叶病检测</h1>
      </div>
    </div>
    <div class="header-center">
      <el-menu :default-active="activeIndex" class="el-menu-demo" mode="horizontal" @select="handleSelect">
        <el-menu-item index="1" :class="[activeIndex == 1 ? 'active' : '']">首&nbsp;&nbsp;&nbsp;&nbsp;页</el-menu-item>
        <el-menu-item index="2" :class="[activeIndex == 2 ? 'active' : '']">病害百科</el-menu-item>
        <el-menu-item index="3" :class="[activeIndex == 3 ? 'active' : '']">病害检测</el-menu-item>
        <el-menu-item index="4" :class="[activeIndex == 4 ? 'active' : '']">检测记录</el-menu-item>
      </el-menu>
    </div>
    <div class="header-right">
      你好，{{username}}
      <!-- 右边用户图标和下拉菜单 -->
      <el-button style="margin-left: 15px;" type="primary" @click="out">登出</el-button>
    </div>
  </header>
</template>

<script setup>
  import { ref,defineEmits,onMounted,watch } from 'vue'
  import { useRouter } from "vue-router";
  import { defineProps } from 'vue';
  
  const props = defineProps({
    nume: String
  })

  const router = useRouter();

  const activeIndex = ref('1')
  const emitEvents  = defineEmits(["updateNum"]);
  const handleSelect = e => {
    activeIndex.value = e
    emitEvents('updateNum',e)
  }

  watch(()=>props.nume,(newValue,oldValue)=>{
    activeIndex.value = newValue
  })

  const out = () => {
    localStorage.removeItem('loginResult')
    router.push({path: '/'})
  }

  onMounted( () => {
  });
  
  const username = localStorage.getItem('loginResult') ? JSON.parse(localStorage.getItem('loginResult'))['user_name'] : ''
</script>

<style scoped>
.header {
  display: flex;
  align-items: center;
  background-color: #f0f2f5;
  width:100%;
  height: 20%;
}
.header-left {
  flex: 30%;
}

.content-logo{
  height: 55px;
  margin: 10px 0;
  cursor: pointer;
}

.content-logo img {
  height: 100%;
  float: left;
}

.header-center {
  flex: 70%;
  text-align: center;
  justify-content: space-around;
}


.header-right {
  flex: 20%;
  margin: 10px 40px 10px 0 ;
  text-align: right;
}
.active{
  background: #ecf5ff;
}

</style>