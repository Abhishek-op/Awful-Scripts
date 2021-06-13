import re
pas = input("Enter password: ")
def check(password):
  num_check = re.search(r"\d", password) is None
  lowercase_check = re.search(r"[a-z]", password) is None
  length_check = len(password) < 8
  uppercase_check = re.search(r"[A-Z]", password) is None
  symbol_check = re.search(r"[ !#$%&'()*+,-./[\\\]^_`{|}~" + r'"]', password) is None
  strong = not (length_check or num_check or uppercase_check or lowercase_check or symbol_check)
  return {
        'Password is Strong': strong,
        'Password is short': length_check,
        'Password does not have number': num_check,
        'Password does not use uppercase alphabets': uppercase_check,
        'Password does not use lowecase alphabets': lowercase_check,
        'Password does not contain special character': symbol_check,
    }


dict = check(pas)

for statement, condition in dict.items():
    if condition:
        print(statement)
