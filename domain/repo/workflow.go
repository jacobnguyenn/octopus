package repo

import "context"

type IWorkflowRepo interface {
	// input ?
	Start(ctx context.Context, workflowId string) error
	End(ctx context.Context, workflowId string) error
}
