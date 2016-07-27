package database

type FakeSession struct{}

func (self FakeSession) Clone() Session {
	return MgoSession{}
}

func (self FakeSession) Close() {}

func (self FakeSession) Schema(name string) Schema {
	return MgoSchema{}
}

func (self FakeSession) DefaultSchema() Schema {
	return MgoSchema{}
}
