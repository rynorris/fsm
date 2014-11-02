/*
 * General Purpose FSM library.
 *
 */
package fsm

import "fmt"

const (
	NO_INPUT Input = -1
)

type Input int

type Action func() Input

func NO_ACTION() Input { return NO_INPUT }

type Outcome struct {
	State  int
	Action Action
}

type State struct {
	Index    int
	Outcomes map[Input]Outcome
}

type FSM struct {
	states  map[int]State
	current int
}

type InvalidInputError struct {
	StateIndex int
	Input      Input
}

func (err InvalidInputError) Error() string {
	return fmt.Sprintf("input invalid in current state.  (State: %v, Input: %v)", err.StateIndex, err.Input)
}

type ImpossibleStateError int

func (err ImpossibleStateError) Error() string {
	return fmt.Sprintf("FSM in impossible state: %d", err)
}

// Define an FSM from a list of States.
func Define(states ...State) *FSM {
	stateMap := map[int]State{}
	for _, s := range states {
		stateMap[s.Index] = s
	}

	return &FSM{
		states:  stateMap,
		current: states[0].Index,
	}
}

// Spin the FSM one time.
func (f *FSM) Spin(i Input) error {
	s, ok := f.states[f.current]

	if !ok {
		return ImpossibleStateError(f.current)
	}

	do, ok := s.Outcomes[i]

	if !ok {
		return InvalidInputError{f.current, i}
	}

	next := do.Action()
	f.current = do.State
	if next != NO_INPUT {
		return f.Spin(next)
	}

	return nil
}
