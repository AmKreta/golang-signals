package runner

type Runner struct {
	Execute      *func()
	Dependencies map[*map[*Runner]bool]bool
}

type TRunnerArray = *[]*Runner

var RunnerArray []*Runner = []*Runner{}

func NewRunner(execute *func(), dependencies map[TRunnerArray]bool) *Runner {
	runner := &Runner{
		Execute:      execute,
		Dependencies: make(map[*map[*Runner]bool]bool),
	}

	return runner
}
