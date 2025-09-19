package workflows

type WorkflowDependencies struct {
	ProjectName string
	Database    string
}

type Workflow interface {
	AddStep(step Step)
	Handle(dependencies WorkflowDependencies) error
}
