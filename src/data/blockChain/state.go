package blockChain

type State struct {
	stateMap map[string]string
}

var stateInstance *State

func GetStateInstance() *State {
	if stateInstance == nil {
		stateInstance = &State{
			stateMap: make(map[string]string),
		}
	}
	return stateInstance
}

func (s *State) GetState(key string) string {
	return s.stateMap[key]
}

func (s *State) SetState(key, value string) {
	s.stateMap[key] = value
}
