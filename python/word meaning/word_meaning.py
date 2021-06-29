import requests
import sys


def fetch_meaning(user_defined_word):
    url = "https://api.dictionaryapi.dev/api/v2/entries/en/'
    response = requests.get(url+user_defined_word)

    if response.status_code in range(200, 299):
        meanings_response = response.json()[0]['meanings']
        for meanings in meanings_response:
            for meaning in meanings['definitions']:
                print(f'Part of Speech: {meanings["partOfSpeech"]}')
                print(f'Meaning: {meaning["definition"]} \n')
    elif response.status_code in range(400, 499):
        sys.exit("Please check your internet or the word you typed")
    elif response.status_code in range(500, 599):
        sys.exit("Oops a server side error has occured!")


if __name__ == "__main__":
    word = input("Please enter a word: ")
    fetch_meaning(word)
