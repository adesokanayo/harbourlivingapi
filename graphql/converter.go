package graphql

const (
	InvalidUserType UserTypes = iota
	AttendeeUserType
	HostUserType
	SponsorUserType
	AdminUserType
)

type UserTypes int

func ConvertDbToUserRole(v int32) UserTypeOptions {
	switch v {
	case 1:
		return UserTypeOptionsAttendee

	case 2:
		return UserTypeOptionsHost

	case 3:
		return UserTypeOptionsSponsor

	case 4:
		return UserTypeOptionsAdmin
	}

	return "invalid usertype"
}

func ConvertUserTypeOptionsToDb(v UserTypeOptions) int {
	switch v {
	case UserTypeOptionsAttendee:
		return 1

	case UserTypeOptionsHost:
		return 2

	case UserTypeOptionsSponsor:
		return 3

	case UserTypeOptionsAdmin:
		return 4
	}
	return 0
}

func ConvertDbToEventTypeOptions(v int32) EventTypeOptions {
	switch v {
	case 1:
		return EventTypeOptionsFree
	case 2:
		return EventTypeOptionsPaid
	}
	return "invalid eventype"

}

func ConvertEventTypeOptionsToDb(v EventTypeOptions) int {
	switch v {
	case EventTypeOptionsFree:
		return 1

	case EventTypeOptionsPaid:
		return 2
	}
	return 0
}

func ConvertDbToStatusOptions(v int32) StatusOptions {
	switch v {
	case 1:
		return StatusOptionsDraft
	case 2:
		return StatusOptionsPublished

	case 3:
		return StatusOptionsApproved

	case 4:
		return StatusOptionsRejected

	case 5:
		return StatusOptionsCompleted
	}
	return "invalid usertype"

}

func ConvertStatusOptionsToDb(v StatusOptions) int {
	switch v {
	case StatusOptionsDraft:
		return 1

	case StatusOptionsPublished:
		return 2

	case StatusOptionsApproved:
		return 3
	case StatusOptionsRejected:
		return 4
	case StatusOptionsCompleted:
		return 5
	}
	return 0
}
