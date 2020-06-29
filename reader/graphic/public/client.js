const container = document.querySelector("#container");
let ctx;
const width = 300;
const height = 300;

const radToDeg = (radian) => (radian * 180) / Math.pi;

const convertCoords = (distance, angle) => {
  const cartesianX = distance * Math.cos(angle);
  const cartesianY = distance * Math.sin(angle);

  const canvasX = width / 2 + cartesianX;
  const canvasY = height / 2 + cartesianY;

  return { x: canvasX, y: canvasY };
};

const initCanvas = () => {
  const canvas = document.createElement("canvas");
  canvas.width = width;
  canvas.height = height;
  container.appendChild(canvas);
  ctx = canvas.getContext("2d");
};

const receiveRecord = (data) => {
  console.log(data);
  const [distance, angle] = data.split(" ");
  const { x, y } = convertCoords(Number(distance), Number(angle));
  console.log(x, y);

  ctx.fillRect(x, y, 1, 1);
};

let socket;
try {
  socket = new WebSocket("ws://localhost:8080/ws");
} catch (err) {
  console.error(err);
}

socket.onerror = console.error;

socket.onopen = function () {
  console.log("Connected");
  this.onclose = function () {
    console.log("Closed");
  };
  this.onmessage = function (event) {
    receiveRecord(event.data);
  };
};

initCanvas();
