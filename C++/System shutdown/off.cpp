#include <stdio.h>
#include<stdlib.h>
#include <iostream>
int shutdown()
{
	system("c:\\windows\\system32\\shutdown /i");
	return 0;
}
int restart()
{
	system("c:\\windows\\system32\\shutdown /r");
	return 0;
}
int logoff()
{
	system("c:\\windows\\system32\\shutdown /l");
	return 0;
}
int main()
{
  std::cout << "welcome to awfulscripts ";
  std::cout << "Press 1 to shutdown ";
  std::cout << "Press 2 to restart ";
  std::cout << "Press 3 to log off ";
  std::cout << "Choose:";
  int a;
  std::cin >> a;
  if( a == 1 ) {
      shutdown();
   } else if( a == 2 ) {
     restart();
      
     
   } else if( a == 3 ) {
      logoff();
     
   } else {
     std::cout << "Choose correct option ";
   }
}