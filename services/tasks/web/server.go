package web

import (
	"context"
	"fmt"
	"net"

	metav1 "github.com/BradErz/monorepo/gen/go/org/meta/v1"
	tasksv1 "github.com/BradErz/monorepo/gen/go/org/tasks/v1"
	"github.com/BradErz/monorepo/pkg/xgrpc"
	"github.com/google/uuid"
	"github.com/sirupsen/logrus"
	"google.golang.org/grpc/reflection"
	"google.golang.org/protobuf/types/known/timestamppb"
)

type Server struct {
	le    *logrus.Entry
	addr  string
	tasks []*tasksv1.Task
}

func New(le *logrus.Entry, addr string) (*Server, error) {
	return &Server{
		le:    le,
		addr:  addr,
		tasks: make([]*tasksv1.Task, 0),
	}, nil
}

func (srv *Server) ListTasks(ctx context.Context, req *tasksv1.ListTasksRequest) (*tasksv1.ListTasksResponse, error) {
	return &tasksv1.ListTasksResponse{
		Tasks: srv.tasks,
	}, nil
}

func (srv *Server) CreateTask(ctx context.Context, req *tasksv1.CreateTaskRequest) (*tasksv1.CreateTaskResponse, error) {
	task := &tasksv1.Task{
		Id:     uuid.New().String(),
		Name:   req.GetName(),
		Labels: req.GetLabels(),
		Created: &metav1.Created{
			By: "me",
			At: timestamppb.Now(),
		},
	}
	srv.tasks = append(srv.tasks, task)
	return &tasksv1.CreateTaskResponse{
		Task: task,
	}, nil
}

func (srv *Server) Run() error {
	lis, err := net.Listen("tcp", srv.addr)
	if err != nil {
		return fmt.Errorf("failed to create tcp listener: %w", err)
	}

	gSrv, err := xgrpc.NewServer(srv.le)
	if err != nil {
		return fmt.Errorf("failed to create grpc.Server: %w", err)
	}
	tasksv1.RegisterTasksServiceServer(gSrv, srv)
	reflection.Register(gSrv)

	return gSrv.Serve(lis)
}
