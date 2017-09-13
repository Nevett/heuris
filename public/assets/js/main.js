var ws = new WebSocket("ws://" + document.location.host);
ws.onmessage = function(msg) {
    var channel = JSON.parse(msg.data);

    var tableBody = document.querySelector('tbody');
    var row = tableBody.querySelector("." + channel.name);

    if (row == null) {
        row = createRow(channel);
    } else {
        modifyRow(row, channel);
    }

    tableBody.insertBefore(row, tableBody.firstChild);
};

function createRow(channel) {
    var row = document.createElement('tr');
    row.className = channel.name;

    var nameCell = document.createElement('td')
    nameCell.className = 'name';
    nameCell.innerHTML = channel.name;
    row.appendChild(nameCell);

    var clientsCell = document.createElement('td')
    clientsCell.className = 'num-subscribers';
    clientsCell.innerHTML = channel.numSubscribers;
    row.appendChild(clientsCell);

    var messagesPublishedCell = document.createElement('td')
    messagesPublishedCell.className = 'num-messages-published';
    messagesPublishedCell.innerHTML = channel.numMessagesPublished;
    row.appendChild(messagesPublishedCell);

    return row;
}

function modifyRow(row, channel) {
    var clientsCell = row.querySelector(".num-subscribers");
    var messagesPublishedCell = row.querySelector(".num-messages-published");

    clientsCell.innerHTML = channel.numSubscribers;
    messagesPublishedCell.innerHTML = channel.numMessagesPublished;
}
