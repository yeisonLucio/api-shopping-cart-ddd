package infrastructure

type MysqlUser struct {
}

func NewMysqlUser() *MysqlUser {
	return &MysqlUser{}
}

func (mu *MysqlUser) Create() error {
	return nil
}
