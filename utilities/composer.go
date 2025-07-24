package utilities

import (
	"bytes"
	"fmt"
	"os/exec"
)

// CommandRunner abstracts exec.Command for testability
type CommandRunner interface {
	Command(name string, arg ...string) Cmd
	LookPath(file string) (string, error)
}

// Cmd abstracts exec.Cmd for testability
type Cmd interface {
	Run() error
	StderrPipe() (*bytes.Buffer, error)
	SetStderr(buf *bytes.Buffer)
}

// ComposerRunner implements CommandRunner using os/exec
type ComposerRunner struct{}

func (ComposerRunner) Command(name string, arg ...string) Cmd {
	return &ComposerCommand{cmd: exec.Command(name, arg...)}
}
func (ComposerRunner) LookPath(file string) (string, error) {
	return exec.LookPath(file)
}

type ComposerCommand struct {
	cmd *exec.Cmd
}

func (c *ComposerCommand) Run() error                         { return c.cmd.Run() }
func (c *ComposerCommand) StderrPipe() (*bytes.Buffer, error) { return &bytes.Buffer{}, nil }
func (c *ComposerCommand) SetStderr(buf *bytes.Buffer)        { c.cmd.Stderr = buf }

func RunComposer(runner CommandRunner, args []string) (string, error) {
	// validate if Composer is installed
	if err := validateIfComposerIsInstalled(runner); err != nil {
		return "", err
	}

	// default arguments for composer command
	args = append(args, "--no-interaction", "--no-progress")
	// execute the composer command
	cmd := runner.Command("composer", args...)
	var stderr bytes.Buffer
	cmd.SetStderr(&stderr)

	if err := cmd.Run(); err != nil {
		return stderr.String(), fmt.Errorf("%w", err)
	}
	return stderr.String(), nil
}

func validateIfComposerIsInstalled(runner CommandRunner) error {
	if _, err := runner.LookPath("composer"); err != nil {
		return fmt.Errorf("composer is not installed or not found in PATH. please install composer to use this tool")
	}
	return nil
}
