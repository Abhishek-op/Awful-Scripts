import java.util.Scanner;
import java.util.Random;
import java.util.Arrays;
public class Main {

	public static void main(String[] args) {
		System.out.println(gameTurn());
	}
	public static String gameTurn() {
		Scanner scan = new Scanner(System.in);
		int guess = 10;
		boolean game = true;
		String[] randomWord = {"goodluck","guess","word","computer"};	
		String wordToGuess = randomWordArray(randomWord);
		String wordToPrint = wordToStarString(wordToGuess);
		String result = "";
    System.out.println("\t ¦¦¦¦¦¦  ¦¦    ¦¦ ¦¦¦¦¦¦¦ ¦¦¦¦¦¦¦ ¦¦¦¦¦¦¦       ¦¦     ¦¦  ¦¦¦¦¦¦  ¦¦¦¦¦¦  ¦¦¦¦¦¦   ");
    System.out.println("\t¦¦       ¦¦    ¦¦ ¦¦      ¦¦      ¦¦            ¦¦     ¦¦ ¦¦    ¦¦ ¦¦   ¦¦ ¦¦   ¦¦ ");
    System.out.println("\t¦¦   ¦¦¦ ¦¦    ¦¦ ¦¦¦¦¦   ¦¦¦¦¦¦¦ ¦¦¦¦¦¦¦ ¦¦¦¦¦ ¦¦  ¦  ¦¦ ¦¦    ¦¦ ¦¦¦¦¦¦  ¦¦   ¦¦ ");
    System.out.println("\t¦¦    ¦¦ ¦¦    ¦¦ ¦¦           ¦¦      ¦¦       ¦¦ ¦¦¦ ¦¦ ¦¦    ¦¦ ¦¦   ¦¦ ¦¦   ¦¦ ");
    System.out.println("\t ¦¦¦¦¦¦   ¦¦¦¦¦¦  ¦¦¦¦¦¦¦ ¦¦¦¦¦¦¦ ¦¦¦¦¦¦¦        ¦¦¦ ¦¦¦   ¦¦¦¦¦¦  ¦¦   ¦¦ ¦¦¦¦¦¦  ");
    System.out.println("      ");

		System.out.println("\nWelcome player to a game of Guess!");
		System.out.println("\nYou have 10 tries to guess the right word");
		System.out.println("Good luck!");
		
		while(game) {
			System.out.println("\nWrong guesses left: " + guess);		
			System.out.println("Your word to guess is: " + wordToPrint);
			System.out.print("Enter you guess: ");
			
			String guessTry = scan.next();
			guessTry.toLowerCase();
			if (guessTry.equals(wordToGuess)) {
				System.out.println("You guessed the whole word!");
				System.out.println("Well-done!!");
				game = false;
			}
			else {
				char charGuess = guessTry.charAt(0);
			
				if(Character.isLetter(charGuess)) {
				
					int[] indexCharInWord = indexCharInWord(wordToGuess, charGuess);
					if(indexCharInWord.length != 0) {
						for(int i = 0; i < indexCharInWord.length;i++) {
							wordToPrint = starToCharInWord(wordToGuess, wordToPrint, indexCharInWord[i]);	
						}	
					}
					else {
						System.out.println("Wrong guess, try again");
						guess--;
					}
				}
				else {
					System.out.println("Please try again and use letters only!");
				}
				if(!checkForStar(wordToPrint)) {
					result = "You have saved the stickman!";
					game = false;
				}
				if(guess == 0) {
					result = "Ow no you failed!!";
					game = false;
				}
			}
		}	
		scan.close();		
		return result;
	}
	public static String randomWordArray (String[] stringArray) {
		Random r = new Random();
		String result = stringArray[r.nextInt(stringArray.length)];
		return result;
	}
	public static String wordToStarString (String selectedWord) {
		String result = "";
		for(int i = 0; i < selectedWord.length(); i++) {		
			result = result + "*";
		}	
		return result;
	}
	public static int[] indexCharInWord (String str, char a) {
		String whereChar = "";	
		for(int i = 0; i < str.length(); i++) {
			if(str.charAt(i) == a) {
				whereChar = whereChar + i;
			}
		}
		int[] index = stringToIntArray(whereChar);
		return index;
	}
	public static int[] stringToIntArray(String str) {
		int[] index = new int[str.length()];
		for(int i = 0; i < index.length; i++) {
			index[i] = Character.getNumericValue(str.charAt(i));
		}
		return index;
	}
	public static String starToCharInWord (String guessStr, String printStr, int index) {
		String toChange = new String(printStr);
		String changeWith = new String(guessStr);
		
		char[] changeWith1 = changeWith.toCharArray();
		char[] toChange1 = toChange.toCharArray();
		
		toChange1[index] = changeWith1[index];
		String result = new String(toChange1);
		
		return result;
	}
	public static boolean checkForStar(String wordToCheck) {
		boolean starsLeft = true;
		int[] no = new int[0];
		if(Arrays.equals(indexCharInWord(wordToCheck,'*'),no)) {
			starsLeft = false;
		}
		return starsLeft;
	}
	
}