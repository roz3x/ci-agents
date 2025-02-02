package restapi

import (
	"bytes"
	"context"
	"io"
	"sort"
	"strings"
	"time"

	check "github.com/erikh/check"
	"github.com/golang/mock/gomock"
	"github.com/tinyci/ci-agents/errors"
	"github.com/tinyci/ci-agents/mocks/github"
	"github.com/tinyci/ci-agents/types"

	gh "github.com/google/go-github/github"
)

var ctx = context.Background()

func (us *uisvcSuite) TestCapabilities(c *check.C) {
	client := github.NewMockClient(gomock.NewController(c))
	_, doneChan, tc, utc, err := MakeUIServer(client)
	c.Assert(err, check.IsNil)
	defer close(doneChan)

	c.Assert(utc.AddCapability(ctx, "erikh2", "modify:user"), check.NotNil)
	c.Assert(tc.AddCapability(ctx, "erikh2", "modify:user"), check.IsNil)
	c.Assert(utc.AddCapability(ctx, "erikh2", "modify:ci"), check.IsNil)

	props, err := utc.GetUserProperties(ctx)
	c.Assert(err, check.IsNil)
	caps := []string{}
	for _, cap := range props["capabilities"].([]interface{}) {
		caps = append(caps, cap.(string))
	}

	sort.Strings(caps)
	c.Assert(caps, check.DeepEquals, []string{"modify:ci", "modify:user"})

	c.Assert(utc.RemoveCapability(ctx, "erikh2", "modify:user"), check.IsNil)
	c.Assert(utc.RemoveCapability(ctx, "erikh2", "modify:ci"), check.NotNil)
}

func (us *uisvcSuite) TestErrors(c *check.C) {
	client := github.NewMockClient(gomock.NewController(c))
	_, doneChan, tc, _, err := MakeUIServer(client)
	c.Assert(err, check.IsNil)
	defer close(doneChan)

	errs, err := tc.Errors(ctx)
	c.Assert(err, check.IsNil)
	c.Assert(len(errs), check.Equals, 0)
}

func (us *uisvcSuite) TestLogAttach(c *check.C) {
	client := github.NewMockClient(gomock.NewController(c))
	_, doneChan, tc, _, err := MakeUIServer(client)
	c.Assert(err, check.IsNil)
	defer close(doneChan)

	c.Assert(us.assetsvcClient.Write(context.Background(), 1, bytes.NewBufferString("this is a log")), check.IsNil)
	time.Sleep(100 * time.Millisecond)

	pr, pw := io.Pipe()
	buf := bytes.NewBuffer(nil)

	finished := make(chan struct{})
	go func() {
		defer close(finished)
		_, err := io.Copy(buf, pr)
		c.Assert(err, check.IsNil)
	}()

	c.Assert(tc.LogAttach(ctx, 1, pw), check.IsNil)
	pw.Close()
	<-finished
	c.Assert(strings.HasPrefix(buf.String(), "this is a log"), check.Equals, true)
}

func (us *uisvcSuite) TestTokenEndpoints(c *check.C) {
	client := github.NewMockClient(gomock.NewController(c))
	_, doneChan, tc, _, err := MakeUIServer(client)
	c.Assert(err, check.IsNil)
	defer close(doneChan)

	c.Assert(tc.DeleteToken(ctx), check.IsNil)
	_, err = tc.Errors(ctx)
	c.Assert(err, check.ErrorMatches, ".*invalid authentication")
}

func (us *uisvcSuite) TestDeleteToken(c *check.C) {
	client := github.NewMockClient(gomock.NewController(c))
	_, doneChan, tc, _, err := MakeUIServer(client)
	c.Assert(err, check.IsNil)
	defer close(doneChan)

	c.Assert(tc.DeleteToken(ctx), check.IsNil)
	_, err = tc.Errors(ctx)
	c.Assert(err, check.ErrorMatches, ".*invalid authentication")
}

func (us *uisvcSuite) TestSubmit(c *check.C) {
	client := github.NewMockClient(gomock.NewController(c))
	_, doneChan, tc, utc, err := MakeUIServer(client)
	c.Assert(err, check.IsNil)
	defer close(doneChan)

	client.EXPECT().GetRepository(gomock.Any(), "erikh/not-real").Return(nil, errors.New("not found"))

	c.Assert(tc.Submit(ctx, "erikh/not-real", "master", true), check.ErrorMatches, ".* not found")

	client.EXPECT().MyRepositories(gomock.Any()).Return([]*gh.Repository{{FullName: gh.String("erikh/parent")}}, nil)

	repos, err := tc.LoadRepositories(ctx, "")
	c.Assert(err, check.IsNil)
	c.Assert(len(repos), check.Not(check.Equals), 0)

	c.Assert(us.datasvcClient.Client().EnableRepository(ctx, "erikh", "erikh/parent"), check.IsNil)

	client.EXPECT().GetRepository(gomock.Any(), "erikh/not-real").Return(nil, errors.New("not found"))
	c.Assert(tc.Submit(ctx, "erikh/not-real", "master", true), check.ErrorMatches, ".* not found")

	sub := &types.Submission{
		Parent:  "erikh/parent",
		Fork:    "erikh/test",
		BaseSHA: "aaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaa",
		HeadSHA: "bbbbbbbbbbbbbbbbbbbbbbbbbbbbbbbbbbbbbbbb",
		Manual:  true,
		All:     true,
	}

	c.Assert(us.queuesvcClient.SetMockSubmissionSuccess(client.EXPECT(), sub, "../"), check.IsNil)

	client.EXPECT().GetSHA(gomock.Any(), "erikh/parent", "heads/master").Return("", errors.New("not found"))
	client.EXPECT().GetSHA(gomock.Any(), "erikh/test", "heads/master").Return("bbbbbbbbbbbbbbbbbbbbbbbbbbbbbbbbbbbbbbbb", nil)

	c.Assert(tc.Submit(ctx, "erikh/test", "master", true), check.ErrorMatches, ".*not found")

	client.EXPECT().GetSHA(gomock.Any(), "erikh/parent", "heads/master").Return("aaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaa", nil)
	client.EXPECT().GetSHA(gomock.Any(), "erikh/test", "heads/master").Return("", errors.New("not found"))

	c.Assert(tc.Submit(ctx, "erikh/test", "master", true), check.ErrorMatches, ".*not found")

	c.Assert(us.queuesvcClient.SetMockSubmissionSuccess(client.EXPECT(), sub, "../"), check.IsNil)
	client.EXPECT().GetSHA(gomock.Any(), "erikh/parent", "heads/master").Return("aaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaa", nil)
	client.EXPECT().GetSHA(gomock.Any(), "erikh/test", "heads/master").Return("bbbbbbbbbbbbbbbbbbbbbbbbbbbbbbbbbbbbbbbb", nil)
	client.EXPECT().ClearStates(gomock.Any(), "erikh/parent", "bbbbbbbbbbbbbbbbbbbbbbbbbbbbbbbbbbbbbbbb").Return(nil)

	c.Assert(utc.Submit(ctx, "erikh/test", "master", true), check.NotNil)
	c.Assert(tc.Submit(ctx, "erikh/test", "master", true), check.IsNil)

	tasks, err := tc.Tasks(ctx, "erikh/test", "bbbbbbbbbbbbbbbbbbbbbbbbbbbbbbbbbbbbbbbb", 0, 200)
	c.Assert(err, check.IsNil)
	c.Assert(len(tasks), check.Not(check.Equals), 0)
	count, err := tc.TaskCount(ctx, "erikh/test", "bbbbbbbbbbbbbbbbbbbbbbbbbbbbbbbbbbbbbbbb")
	c.Assert(err, check.IsNil)
	c.Assert(int(count), check.Equals, len(tasks))

	for _, task := range tasks {
		c.Assert(task.Parent.HookSecret, check.Equals, "")
		c.Assert(task.Ref.Repository.HookSecret, check.Equals, "")
		runs, err := tc.RunsForTask(ctx, task.ID, 0, 200)
		c.Assert(err, check.IsNil)
		count, err := tc.RunsForTaskCount(ctx, task.ID)
		c.Assert(err, check.IsNil)
		c.Assert(len(runs), check.Equals, int(count))
	}

	runs, err := tc.Runs(ctx, "erikh/test", "bbbbbbbbbbbbbbbbbbbbbbbbbbbbbbbbbbbbbbbb", 0, 200)
	c.Assert(err, check.IsNil)
	c.Assert(len(runs), check.Not(check.Equals), 0)

	count, err = tc.RunsCount(ctx, "erikh/test", "bbbbbbbbbbbbbbbbbbbbbbbbbbbbbbbbbbbbbbbb")
	c.Assert(err, check.IsNil)
	c.Assert(len(runs), check.Equals, int(count))

	for _, run := range runs {
		c.Assert(run.Task.Parent.HookSecret, check.Equals, "")
		c.Assert(run.Task.Ref.Repository.HookSecret, check.Equals, "")
	}

	runs, err = tc.RunsForTask(ctx, runs[0].Task.ID, 0, 200)
	c.Assert(err, check.IsNil)

	for _, run := range runs {
		c.Assert(run.Task.Parent.HookSecret, check.Equals, "")
		c.Assert(run.Task.Ref.Repository.HookSecret, check.Equals, "")
	}

	for i := 0; i < len(runs); i++ {
		client.EXPECT().ErrorStatus(
			gomock.Any(),
			"erikh",
			"parent",
			gomock.Any(),
			"bbbbbbbbbbbbbbbbbbbbbbbbbbbbbbbbbbbbbbbb",
			gomock.Any(),
			gomock.Any()).Return(nil)
	}

	c.Assert(runs[0].Task.Canceled, check.Equals, false)
	c.Assert(tc.CancelRun(ctx, runs[0].ID), check.IsNil)
	run, err := tc.GetRun(ctx, runs[0].ID)
	c.Assert(err, check.IsNil)
	c.Assert(run.Task.Canceled, check.Equals, true)
}

func (us *uisvcSuite) TestAddDeleteCI(c *check.C) {
	client := github.NewMockClient(gomock.NewController(c))
	h, doneChan, tc, utc, err := MakeUIServer(client)
	c.Assert(err, check.IsNil)
	defer close(doneChan)

	c.Assert(tc.AddToCI(ctx, "f7u12"), check.NotNil)

	client.EXPECT().MyRepositories(gomock.Any()).Return([]*gh.Repository{{FullName: gh.String("erikh/test")}}, nil)

	repos, err := tc.LoadRepositories(ctx, "")
	c.Assert(err, check.IsNil)
	c.Assert(len(repos), check.Not(check.Equals), 0)

	for _, repo := range repos {
		c.Assert(repo.HookSecret, check.Equals, "")
	}

	c.Assert(tc.AddToCI(ctx, "erikh/not-real"), check.NotNil) // not saved

	client.EXPECT().TeardownHook(gomock.Any(), "erikh", "test", h.HookURL).Return(errors.New("wat's up"))

	c.Assert(tc.AddToCI(ctx, "erikh/test"), check.ErrorMatches, "wat's up")

	client.EXPECT().TeardownHook(gomock.Any(), "erikh", "test", h.HookURL).Return(nil)
	client.EXPECT().SetupHook(gomock.Any(), "erikh", "test", h.HookURL, gomock.Any()).Return(errors.New("yep"))

	c.Assert(tc.AddToCI(ctx, "erikh/test"), check.ErrorMatches, "yep")

	client.EXPECT().TeardownHook(gomock.Any(), "erikh", "test", h.HookURL).Return(nil)
	client.EXPECT().SetupHook(gomock.Any(), "erikh", "test", h.HookURL, gomock.Any()).Return(nil)

	c.Assert(utc.AddToCI(ctx, "erikh/test"), check.ErrorMatches, ".*capability.*")
	c.Assert(tc.AddToCI(ctx, "erikh/test"), check.IsNil)

	client.EXPECT().MyRepositories(gomock.Any()).Return([]*gh.Repository{{FullName: gh.String("erikh/test")}}, nil)

	visible, err := tc.Visible(ctx, "erikh/test")
	c.Assert(err, check.IsNil)
	c.Assert(len(visible), check.Equals, 1)

	for _, repo := range visible {
		c.Assert(repo.HookSecret, check.Equals, "")
	}

	c.Assert(tc.DeleteFromCI(ctx, "erikh/not-real"), check.NotNil)

	client.EXPECT().TeardownHook(gomock.Any(), "erikh", "test", h.HookURL).Return(errors.New("wat's up"))
	c.Assert(tc.DeleteFromCI(ctx, "erikh/test"), check.NotNil)

	client.EXPECT().TeardownHook(gomock.Any(), "erikh", "test", h.HookURL).Return(nil)
	c.Assert(utc.DeleteFromCI(ctx, "erikh/test"), check.NotNil)
	c.Assert(tc.DeleteFromCI(ctx, "erikh/test"), check.IsNil)
	c.Assert(tc.DeleteFromCI(ctx, "erikh/test"), check.ErrorMatches, "repo is not enabled")
}

func (us *uisvcSuite) TestSubscriptions(c *check.C) {
	client := github.NewMockClient(gomock.NewController(c))
	_, doneChan, tc, _, err := MakeUIServer(client)
	c.Assert(err, check.IsNil)
	defer close(doneChan)

	c.Assert(tc.Subscribe(ctx, "erikh/test"), check.NotNil)
	c.Assert(tc.Unsubscribe(ctx, "erikh/test"), check.NotNil)

	client.EXPECT().MyRepositories(gomock.Any()).Return([]*gh.Repository{{FullName: gh.String("erikh/test")}}, nil)

	repos, err := tc.LoadRepositories(ctx, "")
	c.Assert(err, check.IsNil)
	c.Assert(len(repos), check.Not(check.Equals), 0)

	repos, err = tc.Subscribed(ctx, "")
	c.Assert(err, check.IsNil)
	c.Assert(len(repos), check.Equals, 0)

	c.Assert(tc.Subscribe(ctx, "erikh/test"), check.IsNil)

	repos, err = tc.Subscribed(ctx, "")
	c.Assert(err, check.IsNil)
	c.Assert(len(repos), check.Equals, 1)
	c.Assert(repos[0].Name, check.Equals, "erikh/test")

	c.Assert(tc.Unsubscribe(ctx, "erikh/test"), check.IsNil)

	repos, err = tc.Subscribed(ctx, "")
	c.Assert(err, check.IsNil)
	c.Assert(len(repos), check.Equals, 0)
}

func (us *uisvcSuite) TestVisibility(c *check.C) {
	client := github.NewMockClient(gomock.NewController(c))
	_, doneChan, tc, _, err := MakeUIServer(client)
	c.Assert(err, check.IsNil)
	defer close(doneChan)

	_, err = us.datasvcClient.MakeUser("not-erikh")
	c.Assert(err, check.IsNil)
	_, err = us.datasvcClient.MakeUser("erikh-the-third")
	c.Assert(err, check.IsNil)

	c.Assert(us.datasvcClient.MakeRepo("not-erikh/private-test", "not-erikh", true, ""), check.IsNil)
	c.Assert(us.datasvcClient.MakeRepo("erikh/private-test", "erikh", true, ""), check.IsNil)
	c.Assert(us.datasvcClient.MakeRepo("erikh/public", "erikh", false, ""), check.IsNil)
	c.Assert(us.datasvcClient.MakeRepo("erikh-the-third/public", "erikh-the-third", false, ""), check.IsNil)

	repos, err := tc.Visible(ctx, "")
	c.Assert(err, check.IsNil)
	c.Assert(len(repos), check.Equals, 3)

	for _, repo := range repos {
		c.Assert(repo.Name, check.Not(check.Equals), "not-erikh/private-test")
	}
}
