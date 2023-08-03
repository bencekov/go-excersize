package api

type ServiceInterface interface {
	RemoveVowels(string) (string, error)
	CounterAdd()
	GetCounter() int
}
