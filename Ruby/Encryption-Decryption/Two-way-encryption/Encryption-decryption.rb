require "base64"

class String
def self.encrypt(input_word)
    Base64.encode64(input_word)
  end

  def self.decrypt(input_word)
    Base64.decode64(input_word)
  end

puts "Select 1 for encription, 2 for decryption" 
main = gets.chomp    
if main == "1"
puts "Enter the text"
string = gets.chomp
encrypted = encrypt(string)
puts "here is your encrypted text"
puts encrypted
elsif main == "2"
puts "Enter encrypcted text"
encrypted = gets.chomp
puts "Here is your original string"
puts decrypt(encrypted)
else
  puts "Select a valid option"
end 
end