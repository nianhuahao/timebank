<template>
  <div class="notice">
    <el-tabs
        v-model="activeName"
        type="card"
        @tab-click="handleClick"

    >
      <el-tab-pane label="最新需求" name="first">
        <el-table
            :data="NeedData"
            stripe="true"
            max-height="226"
            style="width: 100%;font-size:13px">
          <el-table-column
              prop="title"
              label="需求标题"
              width="180">
          </el-table-column>
          <el-table-column
              prop="address"
              label="服务地点"
              width="180">
          </el-table-column>
          <el-table-column
              prop="money"
              label="时间币">
          </el-table-column>
          <el-table-column
              prop="time"
              label="发布时间">
          </el-table-column>
          <el-table-column
              prop="state"
              label="状态">
          </el-table-column>
        </el-table>
      </el-tab-pane>
      <el-tab-pane label="最新活动" name="second">
        <el-table
            :data="activityData"
            stripe="true"
            style="width: 100%;font-size:13px"
            max-height="226"
        >
          <el-table-column
              prop="title"
              label="活动标题"
              width="220">
          </el-table-column>
          <el-table-column
              prop="time"
              label="活动时间"
              width="180">
          </el-table-column>
          <el-table-column
              prop="count"
              label="所需人数">
          </el-table-column>
          <el-table-column
              prop="state"
              label="状态">
          </el-table-column>
        </el-table>
      </el-tab-pane>
      <el-tab-pane label="服务动态" name="third">
        <el-table
            :data="stateData"
            stripe="true"
            max-height="226"
            style="width: 100%;font-size:13px">
          <el-table-column
              prop="title"
              label="活动标题"
              width="220">
          </el-table-column>
          <el-table-column
              prop="time"
              label="活动时间"
              width="180">
          </el-table-column>
          <el-table-column
              prop="name"
              label="承接人">
          </el-table-column>
          <el-table-column
              prop="state"
              label="状态">
          </el-table-column>
        </el-table>
      </el-tab-pane>
    </el-tabs>
  </div>
</template>

<script>
export default {
  data() {
    return {
      activeName: 'first',
      NeedData: [
        {
          title:'时间银行志愿服务活动',
          address:'黄阁镇东里村联合围街34号',
          money: '12.0',
          time:'2022年4月',
          state:'发布中'
        },
        {
          title:'时间银行志愿服务活动',
          address:'黄阁镇东里村联合围街34号',
          money: '12.0',
          time:'2022年4月',
          state:'发布中'
        },
        {
          title:'时间银行志愿服务活动',
          address:'黄阁镇东里村联合围街34号',
          money: '12.0',
          time:'2022年4月',
          state:'发布中'
        },
        {
          title:'时间银行志愿服务活动',
          address:'黄阁镇东里村联合围街34号',
          money: '12.0',
          time:'2022年4月',
          state:'发布中'
        },
        {
          title:'时间银行志愿服务活动',
          address:'黄阁镇东里村联合围街34号',
          money: '12.0',
          time:'2022年4月',
          state:'发布中'
        }
      ],
      activityData:[
        {
          title:'“安全排查，青年助力”宣传活动',
          time: '2022年4月',
          state:'发布中',
          count:'20人'
        },
        {
          title:'“安全排查，青年助力”宣传活动',
          time: '2022年4月',
          state:'发布中',
          count:'20人'
        },
        {
          title:'“安全排查，青年助力”宣传活动',
          time: '2022年4月',
          state:'发布中',
          count:'20人'
        }
      ],
      stateData:[
        {
          title:'“安全排查，青年助力”宣传活动',
          time: '2022年4月',
          state:'发布中',
          name:'啊哈'
        },
        {
          title:'“安全排查，青年助力”宣传活动',
          time: '2022年4月',
          state:'发布中',
          name:'啊哈'
        },
        {
          title:'“安全排查，青年助力”宣传活动',
          time: '2022年4月',
          state:'发布中',
          name:'啊哈'
        },
        {
          title:'“安全排查，青年助力”宣传活动',
          time: '2022年4月',
          state:'发布中',
          name:'啊哈'
        }
      ]
    };
  },
  methods: {
    handleClick(tab, event) {
      console.log(tab, event);
    }
  }
};
</script>

<style >
.notice{
  margin-top: 10px;
  border: 1px solid #dddddd;
  width: 668px;
  height: 282px;
  color: #ec7f30;
}
.notice .el-tabs .el-tab-pane{
  background-color: #f9f9f9;
  color: #303033;
}
</style>
