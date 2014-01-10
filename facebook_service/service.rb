require "rubygems"
require 'sinatra'
require 'json'
require 'bunny'

set :port, 8080
puts "Starting server on port 8080"

connection = Bunny.new
connection.start
channel  = connection.create_channel
exchange = channel.topic("facebook_events", :durable => true)
channel.queue("location", :durable => true).bind(exchange, :routing_key => "location")

get '/event' do
  raw = request.body.read
  data = JSON.parse raw
  field = data["entry"][0]["changed_fields"][0]
  exchange.publish(raw, :routing_key => field)
  '{"success":true}'
end
