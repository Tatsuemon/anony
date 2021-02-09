package testutils

// UserServiceMock
type UserServiceMock struct {
	FakeExistsID             func(id string) (bool, error)
	FakeExistsName           func(name string) (bool, error)
	FakeExistsEmail          func(email string) (bool, error)
	FakeExistsDuplicatedUser func(name, email string) (bool, error)
}

func (m UserServiceMock) ExistsID(id string) (bool, error) {
	return m.FakeExistsID(id)
}
func (m UserServiceMock) ExistsName(name string) (bool, error) {
	return m.FakeExistsName(name)
}
func (m UserServiceMock) ExistsEmail(email string) (bool, error) {
	return m.FakeExistsEmail(email)
}
func (m UserServiceMock) ExistsDuplicatedUser(name, email string) (bool, error) {
	return m.FakeExistsDuplicatedUser(name, email)
}
