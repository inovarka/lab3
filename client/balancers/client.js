const http = require("../common/http");

const Client = (baseUrl) => (
  (client = http.Client(baseUrl)),
  {
    ListBalancers: () => client.get("/balancers"),
    UpdateMachine: (id, isWorking) =>
      client.patch(`/balancers`, { id, isWorking }),
  }
);

module.exports = { Client };