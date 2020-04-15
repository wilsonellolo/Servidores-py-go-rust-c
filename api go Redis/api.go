package main
import (
        "fmt"
        "io/ioutil"
        "github.com/go-redis/redis"
		"strings"
		"strconv"
		"time"
)
//go mod init github.com/my/repo
//go get github.com/go-redis/redis/v7
func main() {
		 numero :=1
		 //CONEXXION A REDIS
		 rdb := redis.NewClient(&redis.Options{
			Addr:     "localhost:6379",
			Password: "", // no password set
			DB:       0,  // use default DB
		})
		for numero<2{
			//PARA LEER ARCHIVOS
			datos, err := ioutil.ReadFile("/proc/meminfo")
			if err != nil {
					panic(err)
			}
			datos1 := string(datos)
			//para leer memoria ram
			datosMemoria := strings.Split(datos1, "\n")
			//Memoria total
			datoMemoriaTotal1 := datosMemoria[0]
			datoMemoriaTotal2 := strings.Replace(datoMemoriaTotal1, " ", "", -1)
			datoMemoriaTotal3 :=  strings.Split(datoMemoriaTotal2, ":")
			datoMemoriaTotal4 :=  datoMemoriaTotal3[1]
			datoMemoriaTotal5 :=  strings.Split(datoMemoriaTotal4, "k")
			datoMemoriaTotal6 :=  datoMemoriaTotal5[0]
			//Memoria libre
			datoMemoriaLibre1 := datosMemoria[1]
			datoMemoriaLibre2 := strings.Replace(datoMemoriaLibre1, " ", "", -1)
			datoMemoriaLibre3 :=  strings.Split(datoMemoriaLibre2, ":")
			datoMemoriaLibre4 :=  datoMemoriaLibre3[1]
			datoMemoriaLibre5 :=  strings.Split(datoMemoriaLibre4, "k")
			datoMemoriaLibre6 :=  datoMemoriaLibre5[0]
			totalFinal, _ := strconv.ParseFloat(datoMemoriaTotal6, 64)
			libreFinal, _ := strconv.ParseFloat(datoMemoriaLibre6, 64)
			MemoriaUsada :=((totalFinal-libreFinal)/totalFinal)*100
			tiempo:=time.Now().UnixNano()
			fmt.Println("momento->",tiempo)
			fmt.Println("ram a ",MemoriaUsada,"% de uso")
			if err := rdb.RPush("tiempo", tiempo).Err(); err != nil {
				panic(err)
			}
			if err := rdb.RPush("memoria", MemoriaUsada).Err(); err != nil {
				panic(err)
			}
			/*LLENANDO CPU PORCNETAJE*/
			//PARA LEER ARCHIVOS
			datosStatArchivo, err := ioutil.ReadFile("/proc/stat")
			if err != nil {
					panic(err)
			}
			datosStat:=string(datosStatArchivo)
			//para leer tiempos del cpu
			datosDelCPU := strings.Split(datosStat, "\n")
			//obteniendo la primerla liena
			datoTiemposCPU := datosDelCPU[0]
			arregloDeDatos := strings.Split(datoTiemposCPU, " ")
			contador:=0
			sumaTotal:=0.0
			oceo,_:=strconv.ParseFloat(string(arregloDeDatos[5]), 64)
			for contador< len(arregloDeDatos){
				if contador>=2 {
					sumatot,_:=strconv.ParseFloat(string(arregloDeDatos[contador]), 64)
					sumaTotal += sumatot
				}
				contador++
			}
			tiempoUsado:=sumaTotal-oceo
			porcentajedecpu:=(tiempoUsado/sumaTotal)*100
			fmt.Println("cpu a ", porcentajedecpu,"% de uso")
			if err := rdb.RPush("cpu", porcentajedecpu).Err(); err != nil {
				panic(err)
			}
			time.Sleep(5 * time.Second)
		}
}





package main	

import (
    "fmt"
    "time"
)


func main() {

	time.Sleep(5 * time.Second)
}


/*cosas para redis*/

err := rdb.Set("key", "value", 0).Err()
if err != nil {
		panic(err)
}

//AGREGANDO DATOS A UNA LISTA REDIS LLAMADA queue
if err := rdb.RPush("queue", "message1").Err(); err != nil {
		panic(err)
}


//Obteniendo los datos de la lista llamada queue
result3, err := rdb.LRange("queue", 0, -1).Result()
fmt.Println(result3)


/**********************/

package main

import (
        "fmt"
        "io/ioutil"
        "github.com/go-redis/redis"
                "strings"
                "strconv"
                "time"
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


		//Obteniendo los datos de la lista llamada queue
		result3, err := rdb.LRange("queue", 0, -1).Result()


		fmt.Println(result3)
		//PARA LEER ARCHIVOS
		datos, err := ioutil.ReadFile("/proc/meminfo")
		if err != nil {
				panic(err)
		}
				datos1 := string(datos)

		//para leer memoria ram
				datosMemoria := strings.Split(datos1, "\n")

		//Memoria total
		datoMemoriaTotal1 := datosMemoria[0]
		datoMemoriaTotal2 := strings.Replace(datoMemoriaTotal1, " ", "", -1)
		datoMemoriaTotal3 :=  strings.Split(datoMemoriaTotal2, ":")
		datoMemoriaTotal4 :=  datoMemoriaTotal3[1]
		datoMemoriaTotal5 :=  strings.Split(datoMemoriaTotal4, "k")
		datoMemoriaTotal6 :=  datoMemoriaTotal5[0]


        //Memoria libre
        datoMemoriaLibre1 := datosMemoria[1]
        datoMemoriaLibre2 := strings.Replace(datoMemoriaLibre1, " ", "", -1)
        datoMemoriaLibre3 :=  strings.Split(datoMemoriaLibre2, ":")
        datoMemoriaLibre4 :=  datoMemoriaLibre3[1]
        datoMemoriaLibre5 :=  strings.Split(datoMemoriaLibre4, "k")
        datoMemoriaLibre6 :=  datoMemoriaLibre5[0]

        fmt.Println("la memoria total es:"+datoMemoriaTotal6)
                fmt.Println("la memoria libre es:"+datoMemoriaLibre6)

                totalFinal, _ := strconv.ParseFloat(datoMemoriaTotal6, 64)

                libreFinal, _ := strconv.ParseFloat(datoMemoriaLibre6, 64)

                MemoriaUsada :=((totalFinal-libreFinal)/totalFinal)*100

                tiempo:=time.Now().UnixNano()
                fmt.Println(tiempo)
                fmt.Println(MemoriaUsada)

                if err := rdb.RPush("tiempo", tiempo).Err(); err != nil {
                        panic(err)
                }
                if err := rdb.RPush("memoria", MemoriaUsada).Err(); err != nil {
                        panic(err)
                }

}
