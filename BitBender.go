package main

import ."Benders"
import "io/ioutil"
import "os/exec"
import "strconv"
import "os"

var P Parameters

func main() {


	ARGS := os.Args[1:]

  if len(ARGS) == 0{
    BoldRed.Println(BANNER)
    Green.Println(HELP)
    os.Exit(0)
  }else if len(ARGS) == 1 && (ARGS[0] == "--help" || ARGS[0] == "-h") {
    BoldRed.Println(BANNER)
    Green.Println(HELP)
    os.Exit(0)
  }
  
  for i := 0; i < len(ARGS); i++ {


    if ARGS[i] == "^" {
      P.Mode = "^"
      Value, Err := strconv.Atoi(ARGS[i+1])
      if Err != nil{
        BoldRed.Println("[-] ERROR: Invalid XOR key size !")
        os.Exit(1)
      }
      P.KeySize = Value
    }
	  if ARGS[i] == "^=" {
      P.Mode = "^="
      P.Key = []byte(ARGS[i+1])
    }
    if ARGS[i] == "+" {
      P.Mode = "+"
      Value, Err := strconv.Atoi(ARGS[i+1])
      if Err != nil{
        BoldRed.Println("[-] ERROR: Invalid add value !")
        os.Exit(1)
      }
      P.Plus = Value 
    }
    if ARGS[i] == "-" {
      P.Mode = "-"
      Value, Err := strconv.Atoi(ARGS[i+1])
      if Err != nil{
        BoldRed.Println("[-] ERROR: Invalid subtract value !")
        os.Exit(1)
      }
      P.Minus = Value
    }     
    if ARGS[i] == "!" {
      P.Mode = "!"
    }
    if ARGS[i] == "ror" || ARGS[i] == "--ror" {
      P.Mode = "ror"
      Value, Err := strconv.Atoi(ARGS[i+1])
      if Err != nil{
        BoldRed.Println("[-] ERROR: Invalid retation value !")
        os.Exit(1)
      }
      P.RotValue = uint(Value)
    }
    if ARGS[i] == "rol" || ARGS[i] == "--rol" {
      P.Mode = "rol"
      Value, Err := strconv.Atoi(ARGS[i+1])
      if Err != nil{
        BoldRed.Println("[-] ERROR: Invalid rotation value !")
        os.Exit(1)
      }
      P.RotValue = uint(Value)
    }
    if ARGS[i] == "=" || ARGS[i] == "--checksum" {
      P.Mode = "="
    }
  }

  File, Err := ioutil.ReadFile(ARGS[len(ARGS)-1])
  if Err != nil {
    BoldRed.Println("[-] ERROR : Unable to open input file !")
    Red.Println(Err)
    os.Exit(1)
  }

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
    if Err != nil{
      BoldRed.Println("[-] ERROR : Unable to create output file !")
      Red.Println(Err)
      os.Exit(1)   
    }
    BoldYellow.Println("[*] Writing output...")
    XoredFile.Write(_File)
    XoredFile.Close()
    KeyFile,Err2 := os.Create(string(ARGS[len(ARGS)-1]+".key"))
    if Err2 != nil{
      BoldRed.Println("[-] ERROR : Unable create key file !")
      Red.Println(Err2)
      os.Exit(1)   
    }
    KeyFile.Write(P.Key)
    KeyFile.Close()    
    FileOut, Err3 := exec.Command("sh", "-c", string("xxd -i "+ARGS[len(ARGS)-1]+".xor")).Output()
    if Err2 != nil {
      BoldRed.Println("[-] ERROR : Unable to retrieve file hex output !")
      Red.Println(Err3)
      os.Exit(1)      
    }
    BoldGreen.Println(string(FileOut))
    KeyOut, Err4 := exec.Command("sh", "-c", string("xxd -i "+ARGS[len(ARGS)-1]+".key")).Output()
    if Err3 != nil {
      BoldRed.Println("[-] ERROR : Unable to retrieve key hex output !")
      Red.Println(Err4)
      os.Exit(1)      
    }
    BoldBlue.Println(string(KeyOut))
    os.Exit(0)
  }else if P.Mode == "^=" {
    BoldYellow.Print("[*] Key (ASCII): ")
    BoldBlue.Println(string(P.Key))
    BoldYellow.Println("[*] Ciphering...")
    _File := Xor(File,P.Key)
    XoredFile,Err := os.Create(string(ARGS[len(ARGS)-1]+".xor"))
    if Err != nil{
      BoldRed.Println("[-] ERROR : Unable to create output file !")
      Red.Println(Err)
      os.Exit(1)   
    }
    BoldYellow.Println("[*] Writing output...")
    XoredFile.Write(_File)
    XoredFile.Close()
    KeyFile,Err2 := os.Create(string(ARGS[len(ARGS)-1]+".key"))
    if Err2 != nil{
      BoldRed.Println("[-] ERROR : Unable create key file !")
      Red.Println(Err2)
      os.Exit(1)   
    }
    KeyFile.Write(P.Key)
    KeyFile.Close()
    FileOut, Err3 := exec.Command("sh", "-c", string("xxd -i "+ARGS[len(ARGS)-1]+".xor")).Output()
    if Err3 != nil {
      BoldRed.Println("[-] ERROR : Unable to retrieve file hex output !")
      Red.Println(Err3)
      os.Exit(1)      
    }
    BoldGreen.Println(string(FileOut))
    KeyOut, Err4 := exec.Command("sh", "-c", string("xxd -i "+ARGS[len(ARGS)-1]+".key")).Output()
    if Err3 != nil {
      BoldRed.Println("[-] ERROR : Unable to retrieve key hex output !")
      Red.Println(Err4)
      os.Exit(1)      
    }
    BoldBlue.Println(string(KeyOut))
    os.Exit(0)
  }else if P.Mode == "+" {
    BoldYellow.Print("[*] Increment Value: ")
    BoldBlue.Println(P.Plus)
    BoldYellow.Println("[*] Incrementing bytes...")
    _File := Add(File,P.Plus)
    IncFile,Err := os.Create(string(ARGS[len(ARGS)-1]+".add"))
    if Err != nil{
      BoldRed.Println("[-] ERROR : Unable to create output file !")
      Red.Println(Err)
      os.Exit(1)   
    }
    BoldYellow.Println("[*] Writing output...")
    IncFile.Write(_File)
    IncFile.Close()
    FileOut, Err2 := exec.Command("sh", "-c", string("xxd -i "+ARGS[len(ARGS)-1]+".add")).Output()
    if Err2 != nil {
      BoldRed.Println("[-] ERROR : Unable to retrieve file hex output !")
      Red.Println(Err2)
      os.Exit(1)      
    }
    BoldGreen.Println(string(FileOut))
    os.Exit(0)
  }else if P.Mode == "-" {
    BoldYellow.Print("[*] Decrement Value: ")
    BoldBlue.Println(P.Minus)
    BoldYellow.Println("[*] Decrementing bytes...")
    _File := Sub(File,P.Minus)
    DecFile,Err := os.Create(string(ARGS[len(ARGS)-1]+".sub"))
    if Err != nil{
      BoldRed.Println("[-] ERROR : Unable to create output file !")
      Red.Println(Err)
      os.Exit(1)   
    }
    BoldYellow.Println("[*] Writing output...")
    DecFile.Write(_File)
    DecFile.Close()
    FileOut, Err2 := exec.Command("sh", "-c", string("xxd -i "+ARGS[len(ARGS)-1]+".sub")).Output()
    if Err2 != nil {
      BoldRed.Println("[-] ERROR : Unable to retrieve file hex output !")
      Red.Println(Err2)
      os.Exit(1)      
    }
    BoldGreen.Println(string(FileOut))
    os.Exit(0)
  }else if P.Mode == "!" {
    BoldYellow.Println("[*] Reversing bytes...")
    _File := Not(File)
    NotFile,Err := os.Create(string(ARGS[len(ARGS)-1]+".not"))
    if Err != nil{
      BoldRed.Println("[-] ERROR : Unable to create output file !")
      Red.Println(Err)
      os.Exit(1)   
    }
    BoldYellow.Println("[*] Writing output...")
    NotFile.Write(_File)
    NotFile.Close()
    FileOut, Err2 := exec.Command("sh", "-c", string("xxd -i "+ARGS[len(ARGS)-1]+".not")).Output()
    if Err2 != nil {
      BoldRed.Println("[-] ERROR : Unable to retrieve file hex output !")
      Red.Println(Err2)
      os.Exit(1)      
    }
    BoldGreen.Println(string(FileOut))
    os.Exit(0)    
  }else if P.Mode == "ror" {
    BoldYellow.Print("[*] Rotation Value : ")
    BoldBlue.Print(P.RotValue)
    BoldBlue.Println(">>")
    BoldYellow.Println("[*] Rotating bytes...")
    _File := Ror(File,P.RotValue)
    RorFile,Err := os.Create(string(ARGS[len(ARGS)-1]+".ror"))
    if Err != nil{
      BoldRed.Println("[-] ERROR : Unable to create output file !")
      Red.Println(Err)
      os.Exit(1)   
    }
    BoldYellow.Print("[*] Writing output...")
    RorFile.Write(_File)
    RorFile.Close()
    FileOut, Err2 := exec.Command("sh", "-c", string("xxd -i "+ARGS[len(ARGS)-1]+".ror")).Output()
    if Err2 != nil {
      BoldRed.Println("[-] ERROR : Unable to retrieve file hex output !")
      Red.Println(Err2)
      os.Exit(1)      
    }
    BoldGreen.Println(string(FileOut))
    os.Exit(0)    
  }else if P.Mode == "rol" {
    BoldYellow.Print("[*] Rotation Value : ")
    BoldBlue.Print("<<")
    BoldBlue.Println(P.RotValue)
    BoldYellow.Println("[*] Rotating bytes...")
    _File := Ror(File,P.RotValue)
    RolFile,Err := os.Create(string(ARGS[len(ARGS)-1]+".rol"))
    if Err != nil{
      BoldRed.Println("[-] ERROR : Unable to create output file !")
      Red.Println(Err)
      os.Exit(1)   
    }
    BoldYellow.Print("[*] Writing output...")
    RolFile.Write(_File)
    RolFile.Close()
    FileOut, Err2 := exec.Command("sh", "-c", string("xxd -i "+ARGS[len(ARGS)-1]+".rol")).Output()
    if Err2 != nil {
      BoldRed.Println("[-] ERROR : Unable to retrieve file hex output !")
      Red.Println(Err2)
      os.Exit(1)      
    }
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