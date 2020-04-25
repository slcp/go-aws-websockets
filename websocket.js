const WebSocket = require("ws");

var ws;
try {
  ws = new WebSocket(process.env.WEBSOCKET_ENDPOINT);
} catch (e) {
  console.log("helloe");
  console.log(e);
}

ws.on("open", () => {
  ws.send(
    JSON.stringify({
      action: "customaction",
      data: {
        hello: "world",
      },
    })
  );
  console.log("connected");
});
ws.on("message", (data) => console.log(`From server: ${data}`));
ws.on("close", () => {
  console.log("disconnected");
  process.exit();
});
