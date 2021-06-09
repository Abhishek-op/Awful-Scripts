#made by Abhishek-op
#One way encryption

# take text
puts "Enter the text"
string = gets.chomp

#take key
puts "Enter the Key"
key = gets.chomp

#encryption
encrypted = key.crypt(string)

#output
puts "here is your encrypted text"
puts encrypted
