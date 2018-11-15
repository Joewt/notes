<?php
$str = "A 'quote' is <b>bold</b>";
//echo htmlentities($str);


echo html_entity_decode(htmlentities($str));
