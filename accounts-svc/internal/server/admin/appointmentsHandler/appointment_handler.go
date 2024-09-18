package appointmentshandler

import (
	"context"

	ucinterface "github.com/AbdulRahimOM/gov-services-app/accounts-svc/internal/usecase/interface"
	requests "github.com/AbdulRahimOM/gov-services-app/internal/common-dto/request"
	pb "github.com/AbdulRahimOM/gov-services-app/internal/pb/generated"
	stdresponse "github.com/AbdulRahimOM/gov-services-app/internal/std-response/std-response"
	"github.com/sirupsen/logrus"
	"google.golang.org/protobuf/types/known/emptypb"
)

type AppointmentServer struct {
	AppointmentUseCase ucinterface.IAppointmentUC
	pb.UnimplementedAppointmentServiceServer
	getGrpcStatus func(respCode string, errMsg string) error
}

func NewAppointmentServer(appointmentUseCase ucinterface.IAppointmentUC,logger *logrus.Entry) *AppointmentServer {
	return &AppointmentServer{
		AppointmentUseCase: appointmentUseCase,
		getGrpcStatus: stdresponse.NewGetGrpcStatusForService("accounts-svc", logger),
	}
}

// CreateChildOffice
func (s *AppointmentServer) CreateChildOffice(ctx context.Context, req *pb.CreateChildOfficeRequest) (*pb.CreateChildOfficeResponse, error) {

	proposedChildOffice := requests.ProposedOffice{
		Name:    req.ProposedChildOffice.Name,
		Address: req.ProposedChildOffice.Address,
	}

	childOfficeID, responseCode, err := s.AppointmentUseCase.CreateChildOffice(req.AdminID, &proposedChildOffice)
	if err != nil {
		return nil, stdresponse.GetGrpcStatus(responseCode, err.Error())
	} else {
		return &pb.CreateChildOfficeResponse{
			ChildOfficeID: childOfficeID,
		}, nil
	}
}

// AppointAttender
func (s *AppointmentServer) AppointAttender(ctx context.Context, req *pb.AttenderAppointmentRequest) (*emptypb.Empty, error) {
	appointee := requests.Appointee{
		FirstName:   req.Appointee.FirstName,
		LastName:    req.Appointee.LastName,
		Email:       req.Appointee.Email,
		PhoneNumber: req.Appointee.PhoneNumber,
	}

	_, responseCode, err := s.AppointmentUseCase.AppointAttender(req.Appointer.Id, &appointee)
	if err != nil {
		return nil, stdresponse.GetGrpcStatus(responseCode, err.Error())
	} else {
		return &emptypb.Empty{}, nil
	}
}

// AppointChildOfficeHead
func (s *AppointmentServer) AppointChildOfficeHead(ctx context.Context, req *pb.OfficeHeadAppointmentRequest) (*emptypb.Empty, error) {
	appointee := requests.Appointee{
		FirstName:   req.Appointee.FirstName,
		LastName:    req.Appointee.LastName,
		Email:       req.Appointee.Email,
		PhoneNumber: req.Appointee.PhoneNumber,
	}

	_, responseCode, err := s.AppointmentUseCase.AppointChildOfficeHead(req.Appointer.Id,req.ChildOfficeID, &appointee)
	if err != nil {
		return nil, stdresponse.GetGrpcStatus(responseCode, err.Error())
	} else {
		return &emptypb.Empty{}, nil
	}
}

// AppointChildOfficeDeputyHead
func (s *AppointmentServer) AppointChildOfficeDeputyHead(ctx context.Context, req *pb.OfficeHeadAppointmentRequest) (*emptypb.Empty, error) {
	appointee := requests.Appointee{
		FirstName:   req.Appointee.FirstName,
		LastName:    req.Appointee.LastName,
		Email:       req.Appointee.Email,
		PhoneNumber: req.Appointee.PhoneNumber,
	}

	_, responseCode, err := s.AppointmentUseCase.AppointChildOfficeDeputyHead(req.Appointer.Id,req.ChildOfficeID, &appointee)
	if err != nil {
		return nil, stdresponse.GetGrpcStatus(responseCode, err.Error())
	} else {
		return &emptypb.Empty{}, nil
	}
}
