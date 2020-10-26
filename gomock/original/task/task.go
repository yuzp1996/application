package task

type Task interface {
	Do(int) (string, error)
}
