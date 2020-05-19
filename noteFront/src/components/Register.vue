<template>
    <div class="regiser">
        <div class="error_box" v-if="showError"><el-alert title="提示" description="请填写用户名称" type="error" show-icon class="errorBox"></el-alert></div>
        <div class="regiser_box">
            <div class="">
                <el-menu :default-active="activeIndex" class="el-menu-demo form_title" mode="horizontal"  router>
                    <el-menu-item index="1" route="/login">登录</el-menu-item>
                    <el-menu-item index="2" route="/register">注册</el-menu-item>
                  </el-menu>
            </div>
            <el-form ref="registerFormref" :model="registerForm" :rules="registerFormRules"  label-width="0px" class="register_form">
                <el-form-item prop="email">
                    <el-input v-model="registerForm.email" prefix-icon="el-icon-user-solid" placeholder="请输入邮箱" type="email" ></el-input>
                </el-form-item>
                <el-form-item prop="password">
                    <el-input v-model="registerForm.password" prefix-icon="el-icon-s-goods" placeholder="请输入密码" type="password" ></el-input>
                </el-form-item>
                <el-form-item prop="repassword">
                    <el-input v-model="registerForm.repassword" prefix-icon="el-icon-s-goods" placeholder="请输入确认密码" type="password" ></el-input>
                </el-form-item>
                <el-form-item class="regBtn">
                    <el-button :center="true" type="primary" @click="saveRegisterForm('registerFormref')" class="regBtn">注册</el-button>
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
        var validatePass = (rule, value, callback) => {
            if (value === '') {
                callback(new Error('请输入密码'));
            } else {
                if (this.registerForm.repassoword !== '') {
                    this.$refs.registerForm.validateField('repassoword');
                }
                callback();
            }
        };
        var validatePass2 = (rule, value, callback) => {
            if (value === '') {
                callback(new Error('请再次输入密码'));
            } else if (value !== this.registerForm.password) {
                callback(new Error('两次输入密码不一致!'));
            } else {
                callback();
            }
        };
        return {
            showError:false,
            activeIndex: '2',
            registerForm : {
                email:"",
                password:"",
                repassoword:"",
            },
            registerFormRules : {
                //验证邮箱
                email : [
                    {required:true, message:'请输入邮箱', trigger:'blur'},
                ],
                //验证密码
                password : [
                    {validator: validatePass, trigger: 'blur'}
                ],
                repassword : [
                    {validator: validatePass2, trigger: 'blur'}
                ],
            },
        }
    },
    methods :{
        saveRegisterForm(formname){
            var _this = this
            this.$refs[formname].validate((valid) => {
                if(valid){
                    this.$axios({
                        method:'post',
                        url:'/register',
                        responseType:'Json',
                        data:{email:_this.registerForm.email,password:_this.registerForm.password},
                    }).then(function(response){
                        if( response.data.code != '200'){
                            _this.$message(response.data.message);
                        }else{
                            _this.$message(({
                                dangerouslyUseHTMLString:true,
                                message:'<el-link type="success">注册成功,即将跳转至登录页面。</el-link>',
                                type: 'success',
                                duration:3000,
                                showClose:true,
                                offset:200,
                                onClose:function(){
                                    _this.$router.push({path: '/login'});
                                }
                            }));
                        }
                    }).catch(function(error){
                        // console.log('error')
                    });
                }
            });
            return false;
        },
        resetForm(){

        },
    }
}
</script>

<style lang="less" scoped>
    .error_box{
        width:450px;
        height:60px;
        position:absolute;
        left:50%;
        top:30%;
        transform: translate(-50%, -50%);
    }
    .regiser{
        height: 100%;
        background-color: #2b4b6b;
    }
    .form_title{
        border-radius: 6px;
    }
    .regiser_box{
        border-radius: 1px solid red ;
        width:450px;
        height:330px;
        background-color: #fff;
        border-radius: 6px;
        position:absolute;
        left:50%;
        top:50%;
        transform: translate(-50%, -50%);
    }
    .register_form{
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