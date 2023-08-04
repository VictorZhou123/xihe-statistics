package app

const (
	ErrorEmptyGitLabProjectIdPage = "empty page of project id"

	ErrorAccessOverAllowed = "error_access_over_allowed"

	ErrorAccessConcurrentUpdating = "error_access_concurrent_updating"
)

type errorEmptyGitLabProjectIdPage struct {
	error
}

func IsErrorEmptyProjectIdPage(err error) bool {
	_, ok := err.(errorEmptyGitLabProjectIdPage)

	return ok
}
