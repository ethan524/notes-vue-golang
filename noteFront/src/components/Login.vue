<template>
    <div class="login">
        <div class="login_box">
            <div class="">
                <el-menu :default-active="activeIndex" class="el-menu-demo form_title" mode="horizontal"  router>
                    <el-menu-item index="1" route="/login">登录</el-menu-item>
                    <el-menu-item index="2" route="/register">注册</el-menu-item>
                  </el-menu>
            </div>
            <el-form ref="loginFormRef" :model="loginForm" :rules="loginFormRules"  label-width="0px" class="login_form">
                <el-form-item prop="username">
                    <el-input v-model="loginForm.username" prefix-icon="el-icon-user-solid" placeholder="请输入用户名" ></el-input>
                </el-form-item>
                <el-form-item prop="password">
                    <el-input v-model="loginForm.password" prefix-icon="el-icon-s-goods" placeholder="请输入密码" type="password" ></el-input>
                </el-form-item>
                <el-form-item class="regBtn">
                    <el-button :center="true" type="primary" @click="saveLoginForm" class="regBtn">登录</el-button>
                    <el-button :center="true" type="info" @click="resetForm" class="regBtn">重置</el-button>
                </el-form-item>
            </el-form>
        </div>
    </div>
</template>
<script>
export default {
    name:"Register",
    data(){
        return {
            activeIndex: '1',
            loginForm : {
                username:"",
                password:"",
                repassoword:"",
            },
            loginFormRules : {
                //验证用户名
                username : [
                    {required:true, message:'请输入用户名', trigger:'blur'},
                    {min:2, max:20, message:'长度在2~20个字符', trigger:'blur'}
                ],
                //验证密码
                password : [
                    {required:true, message:'请输入密码', trigger:'blur'}
                ],
                repassword : [
                    {required:true, message:'请输入确认密码', trigger:'blur'}
                ],
            }
        }
    },
    methods :{
        saveLoginForm(){

        },
        resetForm(){

        }
    }
}
</script>

<style lang="less" scoped>
    .login{
        height: 100%;
        background-color: #2b4b6b;
    }
    .form_title{
        border-radius: 6px;
    }
    .login_box{
        border-radius: 1px solid red ;
        width:450px;
        height:300px;
        background-color: #fff;
        border-radius: 6px;
        position:absolute;
        left:50%;
        top:50%;
        transform: translate(-50%, -50%);
    }
    .login_form{
        position: absolute;
        width:100%;
        padding:0 20px;
        box-sizing: border-box;
        left:50%;
        top:60%;
        transform: translate(-50%, -50%);
    }
    .regBtn{
        text-align: center;
    }
</style>