package git

import (
	gitPb "alfred/modules/GitService/v1/pb"
	"bufio"
	"context"
	"encoding/json"
	"fmt"
	"go.uber.org/cadence/activity"
	"go.uber.org/cadence/workflow"
	"go.uber.org/zap"
	"io/ioutil"
	"log"
	"net/http"
	"os"
	"os/exec"
	"strings"
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

	request := PublishRequest{
		Method: "publish",
		Params: &Params{
			Channel: "clone-logs",
			Data: &Data{
				Log: "Some logs will be present here",
			},
		},
	}

	cmd := exec.Command(name, args...)
	stdout, _ := cmd.StdoutPipe()
	if err := cmd.Start(); err != nil {
		return err
	}

	oneByte := make([]byte, 100)
	num := 1
	for {
		_, err := stdout.Read(oneByte)
		if err != nil {
			log.Println(err)
			break
		}
		r := bufio.NewReader(stdout)
		line, _, _ := r.ReadLine()
		fmt.Println(string(line))
		num = num + 1
	}

	if err := cmd.Wait(); err != nil {
		return err
	}

	if err := Publish(&request); err != nil {
		return err
	}

	return nil
}

type PublishRequest struct {
	Method string  `json:"method"`
	Params *Params `json:"params"`
}

type Params struct {
	Channel string `json:"channel"`
	Data    *Data  `json:"data"`
}

type Data struct {
	Log string `json:"log"`
}

func Publish(request *PublishRequest) error {
	publishBody, err := json.Marshal(request)
	if err != nil {
		return err
	}
	url := "http://localhost:9000/api"
	method := "POST"
	payload := strings.NewReader(string(publishBody))

	client := &http.Client{}
	req, err := http.NewRequest(method, url, payload)
	if err != nil {
		return err
	}

	req.Header.Add("Content-Type", "application/json")
	req.Header.Add("Authorization", "apikey abc")
	res, err := client.Do(req)
	if err != nil {
		return err
	}

	defer res.Body.Close()
	_, err = ioutil.ReadAll(res.Body)
	if err != nil {
		return err
	}

	if res.StatusCode != 200 {
		return err
	}
	return nil
}
