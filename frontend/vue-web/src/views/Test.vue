<template>
  <div class="test">
    
    <el-row>
      <el-col :span="6"  v-for="(data, index) in ydata" :key="data"  :offset="1" >
      <el-card :body-style="{ padding: '0px' }">
      <img v-bind:src=data class="image">
      <div style="padding:14px;">
        <span>{{xdata[index]}}</span>
        <div class="bottom clearfix">
          <time class="time">{{ currentDate }}</time>
          <el-button type="text" class="button">操作按鈕</el-button>
        </div>
      </div>
      </el-card>
      </el-col>
      
    </el-row>
    
  </div>
</template>

<script>
import axios from 'axios'
 export default {
  name: 'test',
  data () {
    return {
      msg: 'message ftom test',
      site: "localhost",
      url: "http://localhost:8081/",
      xdata: null,
      ydata: null,
    }
  },
  methods: {
    details: function() {
      return this.site
    },
  },
  computed: {
    getImageUrl(data){
      return "\"\"" + data + "\"\""
    },
  },
  mounted () {
      axios.get('http://localhost:8000/v1/line').then(response => (this.xdata = response.data.legend_data,this.ydata = response.data.xAxis_data))

  }
 }
</script>

<style lang="scss" scoped>
.test {
  color: red;
}

.time {
    font-size:13px;
    color:#999;
}
.bottom {
    margin-top:13px;
    line-height:12px;
}
.button {
    padding:0;
    float: right;
}
.image {
    width:100%;
    display: block;
}
.clearfix:before,
.clearfix:after {
      display: table;
      content:"";
}
.clearfix:after {
      clear: both
}
</style>