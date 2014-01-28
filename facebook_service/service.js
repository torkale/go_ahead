var amqp = require('amqp');
var http = require('http');
var cluster = require('cluster');

var connection = amqp.createConnection({url:"amqp://guest:guest@localhost:5672"});
connection.on('ready', function () {
  connection.queue('location', {durable: true}, function(q){
      q.bind('location');
  });
});

var app = function(request,res){
  var body = '';

  request.on('data', function (data) {
    body += data;
  });

  request.on('end', function () {
    var user = JSON.parse(body);
    connection.publish(user.entry[0].changed_fields[0], JSON.stringify(user));
    res.writeHeader(200, {"Content-Type": "application/json"});
    res.write('{"success":true}')
    res.end();
  });
};

if (cluster.isMaster) {
  // Fork workers.
  for (var i = 0; i < 4; i++) {
    cluster.fork();
  }
} else {
  console.log("Starting server on port 8080")
  http.createServer(app).listen(8080);
}
