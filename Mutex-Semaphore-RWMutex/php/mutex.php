<?php

use parallel\Runtime;
use parallel\Sync;

$threads = 3;
$runtimes = [];
$futures = [];
$counter = new Sync(0);

for($i=0; $i<$threads; $i++) {
	$runtimes[$i] = new Runtime();
	$futures[$i] = $runtimes[$i]->run(
		function($counter, $threadId) {
			for($j=0; $j<5; $j++) {
				usleep(rand(10000, 50000));
				$counter->set($counter->get() + 1);
				echo "Thread {$threadId} - {$j} working\n"; 
			}
	}, [$counter, $i]);
}

foreach($futures as $future) {
	$future->value();
}

echo "Final Counter: " . $counter->get();
?>
