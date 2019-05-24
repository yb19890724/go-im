<template>
    <div>
        <h1>im登录</h1>
        <div class="mui-content" id="pageapp">
            <form id='login-form' class="mui-input-group">
                <div class="mui-input-row">
                    <label>账号</label>
                    <input v-model="user.username" placeholder="请输入用户名" type="text" class="mui-input-clear mui-input" >
                </div>
                <div class="mui-input-row">
                    <label>密码</label>
                    <input v-model="user.password" placeholder="请输入密码"  type="password" class="mui-input-clear mui-input" >
                </div>
            </form>
            <div class="mui-content-padded">
                <button @click="login"  type="button"  class="mui-btn mui-btn-block mui-btn-primary">登录</button>
                <div class="link-area"><a id='reg' href="register.shtml">注册账号</a> <span class="spliter">|</span> <a  id='forgetPassword'>忘记密码</a>
                </div>
            </div>
            <div class="mui-content-padded oauth-area">
            </div>
        </div>
    </div>
</template>

<script>
    export default {
        data:function(){
            return {
                user:{
                    username:"",
                    password:""
                }
            }
        },
        methods:{
            login:function(){
                this.$http.post("http://localhost:9090/login",JSON.stringify(this.user)).then((response) => {
                    if(response.status==200 && response.data.data!=""){
                        alert( response.data.msg)
                        localStorage.setItem("currentUserToken",response.data.data.token);
                        this.$router.push({name:"home"})
                        return false;
                    }

                });
            },
        }

    }
</script>

<style scoped>
    @import "../../assets/css/mui.min.css";

    .area {
        margin: 20px auto 0px auto;
    }

    .mui-input-group {
        margin-top: 10px;
    }

    .mui-input-group:first-child {
        margin-top: 20px;
    }

    .mui-input-group label {
        width: 22%;
    }

    .mui-input-row label~input,
    .mui-input-row label~select,
    .mui-input-row label~textarea {
        width: 78%;
    }

    .mui-checkbox input[type=checkbox],
    .mui-radio input[type=radio] {
        top: 6px;
    }

    .mui-content-padded {
        margin-top: 25px;
    }

    .mui-btn {
        padding: 10px;
    }

    .link-area {
        display: block;
        margin-top: 25px;
        text-align: center;
    }

    .spliter {
        color: #bbb;
        padding: 0px 8px;
    }

    .oauth-area {
        position: absolute;
        bottom: 20px;
        left: 0px;
        text-align: center;
        width: 100%;
        padding: 0px;
        margin: 0px;
    }

    .oauth-area .oauth-btn {
        display: inline-block;
        width: 50px;
        height: 50px;
        background-size: 30px 30px;
        background-position: center center;
        background-repeat: no-repeat;
        margin: 0px 20px;
        /*-webkit-filter: grayscale(100%); */
        border: solid 1px #ddd;
        border-radius: 25px;
    }

    .oauth-area .oauth-btn:active {
        border: solid 1px #aaa;
    }

    .oauth-area .oauth-btn.disabled {
        background-color: #ddd;
    }
</style>