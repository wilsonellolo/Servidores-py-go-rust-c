package main

import (
	"fmt"
	"io/ioutil"

	"github.com/go-redis/redis"
)

//go mod init github.com/my/repo
//go get github.com/go-redis/redis/v7

func main() {
	//CONEXXION A REDIS
	rdb := redis.NewClient(&redis.Options{
		Addr:     "localhost:6379",
		Password: "", // no password set
		DB:       0,  // use default DB
	})
	err := rdb.Set("key", "value", 0).Err()
	if err != nil {
		panic(err)
	}

	//AGREGANDO DATOS A UNA LISTA REDIS LLAMADA queue
	if err := rdb.RPush("queue", "message1").Err(); err != nil {
		panic(err)
	}
	if err := rdb.RPush("queue", "message2").Err(); err != nil {
		panic(err)
	}

	if err := rdb.RPush("queue", "message3").Err(); err != nil {
		panic(err)
	}

	if err := rdb.RPush("queue", "message4").Err(); err != nil {
		panic(err)
	}

	//Obteniendo los datos de la lista llamada queue
	result3, err := rdb.LRange("queue", 0, -1).Result()
	fmt.Println(result3)

	//PARA LEER ARCHIVOS
	datos, err := ioutil.ReadFile("C:\\Users\\Wilson Palma\\Desktop\\fuera.txt")
	if err != nil {
		panic(err)
	}

	fmt.Println(string(datos))
}
