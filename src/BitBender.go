package main

import "io/ioutil"
import "strconv"
import "flag"
import "fmt"
import "os"

var ARGS []string
var PAR Parameters

func main() {

	ARGS = os.Args[1:]
	if len(ARGS) < 1 {
		fmt.Println(HELP)
		flag.PrintDefaults()
		os.Exit(1)
	}

	PAR.RC4 = flag.Bool("rc4", false, "Cipher with RC4 algorithm.")
	PAR.XOR = flag.Bool("xor", false, "Cipher with logical XOR operation.")
	PAR.ROR = flag.Bool("rot", false, "Cipher with logical ROT operation.")
	PAR.ROL = flag.Bool("rol", false, "Cipher with logical ROL operation.")
	PAR.INC = flag.Bool("inc", false, "Cipher with logical INC operation.")
	PAR.DEC = flag.Bool("dec", false, "Cipher with logical DEC operation.")
	PAR.NOT = flag.Bool("not", false, "Cipher with logical NOT operation.")
	PAR.CHK = flag.Bool("chk", false, "Calculate the checksum.")
	PAR.KeySize = flag.Int("k", 1, "Size of the random key.")
	PAR.key = flag.String("K", "", "Cipher key.")

	flag.Parse()

	file, err := ioutil.ReadFile(ARGS[len(ARGS)-1])
	ParseError(err, "Unable to open input file !")

	BoldYellow.Print("[*] Key Size: ")
	if *PAR.key == "" {
		BoldBlue.Println(*PAR.KeySize)
	} else {
		BoldBlue.Println(len(*PAR.key))
		*PAR.KeySize = len(*PAR.key)
	}

	if *PAR.RC4 {
		if *PAR.key == "" {
			*PAR.key = string(GenerateKey(*PAR.KeySize))
		}
		out(RC4(file, []byte(*PAR.key)), ".rc4")
		xxd(".rc4")
		os.Exit(0)
	} else if *PAR.XOR {
		if *PAR.key == "" {
			*PAR.key = string(GenerateKey(*PAR.KeySize))
		}
		out(xor(file, []byte(*PAR.key)), ".xor")
		xxd(".xor")
		os.Exit(0)
	} else if *PAR.INC {
		if *PAR.key == "" {
			*PAR.key = string(GenerateKey(*PAR.KeySize))
		}
		out(inc(file, *PAR.KeySize), ".inc")
		xxd(".inc")
		os.Exit(0)
	} else if *PAR.DEC {
		if *PAR.key == "" {
			*PAR.key = string(GenerateKey(*PAR.KeySize))
		}
		out(dec(file, *PAR.KeySize), ".dec")
		xxd(".dec")
		os.Exit(0)
	} else if *PAR.ROR {
		if *PAR.key == "" {
			*PAR.key = string(GenerateKey(*PAR.KeySize))
		}
		out(ror(file, uint(*PAR.KeySize)), ".dec")
		xxd(".dec")
		os.Exit(0)
	} else if *PAR.ROL {
		if *PAR.key == "" {
			*PAR.key = string(GenerateKey(*PAR.KeySize))
		}
		out(rol(file, uint(*PAR.KeySize)), ".rol")
		xxd(".rol")
		os.Exit(0)
	} else if *PAR.NOT {
		if *PAR.key == "" {
			*PAR.key = string(GenerateKey(*PAR.KeySize))
		}
		out(not(file), ".not")
		xxd(".not")
		os.Exit(0)
	} else if *PAR.CHK {
		var checksum int64 = checksum(file)
		BoldGreen.Print("[#] ")
		BoldYellow.Print("Checksum : ")
		BoldRed.Print("0x", strconv.FormatInt(checksum, 16), "\n")
		os.Exit(0)
	} else {
		BoldRed.Println("[-] ERROR : Choose a valid operation mode !")
		os.Exit(1)
	}

}
