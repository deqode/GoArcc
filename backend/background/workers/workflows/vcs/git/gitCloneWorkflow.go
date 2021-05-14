package git

import (
	gitPb "alfred/modules/git/v1/pb"
	"bufio"
	"context"
	"go.uber.org/cadence/activity"
	"go.uber.org/cadence/workflow"
	"go.uber.org/zap"
	"io"
	"log"
	"os"
	"os/exec"
	"sync"
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
		HeartbeatTimeout:       time.Minute * 20,
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
	wd, _ := os.Getwd()
	cmd.Dir = wd + "/repositories"
	stderr, _ := cmd.StderrPipe()
	stdout, _ := cmd.StdoutPipe()
	if err := cmd.Start(); err != nil {
		return err
	}
	wg := new(sync.WaitGroup)
	wg.Add(2)

	go reader(wg, stderr)
	go reader(wg, stdout)
	wg.Wait()

	err := cmd.Wait()
	if err != nil {
		return err
	}
	return nil
}

//read stdout and std err and publish
func reader(wg *sync.WaitGroup, closer io.ReadCloser) {
	request := PublishRequest{
		Method: "publish",
		Params: &Params{
			Channel: "clone-logs",
			Data: &Data{
				Log: "Some logs will be present here",
			},
		},
	}
	defer wg.Done()
	in := bufio.NewScanner(closer)
	for in.Scan() {
		request.Params.Data.Log = in.Text()
		log.Println(in.Text())
		err := Publish(&request)
		if err != nil {
			log.Println(err)
		}
	}
	if err := in.Err(); err != nil {
		log.Println("error:", err)
	}
}
