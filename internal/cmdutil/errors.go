package cmdutil

const ActionRequestedExitCode = 1
const FatalErrExitCode = 2

// Error can be used by a command to change the exit status of the CLI.
type Error struct {
	Msg string
	// Status
	ExitCode int
}

func (e *Error) Error() string { return e.Msg }
