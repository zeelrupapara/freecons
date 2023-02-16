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
      <!-- ========================================================================= -->
      <div class="row">
        <div class="col col-lg-8">
          <stats-card> </stats-card>
        </div>
        <div class="col col-lg-1 text-right">
          <stats-card>
            <div slot="content">
              <p class="card-category">Select date for view data of that day</p>
            </div>
          </stats-card>
        </div>
      </div>
      <div class="row">
        <div class="col-md-8">
          <chart-card
            :chart-data="lineChart.data"
            :chart-options="lineChart.options"
            :responsive-options="lineChart.responsiveOptions"
          >
            <template slot="header">
              <h4 class="card-title">Inserted Icons Category</h4>
              <p class="card-category">Total Inserted Icons Per 5 Minutes</p>
            </template>
            <template slot="footer">
              <div class="legend">
                <i class="fa fa-circle text-info"></i> Total
                <i class="fa fa-circle text-success"></i> Activated
                <i class="fa fa-circle text-danger"></i> Errors
              </div>
              <hr />
              <div class="stats">
                <i class="fa fa-history"></i> Updated 3 minutes ago
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
              <div class="legend">
                <i class="fa fa-circle text-info"></i> Total
                <i class="fa fa-circle text-success"></i> Activated
                <i class="fa fa-circle text-danger"></i> Errors
              </div>
              <hr />
              <div class="stats">
                <i class="fa fa-clock-o"></i> Campaign sent 2 days ago
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

export default {
  components: {
    LTable,
    ChartCard,
    StatsCard,
  },
  data() {
    return {
      connection: null,
      countsDetails: {},
      editTooltip: "Edit Task",
      deleteTooltip: "Remove",
      pieChart: {
        data: {
          labels: ["total", "active", "error"],
          series: [
            countsData.total_icons,
            countsData.total_active_icons,
            countsData.total_error_icons,
          ],
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
  },
  mounted() {
    console.log("Started conecting to websocket");
    this.connection = new WebSocket(
      "ws://localhost:8080/api/v1/dashboard/counts"
    );
    this.connection.onopen = (event) => {
      console.log(event);
      console.log("Connected to websocket");
    };
    this.connection.onmessage = (event) => {
      this.countsDetails = JSON.parse(event.data);
      console.log(this.countsDetails);
    };
  },
};
</script>
<style></style>
