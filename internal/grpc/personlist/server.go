package personsgrpc

import (
	"context"
	"persons/service/internal/domain/models"
	pb "persons/service/internal/grpc/gen"
	"persons/service/internal/services/personlist"
	"time"

	"google.golang.org/grpc"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
)

type Server struct {
	pb.UnimplementedPersonsServer
	personList *personlist.PersonList
}

func Register(s *grpc.Server, personList personlist.PersonList) {
	pb.RegisterPersonsServer(s, &Server{personList: &personList})
}

func (s *Server) AddPerson(ctx context.Context, in *pb.Person) (*pb.PersonID, error) {
	birthdate, err := time.Parse("2006-01-02", in.Birthdate)
	if err != nil {
		return nil, status.Error(codes.InvalidArgument, "Failed to parse birthdate")
	}
	id, err := s.personList.AddPerson(
		ctx,
		&models.Person{
			ID:        in.ID,
			Name:      in.Name,
			Birthdate: birthdate,
		},
	)
	if err != nil {
		return nil, status.Error(codes.Internal, "")
	}
	return &pb.PersonID{Value: id}, status.New(codes.OK, "").Err()
}

func (s *Server) GetPerson(ctx context.Context, in *pb.PersonID) (*pb.Person, error) {
	person, err := s.personList.GetPerson(ctx, in.Value)
	if err != nil {
		return nil, status.Error(codes.Internal, "Failed to get person")
	}
	return &pb.Person{
		ID:        person.ID,
		Name:      person.Name,
		Birthdate: person.Birthdate.Format("2006-01-02"),
	}, status.New(codes.OK, "").Err()
}

func (s *Server) EditPerson(ctx context.Context, in *pb.Person) (*pb.Empty, error) {
	birthdate, err := time.Parse("2006-01-02", in.Birthdate)
	if err != nil {
		return nil, status.Error(codes.InvalidArgument, "Failed to parse birthdate")
	}
	err = s.personList.EditPerson(ctx,
		&models.Person{
			ID:        in.ID,
			Name:      in.Name,
			Birthdate: birthdate,
		},
	)
	if err != nil {
		return nil, status.Error(codes.Internal, "Failed to edit person")
	}
	return nil, status.New(codes.OK, "").Err()
}

func (s *Server) RemovePerson(ctx context.Context, in *pb.PersonID) (*pb.Empty, error) {
	err := s.personList.RemovePerson(ctx, in.Value)
	if err != nil {
		return nil, status.Error(codes.Internal, "Failed to remove person")
	}
	return nil, status.New(codes.OK, "").Err()
}

func (s *Server) GetList(in *pb.Empty, stream grpc.ServerStreamingServer[pb.Person]) error {
	personList, err := s.personList.GetList()
	if err != nil {
		return status.Error(codes.Internal, "Failed to get person list")
	}
	for _, person := range personList {
		err = stream.Send(&pb.Person{
			ID:        person.ID,
			Name:      person.Name,
			Birthdate: person.Birthdate.Format("2006-01-02"),
		})
		if err != nil {
			return status.Error(codes.Aborted, "")
		}
	}
	return status.New(codes.OK, "").Err()
}
