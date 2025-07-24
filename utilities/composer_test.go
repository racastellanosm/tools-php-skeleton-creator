package utilities

import (
	"bytes"
	"errors"
	"os/exec"
	"reflect"
	"testing"
)

type mockCmd struct {
	runErr error
	stderr string
}

func (m *mockCmd) Run() error                         { return m.runErr }
func (m *mockCmd) StderrPipe() (*bytes.Buffer, error) { return bytes.NewBufferString(m.stderr), nil }
func (m *mockCmd) SetStderr(buf *bytes.Buffer)        { buf.WriteString(m.stderr) }

type mockRunner struct {
	lookPathErr error
	cmd         *mockCmd
}

func (m *mockRunner) Command(name string, arg ...string) Cmd { return m.cmd }
func (m *mockRunner) LookPath(file string) (string, error) {
	if m.lookPathErr != nil {
		return "", m.lookPathErr
	}
	return "/usr/bin/composer", nil
}

func TestRunComposer_Success(t *testing.T) {
	runner := &mockRunner{cmd: &mockCmd{runErr: nil, stderr: ""}}
	output, err := RunComposer(runner, []string{"install"})
	if err != nil {
		t.Fatalf("expected no error, got %v", err)
	}
	if output != "" {
		t.Errorf("expected empty stderr, got %q", output)
	}
}

func TestRunComposer_ComposerNotInstalled(t *testing.T) {
	runner := &mockRunner{lookPathErr: errors.New("not found"), cmd: &mockCmd{}}
	_, err := RunComposer(runner, []string{"install"})
	if err == nil || err.Error() != "composer is not installed or not found in PATH. please install composer to use this tool" {
		t.Errorf("expected composer not installed error, got %v", err)
	}
}

func TestRunComposer_CommandFails(t *testing.T) {
	runner := &mockRunner{cmd: &mockCmd{runErr: errors.New("fail"), stderr: "some error"}}
	output, err := RunComposer(runner, []string{"install"})
	if err == nil {
		t.Fatal("expected error, got nil")
	}
	if output != "some error" {
		t.Errorf("expected stderr 'some error', got %q", output)
	}
}

func TestComposerRunner_Command_ReturnsComposerCommand(t *testing.T) {
	runner := ComposerRunner{}
	cmd := runner.Command("echo", "hello")
	if _, ok := cmd.(*ComposerCommand); !ok {
		t.Errorf("expected ComposerCommand, got %T", cmd)
	}
}

func TestComposerRunner_LookPath(t *testing.T) {
	runner := ComposerRunner{}
	path, err := runner.LookPath("go")
	if err != nil {
		t.Errorf("expected to find 'go' in PATH, got error: %v", err)
	}
	if path == "" {
		t.Error("expected non-empty path for 'go' executable")
	}
}

func TestComposerCommand_SetStderr(t *testing.T) {
	cmd := &ComposerCommand{cmd: exec.Command("echo", "test")}
	buf := &bytes.Buffer{}
	cmd.SetStderr(buf)
	if !reflect.DeepEqual(cmd.cmd.Stderr, buf) {
		t.Error("SetStderr did not set the Stderr field correctly")
	}
}

func TestComposerCommand_StderrPipe(t *testing.T) {
	cmd := &ComposerCommand{cmd: exec.Command("echo", "test")}
	buf, err := cmd.StderrPipe()
	if err != nil {
		t.Errorf("expected no error, got %v", err)
	}
	if buf == nil {
		t.Error("expected non-nil buffer from StderrPipe")
	}
}
