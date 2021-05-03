package git

import (
	gitPb "alfred/modules/GitService/v1/pb"
	"context"
	"fmt"
	"go.uber.org/cadence/activity"
	"go.uber.org/cadence/workflow"
	"go.uber.org/zap"
	"os"
	"os/exec"
	"time"
)

// This is registration process where you register all your workflows
// and activity function handlers.

func init() {
	workflow.Register(GitCloneWorkflow)
	activity.Register(GitCloneActivity)
}

func GitCloneActivity(ctx context.Context, req *gitPb.CloneRepositoryRequest) (*gitPb.GetCloningStatusResponse, error) {
	logger := activity.GetLogger(ctx)
	logger.Info("repository cloning activity started")

	wd, err := os.Getwd()
	if err != nil {
		return nil, err
	}
	repoWD := wd + "/repositories"
	if err := os.Chdir(repoWD); err != nil {
		return nil, err
	}
	fmt.Println(repoWD)
	args := []string{"clone", req.RepositoryUrl}
	if err := runCommand("git", args...); err != nil {
		return nil, err
	}
	return &gitPb.GetCloningStatusResponse{
		Status: true,
	}, nil
}

func GitCloneWorkflow(ctx workflow.Context, req *gitPb.CloneRepositoryRequest) (*gitPb.GetCloningStatusResponse, error) {
	ctx = workflow.WithActivityOptions(ctx, workflow.ActivityOptions{
		ScheduleToStartTimeout: time.Minute,
		StartToCloseTimeout:    time.Minute,
		HeartbeatTimeout:       time.Second * 20,
	})

	logger := workflow.GetLogger(ctx)
	logger.Info("repository clone workflow started")
	var activityResult gitPb.GetCloningStatusResponse
	err := workflow.ExecuteActivity(ctx, GitCloneActivity, req).Get(ctx, &activityResult)
	if err != nil {
		logger.Error("Activity failed.", zap.Error(err))
		return nil, err
	}
	return &activityResult, nil
}

func runCommand(name string, args ...string) error {
	cmd := exec.Command(name, args...)
	cmd.Stderr = os.Stdout
	return cmd.Run()

}
