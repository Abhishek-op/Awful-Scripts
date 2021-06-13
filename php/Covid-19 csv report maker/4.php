<?php
echo <<<EOL
\e[0;34:40m
\e[0m\n
EOL;
echo PHP_EOL;
echo "Enter your country name:";
echo PHP_EOL;
$a = readline();
echo PHP_EOL;
$url = 'https://api.covid19api.com/total/country/';
$api_url = $url.$a;
$json_data = file_get_contents($api_url);
$array = json_decode($json_data, true);
$csvFileName = 'Report.csv';
$fcsv = fopen($csvFileName, 'w');
$header = false;
foreach ($array as $line) {	
    if (empty($header)) {
        $header = array_keys($line);
        fputcsv($fcsv, $header);
        $header = array_flip($header);		
    }
	
	$line_array = array();
	
	foreach($line as $value) {
		array_push($line_array, $value);
	}

    fputcsv($fcsv, $line_array);
}

fclose($fcsv);
echo "Covid 19 report for $a is created as report.csv";
?>