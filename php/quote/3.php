<?php
echo <<<EOL
\e[0;34:40m
\e[0m\n
EOL;
echo PHP_EOL;
echo "Made by Abhishek-op ";
echo PHP_EOL;
echo "\e[0;34:40mQuote:\e[0m\n";
echo PHP_EOL;
$url = 'http://api.forismatic.com/api/1.0/?method=getQuote&format=json&lang=en';
$api_url = $url.$a;
$json_data = file_get_contents($api_url);
$response_data = json_decode($json_data, true);
print $response_data['quoteText'];
echo PHP_EOL;
echo "\e[0;34:40mAuthor:\e[0m\n";
print $response_data['quoteAuthor'];
echo PHP_EOL;
?>