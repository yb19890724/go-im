<template>
    <div>

    </div>
</template>

<script>
    export default {
        name: "HeadView"
    }
</script>

<style scoped>
    .flex-container{
        display:flex;
        flex-direction:row;
        width:100%;
        padding-top: 10px;
        position: fixed;
        bottom: 0px;
        background-color: #FFFFFF;
    }
    .item-1{
        height:50px;
        height:50px;
        padding: 5px 5px 5px 5px;
    }
    .item-2{
        margin-right:auto;
        height:50px;
        width: 100%;
    }
    .txt{
        margin-right:auto;
    }
    .item-3{
        height:50px;
        height:50px;
        padding: 5px 5px 5px 5px;
    }
    .item-4{
        height:50px;
        height:50px;
        padding: 5px 5px 5px 5px;
    }

    li.chat{
        justify-content: flex-start;
        align-items: flex-start;
        display: flex;

    }
    .chat.other{
        flex-direction: row;
    }
    .chat.mine{
        flex-direction: row-reverse;
    }
    img.avatar{
        width: 64px;
        height:64px;
    }
    .other .avatar{
        margin-left:10px;
    }
    .mine .avatar{
        margin-right:10px;
    }
    .other span{
        border: 10px solid;
        border-color: transparent #FFFFFF transparent transparent ;
        margin-top: 10px;
    }
    .mine span{
        border: 10px solid;
        border-color: transparent  transparent transparent #32CD32;
        margin-top: 10px;
    }
    .other>.content{
        background-color: #FFFFFF;

    }
    .mine>.content{
        background-color: #32CD32;

    }
    div.content{
        min-width: 60px;
        clear: both;
        display: inline-block;
        padding: 16px 16px 16px 10px;
        margin: 0 0 20px 0;
        font: 16px/20px 'Noto Sans', sans-serif;
        border-radius: 10px;

        min-height: 64px;
    }
    .content>img.pic{
        width: 100%;
        margin:3px 3px 3px 3px;
    }
    .content>img.audio{
        width: 32px;
        color: white;
    }
    #panels{
        background-color: #FFFFFF;
        display: flex;
        position: fixed;
        bottom: 50px;
    }
    .doutures{
        flex-direction: row;
        flex-wrap: wrap;
        display: flex;
    }
    .doutures img{
        margin: 10px 10px 10px 10px;
    }
    .doutupkg{
        flex-direction: row;
        flex-wrap: wrap;
        display: flex;
    }
    .plugins{
        flex-direction: row;
        flex-wrap: wrap;
        display: flex;
    }
    .plugin{
        padding: 10px 10px 10px 20px;
        margin-left: 10px;
        margin-right: 10px;
    }
    .plugin img{
        width: 40px;
    }
    .plugin p{
        text-align: center;
        font-size: 16px;
    }
    .doutupkg img{
        width: 32px;
        height: 32px;
        margin: 5px 5px 5px 5px;
    }
    .upload{
        width: 64px;
        height: 64px;
        position: absolute;
        top: 1px;
        opacity:0;
    }
    .tagicon{
        width: 32px;
        height:32px;
    }

    .small{
        width: 32px;
        height:32px;
    }
    .middle{
        width: 64px;
        height:64px;
    }
    .large{
        width: 96px;
        height:96px;
    }
    .res image{
        width: 32px;
        height:32px;
    }
</style>