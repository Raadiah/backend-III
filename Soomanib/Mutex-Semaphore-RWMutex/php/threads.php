<?php

use parallel\Runtime;

$runtime1 = new Runtime();
$runtime2 = new Runtime();

$future1 = $runtime1->run(function() {
	echo "Task 1 started\n";
	sleep(4);
	return "Thread 1 running completed\n";
});

$future2 = $runtime2->run(function() {
	echo "Task 2 started\n";
	sleep(1);
	return "Thread 2 running completed\n";
});

echo "Main script running..\n";
$thread2Val = $future1->value();
$thread1Val = $future2->value();

echo $thread1Val;
echo $thread2Val;
?>
