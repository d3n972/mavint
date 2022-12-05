package main

import (
	"fmt"
	"github.com/d3n972/mavint/domain/repository"
	"time"
)

func main() {
	q := repository.NewAndSpecification(
		repository.NewEqualsSpecification("test", 1),
		repository.NewRelationSpecification(
			"watch_until", repository.GreaterOrEq, time.Now().UTC().Format(time.RFC3339)),
		repository.NewRelationSpecification(
			"asdasd", repository.GreaterOrEq, time.Now().UTC().Format(time.RFC3339)),
		repository.NewRelationSpecification(
			"reads", repository.GreaterOrEq, time.Now().UTC().Format(time.RFC3339)),

		repository.NewOrSpecification(
			repository.NewRelationSpecification(
				"as", repository.GreaterOrEq, 24.5123),
			repository.NewOrSpecification(repository.NewRelationSpecification(
				"asdasdad", repository.GreaterOrEq, time.Now().UTC().Format(time.RFC3339)),
			),
		))
	fmt.Printf("%s\n", q.Query())
}
