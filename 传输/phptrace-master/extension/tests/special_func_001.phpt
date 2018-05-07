--TEST--
Trace special function >= 7.0
--INI--
trace.dotrace=1
--SKIPIF--
<?php
require 'skipif.inc';
for_verion_gte('7.0');
?>
--FILE--
<?php
function call_normal($arg = null) {}

// Call in internal
call_user_func('call_normal');

// strlen
strlen('hello');

// Closure
call_user_func(function () {
    call_normal();
});
?>
--EXPECTF--
> cli php %s
> {main}() called at [%s:2]
    > call_normal() called at [%s:5]
    < call_normal() = NULL called at [%s:5] ~ %fs %fs
    > {closure:%s:11-13}() called at [%s:13]
        > call_normal() called at [%s:12]
        < call_normal() = NULL called at [%s:12] ~ %fs %fs
    < {closure:%s:11-13}() = NULL called at [%s:13] ~ %fs %fs
< {main}() = 1 called at [%s:2] ~ %fs %fs
< cli php %s
