<template>
  <div class="content">
    <div class="container-fluid">
      <div class="row">
        <!-- TODO: total count -->
        <div class="col-xl-4 col-md-4">
          <stats-card>
            <div slot="header" class="icon-warning">
              <i class="nc-icon nc-chart text-warning"></i>
            </div>
            <div slot="content">
              <p class="card-category">Total Icons</p>
              <h4 class="card-title">{{ countsData.total_icons }}</h4>
            </div>
            <div slot="footer">
              <i class="fa fa-refresh"></i>{{ countsData.total_icons_time }}
            </div>
          </stats-card>
        </div>

        <!-- TODO : total activeted icons 200 -->
        <div class="col-xl-4 col-md-4">
          <stats-card>
            <div slot="header" class="icon-success">
              <i class="nc-icon nc-light-3 text-success"></i>
            </div>
            <div slot="content">
              <p class="card-category">Total Activated Icons</p>
              <h4 class="card-title">{{ countsData.total_active_icons }}</h4>
            </div>
            <div slot="footer">
              <i class="fa fa-calendar-o"></i
              >{{ countsData.total_active_icons_time }}
            </div>
          </stats-card>
        </div>

        <!-- TODO: total errors icons table count -->
        <div class="col-xl-4 col-md-4">
          <stats-card>
            <div slot="header" class="icon-danger">
              <i class="nc-icon nc-vector text-danger"></i>
            </div>
            <div slot="content">
              <p class="card-category">Total Errors Icons</p>
              <h4 class="card-title">{{ countsData.total_error_icons }}</h4>
            </div>
            <div slot="footer">
              <i class="fa fa-clock-o"></i
              >{{ countsData.total_error_icons_time }}
            </div>
          </stats-card>
        </div>
      </div>

      <div class="row">
        <div class="col-md-8">
          <chart-card
            :chart-data="lineChartData"
            :chart-options="lineChart.options"
            :responsive-options="lineChart.responsiveOptions"
          >
            <template slot="header">
              <h4 class="card-title mb-4">Icons Analitics</h4>
              <div class="mb-4">
                <select class="custom-select" @change="getDataFromThatDay">
                  <option v-for="date in dates" :key="date" :value="date">
                    {{ date }}
                  </option>
                </select>
              </div>
            </template>
            <template slot="footer">
              <div class="legend">
                <i class="fa fa-circle text-info"></i> Total
                <i class="fa fa-circle text-danger"></i> Activated
                <i class="fa fa-circle text-warning"></i> Errors
              </div>
            </template>
          </chart-card>
        </div>

        <div class="col-md-4">
          <chart-card :chart-data="pieChart.data" chart-type="Pie">
            <template slot="header">
              <h4 class="card-title">Icons Statistics</h4>
              <p class="card-category">Inserted Icons Analitics</p>
            </template>
            <template slot="footer">
              <div class="pb-5">
                <div class="pb-5"></div>
              </div>
              <div class="legend">
                <i class="fa fa-circle text-info"></i> Total
                <i class="fa fa-circle text-danger"></i> Activated
                <i class="fa fa-circle text-warning"></i> Errors
              </div>
            </template>
          </chart-card>
        </div>
      </div>
    </div>
  </div>
</template>
<script>
import ChartCard from "src/components/Cards/ChartCard.vue";
import StatsCard from "src/components/Cards/StatsCard.vue";
import LTable from "src/components/Table.vue";
import axios from "axios";
export default {
  components: {
    LTable,
    ChartCard,
    StatsCard,
  },
  data() {
    return {
      connection: null,
      dates: [],
      countsDetails: {},
      editTooltip: "Edit Task",
      deleteTooltip: "Remove",
      pieChart: {
        data: {
          labels: [],
          series: [],
        },
      },
      lineChart: {
        data: {
          labels: [],
          series: [[], [], []],
        },
        options: {
          showArea: false,
          axisX: {
            showGrid: false,
          },
          lineSmooth: true,
          showLine: true,
          showPoint: true,
          fullWidth: true,
          chartPadding: {
            right: 100,
          },
        },
        responsiveOptions: [
          [
            "screen and (max-width: 640px)",
            {
              axisX: {
                labelInterpolationFnc(value) {
                  return value[0];
                },
              },
            },
          ],
        ],
      },
    };
  },
  computed: {
    countsData() {
      return this.countsDetails;
    },
    lineChartData() {
      console.log("hello");
      return this.lineChart.data;
    },
  },
  async mounted() {
    await this.getPieChatDate();
    await this.getLineChart();
    await this.getCounts();
  },
  methods: {
    getPieChatDate() {
      axios.get("/api/v1/dashboard/piechart").then((response) => {
        this.pieChart.data.series = response.data.data;
        this.pieChart.data.labels = response.data.labels;
      });
    },
    getLineChart() {
      axios.get("/api/v1/dashboard/linechart").then((response) => {
        this.lineChart.data.series = response.data.data;
        this.lineChart.data.labels = response.data.labels;
        this.dates = response.data.dates;
      });
    },
    getCounts() {
      axios.get("/api/v1/dashboard/counts").then((response) => {
        this.countsDetails = response.data;
      });
    },
    getDataFromThatDay(event) {
      axios
        .get("/api/v1/dashboard/linechart?date=" + event.target.value)
        .then((response) => {
          this.lineChart.data.series = response.data.data;
          this.lineChart.data.labels = response.data.labels;
        });
    },
  },
};
</script>
<style></style>
