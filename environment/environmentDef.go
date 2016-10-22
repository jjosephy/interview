package environment

import (
    "github.com/jjosephy/interview/repository"
)

type Environment struct {
    Path string
    Type string
    Repository repository.InterviewRepository
}

func NewEnvironment(path string, env string, repo repository.InterviewRepository) (*Environment) {
    return &Environment{
        Path: path,
        Type: env,
        Repository: repo,
    }
}
