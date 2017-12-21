package main
import (
	"os"
	"log"
)

func main(){
	_, err:=os.Open("nothing")
	if err!=nil{
		log.Fatal(err)
	}
}