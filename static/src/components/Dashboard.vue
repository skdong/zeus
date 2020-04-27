<template>
  <div class="container">
    <div class="header ">
      <img :src="img_src" class="logo" />
      <div class="header-right">
         <el-row>
            <el-col :span="16">
               <el-date-picker
                v-model="timerange"
                type="datetimerange"
                range-separator="至"
                start-placeholder="开始日期"
                end-placeholder="结束日期">
              </el-date-picker> 
            </el-col>
            <el-col :span="2" style="font-size:32px;cursor:pointer">
                <i class="el-icon-download" title="导出数据" @click="exportData()"></i>
            </el-col>
            <el-col :span="6" style="text-align:right;padding-right:10px">
                <span>{{currentDate}}</span>
                <span>{{this.currentTime}}</span>
            </el-col>
        </el-row>
        
      </div>
    </div>

    <div class="content">
      <div class="windspeed" id="wind-speed-gauge"></div>

      <div class="winddirection" id="wind-direction-gauge"></div>

      <div class="windaverage">
          <div class="windspeedaverage">
             <label class="title">平均风速</label>
             <span>{{averageSpeed}}</span>
             <span>m/s</span>
          </div>

          <div class="winddirectionaverage">
             <label class="title">平均风向</label>
             <span>{{averageDirection}}</span>
             <span>°</span>
          </div>

      </div>

      <div class="runtimewindspeed" id="runtime-wind-speed"></div>
      <div class="runtimewinddirection" id="runtime-wind-direction"></div>

      <div class="windtable" id="windtable">
        <label style="padding:5px;font-weight:bold;font-size:18px;color:#333">最新1000条数据</label>
        <el-table
            :data="runtimeData"
            stripe
            height="400"
            style="width: 1020px">
            <el-table-column prop="time" label="时间"  width="340"></el-table-column>
            <el-table-column prop="speed" label="风速" width="340"></el-table-column>
            <el-table-column prop="direction" label="风向" width="340"></el-table-column>
          </el-table>
      </div>
      
    </div>
    
    
  </div>
</template>

<script>

import Vue from 'vue'
import echarts from 'echarts'
import XLSX from 'xlsx'

export default {
  name: 'Dashboard',
  data () {
    return {
      img_src:require("../assets/image/logo.png"),
      currentDate:"2020/04/23",
      currentTime:"星期六 9:20:25",
      averageSpeed:0,
      averageDirection:0,
      timerange:"",
      currentSpeed:"",
      currentDirection:"",
      runtimeData:[],
      runtimeSpeedData:[],
      runtimeDirectionData:[],
      queryUrl:"/winds",
      socketUrl:"ws://101.132.38.140:8080/ws/winds"
    }
  },
  created:function(){
    this.initWebsocket();
  },
  mounted:function(){
    this.initTime(); 

    this.drawWindSpeedGuage();
    this.drawWindDirectionGuage();
    this.drawRuntimeWindSpeed();
    this.drawRuntimeWindDirection();

    this.queryData();
    
    
  },
  methods:{
  
    testWebsocket:function(){

      this.currentSpeed = Math.random()*60;
      this.currentDirection = Math.random()*360;
      //改变仪表盘的值
      this.speedGuageOption.series[0].data[0].value = this.currentSpeed;
      this.speedGuageChart.setOption(this.speedGuageOption, true);

      this.directionGuageOption.series[0].data[0].value = this.currentDirection;
      this.directionGuageChart.setOption(this.directionGuageOption, true);
      

      //计算平均值
      this.runtimeData.push({"time":new Date().toString(),"speed":this.currentSpeed,"direction":this.currentDirection});

      //this.$set(runtimeData,runtimeData)
      var speedSum = 0;
      var directionSum = 0;
      for(var i in this.runtimeData){
      
         speedSum += this.runtimeData[i].speed;
         directionSum += this.runtimeData[i].direction;
      }

      this.averageSpeed = speedSum/this.runtimeData.length;
      this.averageDirection = directionSum/this.runtimeData.length;

      this.averageSpeed = Math.round(this.averageSpeed*100)/100;
      this.averageDirection = Math.round(this.averageDirection*100)/100;

      //实时数据
      if(this.runtimeSpeedData.length>1000){
        this.runtimeSpeedData.shift();
      }
      this.runtimeSpeedData.push({"name":new Date(),value:[new Date(),this.currentSpeed]});
      this.runtimeSpeedChart.setOption({
                series: [{
                    data: this.runtimeSpeedData
                }]
            });

      if(this.runtimeDirectionData.length>1000){
        this.runtimeDirectionData.shift();
      }
      this.runtimeDirectionData.push({"name":new Date(),value:[new Date(),this.currentDirection]});
      this.runtimeDirectionChart.setOption({
                series: [{
                    data: this.runtimeDirectionData
                }]
            });

    },
    initTime:function(){
      setInterval(this.getTime,500);
    },
    getTime:function(){
      var aData = new Date();
      var month =aData.getMonth() < 9 ? "0" + (aData.getMonth() + 1) : aData.getMonth() + 1;
      var date = aData.getDate() <= 9 ? "0" + aData.getDate() : aData.getDate();
      this.currentDate = aData.getFullYear() + "/" + month + "/" + date;

      let hh = aData.getHours();
      let mf = aData.getMinutes()<10 ? '0'+aData.getMinutes() : aData.getMinutes();
      let ss = aData.getSeconds()<10 ? '0'+aData.getSeconds() : aData.getSeconds();

      var week = "星期" + "日一二三四五六".charAt(aData.getDay());
      this.currentTime = week+' '+hh+':'+mf+':'+ss;
      //console.log(this.currentTime)

    },
    queryData:function(){

      //setInterval(this.testWebsocket,1000);

    },
    initWebsocket:function(){

      const wsuri = this.socketUrl;
      this.websock = new WebSocket(wsuri);
      this.websock.onmessage = this.websocketonmessage;
      this.websock.onopen = this.websocketonopen;
      this.websock.onerror = this.websocketonerror;
      this.websock.onclose = this.websocketclose;
    },
    websocketonmessage:function(e){
      const redata = JSON.parse(e.data);
      this.currentSpeed = redata.Speed;
      this.currentDirection = redata.Direction;

      //解析时间
      var time = redata.CreateAt;
      time = time.replace(/-/g,'/');
      time = time.replace(/T/g,' ');
      time = time.split(".")[0];
      
      this.runtimeData.push({"time":time,"speed":this.currentSpeed,"direction":this.currentDirection});

      //改变仪表盘的值
      this.speedGuageOption.series[0].data[0].value = this.currentSpeed;
      this.speedGuageChart.setOption(this.speedGuageOption, true);

      this.directionGuageOption.series[0].data[0].value = this.currentDirection;
      this.directionGuageChart.setOption(this.speedGuageOption, true);

      //计算平均值
      var speedSum = 0;
      var directionSum = 0;
      for(var i in this.runtimeData){
         speedSum += this.runtimeData[i].speed;
         directionSum += this.runtimeData[i].direction;
      }
      this.averageSpeed = speedSum/this.runtimeData.length;
      this.averageDirection = directionSum/this.runtimeData.length;

      this.averageSpeed = Math.round(this.averageSpeed*100)/100;
      this.averageDirection = Math.round(this.averageDirection*100)/100;

      //实时数据
      if(this.runtimeSpeedData.length>1000){
        this.runtimeSpeedData.shift();
      }
      this.runtimeSpeedData.push({"name":time,value:[time,this.currentSpeed]});
      this.runtimeSpeedChart.setOption({
                series: [{
                    data: this.runtimeSpeedData
                }]
            });

      if(this.runtimeDirectionData.length>1000){
        this.runtimeDirectionData.shift();
      }
      this.runtimeDirectionData.push({"name":time,value:[time,this.currentDirection]});
      this.runtimeDirectionChart.setOption({
                series: [{
                    data: this.runtimeDirectionData
                }]
            });

    },
    websocketonopen:function(){
      console.log("websocket link is open");
      this.websocket.send(JSON.stringify({"devices": ["Q"], "limit": 50, "enable": true, "close": true}));
    },
    websocketonerror:function(){
      this.initWebsocket();
    },
    websocketclose:function(){
      console.log("websocket link is closed");
    },
    drawWindSpeedGuage:function(){

      let myChart = echarts.init(document.getElementById('wind-speed-gauge'));

     
      var option = {
          tooltip: {
              formatter: '{a} : {b}m/s'
          },
          toolbox: {
              feature: {
                  
                  saveAsImage: {}
              }
          },
          series: [
              {
                  name: '风速',
                  type: 'gauge',
                  z: 3,
                  min: 0,
                  max: 60,
                  splitNumber: 12,
                  radius: '100%',
                  axisLine: {            // 坐标轴线
                      lineStyle: {       // 属性lineStyle控制线条样式
                          width: 10
                      }
                  },
                  axisTick: {            // 坐标轴小标记
                      length: 15,        // 属性length控制线长
                      lineStyle: {       // 属性lineStyle控制线条样式
                          color: 'auto'
                      }
                  },
                  splitLine: {           // 分隔线
                      length: 20,         // 属性length控制线长
                      lineStyle: {       // 属性lineStyle（详见lineStyle）控制线条样式
                          color: 'auto'
                      }
                  },
                  axisLabel: {
                      backgroundColor: 'auto',
                      borderRadius: 2,
                      color: '#eee',
                      padding: 3,
                      textShadowBlur: 2,
                      textShadowOffsetX: 1,
                      textShadowOffsetY: 1,
                      textShadowColor: '#222'
                  },
                  title: {
                      // 其余属性默认使用全局文本样式，详见TEXTSTYLE
                      fontWeight: 'bolder',
                      fontSize: 20,
                      fontStyle: 'italic'
                  },
                  detail: {
                      // 其余属性默认使用全局文本样式，详见TEXTSTYLE
                      formatter: function (value) {
                          value = (value + '').split('.');
                          value.length < 2 && (value.push('00'));
                          return ('00' + value[0]).slice(-2)
                              + '.' + (value[1] + '00').slice(0, 2);
                      },
                      fontWeight: 'bolder',
                      borderRadius: 3,
                      backgroundColor: '#444',
                      borderColor: '#aaa',
                      shadowBlur: 5,
                      shadowColor: '#333',
                      shadowOffsetX: 0,
                      shadowOffsetY: 3,
                      borderWidth: 2,
                      textBorderColor: '#000',
                      textBorderWidth: 2,
                      textShadowBlur: 2,
                      textShadowColor: '#fff',
                      textShadowOffsetX: 0,
                      textShadowOffsetY: 0,
                      fontFamily: 'Arial',
                      width: 100,
                      color: '#eee',
                      rich: {}
                  },
                  data: [{value: 40, name: '风速（m/s）'}]
              }
          ]
      };
      option.series[0].data[0].value = this.currentSpeed;
      myChart.setOption(option, true);
      
      this.speedGuageChart = myChart;
      this.speedGuageOption = option;
      /*setInterval(function () {
          option.series[0].data[0].value = (Math.random() * 100).toFixed(2) - 0;
          myChart.setOption(option, true);
      },2000);*/

    },
    
    drawWindDirectionGuage:function(){

      let myChart = echarts.init(document.getElementById('wind-direction-gauge'));

     
      var option = {
          tooltip: {
              formatter: '{a} : {b}°'
          },
          toolbox: {
              feature: {
                  
                  saveAsImage: {}
              }
          },
          series: [
              {
                  name: '风向',
                  type: 'gauge',
                  z: 3,
                  min: 0,
                  max: 360,
                  splitNumber: 12,
                  radius: '100%',
                  axisLine: {            // 坐标轴线
                      lineStyle: {       // 属性lineStyle控制线条样式
                          width: 10
                      }
                  },
                  axisTick: {            // 坐标轴小标记
                      length: 15,        // 属性length控制线长
                      lineStyle: {       // 属性lineStyle控制线条样式
                          color: 'auto'
                      }
                  },
                  splitLine: {           // 分隔线
                      length: 20,         // 属性length控制线长
                      lineStyle: {       // 属性lineStyle（详见lineStyle）控制线条样式
                          color: 'auto'
                      }
                  },
                  axisLabel: {
                      backgroundColor: 'auto',
                      borderRadius: 2,
                      color: '#eee',
                      padding: 3,
                      textShadowBlur: 2,
                      textShadowOffsetX: 1,
                      textShadowOffsetY: 1,
                      textShadowColor: '#222'
                  },
                  title: {
                      // 其余属性默认使用全局文本样式，详见TEXTSTYLE
                      fontWeight: 'bolder',
                      fontSize: 20,
                      fontStyle: 'italic'
                  },
                  detail: {
                      // 其余属性默认使用全局文本样式，详见TEXTSTYLE
                      formatter: function (value) {
                          value = (value + '').split('.');
                          value.length < 2 && (value.push('00'));
                          return ('00' + value[0]).slice(-2)
                              + '.' + (value[1] + '00').slice(0, 2);
                      },
                      fontWeight: 'bolder',
                      borderRadius: 3,
                      backgroundColor: '#444',
                      borderColor: '#aaa',
                      shadowBlur: 5,
                      shadowColor: '#333',
                      shadowOffsetX: 0,
                      shadowOffsetY: 3,
                      borderWidth: 2,
                      textBorderColor: '#000',
                      textBorderWidth: 2,
                      textShadowBlur: 2,
                      textShadowColor: '#fff',
                      textShadowOffsetX: 0,
                      textShadowOffsetY: 0,
                      fontFamily: 'Arial',
                      width: 100,
                      color: '#eee',
                      rich: {}
                  },
                  data: [{value: 40, name: '风向（°）'}]
              }
          ]
      };
      option.series[0].data[0].value = this.currentDirection;
      myChart.setOption(option, true);

      this.directionGuageChart = myChart;
      this.directionGuageOption = option;

      /*setInterval(function () {
          option.series[0].data[0].value = (Math.random() * 100).toFixed(2) - 0;
          myChart.setOption(option, true);
      },2000);*/

    },
    drawRuntimeWindSpeed:function(){
    
        let myChart = echarts.init(document.getElementById('runtime-wind-speed'));

        var option = {
            title: {
                text: '实时风速数据'
            },
            tooltip: {
                trigger: 'axis',
                formatter: function (params) {
                    params = params[0];
                    var date = new Date(params.name);
                    return date.getHours() + ':' + (date.getMinutes() + 1) + ':' + date.getSeconds() + ' : ' + params.value[1];
                    
                },
                axisPointer: {
                    animation: false
                }
            },
            xAxis: {
                type: 'time',
                splitLine: {
                    show: false
                }
            },
            yAxis: {
                type: 'value',
                boundaryGap: [0, '100%'],
                splitLine: {
                    show: false
                },
                min:0,
                max:60
            },
            series: [{
                name: '风速数据',
                type: 'line',
                showSymbol: false,
                hoverAnimation: false,
                data: this.runtimeSpeedData
            }]
        };
        myChart.setOption(option);
       
        this.runtimeSpeedChart = myChart;
        this.runtimeSpeedData = [];

        
    },
    randomData:function(){
         this.now = new Date(+this.now + 60 * 1000);
         var now = this.now;
        var value =  Math.random() * 60;
        var valueName = now.getFullYear() + '/' + (now.getMonth() + 1) + '/' + now.getDate() + ' ' + (now.getHours() >= 10 ? now.getHours() : '0' + now.getHours()) + ':' + (now.getMinutes() >= 10 ? now.getMinutes() : '0' + now.getMinutes());
        
        return {
            name: now.toString(),
            value: [ 
                valueName,
                Math.round(value)
            ]
        }



        
    },
    drawRuntimeWindDirection:function(){
        let myChart = echarts.init(document.getElementById('runtime-wind-direction'));
        var option = {
            title: {
                text: '实时风向数据'
            },
            tooltip: {
                trigger: 'axis',
                formatter: function (params) {
                    params = params[0];
                    var date = new Date(params.name);
                    return date.getHours() + ':' + (date.getMinutes() + 1) + ':' + date.getSeconds() + ' : ' + params.value[1];
                    
                },
                axisPointer: {
                    animation: false
                }
            },
            xAxis: {
                type: 'time',
                splitLine: {
                    show: false
                }
            },
            yAxis: {
                type: 'value',
                boundaryGap: [0, '100%'],
                splitLine: {
                    show: false
                },
                min:0,
                max:360
            },
            series: [{
                name: '风向数据',
                type: 'line',
                showSymbol: false,
                hoverAnimation: false,
                data: this.runtimeDirectionData
            }]
        };
        myChart.setOption(option);
       
        this.runtimeDirectionChart = myChart;
        this.runtimeDirectionData = [];

        
        
    },
    exportData:function(){
    console.log(111);
      if(this.selectedTime()){
         this.queryExportData(); 

      }else{
        this.$message('请先选择时间范围！');
      }
    },
    selectedTime:function(){
        var range = this.timerange;
        if(Array.isArray(range)&& range.length == 2)
            return true;
        return false;
    },
    queryExportData:function(){
       var start = this.timerange[0];
       var end = this.timerange[1];
       start = this.getFormatTime(start);
       end = this.getFormatTime(end);
       
       var _this = this;
       _this.$axios.get(this.queryUrl,{
        params:{"start_at":start,"end_at":end}
       }).then(function(response){

          _this.exportFun(response.data);
       })
    },
    exportFun:function(data){
      /* 创建worksheet */
      var ws = XLSX.utils.json_to_sheet(data.winds);
console.log(1234);
      /* 新建空workbook，然后加入worksheet */
      var wb = XLSX.utils.book_new();
      XLSX.utils.book_append_sheet(wb, ws, "data");

      /* 生成xlsx文件 */
      XLSX.writeFile(wb, "风速风向数据.xlsx");
    },
    getFormatTime:function(aData){
      var month =aData.getMonth() < 9 ? "0" + (aData.getMonth() + 1) : aData.getMonth() + 1;
      var date = aData.getDate() <= 9 ? "0" + aData.getDate() : aData.getDate()

      let hh = aData.getHours();
      let mf = aData.getMinutes()<10 ? '0'+aData.getMinutes() : aData.getMinutes();
      let ss = aData.getSeconds()<10 ? '0'+aData.getSeconds() : aData.getSeconds();

      
      return aData.getFullYear() + "-" + month + "-" + date+' '+hh+':'+mf+':'+ss;
    }
  }
}
</script>

<!-- Add "scoped" attribute to limit CSS to this component only -->
<style scoped>
.container {
  width:100%;
  height:100%;
  background-color:#FFFAF0;
  overflow:auto;
}
.header{
    width: 100%;
    height: 49px;
    background-color: rgba(15,66,95,.9);
    position: relative;
    z-index: 999;
}
.clearfloat{
  display:block;
  clear:both;
  content:"";
  visibility:hidden;
  height:0;
}
.logo{
  width: auto;
  height: 44px;
  padding: 2.5px 12px;
}
.header-right{
  width:50%;
  float:right;
  color:#eee;
  font-size:16px;
  padding:5px;
  
}
.header-right span{
  display:block;
}

.content{
  width:1040px;
  margin:0 auto;
  background-color:white;
  padding:0px 10px;
  overflow:auto;
}
.content div{
  height:350px;
  margin:5px;
  background-color:#fafafa;
  float:left;
}
.windspeed, .winddirection{
  width:400px;
  text-align:center;

}
.windspeed div, .winddirection div{
  margin:0 auto;
}

.content div.windaverage{
  background-color:white;
  width:200px;
  height:350px;
}
.windspeedaverage,.winddirectionaverage{
   width:100%;
   background-color:#fafafa;
   text-align:center;
   vertical-align:middle;
   position:relative;
}
.windspeedaverage label,.winddirectionaverage label{
  position:absolute;
  left:10px;
  top:20px;
  font-size:12px;

}
.windspeedaverage span,.winddirectionaverage span{
  font-size:28px;
  line-height:170px;
  fong-weight:bold;
  font-style:italic;
}
.content .windspeedaverage {
  height:170px;
  
  margin-bottom:10px;
  margin-top:0px;
}
.content .winddirectionaverage{
  height:170px;
  
}
.runtimewindspeed,.runtimewinddirection{
  width:508px;
  height:350px;
}

.content div.windtable{
  overflow:scroll;
}



</style>
