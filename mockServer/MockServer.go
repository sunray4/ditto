package mockServer

import "encoding/json"

type MockServer struct {
	Username string
	ServerName string
	ServerCode string
	Commands []Command
}

type Command struct {
	Id int
	Method string
	Path string
	Expectation any
}

func CreateMockServer(username string, servername string) MockServer {

	mockServer := MockServer{
		Username: username,
		ServerName: servername,
		ServerCode: GenerateCode(14),
		Commands: []Command{},
	}
//save to database under username
	return mockServer
}

func (m *MockServer) createCommand(method string, path string, expectation any) Command {
	return Command{
		Id: len(m.Commands) + 1,
		Method: method,
		Path: path,
		Expectation: expectation,
	}
}

func (m *MockServer) AddCommand(method string, path string, expectation any) bool {
	command := m.createCommand(method, path, expectation)
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

func (m *MockServer) GetServer() ([]byte, error) {
	serverInfo := map[string]interface{}{
		"ServerCode": m.ServerCode,
		"Username":   m.Username,
		"ServerName": m.ServerName,
		"Commands":   m.Commands,
	}
	return json.Marshal(serverInfo)
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