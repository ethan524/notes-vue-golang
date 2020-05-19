<template>
    <el-container>
          <el-header class="header_contaniner">
            <el-row>
              <el-col class="header_box" :span="14" :push="5">
                <div class="grid-content bg-purple">
                  <el-menu :default-active="activeIndex" class="el-menu-demo" mode="horizontal" @select="handleSelect" router>
                    <el-menu-item index="1" route="/index"><i class="el-icon-s-home"></i>主页</el-menu-item>
                    <el-menu-item index="2" route="/follow"><i class="el-icon-s-ticket"></i>关注</el-menu-item>
                    <el-submenu index="3">
                      <template slot="title"><i class="el-icon-message-solid"></i>消息</template>
                      <el-menu-item index="3-1" route="/comment"><i class="el-icon-s-comment"></i>评论</el-menu-item>
                      <el-menu-item index="3-2" route="/likes" ><i class="el-icon-star-on"></i>喜欢和赞</el-menu-item>
                    </el-submenu>
                    <el-menu-item index="3" disabled>
                      <el-input v-model="header_search_input" placeholder="请输入内容" suffix-icon="el-icon-search"></el-input>
                    </el-menu-item>
                    <el-menu-item index="5" style="float:right;" route="/riegister">注册</el-menu-item>
                    <el-menu-item index="4" style="float:right;" route="/login">登录</el-menu-item>
                  </el-menu>
                  <div class="line"></div>
                </div>
              </el-col>
            </el-row>
          </el-header>
          <el-main>
            <router-view/>
          </el-main>
      </el-container>
</template>
<script>
export default {
  name: 'Home',
  data() {
    return {
      activeIndex: '1',
      header_search_input:'',
    };
  },
  methods: {
    handleSelect(key, keyPath) {
      console.log(key, keyPath);
    }
  }
}
</script>

<style lang="less" scoped>
html,body{
  height:100%;
  margin:0;
  padding:0;
  font-family: "Helvetica Neue",Helvetica,"PingFang SC","Hiragino Sans GB","Microsoft YaHei","微软雅黑",Arial,sans-serif;
}
.header_contaniner{
  height:55px;
  margin:0;
  padding:0;
}
.header_box{
  border-bottom: 1px solid gray;
  text-align: center;
  height:55px;
  margin:0 auto;
  line-height: 55px;
}
.header_search_inputs{
  border-radius: 20px;
}
.el-input__inner{
  border-radius: 20px;
  border:1px solid rgb(78,78,78);
}
</style>