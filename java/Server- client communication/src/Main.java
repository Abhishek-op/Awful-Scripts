import java.net.*;  
import java.io.*;  
// Server class
class Main{  
public static void main(String args[])throws Exception{ // creating a socket at port 3333
ServerSocket ss=new ServerSocket(3333);
System.out.println("[+]Server ai started at port 3333");
//accept connection
Socket s=ss.accept(); 
System.out.println("[-]Type your message: "); 
//Server-client 
DataInputStream din=new DataInputStream(s.getInputStream());  
DataOutputStream dout=new DataOutputStream(s.getOutputStream());  
BufferedReader br=new BufferedReader(new InputStreamReader(System.in));  

System.out.println("");
String str="",str2="";  

while(!str.equals("stop")){  
str=din.readUTF();  
System.out.println("[*]client says: "+str);
System.out.println("[-]Type your message: ");
str2=br.readLine();  
dout.writeUTF(str2);  
dout.flush();  
}  
din.close();  
s.close();  
ss.close();  
}}  