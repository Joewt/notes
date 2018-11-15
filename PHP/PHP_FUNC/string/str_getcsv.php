<?php
$csv = array_map('str_getcsv', file('data.csv'));

print_r($csv);
