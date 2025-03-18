<?php
use parallel\Runtime;
use parallel\Sync;

$THREADS = 15;

$runtimes = [];
$futures = [];
$sync = new Sync(1);

for($i=0; $i<$THREADS; $i++) {
	$runtimes[] = new Runtime();
}

for($i=0; $i<$THREADS; $i++) {
	//echo "Task {$i} waiting to enter the room\n";
	$futures[] = $runtimes[$i]->run(function($sync, $taskId) {
		while ($sync->get() == 0) {
			continue;
		}

		$sync->set($sync->get()-1);
		echo "Task {$taskId} entered the room\n";
		echo "Task {$taskId} exited\n";
		$sync->set($sync->get() + 1);
	}, [$sync, $i]);
}

foreach ($futures as $future) {
	$future->value();
}
?>
