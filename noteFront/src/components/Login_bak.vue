<template>
    <div class='login_container'>
        <div class="login_box">
            <!-- 头像区域 -->
            <div class="avatar_box">
                <img src="../assets/logo.png" alt="">
            </div>
            <!-- 登录表单区域 -->
            <el-form ref="loginFormRef" :model="loginForm" :rules="loginFormRules" label-width="0px" class="login_form">
                <el-form-item prop="username">
                    <el-input v-model="loginForm.username" prefix-icon="el-icon-user-solid" placeholder="请输入用户名" ></el-input>
                </el-form-item>
                <el-form-item prop="password">
                    <el-input v-model="loginForm.password" prefix-icon="el-icon-s-goods" placeholder="请输入密码" type="password" ></el-input>
                </el-form-item>
                <el-form-item >
                    <el-button type="primary" @click="loginForm">登录</el-button>
                    <el-button type="info" @click="resetLoginForm">重置</el-button>
                </el-form-item>
            </el-form>
        </div>
    </div>
</template>

<script>
export default {
    name:'Login',
    data(){
        return {
            //这是登录表单的数据绑定对象
            loginForm:{
                username : '',
                password : '',
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
            }
        }
    },
    methods :{
        resetLoginForm(){
            console.log(this)
        },
        loginForm(){
            this.$axios({
                method:"POST",
            })
        },
    }
}
</script>

<style scoped lang="less">
    .login_container{
        background-color: #2b4b6b;
        height:100%;
    }
    .login_box{
        width:450px;
        height:300px;
        background-color: #fff;
        border-radius: 3px;
        position:absolute;
        left:50%;
        top:50%;
        transform: translate(-50%, -50%);

        .avatar_box{
            height: 130px;
            width:130px;
            border:1px solid #eee;
            border-radius: 50%;
            padding:10px;
            box-shadow: 0 0 10px #ddd;
            position: absolute;
            left:50%;
            transform: translate(-50%, -50%);
            background-color: #fff;
            img{
                width: 100%;
                height: 100%;
                border-radius: 50%;
                background-color: #eee;
            }
        }
    }

    .login_form{
        position: absolute;
        bottom: 0;
        width:100%;
        padding:0 20px;
        box-sizing: border-box;
    }
</style>