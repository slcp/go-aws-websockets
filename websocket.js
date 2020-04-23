// const io = require("socket.io-client");

// const socket = io("wss://tok1tuxpge.execute-api.eu-west-2.amazonaws.com/dev");
// console.log('after')
// socket.on("connect_error", (data) => {
//   console.log('connect error', data);
// });
// socket.on("connect", (data) => {
//   console.log('connect', data);
// });
const WebSocket = require("ws");

// try {
  const ws = new WebSocket(
    "wss://tok1tuxpge.execute-api.eu-west-2.amazonaws.com/dev"
  );
// } catch (e) {
//     console.log('helloe')
//   console.log(e);
// }

ws.on("open", () => console.log("connected"));
ws.on("message", (data) => console.log(`From server: ${data}`));
ws.on("close", () => {
  console.log("disconnected");
  process.exit();
});

// ws.on("open", function open() {
//   ws.send("something");
// });

// ws.on("message", function incoming(data) {
//   console.log(data);
// });
