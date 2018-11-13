<?php

/**

排序函数属性
函数名称	排序依据	数组索引键保持	排序的顺序	相关函数
array_multisort()	值	键值关联的保持，数字类型的不保持	第一个数组或者由选项指定	array_walk()
asort()	值	是	由低到高	arsort()
arsort()	值	是	由高到低	asort()
krsort()	键	是	由高到低	ksort()
ksort()	键	是	由低到高	asort()
natcasesort()	值	是	自然排序，大小写不敏感	natsort()
natsort()	值	是	自然排序	natcasesort()
rsort()	值	否	由高到低	sort()
shuffle()	值	否	随机	array_rand()
sort()	值	否	由低到高	rsort()
uasort()	值	是	由用户定义	uksort()
uksort()	键	是	由用户定义	uasort()
usort()	值	否	由用户定义	uasort()

**/


/**
bool sort ( array &$array [, int $sort_flags = SORT_REGULAR ] )
array
要排序的数组。

sort_flags
可选的第二个参数 sort_flags 可以用以下值改变排序的行为：

排序类型标记：

SORT_REGULAR - 正常比较单元（不改变类型）
SORT_NUMERIC - 单元被作为数字来比较
SORT_STRING - 单元被作为字符串来比较
SORT_LOCALE_STRING - 根据当前的区域（locale）设置来把单元当作字符串比较，可以用 setlocale() 来改变。
SORT_NATURAL - 和 natsort() 类似对每个单元以“自然的顺序”对字符串进行排序。 PHP 5.4.0 中新增的。
SORT_FLAG_CASE - 能够与 SORT_STRING 或 SORT_NATURAL 合并（OR 位运算），不区分大小写排序字符串。

**/

function print_arr($arr)
{
    foreach($arr as $k => $v){
        echo $k."=>".$v." ";
    }
    echo "\n";
}

$arr_number = array(1,2,3,1,5,10,11,23,123,14,14,5,1,5,1,1,2,231231,3123);
$arr_string = array("lemon", "orange", "banana", "apple");
print_arr($arr_number);
sort($arr_number);
print_arr($arr_number);

print_arr($arr_string);
sort($arr_string);
print_arr($arr_string);


/**
asort — 对数组进行排序并保持索引关系
bool asort ( array &$array [, int $sort_flags = SORT_REGULAR ] )
**/

$arr_number = array(1,2,3,1,5,10,11,23,123,14,14,5,1,5,1,1,2,231231,3123);
$arr_string = array("lemon", "orange", "banana", "apple");
asort($arr_number);
print_arr($arr_number);


/**
rsort — 对数组逆向排序
bool rsort ( array &$array [, int $sort_flags = SORT_REGULAR ] )
**/


$arr_number = array(1,2,3,1,5,10,11,23,123,14,14,5,1,5,1,1,2,231231,3123);
rsort($arr_number);
print_arr($arr_number);

/**
arsort - 逆向排序并保持索引关系
**/

$arr_number = array(1,2,3,1,5,10,11,23,123,14,14,5,1,5,1,1,2,231231,3123);
arsort($arr_number);
print_arr($arr_number);


/**
ksort — 对数组按照键名排序
bool ksort ( array &$array [, int $sort_flags = SORT_REGULAR ] )
**/

$fruits = array("d"=>"lemon", "a"=>"orange", "b"=>"banana", "c"=>"apple");
ksort($fruits);
print_arr($fruits);


/**
krsort — 对数组按照键名逆向排序
bool krsort ( array &$array [, int $sort_flags = SORT_REGULAR ] )
**/

$fruits = array("d"=>"lemon", "a"=>"orange", "b"=>"banana", "c"=>"apple");
krsort($fruits);
print_arr($fruits);


/**
uasort — 使用用户自定义的比较函数对数组中的值进行排序并保持索引关联
uksort — 使用用户自定义的比较函数对数组中的键名进行排序
usort — 使用用户自定义的比较函数对数组中的值进行排序
**/




/**
natsort — 用“自然排序”算法对数组排序
**/


$array1 = $array2 = array("img12.png", "img10.png", "img2.png", "img1.png");

asort($array1);
echo "Standard sorting\n";
print_r($array1);

natsort($array2);
echo "\nNatural order sorting\n";
print_r($array2);



/**
natcasesort — 用“自然排序”算法对数组进行不区分大小写字母的排序
**/


$array1 = $array2 = array('IMG0.png', 'img12.png', 'img10.png', 'img2.png', 'img1.png', 'IMG3.png');

sort($array1);
echo "Standard sorting\n";
print_r($array1);

natcasesort($array2);
echo "\nNatural order sorting (case-insensitive)\n";
print_r($array2);

