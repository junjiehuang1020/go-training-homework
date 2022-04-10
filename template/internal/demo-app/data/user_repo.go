package data

type UserRepository struct {
	id      int32
	name    string
	email   string
	isVip   bool
	addr    string
	company string
}

func (u *UserRepository) FindUserByid() {

}
