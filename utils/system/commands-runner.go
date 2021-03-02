package system

// Action is a callback for a function call result
type Action func()

// CommandID is just an identificator of commands
type CommandID string

// Command is a command descriptor
type Command struct {
	ID CommandID

	WelcomeMessage string
	GoodbyeMessage string
	ErrorAction    Action
	SuccessAction  Action

	Command   string
	Arguments []string
}

// ICommandRunner is an interface of commands runner
type ICommandRunner interface {
	Run(command Command)
}
