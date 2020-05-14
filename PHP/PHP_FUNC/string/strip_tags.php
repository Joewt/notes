<?php
/**
string strip_tags ( string $str [, string $allowable_tags ] )
**/


$text = '<p>Test paragraph.</p><!-- Comment --> <a href="#fragment">Other text</a>';

echo strip_tags($text);
echo "\n";


$text_php = '<?php echo strlen("hello world");?> fasdf <?php echo 1;?>';

echo strip_tags($text_php);
