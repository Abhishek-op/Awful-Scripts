import java.net.*;  
import java.io.*;  
//main class
class client{  
public static void main(String args[])throws Exception{  //creating socket at localhost port 3333
Socket s=new Socket("localhost",3333);  
System.out.println("[+]Connected to server at localhost:3333");
System.out.println("[-]Type your message: "); 
DataInputStream din=new DataInputStream(s.getInputStream());  
DataOutputStream dout=new DataOutputStream(s.getOutputStream());  
BufferedReader br=new BufferedReader(new InputStreamReader(System.in));  
String str="",str2="";  
while(!str.equals("stop")){  
str=br.readLine();  
dout.writeUTF(str);  
dout.flush();  
str2=din.readUTF();  
System.out.println("[*]Server says: "+str2);
System.out.printf("[-]Type your message: ");   
}  
  
dout.close();  
s.close();  
}}  