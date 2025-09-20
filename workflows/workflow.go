package workflows

import "github.com/racastellanosm/tools-php-skeleton-creator/workflows/steps"

type Workflow interface {
	AddStep(step steps.Step)
	Handle(parameters steps.StepParameters) error
}
