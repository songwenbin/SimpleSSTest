package loginserver

type Contract struct {
	ContractAddress string
	Price           int
	PublicKey       string
}

type ClearContent struct {
	Content   UserContent
	PublicKey string
	Sig       string
}

type UserContent struct {
	TimeStamp string
	Action    string
}
