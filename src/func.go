package main

import "crypto/rc4"
import "math/rand"
import "os/exec"
import "time"
import "fmt"
import "os"

func RC4(data []byte, key []byte) []byte {
	BoldYellow.Println("[*] Ciphering with RC4...")
	c, e := rc4.NewCipher(key)
	ParseError(e, "While RC4 encryption !")
	dst := make([]byte, len(data))
	c.XORKeyStream(dst, data)
	return dst
}

func xor(data []byte, key []byte) []byte {
	BoldYellow.Println("[*] Ciphering...")
	for i := 0; i < len(data); i++ {
		data[i] = (data[i] ^ (key[(i % len(key))]))
	}
	return data
}

func GenerateKey(size int) []byte {
	BoldYellow.Println("[*] Generating random key...")
	key := make([]byte, size)
	rand.Seed(time.Now().UTC().UnixNano())
	for i := 0; i < size; i++ {
		key[i] = byte(rand.Intn(255))
	}
	BoldYellow.Print("[*] Key (ASCII): ")
	BoldBlue.Println(string(key))
	return key
}

func inc(data []byte, val int) []byte {
	BoldYellow.Println("[*] Ciphering...")
	for i := 0; i < len(data); i++ {
		data[i] = (data[i] + byte(val))
	}
	return data
}

func dec(data []byte, val int) []byte {
	BoldYellow.Println("[*] Ciphering...")
	for i := 0; i < len(data); i++ {
		data[i] = (data[i] - byte(val))
	}
	return data
}

func not(data []byte) []byte {
	BoldYellow.Println("[*] Ciphering...")
	for i := 0; i < (len(data)); i++ {
		data[i] = ^data[i]
	}
	return data
}

func ror(data []byte, val uint) []byte {
	BoldYellow.Println("[*] Rotating...")
	for i := 0; i < len(data); i++ {
		data[i] = (data[i] >> uint(val))
	}

	return data
}

func rol(File []byte, Val uint) []byte {
	BoldYellow.Println("[*] Rotating...")
	for i := 0; i < len(File); i++ {
		File[i] = (File[i] << uint(Val))
	}

	return File
}

func checksum(File []byte) int64 {
	BoldYellow.Println("[*] Calculating checsum...")
	var Checksum int64 = 0
	if len(File) > 1 {
		for i := 0; i < len(File); i++ {
			Checksum += int64(File[i])
		}
	} else {
		return 0
	}

	return Checksum
}

func ParseError(err error, msg string) {
	if err != nil {

		fmt.Println("\n")
		BoldRed.Println("\n[-] ERROR: " + msg + "\n")
		os.Exit(1)
	}
}

func xxd(suffix string) {
	xxd, err := exec.Command("xxd", "-i", string(ARGS[len(ARGS)-1]+suffix)).Output()
	ParseError(err, "Unable to retrieve file hex output !")
	BoldGreen.Println(string(xxd))
	xxd, err = exec.Command("xxd", "-i", string(ARGS[len(ARGS)-1]+".key")).Output()
	ParseError(err, "Unable to retrieve key hex output !")
	BoldBlue.Println(string(xxd))
}

func out(data []byte, suffix string) {
	BoldYellow.Println("[*] Writing output...")
	file, err := os.Create(string(ARGS[len(ARGS)-1] + suffix))
	ParseError(err, "Unable to create output file !")
	file.Write(data)
	file.Close()
	KeyFile, err2 := os.Create(string(ARGS[len(ARGS)-1] + ".key"))
	ParseError(err2, "Unable create key file !")
	KeyFile.Write([]byte(*PAR.key))
	KeyFile.Close()
}
