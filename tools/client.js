const WebSocket = require('ws')

const channel = process.argv[2] || 'something';

const ws = new WebSocket(`ws://localhost:8080/${channel}`);

ws.on('message', function incoming(data) {
    console.log(data);
});
