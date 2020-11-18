package main

import (
	"fmt"
	"github.com/kirigaikabuto/cassandraExample/moviews"
	"log"
)

func main() {
	config := moviews.CassandraConfig{
		Nodes:    []string{"localhost"},
		Database: "cassandra_example",
	}
	cassandraStore,err := moviews.NewCassandraMovieStore(config)
	if err != nil{
		log.Fatal(err)
	}
	data,err := cassandraStore.List()
	if err != nil{
		log.Fatal(err)
	}
	fmt.Println(data)
}
