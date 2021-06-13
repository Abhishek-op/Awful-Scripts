<?php
echo <<<EOL
\e[0;34:40m

 __      __                .___           _____                       .__                
/  \    /  \___________  __| _/          /     \   ____ _____    ____ |__| ____    ____  
\   \/\/   /  _ \_  __ \/ __ |  ______  /  \ /  \_/ __ \\__  \  /    \|  |/    \  / ___\ 
 \        (  <_> )  | \/ /_/ | /_____/ /    Y    \  ___/ / __ \|   |  \  |   |  \/ /_/  >
  \__/\  / \____/|__|  \____ |         \____|__  /\___  >____  /___|  /__|___|  /\___  / 
       \/                   \/                 \/     \/     \/     \/        \//_____/  
\e[0m\n
EOL;
echo PHP_EOL;
echo "Made by Abhishek-op ";
echo PHP_EOL;
echo PHP_EOL;
$a = readline("Enter the word:  ");
echo PHP_EOL;
echo PHP_EOL;
$url = 'https://api.dictionaryapi.dev/api/v2/entries/en/';
$api_url = $url.$a;
$json_data = file_get_contents($api_url);
$response_data = json_decode($json_data, true);
foreach ($response_data as $key => $value) {
    foreach ($value["meanings"] as $key => $value) {
        echo "\e[0;34:40mPart-Of-Speech\e[0m\n";
        echo PHP_EOL;
        echo PHP_EOL;
        echo $value["partOfSpeech"];
        echo PHP_EOL;
        echo PHP_EOL;
        foreach ($value["definitions"] as $key => $value) {
            echo "\e[0;34:40mDifinition:\e[0m\n";
            echo PHP_EOL;
            echo PHP_EOL;
            echo $value["definition"];
            echo PHP_EOL;
            echo PHP_EOL;
            echo "\e[0;34:40mExample:\e[0m\n";
            echo PHP_EOL;
            echo PHP_EOL;
            echo $value["example"];
            echo PHP_EOL;
            echo PHP_EOL;
            echo "\e[0;34:40mSynonyms:\e[0m\n";
            echo PHP_EOL;
            echo PHP_EOL;
            foreach ($value["synonyms"] as $key) {
                echo $key.", ";
                echo PHP_EOL;
            }
            echo PHP_EOL;   
        }
    }
}
?>