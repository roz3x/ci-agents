package queuesvc

import (
	"testing"
	"time"

	check "github.com/erikh/check"
	"github.com/tinyci/ci-agents/api/datasvc"
	"github.com/tinyci/ci-agents/api/logsvc"
	"github.com/tinyci/ci-agents/ci-gen/grpc/handler"
	"github.com/tinyci/ci-agents/config"
	"github.com/tinyci/ci-agents/mocks/github"
	"github.com/tinyci/ci-agents/model"
	"github.com/tinyci/ci-agents/testutil"
	"github.com/tinyci/ci-agents/testutil/testclients"
)

type queuesvcSuite struct {
	datasvcClient  *testclients.DataClient
	queuesvcClient *testclients.QueueClient
	queueDoneChan  chan struct{}
	dataDoneChan   chan struct{}
	logDoneChan    chan struct{}
	model          *model.Model
	dataHandler    *handler.H
	logHandler     *handler.H
	queueHandler   *handler.H
}

var _ = check.Suite(&queuesvcSuite{})

func TestQueueSvc(t *testing.T) {
	check.TestingT(t)
}

func (qs *queuesvcSuite) SetUpTest(c *check.C) {
	testutil.WipeDB(c)

	var err error
	qs.model, err = model.New(testutil.TestDBConfig)
	c.Assert(err, check.IsNil)

	qs.dataHandler, qs.dataDoneChan, err = datasvc.MakeDataServer()
	c.Assert(err, check.IsNil)

	var lj *logsvc.LogJournal

	qs.logHandler, qs.logDoneChan, lj, err = logsvc.MakeLogServer()
	c.Assert(err, check.IsNil)

	go lj.Tail()

	qs.queueHandler, qs.queueDoneChan, err = MakeQueueServer()
	c.Assert(err, check.IsNil)

	qs.datasvcClient, err = testclients.NewDataClient()
	c.Assert(err, check.IsNil)

	qs.queuesvcClient, err = testclients.NewQueueClient(qs.datasvcClient)
	c.Assert(err, check.IsNil)
}

func (qs *queuesvcSuite) TearDownTest(c *check.C) {
	close(qs.logDoneChan)
	close(qs.dataDoneChan)
	close(qs.queueDoneChan)
	time.Sleep(100 * time.Millisecond)
}

func (qs *queuesvcSuite) mkGithubClient(client *github.MockClient) {
	config.SetDefaultGithubClient(client)
}
