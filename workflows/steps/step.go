package steps

// StepParameters holds the parameters for a workflow step
type StepParameters struct {
	ProjectName string
	Database    string
}

// Step is the interface for a workflow step
type Step interface {
	Execute(parameters StepParameters) error
}
