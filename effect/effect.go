package effect

import "signal/runner"

func cleanup(currentRunner *runner.Runner) {
	if currentRunner == nil {
		return
	}

	for dep := range currentRunner.Dependencies {
		// Remove the current runner from the dependencies of signals
		if dep != nil {
			delete(*dep, currentRunner)
			if len(*dep) == 0 {
				delete(currentRunner.Dependencies, dep)
			}
		}

		// clearing the dependencies of the current runner
		// wen just wanted dependencies so that we can remove the current runner from the dependencies of signals
		currentRunner.Dependencies = make(map[*map[*runner.Runner]bool]bool)
	}

}

func CreateEffect(cb func()) {
	var _runner *runner.Runner

	execute := func() {
		runner.RunnerArray = append(runner.RunnerArray, _runner)
		cleanup(_runner)
		cb()
		runner.RunnerArray = runner.RunnerArray[:len(runner.RunnerArray)-1]
	}

	_runner = runner.NewRunner(&execute, make(map[runner.TRunnerArray]bool))
	execute()
}
