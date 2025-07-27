package workflows

type Step interface {
	Execute(projectName string) error
}
