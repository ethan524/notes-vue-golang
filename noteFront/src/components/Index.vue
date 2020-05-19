<template>
    <el-row>
        <el-col :span="14" :push="5">main</el-col>
    </el-row>
</template>
<script>
export default {
    name : "Index",
    methods : {
        demo(){
            this.$axios({
                method:'post',
                url:'http://localhost:8080/test',
                data:{
                    name:"ethan123",
                }
            })
            .then(function(response){
                console.log("response")
                for(var x in response.data){
                    console.log(x+'='+response.data[x])
                }
            })
            .catch(function(error){
                console.log("error")
                console.log(error)
            })
        }
    },
    created : function(){
        this.demo()
    }
}
</script>

<style lang="less" scoped>

</style>