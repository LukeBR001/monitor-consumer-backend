package service

import "src/pkg/repository"

func GetTweetsService() string {
	return repository.GetTweetsRepo()
}
