<template>
    <div>
        <nav class="mui-bar mui-bar-tab" v-if="showMenuBar" >
            <a class="mui-tab-item mui-active" href="#tabbar-concat">
                <span class="mui-icon mui-icon-chat"></span>
                <span class="mui-tab-label">好友</span>
            </a>
            <a class="mui-tab-item" href="#tabbar-group">
                <span class="mui-icon mui-icon-chatboxes"><span class="mui-badge">9</span></span>
                <span class="mui-tab-label">群聊</span>
            </a>
            <a class="mui-tab-item" href="#tabbar-profile">
                <span class="mui-icon mui-icon-person"></span>
                <span class="mui-tab-label">我的</span>
            </a>

        </nav>

    </div>
</template>

<script>
    export default {
        name: "MenuView",
        props:{
            showMenuBar:{
                type: Boolean,
                default(){
                    return true;
                }
            }
        },
    }
</script>

<style scoped>

</style>