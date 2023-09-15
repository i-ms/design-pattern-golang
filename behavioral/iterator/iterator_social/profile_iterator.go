package iteratorsocial

type ProfileIterator interface {
	Next() Profile
	HasNext() bool
}
