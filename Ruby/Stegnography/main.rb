#mADE BY aBHISHEK-OP
class Main
puts "Welcome to Steganography script"
puts "Select 1 for Encoding, 2 for Decoding" 
main = gets.chomp    
if main == "1"
file1, file2 = ARGV
puts "Enter full name with path of image"
img = gets.chomp
puts "Enter full name with path of pdf"
pdf= gets.chomp
sec_file = File.read pdf
nor_file = File.read img
sep = '*------------------------*'
one_file = [nor_file, sep, sec_file]
File.open("output.png", 'wb') do |stg|
one_file.each do |f|
stg.puts f
puts "Output image created"
end
end
elsif main == "2"
puts "Enter path of encoded image"
img = gets.chomp
sep = '*------------------------*'
recov_file = File.read(img).force_encoding("BINARY").split(sep).last
File.open('output.pdf', 'wb') {|file| file.print recov_file}
puts "Output pdf created"
else
  puts "Select a valid option"
end 
end
