package posts

import (
	"github.com/AbdulRahimOM/gov-services-app/internal/project/data/models"
)

var (
	SuperAdmin = models.Post{
		ID:                     1,
		PostName:               "Super Admin",
		OfficeID:               1,
		MaxNumberOfPosts:       1,
		CanAppointSubordinates: true,
		Rank:                   1,
	}

	StateNodalOfficer = models.Post{
		ID:                     2,
		PostName:               "State Nodal Officer",
		OfficeID:               1,
		MaxNumberOfPosts:       1,
		CanAppointSubordinates: true,
		Rank:                   2,
	}

	StateAsstNodalOfficer = models.Post{
		ID:                     3,
		PostName:               "State Asst Nodal Officer",
		OfficeID:               1,
		MaxNumberOfPosts:       1,
		CanAppointSubordinates: true,
		Rank:                   2,
	}

	StatePoliceNodalOfficer = models.Post{
		ID:                     4,
		PostName:               "State Police Nodal Officer",
		OfficeID:               3,
		MaxNumberOfPosts:       1,
		CanAppointSubordinates: true,
		Rank:                   3,
	}

	StatePoliceAsstNodalOfficer = models.Post{
		ID:                     5,
		PostName:               "State Police Deputy Nodal Officer",
		OfficeID:               3,
		MaxNumberOfPosts:       1,
		CanAppointSubordinates: true,
		Rank:                   3,
	}

	StateFireSafetyNodalOfficer = models.Post{
		ID:                     6,
		PostName:               "State Fire and Safety Nodal Officer",
		OfficeID:               4,
		MaxNumberOfPosts:       1,
		CanAppointSubordinates: true,
		Rank:                   3,
	}

	StateFireSafetyAsstNodalOfficer = models.Post{
		ID:                     7,
		PostName:               "State Fire and Safety Deputy Nodal Officer",
		OfficeID:               4,
		MaxNumberOfPosts:       1,
		CanAppointSubordinates: true,
		Rank:                   3,
	}

	StateHealthNodalOfficer = models.Post{
		ID:                     8,
		PostName:               "State Health Nodal Officer",
		OfficeID:               5,
		MaxNumberOfPosts:       1,
		CanAppointSubordinates: true,
		Rank:                   3,
	}

	StateHealthAsstNodalOfficer = models.Post{
		ID:                     9,
		PostName:               "State Health Deputy Nodal Officer",
		OfficeID:               5,
		MaxNumberOfPosts:       1,
		CanAppointSubordinates: true,
		Rank:                   3,
	}

	StateEmergencyServiceNOPost = models.Post{
		ID:                     10,
		PostName:               "State Emergency Service Nodal Officer",
		OfficeID:               6,
		MaxNumberOfPosts:       1,
		CanAppointSubordinates: true,
		Rank:                   3,
	}

	StateEmergencyServiceAddlNOPost = models.Post{
		ID:                     11,
		PostName:               "State Emergency Service Deputy Nodal Officer",
		OfficeID:               6,
		MaxNumberOfPosts:       1,
		CanAppointSubordinates: true,
		Rank:                   3,
	}

	StateKSEBNodalOfficer = models.Post{
		ID:                     12,
		PostName:               "State KSEB Nodal Officer",
		OfficeID:               7,
		MaxNumberOfPosts:       1,
		CanAppointSubordinates: true,
		Rank:                   3,
	}

	StateKSEBAsstNodalOfficer = models.Post{
		ID:                     13,
		PostName:               "State KSEB Deputy Nodal Officer",
		OfficeID:               7,
		MaxNumberOfPosts:       1,
		CanAppointSubordinates: true,
		Rank:                   3,
	}

	StateKWANodalOfficer = models.Post{
		ID:                     14,
		PostName:               "State KWA Nodal Officer",
		OfficeID:               8,
		MaxNumberOfPosts:       1,
		CanAppointSubordinates: true,
		Rank:                   3,
	}

	StateKWAAsstNodalOfficer = models.Post{
		ID:                     14,
		PostName:               "State KWA Deputy Nodal Officer",
		OfficeID:               8,
		MaxNumberOfPosts:       1,
		CanAppointSubordinates: true,
		Rank:                   3,
	}
)
