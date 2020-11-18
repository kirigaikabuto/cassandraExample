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
	s := "sdsdds"
	var updateExample = &moviews.MovieUpdate{
		Id:   1,
		Name: &s,
	}
	fmt.Println(updateExample)
	updatedData, err := cassandraStore.Update(updateExample)
	if err != nil {
		log.Fatal(err)
	}
	fmt.Println(updatedData)
	//example := &moviews.Movie{
	//	Name:   "movie3",
	//	Rating: 12,
	//}
	//newExample, err := cassandraStore.Create(example)
	//if err != nil {
	//	log.Fatal(err)
	//}
	//fmt.Println(newExample)
	data, err := cassandraStore.List()
	if err != nil {
		log.Fatal(err)
	}
	fmt.Println(data)
}
