const WebSocket = require('ws')

if (process.argv.length < 3)
    return;

const ws = new WebSocket(process.argv[2]);

ws.on('message', function(data) {
    console.log(data);
});
