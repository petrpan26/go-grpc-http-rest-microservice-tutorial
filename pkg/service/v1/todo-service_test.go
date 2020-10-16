package v1

import (
	"context"
	"io/ioutil"
	"log"
	"testing"
	"time"

	"github.com/golang/protobuf/ptypes"

	as "github.com/aerospike/aerospike-client-go"
	"github.com/petrpan26/go-grpc-http-rest-microservice-tutorial/pkg/api/v1"
	testcontainers "github.com/testcontainers/testcontainers-go"
	"github.com/testcontainers/testcontainers-go/wait"
)

func Test_toDoServiceServer_Create(t *testing.T) {
	ctx := context.Background()
	req := testcontainers.ContainerRequest{
		Image:        "aerospike:latest",
		ExposedPorts: []string{"3000/tcp"},
		WaitingFor:   wait.ForLog("service ready: soon there will be cake!"),
	}
	asClient, err := testcontainers.GenericContainer(ctx, testcontainers.GenericContainerRequest{
		ContainerRequest: req,
		Started:          true,
	})
	if err != nil {
		t.Error(err)
	}
	defer asClient.Terminate(ctx)
	time.Sleep(5 * time.Second)
	ip, err := asClient.Host(ctx)
	if err != nil {
		t.Error(err)
	}
	port, err := asClient.MappedPort(ctx, "3000")
	if err != nil {
		t.Error(err)
	}
	logg, err := asClient.Logs(ctx)
	if err != nil {
		t.Error(err)
	}
	x, err := ioutil.ReadAll(logg)
	if err != nil {
		t.Error(err)
	}
	log.Print(string(x))
	db, err := as.NewClient(ip, port.Int())
	if err != nil {
		t.Error(err)
	}
	s := NewToDoServiceServer(db)
	log.Print("Done setting up")
	tm := time.Now().In(time.UTC)
	reminder, _ := ptypes.TimestampProto(tm)
	ti := time.Now().In(time.UTC)
	pfx := ti.Format(time.RFC3339Nano)
	req1 := v1.CreateRequest{
		Api: apiVersion,
		ToDo: &v1.ToDo{
			Title:       "title (" + pfx + ")",
			Description: "description (" + pfx + ")",
			Reminder:    reminder,
		},
	}
	res1, err := s.Create(ctx, &req1)
	if len(res1.Id) == 0 {
		t.Error("Oops")
	}
}

// func Test_toDoServiceServer_Read(t *testing.T) {
// 	ctx := context.Background()
// 	db, mock, err := sqlmock.New()
// 	if err != nil {
// 		t.Fatalf("an error '%s' was not expected when opening a stub database connection", err)
// 	}
// 	defer db.Close()
// 	s := NewToDoServiceServer(db)
// 	tm := time.Now().In(time.UTC)
// 	reminder, _ := ptypes.TimestampProto(tm)

// 	type args struct {
// 		ctx context.Context
// 		req *v1.ReadRequest
// 	}
// 	tests := []struct {
// 		name    string
// 		s       v1.ToDoServiceServer
// 		args    args
// 		mock    func()
// 		want    *v1.ReadResponse
// 		wantErr bool
// 	}{
// 		{
// 			name: "OK",
// 			s:    s,
// 			args: args{
// 				ctx: ctx,
// 				req: &v1.ReadRequest{
// 					Api: "v1",
// 					Id:  1,
// 				},
// 			},
// 			mock: func() {
// 				rows := sqlmock.NewRows([]string{"ID", "Title", "Description", "Reminder"}).
// 					AddRow(1, "title", "description", tm)
// 				mock.ExpectQuery("SELECT (.+) FROM ToDo").WithArgs(1).WillReturnRows(rows)
// 			},
// 			want: &v1.ReadResponse{
// 				Api: "v1",
// 				ToDo: &v1.ToDo{
// 					Id:          1,
// 					Title:       "title",
// 					Description: "description",
// 					Reminder:    reminder,
// 				},
// 			},
// 		},
// 		{
// 			name: "Unsupported API",
// 			s:    s,
// 			args: args{
// 				ctx: ctx,
// 				req: &v1.ReadRequest{
// 					Api: "v1",
// 					Id:  1,
// 				},
// 			},
// 			mock:    func() {},
// 			wantErr: true,
// 		},
// 		{
// 			name: "SELECT failed",
// 			s:    s,
// 			args: args{
// 				ctx: ctx,
// 				req: &v1.ReadRequest{
// 					Api: "v1",
// 					Id:  1,
// 				},
// 			},
// 			mock: func() {
// 				mock.ExpectQuery("SELECT (.+) FROM ToDo").WithArgs(1).
// 					WillReturnError(errors.New("SELECT failed"))
// 			},
// 			wantErr: true,
// 		},
// 		{
// 			name: "Not found",
// 			s:    s,
// 			args: args{
// 				ctx: ctx,
// 				req: &v1.ReadRequest{
// 					Api: "v1",
// 					Id:  1,
// 				},
// 			},
// 			mock: func() {
// 				rows := sqlmock.NewRows([]string{"ID", "Title", "Description", "Reminder"})
// 				mock.ExpectQuery("SELECT (.+) FROM ToDo").WithArgs(1).WillReturnRows(rows)
// 			},
// 			wantErr: true,
// 		},
// 	}
// 	for _, tt := range tests {
// 		t.Run(tt.name, func(t *testing.T) {
// 			tt.mock()
// 			got, err := tt.s.Read(tt.args.ctx, tt.args.req)
// 			if (err != nil) != tt.wantErr {
// 				t.Errorf("toDoServiceServer.Read() error = %v, wantErr %v", err, tt.wantErr)
// 				return
// 			}

// 			if err == nil && !reflect.DeepEqual(got, tt.want) {
// 				t.Errorf("toDoServiceServer.Read() = %v, want %v", got, tt.want)
// 			}
// 		})
// 	}
// }

// func Test_toDoServiceServer_Update(t *testing.T) {
// 	ctx := context.Background()
// 	db, mock, err := sqlmock.New()
// 	if err != nil {
// 		t.Fatalf("an error '%s' was not expected when opening a stub database connection", err)
// 	}
// 	defer db.Close()
// 	s := NewToDoServiceServer(db)
// 	tm := time.Now().In(time.UTC)
// 	reminder, _ := ptypes.TimestampProto(tm)

// 	type args struct {
// 		ctx context.Context
// 		req *v1.UpdateRequest
// 	}
// 	tests := []struct {
// 		name    string
// 		s       v1.ToDoServiceServer
// 		args    args
// 		mock    func()
// 		want    *v1.UpdateResponse
// 		wantErr bool
// 	}{
// 		{
// 			name: "OK",
// 			s:    s,
// 			args: args{
// 				ctx: ctx,
// 				req: &v1.UpdateRequest{
// 					Api: "v1",
// 					ToDo: &v1.ToDo{
// 						Id:          1,
// 						Title:       "new title",
// 						Description: "new description",
// 						Reminder:    reminder,
// 					},
// 				},
// 			},
// 			mock: func() {
// 				mock.ExpectExec("UPDATE ToDo").WithArgs("new title", "new description", tm, 1).
// 					WillReturnResult(sqlmock.NewResult(1, 1))
// 			},
// 			want: &v1.UpdateResponse{
// 				Api:     "v1",
// 				Updated: 1,
// 			},
// 		},
// 		{
// 			name: "Unsupported API",
// 			s:    s,
// 			args: args{
// 				ctx: ctx,
// 				req: &v1.UpdateRequest{
// 					Api: "v1",
// 					ToDo: &v1.ToDo{
// 						Id:          1,
// 						Title:       "new title",
// 						Description: "new description",
// 						Reminder:    reminder,
// 					},
// 				},
// 			},
// 			mock:    func() {},
// 			wantErr: true,
// 		},
// 		{
// 			name: "Invalid Reminder field format",
// 			s:    s,
// 			args: args{
// 				ctx: ctx,
// 				req: &v1.UpdateRequest{
// 					Api: "v1",
// 					ToDo: &v1.ToDo{
// 						Id:          1,
// 						Title:       "new title",
// 						Description: "new description",
// 						Reminder: &timestamp.Timestamp{
// 							Seconds: 1,
// 							Nanos:   -1,
// 						},
// 					},
// 				},
// 			},
// 			mock:    func() {},
// 			wantErr: true,
// 		},
// 		{
// 			name: "UPDATE failed",
// 			s:    s,
// 			args: args{
// 				ctx: ctx,
// 				req: &v1.UpdateRequest{
// 					Api: "v1",
// 					ToDo: &v1.ToDo{
// 						Id:          1,
// 						Title:       "new title",
// 						Description: "new description",
// 						Reminder:    reminder,
// 					},
// 				},
// 			},
// 			mock: func() {
// 				mock.ExpectExec("UPDATE ToDo").WithArgs("new title", "new description", tm, 1).
// 					WillReturnError(errors.New("UPDATE failed"))
// 			},
// 			wantErr: true,
// 		},
// 		{
// 			name: "RowsAffected failed",
// 			s:    s,
// 			args: args{
// 				ctx: ctx,
// 				req: &v1.UpdateRequest{
// 					Api: "v1",
// 					ToDo: &v1.ToDo{
// 						Id:          1,
// 						Title:       "new title",
// 						Description: "new description",
// 						Reminder:    reminder,
// 					},
// 				},
// 			},
// 			mock: func() {
// 				mock.ExpectExec("UPDATE ToDo").WithArgs("new title", "new description", tm, 1).
// 					WillReturnResult(sqlmock.NewErrorResult(errors.New("RowsAffected failed")))
// 			},
// 			wantErr: true,
// 		},
// 		{
// 			name: "Not Found",
// 			s:    s,
// 			args: args{
// 				ctx: ctx,
// 				req: &v1.UpdateRequest{
// 					Api: "v1",
// 					ToDo: &v1.ToDo{
// 						Id:          1,
// 						Title:       "new title",
// 						Description: "new description",
// 						Reminder:    reminder,
// 					},
// 				},
// 			},
// 			mock: func() {
// 				mock.ExpectExec("UPDATE ToDo").WithArgs("new title", "new description", tm, 1).
// 					WillReturnResult(sqlmock.NewResult(1, 0))
// 			},
// 			wantErr: true,
// 		},
// 	}
// 	for _, tt := range tests {
// 		t.Run(tt.name, func(t *testing.T) {
// 			tt.mock()
// 			got, err := tt.s.Update(tt.args.ctx, tt.args.req)
// 			if (err != nil) != tt.wantErr {
// 				t.Errorf("toDoServiceServer.Update() error = %v, wantErr %v", err, tt.wantErr)
// 				return
// 			}
// 			if err == nil && !reflect.DeepEqual(got, tt.want) {
// 				t.Errorf("toDoServiceServer.Update() = %v, want %v", got, tt.want)
// 			}
// 		})
// 	}
// }

// func Test_toDoServiceServer_Delete(t *testing.T) {
// 	ctx := context.Background()
// 	db, mock, err := sqlmock.New()
// 	if err != nil {
// 		t.Fatalf("an error '%s' was not expected when opening a stub database connection", err)
// 	}
// 	defer db.Close()
// 	s := NewToDoServiceServer(db)

// 	type args struct {
// 		ctx context.Context
// 		req *v1.DeleteRequest
// 	}
// 	tests := []struct {
// 		name    string
// 		s       v1.ToDoServiceServer
// 		args    args
// 		mock    func()
// 		want    *v1.DeleteResponse
// 		wantErr bool
// 	}{
// 		{
// 			name: "OK",
// 			s:    s,
// 			args: args{
// 				ctx: ctx,
// 				req: &v1.DeleteRequest{
// 					Api: "v1",
// 					Id:  1,
// 				},
// 			},
// 			mock: func() {
// 				mock.ExpectExec("DELETE FROM ToDo").WithArgs(1).
// 					WillReturnResult(sqlmock.NewResult(1, 1))
// 			},
// 			want: &v1.DeleteResponse{
// 				Api:     "v1",
// 				Deleted: 1,
// 			},
// 		},
// 		{
// 			name: "Unsupported API",
// 			s:    s,
// 			args: args{
// 				ctx: ctx,
// 				req: &v1.DeleteRequest{
// 					Api: "v1",
// 					Id:  1,
// 				},
// 			},
// 			mock:    func() {},
// 			wantErr: true,
// 		},
// 		{
// 			name: "DELETE failed",
// 			s:    s,
// 			args: args{
// 				ctx: ctx,
// 				req: &v1.DeleteRequest{
// 					Api: "v1",
// 					Id:  1,
// 				},
// 			},
// 			mock: func() {
// 				mock.ExpectExec("DELETE FROM ToDo").WithArgs(1).
// 					WillReturnError(errors.New("DELETE failed"))
// 			},
// 			wantErr: true,
// 		},
// 		{
// 			name: "RowsAffected failed",
// 			s:    s,
// 			args: args{
// 				ctx: ctx,
// 				req: &v1.DeleteRequest{
// 					Api: "v1",
// 					Id:  1,
// 				},
// 			},
// 			mock: func() {
// 				mock.ExpectExec("DELETE FROM ToDo").WithArgs(1).
// 					WillReturnResult(sqlmock.NewErrorResult(errors.New("RowsAffected failed")))
// 			},
// 			wantErr: true,
// 		},
// 		{
// 			name: "Not Found",
// 			s:    s,
// 			args: args{
// 				ctx: ctx,
// 				req: &v1.DeleteRequest{
// 					Api: "v1",
// 					Id:  1,
// 				},
// 			},
// 			mock: func() {
// 				mock.ExpectExec("DELETE FROM ToDo").WithArgs(1).
// 					WillReturnResult(sqlmock.NewResult(1, 0))
// 			},
// 			wantErr: true,
// 		},
// 	}
// 	for _, tt := range tests {
// 		t.Run(tt.name, func(t *testing.T) {
// 			tt.mock()
// 			got, err := tt.s.Delete(tt.args.ctx, tt.args.req)
// 			if (err != nil) != tt.wantErr {
// 				t.Errorf("toDoServiceServer.Delete() error = %v, wantErr %v", err, tt.wantErr)
// 				return
// 			}
// 			if err == nil && !reflect.DeepEqual(got, tt.want) {
// 				t.Errorf("toDoServiceServer.Delete() = %v, want %v", got, tt.want)
// 			}
// 		})
// 	}
// }

// func Test_toDoServiceServer_ReadAll(t *testing.T) {
// 	ctx := context.Background()
// 	db, mock, err := sqlmock.New()
// 	if err != nil {
// 		t.Fatalf("an error '%s' was not expected when opening a stub database connection", err)
// 	}
// 	defer db.Close()
// 	s := NewToDoServiceServer(db)
// 	tm1 := time.Now().In(time.UTC)
// 	reminder1, _ := ptypes.TimestampProto(tm1)
// 	tm2 := time.Now().In(time.UTC)
// 	reminder2, _ := ptypes.TimestampProto(tm2)

// 	type args struct {
// 		ctx context.Context
// 		req *v1.ReadAllRequest
// 	}
// 	tests := []struct {
// 		name    string
// 		s       v1.ToDoServiceServer
// 		args    args
// 		mock    func()
// 		want    *v1.ReadAllResponse
// 		wantErr bool
// 	}{
// 		{
// 			name: "OK",
// 			s:    s,
// 			args: args{
// 				ctx: ctx,
// 				req: &v1.ReadAllRequest{
// 					Api: "v1",
// 				},
// 			},
// 			mock: func() {
// 				rows := sqlmock.NewRows([]string{"ID", "Title", "Description", "Reminder"}).
// 					AddRow(1, "title 1", "description 1", tm1).
// 					AddRow(2, "title 2", "description 2", tm2)
// 				mock.ExpectQuery("SELECT (.+) FROM ToDo").WillReturnRows(rows)
// 			},
// 			want: &v1.ReadAllResponse{
// 				Api: "v1",
// 				ToDos: []*v1.ToDo{
// 					{
// 						Id:          1,
// 						Title:       "title 1",
// 						Description: "description 1",
// 						Reminder:    reminder1,
// 					},
// 					{
// 						Id:          2,
// 						Title:       "title 2",
// 						Description: "description 2",
// 						Reminder:    reminder2,
// 					},
// 				},
// 			},
// 		},
// 		{
// 			name: "Empty",
// 			s:    s,
// 			args: args{
// 				ctx: ctx,
// 				req: &v1.ReadAllRequest{
// 					Api: "v1",
// 				},
// 			},
// 			mock: func() {
// 				rows := sqlmock.NewRows([]string{"ID", "Title", "Description", "Reminder"})
// 				mock.ExpectQuery("SELECT (.+) FROM ToDo").WillReturnRows(rows)
// 			},
// 			want: &v1.ReadAllResponse{
// 				Api:   "v1",
// 				ToDos: []*v1.ToDo{},
// 			},
// 		},
// 		{
// 			name: "Unsupported API",
// 			s:    s,
// 			args: args{
// 				ctx: ctx,
// 				req: &v1.ReadAllRequest{
// 					Api: "v1",
// 				},
// 			},
// 			mock:    func() {},
// 			wantErr: true,
// 		},
// 	}
// 	for _, tt := range tests {
// 		t.Run(tt.name, func(t *testing.T) {
// 			tt.mock()
// 			got, err := tt.s.ReadAll(tt.args.ctx, tt.args.req)
// 			if (err != nil) != tt.wantErr {
// 				t.Errorf("toDoServiceServer.ReadAll() error = %v, wantErr %v", err, tt.wantErr)
// 				return
// 			}
// 			if !reflect.DeepEqual(got, tt.want) {
// 				t.Errorf("toDoServiceServer.ReadAll() = %v, want %v", got, tt.want)
// 			}
// 		})
// 	}
// }
