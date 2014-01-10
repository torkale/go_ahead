var amqp = require('amqp');
var http = require('http');

var amqp = require('amqp');

var connection = amqp.createConnection();

var connection = amqp.createConnection({url:"amqp://guest:guest@localhost:5672"});
// Wait for connection to become established.
connection.on('ready', function () {
  // Use the default 'amq.topic' exchange
  connection.queue('location', function(q){
      // Catch all messages
      q.bind('#');
  });
});
console.log("Starting server on port 8080")

var app = function(request,res){
  var body = '';
  request.on('data', function (data) {
    body += data;
  });
  request.on('end', function () {
    var user = JSON.parse(body);
    connection.publish(user.entry[0].changed_fields[0], JSON.stringify(user));
    res.writeHeader(200, {"Content-Type": "application/json"});
    res.end();
  });
};
http.createServer(app).listen(8080);

