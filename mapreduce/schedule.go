package mapreduce

import "fmt"

//
// schedule() starts and waits for all tasks in the given phase (Map
// or Reduce). the mapFiles argument holds the names of the files that
// are the inputs to the map phase, one per map task. nReduce is the
// number of reduce tasks. the registerChan argument yields a stream
// of registered workers; each item is the worker's RPC address,
// suitable for passing to call(). registerChan will yield all
// existing registered workers (if any) and new ones as they register.
//
func schedule(jobName string, mapFiles []string, nReduce int, phase jobPhase, registerChan chan string) {
	var njobs int
	var num int // number of inputs (for reduce) or outputs (for map)
	switch phase {
	case mapPhase:
		njobs = len(mapFiles)
		num = nReduce
	case reducePhase:
		ntasks = nReduce
		num = len(mapFiles)
	}

	fmt.Printf("Schedule: %v %v tasks, %d I/O files)\n", njobs, phase, num)
	fmt.Printf("Schedule: %v Complete\n", phase)
}
