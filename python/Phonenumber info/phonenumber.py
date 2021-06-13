
import phonenumbers
from phonenumbers import carrier
from phonenumbers import geocoder
from phonenumbers import timezone
def main():
	try:
		a = input('[_]Enter phone number with country code:')
		d=phonenumbers.parse(a)
		time = str(timezone.time_zones_for_number(d))
		Carrier = carrier.name_for_number(d, 'en')
		Region = geocoder.description_for_number(d, 'en')
		valid = str(phonenumbers.is_possible_number(d))
		print('')
		print('[-]is valid:'+valid)
		print('[-]Carrier:'+Carrier)
		print('[-]Region:'+Region)
		print('[-]Timezone:'+time)
	except phonenumbers.phonenumberutil.NumberParseException:
		print('[×]Enter a valid number')
		main()
if __name__ == '__main__':
    main()     