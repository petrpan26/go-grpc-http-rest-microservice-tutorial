package v1

import (
	"context"
	as "github.com/aerospike/aerospike-client-go"

	"fmt"
	"time"

	"github.com/golang/protobuf/ptypes"
	"github.com/google/uuid"

	"github.com/petrpan26/go-grpc-http-rest-microservice-tutorial/pkg/api/v1"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
)

const (
	// apiVersion is version of API is provided by server
	apiVersion = "v1"
)

// toDoServiceServer is implementation of v1.ToDoServiceServer proto interface
type toDoServiceServer struct {
	db *as.Client
}

// NewToDoServiceServer creates ToDo service
func NewToDoServiceServer(db *as.Client) v1.ToDoServiceServer {
	return &toDoServiceServer{db: db}
}

// checkAPI checks if the API version requested by client is supported by server
func (s *toDoServiceServer) checkAPI(api string) error {
	// API version is "" means use current version of the service
	if len(api) > 0 {
		if apiVersion != api {
			return status.Errorf(codes.Unimplemented,
				"unsupported API version: service implements API version '%s', but asked for '%s'", apiVersion, api)
		}
	}
	return nil
}

// Create new todo task
func (s *toDoServiceServer) Create(ctx context.Context, req *v1.CreateRequest) (*v1.CreateResponse, error) {
	// check if the API version requested by client is supported by server
	if err := s.checkAPI(req.Api); err != nil {
		return nil, err
	}

	reminder, err := ptypes.Timestamp(req.ToDo.Reminder)
	if err != nil {
		return nil, status.Error(codes.InvalidArgument, "reminder field has invalid format-> "+err.Error())
	}

	// insert ToDo entity data
	id := uuid.New().String()
	key, err := as.NewKey("test", "task", id)
	if err != nil {
		return nil, status.Error(codes.Unavailable, "can't create new key "+err.Error())
	}
	title := as.NewBin("title", req.ToDo.Title)
	description := as.NewBin("description", req.ToDo.Description)
	reminder := as.NewBin("reminder", reminder.Unix())

	err := s.db.PutBins(key, title, description, reminder)
	if err != nil {
		return nil, status.Error(codes.Unavailable, "can't add new key "+err.Error())
	}

	return &v1.CreateResponse{
		Api: apiVersion,
		Id:  id,
	}, nil
}

// Read todo task
func (s *toDoServiceServer) Read(ctx context.Context, req *v1.ReadRequest) (*v1.ReadResponse, error) {
	// check if the API version requested by client is supported by server
	if err := s.checkAPI(req.Api); err != nil {
		return nil, err
	}

	// query ToDo by ID
	record, err := s.db.Get(nil, req.Key)
	if err != nil {
		return nil, status.Error(codes.Unavailable, "Can't get todo "+err.Error())
	}

	// get ToDo data
	var td v1.ToDo
	var reminder time.Time
	td.Reminder, err = ptypes.TimestampProto(time.Unix(record.Bins["reminder"]))
	if err != nil {
		return nil, status.Error(codes.Unknown, "reminder field has invalid format-> "+err.Error())
	}

	td.Title = record.Bins["title"]
	td.Description = record.Bins["description"]

	return &v1.ReadResponse{
		Api:  apiVersion,
		ToDo: &td,
	}, nil

}

// Update todo task
func (s *toDoServiceServer) Update(ctx context.Context, req *v1.UpdateRequest) (*v1.UpdateResponse, error) {
	// check if the API version requested by client is supported by server
	if err := s.checkAPI(req.Api); err != nil {
		return nil, err
	}

	reminder, err := ptypes.Timestamp(req.ToDo.Reminder)
	if err != nil {
		return nil, status.Error(codes.InvalidArgument, "reminder field has invalid format-> "+err.Error())
	}

	// check Todo exist
	key, err := as.NewKey("test", "task", req.ToDo.Id)
	// insert ToDo entity data
	if err != nil {
		return nil, status.Error(codes.Unavailable, "can't create new update key "+err.Error())
	}

	exist, err := s.db.Exists(nil, key)

	if err != nil {
		return nil, status.Error(codes.Unavailable, "can't query key "+err.Error())
	}

	if !exist {
		return &v1.CreateResponse{
			Api:     apiVersion,
			Updated: 0,
		}, nil
	}

	title := as.NewBin("title", req.ToDo.Title)
	description := as.NewBin("description", req.ToDo.Description)
	reminder := as.NewBin("reminder", reminder.Unix())

	err := s.db.PutBins(key, title, description, reminder)
	if err != nil {
		return nil, status.Error(codes.Unavailable, "can't update todo "+err.Error())
	}

	return &v1.CreateResponse{
		Api:     apiVersion,
		Updated: 1,
	}, nil
}

// Delete todo task
func (s *toDoServiceServer) Delete(ctx context.Context, req *v1.DeleteRequest) (*v1.DeleteResponse, error) {
	// check if the API version requested by client is supported by server
	if err := s.checkAPI(req.Api); err != nil {
		return nil, err
	}

	// query ToDo by ID
	key, err := as.NewKey("test", "task", req.Id)
	existed, err := s.db.Delete(nil, key)
	if err != nil {
		return nil, status.Error(codes.Unavailable, "can't delete task "+err.Error())
	}

	var deleted = 1

	if !existed {
		deleted = 0
	}

	return &v1.ReadResponse{
		Api:     apiVersion,
		Deleted: deleted,
	}, nil
}

// Read all todo tasks
func (s *toDoServiceServer) ReadAll(ctx context.Context, req *v1.ReadAllRequest) (*v1.ReadAllResponse, error) {
	// check if the API version requested by client is supported by server
	if err := s.checkAPI(req.Api); err != nil {
		return nil, err
	}

	spolicy := as.NewScanPolicy()
	spolicy.ConcurrentNodes = true
	spolicy.Priority = as.HIGH
	spolicy.IncludeBinData = true

	recs, err := s.db.ScanAll(spolicy, "test", "task")

	var reminder time.Time
	list := []*v1.ToDo{}
	if err != nil {
		return nil, status.Error(codes.Unknown, "failed to retrieve data from ToDo-> "+err.Error())
	}

	for rec := range recs.Results() {
		if rec.Err != nil {
			return nil, status.Error(codes.Unknown, "failed to retrieve data from ToDo-> "+err.Error())
		} else {
			var td v1.ToDo
			var reminder time.Time
			td.Reminder, err = ptypes.TimestampProto(time.Unix(res.Record.Bins["reminder"]))
			if err != nil {
				return nil, status.Error(codes.Unknown, "reminder field has invalid format-> "+err.Error())
			}

			td.Title = res.Record.Bins["title"]
			td.Description = res.Record.Bins["description"]
			append(list, &td)
		}
	}

	return &v1.ReadAllResponse{
		Api:   apiVersion,
		ToDos: list,
	}, nil
}
