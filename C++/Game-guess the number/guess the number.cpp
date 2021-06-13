#include <iostream>

using namespace std;


int main()
{
    int secretNum = 7;
    int guess;
    int guesscount = 0;
    int guesslimit = 3;
    bool outofguess = false;

    while(secretNum != guess && !outofguess){
        if (guesscount < guesslimit){
        cout << "Enter your guess: ";
        cin >> guess;
        guesscount++;
        } else{
            outofguess = true;
        }
    }

    if (outofguess){
        cout << "Oh! you loose";

    } else {
        cout << "Hurrey! you win";
    }

    return 0;
}