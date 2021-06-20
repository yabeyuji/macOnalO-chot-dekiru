"use strict"

var app = new Vue({
  el: "#app",
  components: {
    "monitor": httpVueLoader("web/vue/monitor.vue"),
  }, // --- End of components --- //

  data: {
    reserves: [],
    assembles: [],
    completes: [],
    passes: [],

    commonClass: "btn p-3 border add-button-style",
  }, // --- End of data --- //

  created: function () {
    this.wsBase = "ws://localhost:4567/ws";
  }, // --- End of created --- //


  computed: {
  }, // --- End of computed --- //

  methods: {


  }, // --- End of methods --- //


  mounted: function () {
    let _this = this

    console.log("## mounted()");
    this.ws = new WebSocket(this.wsBase);
    this.ws.onopen = function (event) {
      _this.isOnline = true;
      console.log("### websocket.onopen()");
    };

    this.ws.onmessage = function (event) {
      console.log(event);

      const eventData = JSON.parse(event.data);
      _this.reserves = eventData.Reserves;
      _this.assembles = eventData.Assembles;
      _this.completes = eventData.Completes;
      _this.passes = eventData.Passes;

    };

    // websocketでエラーが発生した時
    this.ws.onerror = function (event) {
      console.log("### websocket.onerror()");
    };

    // websocketをクローズした時
    this.ws.onclose = function (event) {
      console.log("### websocket.onclose()");
      _this.isOnline = false;
      _this.timer = setInterval(function () {
        axios
          .get("")
          .then(function (response) {
            window.location.reload();
          })
          .catch(function (error) {
            console.log(error)
          })
      }, 1000);
    }
  }

})
