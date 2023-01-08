package user

import "context"

type (
	UserType         int
	SubscriptionType int
)

const (
	unknownUserType UserType = iota
	lead
	customer
	churned
	lostLead
)

const (
	unknownSubscritpionType SubscriptionType = iota
	basic
	permium
	exclusive
)

type UserAddRequest struct {
	UserType       UserType
	Email          string
	SubType        SubscriptionType
	PaymentDetails PaymentDetails
}

type UserModifyRequest struct {
	ID             string
	UserType       UserType
	Email          string
	SubType        SubscriptionType
	PaymentDetails PaymentDetails
}

type User struct {
	ID             string
	PaymentDetails PaymentDetails
}

type PaymentDetails struct {
	stripeTokenID string
}

type UserManager interface {
	AddUser(ctx context.Context, request UserAddRequest) (User, error)
	ModifyUser(ctx context.Context, request UserModifyRequest) (User, error)
}

type LeadRequest struct {
	Email string
}

type Lead struct {
	ID string
}

type LeadCreator interface {
	CreateLead(ctx context.Context, request LeadRequest) (Lead, error)
}

type Customer struct {
	LeadID string
	UserId string
}

type LeadConvertor interface {
	Convert(ctx context.Context, subType SubscriptionType) (Customer, error)
}

func (l Lead) Convert(ctx context.Context, subType SubscriptionType) (Customer, error) {
	// TODO: implement me
	panic("implement me")
}
