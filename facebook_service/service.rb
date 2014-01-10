require "rubygems"
require 'sinatra'
require 'json'
require 'bunny'
PORT = 8080
puts "Starting server on port #{PORT}"
connection = Bunny.new
connection.start

channel  = connection.create_channel
# topic exchange name can be any string
exchange = channel.topic("facebook_events", :durable => true)
channel.queue("location", :durable => true).bind(exchange, :routing_key => "location")

get '/event' do
  raw = request.body.read
  data = JSON.parse raw
  data["entry"].each do |entry| 
    entry["changed_fields"].each do |field|
      exchange.publish(raw, :routing_key => field)
    end
  end
  return '{"success":true}'
end
