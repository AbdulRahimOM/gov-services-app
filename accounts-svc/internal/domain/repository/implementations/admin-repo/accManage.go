package adminrepo

import (
	"fmt"

	"github.com/AbdulRahimOM/gov-services-app/accounts-svc/internal/domain/dto/request"
	commondto "github.com/AbdulRahimOM/gov-services-app/internal/common-dto"
)

// AdminGetAdmins
func (ur AdminRepository) AdminGetAdmins(sc *request.AdminSearchCriteria) (*[]commondto.Admin, error) {
	var admins []commondto.Admin
	result := ur.DB.Raw(`
					SELECT 
						id, f_name, l_name, email, address, pincode, phone_number, designation, dept_id, rank_id 
					FROM 
						admins
					WHERE
						(rank_id = ? OR ? = 0) AND
						(dept_id = ? OR ? = 0) AND
						(f_name LIKE ?) AND
						(l_name LIKE ?) AND
						(email LIKE ?) AND
						(phone_number LIKE ?) AND
						(designation LIKE ?)
					`,
					sc.RankID,sc.RankID,
					sc.DeptID,sc.DeptID,
					"%"+sc.FirstName+"%",
					"%"+sc.LastName+"%",
					"%"+sc.Email+"%",
					"%"+sc.PhoneNumber+"%",
					"%"+sc.Designation+"%",
					).Scan(&admins)
	if result.Error != nil {
		return nil, fmt.Errorf("@db: failed to get admins: %v", result.Error)
	}
	return &admins, nil
}
