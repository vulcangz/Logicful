package workflow

import (
	"encoding/json"
	"logicful/features/form"
	"logicful/features/workflow/integrations"
	"logicful/models"
	"logicful/service/queue"
)

func Register() {
	registerWorkflowProcessor()
	integrations.RegisterEmail()
	integrations.RegisterWebhook()
}

func registerWorkflowProcessor() {
	queue.Receive("workflow", func(message []byte) error {
		var result models.Submission
		err := json.Unmarshal(message, &result)
		if err != nil {
			return err
		}
		return processWorkflow(result)
	})
}

func processWorkflow(submission models.Submission) error {
	f, err := form.Get(submission.FormId)
	if err != nil {
		return err
	}
	for _, integration := range f.Workflow.Integrations {
		if !integration.Enabled {
			continue
		}
		err := queue.Dispatch(integration.Name, models.Integration{
			Submission: submission,
			Name:       integration.Name,
			Config:     integration.Config,
			Form:       f,
		})
		if err != nil {
			return err
		}
	}
	return nil
}
