在PHP世界中，如果发生灵异事件没有头绪，或者无法进一步排查定位原因时，有两个重要的工具可以进行回溯追踪。这两个工具的效果是类似，但表现形式会有所不同。分别是：

debug_backtrace()，产生一条回溯跟踪
debug_print_backtrace()，打印一条回溯
下面通过一个示例来快速认识这两个强大工具的作用。

创建两个临时的PHP文件，一个是/tmp/test.php入口文件，一个是/tmp/Debug.php类文件。在里面分别放置以下代码：

/tmp/test.php入口文件代码：

<?php
require_once '/tmp/Debug.php';

function test() {
    $obj = new Debug();

    $obj->printBacktrace('first');
    echo PHP_EOL;
    $obj->backtrace('second');
}

test();
/tmp/Debug.php类文件代码：

<?php
class Debug {
    public function printBacktrace($name) {
        debug_print_backtrace();
    }

    public function backtrace($name) {
        var_dump(debug_backtrace());
    }
} 
执行/tmp/test.php文件，可以看到以下输出内容。

$ php /tmp/test.php
#0  Debug->printBacktrace(first) called at [/tmp/test.php:7]
#1  test() called at [/tmp/test.php:12]

array(2) {
  [0] =>
  array(7) {
    'file' =>
    string(13) "/tmp/test.php"
    'line' =>
    int(9)
    'function' =>
    string(9) "backtrace"
    'class' =>
    string(5) "Debug"
    'object' =>
    class Debug#1 (0) {
    }
    'type' =>
    string(2) "->"
    'args' =>
    array(1) {
      [0] =>
      string(6) "second"
    }
  }
  [1] =>
  array(4) {
    'file' =>
    string(13) "/tmp/test.php"
    'line' =>
    int(12)
    'function' =>
    string(4) "test"
    'args' =>
    array(0) {
    }
  }
}
