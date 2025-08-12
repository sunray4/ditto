package mockServer

type MockServer struct {
	Name string
	ServerName string
	ServerCode string
	Commands []Command
}

type Command struct {
	Method string
	Path string
	Expectation any
}

func CreateMockServer(username string, servername string) MockServer {

	mockServer := MockServer{

		ServerName: servername,
		ServerCode: GenerateCode(14),
		Commands: []Command{},
	}
//save to database under username
	return mockServer
}

func (m *MockServer) AddCommand(command Command) bool {
	if m.CheckIfDuplicate(command) {
		return false
	}
	m.Commands = append(m.Commands, command)
	return true
}

func (m *MockServer) DeleteCommand(command Command) bool {
	for i, c := range m.Commands {
		if c.Method == command.Method && c.Path == command.Path {
			m.Commands = append(m.Commands[:i], m.Commands[i+1:]...)
			return true
		}
	}
	return false
}

func (m *MockServer) GetCommands() []Command {
	return m.Commands
}

func (m *MockServer) EditCommand(prevCommand Command, newCommand Command) bool {
	if m.CheckIfDuplicate(newCommand) {
		return false
	}
	for i, c := range m.Commands {
		if c.Method == prevCommand.Method && c.Path == prevCommand.Path {
			m.Commands[i] = newCommand
			return true
		}
	}
	return false
}

func (m *MockServer) CheckIfDuplicate(command Command) bool {
	for _, c := range m.Commands {
		if c.Method == command.Method && c.Path == command.Path {
			return true
		}
	}
	return false
}