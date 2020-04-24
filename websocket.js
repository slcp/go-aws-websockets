const WebSocket = require("ws");

var ws;
try {
  ws = new WebSocket(
    "APIGWWebsocketEndpoint"
  );
} catch (e) {
    console.log('helloe')
  console.log(e);
}

ws.on("open", () => console.log("connected"));
ws.on("message", (data) => console.log(`From server: ${data}`));
ws.on("close", () => {
  console.log("disconnected");
  process.exit();
});
