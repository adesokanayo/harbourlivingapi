package graphql

const (
	InvalidUserType UserTypes = iota
	AttendeeUserType
	HostUserType
	SponsorUserType
	AdminUserType
)

type UserTypes int

func ConvertDbToUserRole(v int32) UserRole {
	switch v {
	case 1:
		return UserRoleAttendee
	case 2:
		return UserRoleHost

	case 3:
		return UserRoleSponsor

	case 4:
		return UserRoleAdmin
	}

	return "invalid usertype"
}

func ConvertUserRoleToDb(v UserRole) int {
	switch v {
	case UserRoleAttendee:
		return 1

	case UserRoleHost:
		return 2

	case UserRoleSponsor:
		return 3

	case UserRoleAdmin:
		return 4
	}
	return 0
}

func ConvertDbToEventTypeOption(v int32) EventTypeOption {
	switch v {
	case 1:
		return EventTypeOptionFree
	case 2:
		return EventTypeOptionPaid
	}
	return "invalid eventype"

}

func ConvertEventTypeOptionToDb(v EventTypeOption) int {
	switch v {
	case EventTypeOptionFree:
		return 1

	case EventTypeOptionPaid:
		return 2
	}
	return 0
}

func ConvertDbToEventStatusOption(v int32) EventStatusOption {
	switch v {
	case 1:
		return EventStatusOptionDraft
	case 2:
		return EventStatusOptionPublished

	case 3:
		return EventStatusOptionApproved

	case 4:
		return EventStatusOptionRejected

	case 5:
		return EventStatusOptionCompleted
	}
	return "invalid usertype"

}

func ConvertEventStatusOptionToDb(v EventStatusOption) int {
	switch v {
	case EventStatusOptionDraft:
		return 1

	case EventStatusOptionPublished:
		return 2

	case EventStatusOptionApproved:
		return 3
	case EventStatusOptionRejected:
		return 4
	case EventStatusOptionCompleted:
		return 5
	}
	return 0
}
