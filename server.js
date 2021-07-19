const WebSocket = require("ws");
const redis = require("redis");

const REDIS_SERVER = "redis://localhost:6379";
const WS_PORT = 3000;

// connection to redis

let redisClient = redis.createClient(REDIS_SERVER);

// subscribe client to the channel

redisClient.subscribe("app:notification");

const server = new WebSocket.Server({ port : WS_PORT });

server.on("connection", function(ws){
    redisClient.on("message", function(channel, msg) {
        console.log(msg);
        ws.send(msg);
    })
});