<template>
  <div class="content">
    <div class="container-fluid">
      <div class="row">
        <div class="col-md-12">
          <card>
            <template slot="header">
              <h4 class="card-title">50 Awesome Free Icons</h4>
            </template>
            <div class="all-icons">
              <div class="row">
                <div class="col col-lg-1" v-for="icon in icons" :key="icon">
                  <div class="p-2">
                    <img
                      class="img-fluid"
                      :src="icon"
                      alt="icon-image"
                      height="50px"
                      width="50px"
                    />
                  </div>
                </div>
              </div>
            </div>
            <div class="d-flex justify-content-between">
              <div>
                <button
                  type="button"
                  class="btn btn-light"
                  @click="previosPage"
                >
                  <svg
                    xmlns="http://www.w3.org/2000/svg"
                    width="16"
                    height="16"
                    fill="currentColor"
                    class="bi bi-caret-left-fill"
                    viewBox="0 0 16 16"
                  >
                    <path
                      d="m3.86 8.753 5.482 4.796c.646.566 1.658.106 1.658-.753V3.204a1 1 0 0 0-1.659-.753l-5.48 4.796a1 1 0 0 0 0 1.506z"
                    />
                  </svg>
                </button>
              </div>
              <div>
                <button type="button" class="btn btn-light" @click="nextPage">
                  <svg
                    xmlns="http://www.w3.org/2000/svg"
                    width="16"
                    height="16"
                    fill="currentColor"
                    class="bi bi-caret-right-fill"
                    viewBox="0 0 16 16"
                  >
                    <path
                      d="m12.14 8.753-5.482 4.796c-.646.566-1.658.106-1.658-.753V3.204a1 1 0 0 1 1.659-.753l5.48 4.796a1 1 0 0 1 0 1.506z"
                    />
                  </svg>
                </button>
              </div>
            </div>
          </card>
        </div>
      </div>
    </div>
  </div>
</template>
<script>
import Card from "src/components/Cards/Card.vue";
import axios from "axios";
export default {
  components: {
    Card,
  },
  data() {
    return {
      page: 1,
      icons: [],
    };
  },
  mounted() {
    axios
      .get(`api/v1/icons/url?page=` + this.page)
      .then((res) => {
        this.icons = res.data;
      })
      .catch((err) => {
        console.log(err);
      });
  },
  methods: {
    nextPage() {
      this.page++;
      axios
        .get(`api/v1/icons/url?page=` + this.page)
        .then((res) => {
          this.icons = res.data;
        })
        .catch((err) => {
          console.log(err);
        });
    },
    previosPage() {
      if (this.page === 1) {
        return;
      }
      this.page--;
      axios
        .get(`api/v1/icons/url?page=` + this.page)
        .then((res) => {
          this.icons = res.data;
        })
        .catch((err) => {
          console.log(err);
        });
    },
  },
};
</script>
<style></style>
