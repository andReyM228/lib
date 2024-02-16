package bus

const (
	//user-service
	SubjectUserServiceGetUserByID = "subject.user.service.get.user.by.id"
	SubjectUserServiceLoginUser   = "subject.user.service.login.user"
	SubjectUserServiceCreateUser  = "subject.user.service.create.user"
	SubjectUserServiceGetCarByID  = "subject.user.service.get.car.by.id"
)

const (
	//tx-service
	SubjectTxServiceIssue    = "subject.tx.service.issue"
	SubjectTxServiceWithdraw = "subject.tx.service.withdraw"
)
