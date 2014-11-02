package fsm

import (
	"testing"
)

const (
	test_input_1 = iota
	test_input_2
	test_input_3
)

const (
	test_state_1 = iota
	test_state_2
	test_state_3
)

func assertState(t *testing.T, fsm *FSM, i Input, s int) {
	oldState := fsm.current

	err := fsm.Spin(i)
	if err != nil {
		t.Fatal(err.Error())
		return
	}

	if fsm.current != s {
		t.Errorf("FSM in wrong state. (Input: %v, State: %v -> %v)", i, oldState, s)
	}
}

func TestStates(t *testing.T) {
	state1 := State{
		Index: test_state_1,
		Outcomes: map[Input]Outcome{
			test_input_1: Outcome{test_state_2, NO_ACTION},
			test_input_2: Outcome{test_state_3, NO_ACTION},
			test_input_3: Outcome{test_state_1, NO_ACTION},
		},
	}
	state2 := State{
		Index: test_state_2,
		Outcomes: map[Input]Outcome{
			test_input_1: Outcome{test_state_1, NO_ACTION},
			test_input_2: Outcome{test_state_1, NO_ACTION},
			test_input_3: Outcome{test_state_1, NO_ACTION},
		},
	}
	state3 := State{
		Index: test_state_3,
		Outcomes: map[Input]Outcome{
			test_input_1: Outcome{test_state_2, NO_ACTION},
			test_input_2: Outcome{test_state_1, NO_ACTION},
			test_input_3: Outcome{test_state_1, NO_ACTION},
		},
	}

	fsm := Define(state1, state2, state3)

	t.Log("1: 1 -> 2")
	assertState(t, fsm, test_input_1, test_state_2)

	t.Log("3: 2 -> 1")
	assertState(t, fsm, test_input_3, test_state_1)

	t.Log("2: 1 -> 3")
	assertState(t, fsm, test_input_2, test_state_3)

	t.Log("1: 3 -> 2")
	assertState(t, fsm, test_input_1, test_state_2)

	t.Log("2: 2 -> 1")
	assertState(t, fsm, test_input_2, test_state_1)

	t.Log("3: 1 -> 1")
	assertState(t, fsm, test_input_3, test_state_1)
}
