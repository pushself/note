--TEST--
Trace Trait
--SKIPIF--
<?php
require 'skipif.inc';
require_debug_mode();
for_verion_gte('5.4');
?>
--FILE--
<?php trace_start();

trait TestTrait
{
    public function callNormal() {}
}

class TestClassA
{
    use TestTrait;
}

class TestClassB
{
    use TestTrait;
}

class TestClassC
{
    use TestTrait {
        TestTrait::callNormal as callAlias;
    }
}

$a = new TestClassA();
$b = new TestClassB();
$c = new TestClassC();

$a->callNormal();
$b->callNormal();
$c->callAlias();

trace_end(); ?>
--EXPECTF--
> TestClassA->callNormal() called at [%s:29]
    < TestClassA->callNormal() = NULL called at [%s:29] ~ %fs %fs
    > TestClassB->callNormal() called at [%s:30]
    < TestClassB->callNormal() = NULL called at [%s:30] ~ %fs %fs
    > TestClassC->callAlias() called at [%s:31]
    < TestClassC->callAlias() = NULL called at [%s:31] ~ %fs %fs
    > trace_end() called at [%s:33]
