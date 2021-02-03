package execution

import (
	"context"
	"errors"
	"fmt"
	"time"

	"github.com/newrelic/newrelic-cli/internal/install/types"
)

type MockRecipeExecutor struct {
	result      bool
	prepareFunc func(context.Context, types.DiscoveryManifest, types.Recipe, bool) (types.RecipeVars, error)
	executeFunc func(context.Context, types.DiscoveryManifest, types.Recipe, types.RecipeVars) error
}

func NewMockRecipeExecutor() *MockRecipeExecutor {
	return &MockRecipeExecutor{
		result:      false,
		prepareFunc: defaultPrepareFunc,
		executeFunc: defaultExecuteFunc,
	}
}

func (m *MockRecipeExecutor) Prepare(ctx context.Context, dm types.DiscoveryManifest, r types.Recipe, y bool) (types.RecipeVars, error) {
	return m.prepareFunc(ctx, dm, r, y)
}

func (m *MockRecipeExecutor) Execute(ctx context.Context, dm types.DiscoveryManifest, r types.Recipe, v types.RecipeVars) error {
	return m.executeFunc(ctx, dm, r, v)
}

func defaultPrepareFunc(ctx context.Context, dm types.DiscoveryManifest, r types.Recipe, y bool) (types.RecipeVars, error) {
	// fmt.Print("\n defaultPrepareFunc \n")

	// time.Sleep(5 * time.Second)

	return types.RecipeVars{}, nil
}

func defaultExecuteFunc(ctx context.Context, dm types.DiscoveryManifest, r types.Recipe, v types.RecipeVars) error {
	fmt.Print("\n defaultPrepareFunc \n")

	return sleepWithContext(ctx, 5*time.Second)
}

func sleepWithContext(ctx context.Context, duration time.Duration) error {
	select {
	case <-ctx.Done(): //context cancelled
		return errors.New("exiting process")
	case <-time.After(duration): //timeout
	}

	return nil
}
