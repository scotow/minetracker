package runner

type Runner interface {
	Run(string) (string, error)
}
