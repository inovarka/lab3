const request = require("./request");

const Client = (baseUrl) => ({
  get: (path) => request(baseUrl + path),
  patch: (path, data) => request(baseUrl + path, "PATCH", data),
});

module.exports = { Client };