require 'net/http'
url = ARGV[0]
loop do
    puts url
    res = Net::HTTP.get_response URI url
    if !res['location'].nil?
        url = res['location']
    else
        break
    end
end