<?php


my_frace_start();
$b = show([1]);

function show($arr)
{
	$a = $arr;
	return show2([2]);
}
function show2($arr)
{
	$a = $arr;
	return [3];
}
$a = my_frace_end();
var_dump($a);


