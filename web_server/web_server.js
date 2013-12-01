var sys = require("sys"), http = require("http");
console.log("Starting server on port 8080")

http.createServer(function(request,response){
    response.writeHeader(200, {"Content-Type": "text/plain"});
    response.write("Request" + request.url);
    response.end();
}).listen(8080);
