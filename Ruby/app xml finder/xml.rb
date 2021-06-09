require 'ruby_apk'

path = ARGV[0]
	
apk = Android::Apk.new(path)
manifest = apk.manifest
puts manifest.to_xml