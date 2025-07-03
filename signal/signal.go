package signal

import (
	"signal/runner"
)

func addSubscription(runner *runner.Runner, subscription *map[*runner.Runner]bool) {
	(*subscription)[runner] = true
	runner.Dependencies[subscription] = true
}

func CreateSignal[T any](value T) (func() T, func(T)) {

	subscription := make(map[*runner.Runner]bool)

	read := func() T {
		currentRunner := runner.RunnerArray[len(runner.RunnerArray)-1]
		if currentRunner != nil {
			addSubscription(currentRunner, &subscription)
		}
		return value
	}

	write := func(newValue T) {
		value = newValue
		for dep := range subscription {
			(*dep.Execute)()
		}
	}

	return read, write
}
