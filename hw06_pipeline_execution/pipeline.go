package hw06pipelineexecution

type (
	In  = <-chan interface{}
	Out = In
	Bi  = chan interface{}
)

type Stage func(in In) (out Out)

func ExecutePipeline(in In, done In, stages ...Stage) Out {
	if in == nil || stages == nil {
		nilChan := make(Bi)
		close(nilChan)
		return nilChan
	}
	for _, stage := range stages {
		in = worker(done, stage(in))
	}
	return in
}

func worker(done In, in In) Bi {
	out := make(Bi)
	go func() {
		defer close(out)
		for {
			select {
			case <-done:
				return
			case v, ok := <-in:
				if !ok {
					return
				}
				out <- v
			}
		}
	}()

	return out
}
