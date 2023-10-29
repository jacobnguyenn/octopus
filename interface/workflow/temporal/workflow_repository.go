package temporal

import (
	"context"
	"log"

	"go.temporal.io/sdk/client"

	"ddd-sample/domain/repo"
)

type workflowRepository struct {
	tC client.Client
}

var (
	_ repo.IWorkflowRepo = (*workflowRepository)(nil)
)

func NewWorkflowRepository() repo.IWorkflowRepo {
	tC, err := newTemporalClient()
	if err != nil {
		panic(err)
	}
	return &workflowRepository{tC: tC}
}

func newTemporalClient() (client.Client, error) {
	return client.Dial(client.Options{})
}

func (w workflowRepository) Start(ctx context.Context, workflowId string) error {
	log.Printf("starting workflow, id %s", workflowId)
	return nil
}

func (w workflowRepository) End(ctx context.Context, workflowId string) error {
	log.Printf("terminating workflow, id %s", workflowId)
	return nil
}

func (w workflowRepository) Status(ctx context.Context, workflowId string) (string, error) {
	// TODO implement me
	return "running", nil
}
