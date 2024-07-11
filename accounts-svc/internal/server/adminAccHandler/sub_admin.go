package adminAccHandler

// AdminAddSubAdmin
// func (s *AdminAccountsServer) AdminAddSubAdmin(c context.Context, req *pb.AdminAddSubAdminRequest) (*pb.AdminAddSubAdminResponse, error) {
// 	newSubAdminID, responseCode, err := s.AdminUseCase.AdminAddSubAdmin(&request.AdminAddSubAdmin{
// 		AdminID: req.AdminId,
// 		NewSubAdmin: request.NewSubAdmin{
// 			FirstName:   req.NewSubAdmin.FirstName,
// 			LastName:    req.NewSubAdmin.LastName,
// 			Email:       req.NewSubAdmin.Email,
// 			PhoneNumber: req.NewSubAdmin.PhoneNumber,
// 			// DeptID:      req.NewSubAdmin.DeptId,
// 			// Designation: req.NewSubAdmin.Designation,
// 			// RankID:      req.NewSubAdmin.RankId,
// 			// OfficeID:    req.NewSubAdmin.OfficeId,
// 			PostID: req.NewSubAdmin.PostID,
// 		},
// 	})
// 	if err != nil {
// 		return nil, stdresponse.GetGrpcStatus(responseCode, err.Error())
// 	} else {
// 		return &pb.AdminAddSubAdminResponse{
// 			NewSubAdminID: newSubAdminID,
// 		}, nil
// 	}
// }
