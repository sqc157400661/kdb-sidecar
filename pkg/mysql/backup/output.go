package backup

type Output interface {
	Put()
	Head()
	Get()
}
