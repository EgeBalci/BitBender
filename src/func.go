package main

import "math/rand"
import "time"
import "fmt"
import "os"




func Xor(Data []byte, Key []byte) ([]byte){
	for i := 0; i < len(Data); i++{
		Data[i] = (Data[i] ^ (Key[(i%len(Key))]))
	}
	return Data
}

func GenerateKey(Size int) ([]byte){
	Key := make([]byte, Size)
	rand.Seed(time.Now().UTC().UnixNano())
	for i := 0; i < Size; i++{
		Key[i] = byte(rand.Intn(255))		
	}
	return Key
}


func Add(Array []byte, Value int) ([]byte){
	for i := 0; i < len(Array); i++ {
		Array[i] = (Array[i] + byte(Value))
	}
	return Array
}


func Sub(Array []byte, Value int) ([]byte){
  for i := 0; i < len(Array); i++ {
    Array[i] = (Array[i] - byte(Value))
  }
  return Array
}

func Not(Shellcode []byte) ([]byte){
	for i := 0; i < (len(Shellcode)); i++ {
		Shellcode[i] = ^Shellcode[i]
	}
	return Shellcode
}

func Ror(File []byte, Val uint) ([]byte){

	for i := 0; i < len(File); i++ {
		File[i] = (File[i] >> uint(Val))	
	}

  return File 
}

func Rol(File []byte, Val uint) ([]byte){

	for i := 0; i < len(File); i++ {
		File[i] = (File[i] << uint(Val))	
	}

  return File 
}

func Checksum(File []byte) (int64){
		
	var Checksum int64 = 0
	if len(File) > 1 {
		for i := 0; i < len(File); i++ {
			Checksum += int64(File[i])
		}		
	}else{
		return 0
	}

	return Checksum
}

func ParseError(err error,msg string) {
	if err != nil {
		
		fmt.Println("\n")
		BoldRed.Println("\n[-] ERROR: "+msg+"\n")
		os.Exit(1)
	}
}
