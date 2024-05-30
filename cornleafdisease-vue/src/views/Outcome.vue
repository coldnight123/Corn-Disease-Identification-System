<script lang="ts" setup>
import Header from '@/components/common/Header.vue'
const fits = ['fill', 'contain', 'cover', 'none', 'scale-down']
const url =
    'https://fuss10.elemecdn.com/e/5d/4a731a90594a4af544c0c25941171jpeg.jpeg'
</script>

<template>
  <div>
    <Header></Header>
  </div>
  <div class="content-wrapper">
    <div class="top-section">
      <el-text tag="b" size="large">
        检测结果
      </el-text>
    </div>
    <div class="bottom-section">
      <div class="left-part">
        <div class="demo-image">
          <div v-for="fit in fits" :key="fit" class="block">
            <span class="demonstration">{{ fit }}</span>
            <el-image style="width: 100px; height: 100px" :src="url" :fit="fit" />
          </div>
        </div>
      </div>
      <div class="right-part">
        <!-- 右边部分的内容 -->
      </div>
    </div>
  </div>
</template>

<style scoped>
body, html {
  margin: 0;
  padding: 0;
  height: 100%;
}
.content-wrapper {
  display: flex;
  flex-direction: column;
  height: 80vh; /* 剩余屏幕高度的80% */
}
.top-section {
  flex: 0 0 10%; /* 上面部分占剩余高度的20% */
  align-items: center;
  justify-content: center;
  display: flex;
}
.bottom-section {
  display: flex; /* 底部部分使用Flexbox布局 */
  flex: 1; /* 占据剩余高度 */
}
.left-part {
  flex: 0 0 60%; /* 左部分占底部宽度的60% */
  background-color: #ddd;
}
.right-part {
  flex: 0 0 40%; /* 右部分占底部宽度的40% */
  background-color: #bbb;
}

.demo-image .block {
  padding: 30px 0;
  text-align: center;
  border-right: solid 1px var(--el-border-color);
  display: inline-block;
  width: 20%;
  box-sizing: border-box;
  vertical-align: top;
}
.demo-image .block:last-child {
  border-right: none;
}
.demo-image .demonstration {
  display: block;
  color: var(--el-text-color-secondary);
  font-size: 14px;
  margin-bottom: 20px;
}

</style>