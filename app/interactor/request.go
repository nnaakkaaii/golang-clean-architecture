package interactor

// UserRequest は...
// interactorに取って都合の良いデータ構造であり (=controllerの事情に依存しない)
// controllerにとっても受け渡しに十分な情報であるべき
type UserRequest struct {
	FirstName  string
	LastName   string
	BirthYear  int
	BirthMonth int
	BirthDay   int
}
