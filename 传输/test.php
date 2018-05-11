<?php
my_trace_start();
$arr = [123,1];
test($arr);
test($arr);
function test($arr)
{
return $arr;
}
$a = my_trace_end();
var_dump($a);
