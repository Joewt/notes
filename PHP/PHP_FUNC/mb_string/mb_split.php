<?php
 $str = 'Ма-
руся';
  print_r(preg_split('//u', $str, null, PREG_SPLIT_NO_EMPTY));
