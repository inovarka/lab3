const { Client } = require("./balancers/client");

const client = Client("http://localhost:8080");

(async () => {
  
  let balancers;
  console.log("=== Scenario 1 ===");
  try {
    balancers = await client.ListBalancers();
    console.log("Balancers:");
    console.table(balancers);
  } catch (err) {
    console.log(`Problem listing balancers: `, err);
  }


})();