package main

import (
	"fmt"
	"github.com/kirigaikabuto/cassandraExample/moviews"
	"log"
)

func main() {
	config := moviews.CassandraConfig{
		Nodes:    []string{"localhost"},
		Database: "keyspace_name",
	}

	cassandraStore, err := moviews.NewCassandraMovieStore(config)
	if err != nil {
		log.Fatal(err)
	}
	example := &moviews.Movie{
		Name:   "movie1",
		Rating: 3,
	}
	newExample, err := cassandraStore.Create(example)
	if err != nil {
		log.Fatal(err)
	}
	fmt.Println(newExample)
	data, err := cassandraStore.List()
	if err != nil {
		log.Fatal(err)
	}
	fmt.Println(data)
}
