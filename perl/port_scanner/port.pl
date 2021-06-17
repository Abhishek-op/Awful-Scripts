#mADE bY aBHISHEK-OP
use IO::Socket;

$| = 1;

$target = $ARGV[0];

$start_port = $ARGV[1];
$end_port = $ARGV[2];
print <<"EOF";
  _                        
 |_) _  ._ _|_             
 |  (_) |   |_             
                           
  __                       
 (_   _  _. ._  ._   _  ._ 
 __) (_ (_| | | | | (/_ |  
                           
     mADE bY aBHISHEK-OP

EOF
foreach ($port = $start_port ; $port <= $end_port ; $port++) 
{
	print "\r[.]Scanning port $port";
	$socket = IO::Socket::INET->new(PeerAddr => $target , PeerPort => $port , Proto => 'tcp' , Timeout => 0.25);
	if( $socket )
	{
		print "\r[+]Port $port is open.\n" ;
	}
	else
	{
		#close
	}
}

print "\n\n[=]Finished Scanning $target\n";

exit (0);