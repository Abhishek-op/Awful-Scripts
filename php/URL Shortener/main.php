<?php
echo <<<EOL
\e[0;34:40m
	 _     _       _                                  
| |   | |     | |                                 
| |   | | ____| |                                 
| |   | |/ ___) |                                 
| |___| | |   | |                                 
 \______|_|   |_|                                 
                                                  
    _    _                                        
   | |  | |                _                      
    \ \ | | _   ___   ____| |_  ____   ____  ____ 
     \ \| || \ / _ \ / ___)  _)|  _ \ / _  )/ ___)
 _____) ) | | | |_| | |   | |__| | | ( (/ /| |    
(______/|_| |_|\___/|_|    \___)_| |_|\____)_|    
                                        
\e[0m\n
EOL;
echo PHP_EOL;
echo "Made by Abhishek-op ";
echo PHP_EOL;
echo PHP_EOL;
$a = readline("Enter the URL:  ");
echo PHP_EOL;
echo PHP_EOL;
$url = 'https://tinyurl.com/api-create.php?url=';
$api_url = $url.$a;
$json_data = file_get_contents($api_url);
echo "Here is your short URL: ";
echo PHP_EOL;
echo $json_data;
echo PHP_EOL;
?>