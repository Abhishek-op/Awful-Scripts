require 'socket'
require 'thread'
require 'timeout'
require 'resolv'

IP = ARGV[0]
host = Resolv.new.getaddress(IP)
puts "[[=]]Portscanner is started"
puts "[[=]]Scanning for "+host
def portscanner(host)
(0..1024).each do |port|
Thread.new {
begin
Timeout.timeout(0.25) do # timeout of running operation
s = TCPSocket.new(host, port) # Create new socket
puts "[+] #{host} | Port #{port} open"
s.close
end
rescue Errno::ECONNREFUSED
next
rescue Timeout::Error
next
end
}.join
end
end

portscanner host