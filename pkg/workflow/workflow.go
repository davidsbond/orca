package workflow

import (
	"context"
	"encoding/json"
)

type (
	Implementation[Input, Output any] struct {
		WorkflowName string
		Action       Action[Input, Output]
	}

	Action[Input any, Output any] func(ctx context.Context, input Input) (Output, error)
)

func (w *Implementation[Input, Output]) Run(ctx context.Context, input []byte) ([]byte, error) {
	var inp Input
	if err := json.Unmarshal(input, &inp); err != nil {
		return nil, err
	}

	output, err := w.Action(ctx, inp)
	if err != nil {
		return nil, err
	}

	return json.Marshal(output)
}

func (w *Implementation[Input, Output]) Name() string {
	return w.WorkflowName
}
