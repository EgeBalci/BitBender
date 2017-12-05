package main

import "io/ioutil"
import "os/exec"
import "strconv"
import "fmt"
import "os"

var P Parameters

func main() {

	ARGS := os.Args[1:]

  if len(ARGS) <= 1 || (ARGS[0] == "--help" || ARGS[0] == "-h"){
    fmt.Println(HELP)
    os.Exit(0)
  }
  
  for i := 0; i < len(ARGS); i++ {

    if ARGS[i] == "^" {
      P.Mode = "^"
      Value, Err := strconv.Atoi(ARGS[i+1])
      ParseError(Err,"Invalid XOR key size !")
      P.KeySize = Value
    }
	  if ARGS[i] == "^=" {
      P.Mode = "^="
      P.Key = []byte(ARGS[i+1])
    }
    if ARGS[i] == "+" {
      P.Mode = "+"
      Value, Err := strconv.Atoi(ARGS[i+1])
      ParseError(Err,"Invalid add value !")
      P.Plus = Value 
    }
    if ARGS[i] == "-" {
      P.Mode = "-"
      Value, Err := strconv.Atoi(ARGS[i+1])
      ParseError(Err,"Invalid subtract value !")
      P.Minus = Value
    }     
    if ARGS[i] == "!" {
      P.Mode = "!"
    }
    if ARGS[i] == "ror" || ARGS[i] == "--ror" {
      P.Mode = "ror"
      Value, Err := strconv.Atoi(ARGS[i+1])
      ParseError(Err,"Invalid retation value !")
      P.RotValue = uint(Value)
    }
    if ARGS[i] == "rol" || ARGS[i] == "--rol" {
      P.Mode = "rol"
      Value, Err := strconv.Atoi(ARGS[i+1])
      ParseError(Err,"Invalid rotation value !")
      P.RotValue = uint(Value)
    }
    if ARGS[i] == "=" || ARGS[i] == "--checksum" {
      P.Mode = "="
    }
  }

  File, Err := ioutil.ReadFile(ARGS[len(ARGS)-1])
  ParseError(Err,"Unable to open input file !")

  if P.Mode == "^" {
    BoldYellow.Print("[*] Key Size: ")
    BoldBlue.Println(P.KeySize)
    BoldYellow.Println("[*] Generating XOR key...")
    P.Key = GenerateKey(P.KeySize)
    BoldYellow.Print("[*] Key (ASCII): ")
    BoldBlue.Println(string(P.Key))
    BoldYellow.Println("[*] Ciphering...")
    _File := Xor(File,P.Key)
    XoredFile,Err := os.Create(string(ARGS[len(ARGS)-1]+".xor"))
    ParseError(Err,"Unable to create output file !")
    BoldYellow.Println("[*] Writing output...")
    XoredFile.Write(_File)
    XoredFile.Close()
    KeyFile,Err2 := os.Create(string(ARGS[len(ARGS)-1]+".key"))
    ParseError(Err2,"Unable create key file !")
    KeyFile.Write(P.Key)
    KeyFile.Close()    
    FileOut, Err3 := exec.Command("sh", "-c", string("xxd -i "+ARGS[len(ARGS)-1]+".xor")).Output()
    ParseError(Err3,"Unable to retrieve file hex output !")
    BoldGreen.Println(string(FileOut))
    KeyOut, Err4 := exec.Command("sh", "-c", string("xxd -i "+ARGS[len(ARGS)-1]+".key")).Output()
    ParseError(Err4,"Unable to retrieve key hex output !")
    BoldBlue.Println(string(KeyOut))
    os.Exit(0)
  }else if P.Mode == "^=" {
    BoldYellow.Print("[*] Key (ASCII): ")
    BoldBlue.Println(string(P.Key))
    BoldYellow.Println("[*] Ciphering...")
    _File := Xor(File,P.Key)
    XoredFile,Err := os.Create(string(ARGS[len(ARGS)-1]+".xor"))
    ParseError(Err,"Unable to create output file !")
    BoldYellow.Println("[*] Writing output...")
    XoredFile.Write(_File)
    XoredFile.Close()
    KeyFile,Err2 := os.Create(string(ARGS[len(ARGS)-1]+".key"))
    ParseError(Err2,"Unable create key file !")
    KeyFile.Write(P.Key)
    KeyFile.Close()
    FileOut, Err3 := exec.Command("sh", "-c", string("xxd -i "+ARGS[len(ARGS)-1]+".xor")).Output()
    ParseError(Err3,"Unable to retrieve file hex output !")
    BoldGreen.Println(string(FileOut))
    KeyOut, Err4 := exec.Command("sh", "-c", string("xxd -i "+ARGS[len(ARGS)-1]+".key")).Output()
    ParseError(Err4,"Unable to retrieve key hex output !")
    BoldBlue.Println(string(KeyOut))
    os.Exit(0)
  }else if P.Mode == "+" {
    BoldYellow.Print("[*] Increment Value: ")
    BoldBlue.Println(P.Plus)
    BoldYellow.Println("[*] Incrementing bytes...")
    _File := Add(File,P.Plus)
    IncFile,Err := os.Create(string(ARGS[len(ARGS)-1]+".add"))
    ParseError(Err,"Unable to create output file !")
    BoldYellow.Println("[*] Writing output...")
    IncFile.Write(_File)
    IncFile.Close()
    FileOut, Err2 := exec.Command("sh", "-c", string("xxd -i "+ARGS[len(ARGS)-1]+".add")).Output()
    ParseError(Err2,"Unable to retrieve file hex output !")
    BoldGreen.Println(string(FileOut))
    os.Exit(0)
  }else if P.Mode == "-" {
    BoldYellow.Print("[*] Decrement Value: ")
    BoldBlue.Println(P.Minus)
    BoldYellow.Println("[*] Decrementing bytes...")
    _File := Sub(File,P.Minus)
    DecFile,Err := os.Create(string(ARGS[len(ARGS)-1]+".sub"))
    ParseError(Err,"Unable to create output file !")    
    BoldYellow.Println("[*] Writing output...")
    DecFile.Write(_File)
    DecFile.Close()
    FileOut, Err2 := exec.Command("sh", "-c", string("xxd -i "+ARGS[len(ARGS)-1]+".sub")).Output()
    ParseError(Err2,"Unable to retrieve file hex output !")
    BoldGreen.Println(string(FileOut))
    os.Exit(0)
  }else if P.Mode == "!" {
    BoldYellow.Println("[*] Reversing bytes...")
    _File := Not(File)
    NotFile,Err := os.Create(string(ARGS[len(ARGS)-1]+".not"))
    ParseError(Err,"Unable to create output file !")
    BoldYellow.Println("[*] Writing output...")
    NotFile.Write(_File)
    NotFile.Close()
    FileOut, Err2 := exec.Command("sh", "-c", string("xxd -i "+ARGS[len(ARGS)-1]+".not")).Output()
    ParseError(Err2,"Unable to retrieve file hex output !")
    BoldGreen.Println(string(FileOut))
    os.Exit(0)    
  }else if P.Mode == "ror" {
    BoldYellow.Print("[*] Rotation Value : ")
    BoldBlue.Print(P.RotValue)
    BoldBlue.Println(">>")
    BoldYellow.Println("[*] Rotating bytes...")
    _File := Ror(File,P.RotValue)
    RorFile,Err := os.Create(string(ARGS[len(ARGS)-1]+".ror"))
    ParseError(Err,"Unable to create output file !")
    BoldYellow.Print("[*] Writing output...")
    RorFile.Write(_File)
    RorFile.Close()
    FileOut, Err2 := exec.Command("sh", "-c", string("xxd -i "+ARGS[len(ARGS)-1]+".ror")).Output()
    ParseError(Err2,"Unable to retrieve file hex output !")
    BoldGreen.Println(string(FileOut))
    os.Exit(0)    
  }else if P.Mode == "rol" {
    BoldYellow.Print("[*] Rotation Value : ")
    BoldBlue.Print("<<")
    BoldBlue.Println(P.RotValue)
    BoldYellow.Println("[*] Rotating bytes...")
    _File := Ror(File,P.RotValue)
    RolFile,Err := os.Create(string(ARGS[len(ARGS)-1]+".rol"))
    ParseError(Err,"Unable to create output file !")
    BoldYellow.Print("[*] Writing output...")
    RolFile.Write(_File)
    RolFile.Close()
    FileOut, Err2 := exec.Command("sh", "-c", string("xxd -i "+ARGS[len(ARGS)-1]+".rol")).Output()
    ParseError(Err2,"Unable to retrieve file hex output !")
    BoldGreen.Println(string(FileOut))
    os.Exit(0)    
  }else if P.Mode == "=" {
    BoldYellow.Println("[*] Calculating checsum...")
    var _Checksum int64 = Checksum(File)
    BoldGreen.Print("[#] ")
    BoldYellow.Print("Checksum : ")
    BoldRed.Print("0x",strconv.FormatInt(_Checksum, 16),"\n")
    os.Exit(0)   
  }else{
    BoldRed.Println("[-] ERROR : Invalid operation mode !")
    os.Exit(1)
  }



}