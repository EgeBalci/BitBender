package Benders

import "github.com/fatih/color"

var Red *color.Color = color.New(color.FgRed)
var BoldRed *color.Color = Red.Add(color.Bold)
var	Blue *color.Color = color.New(color.FgBlue)
var	BoldBlue *color.Color = Blue.Add(color.Bold)
var	Yellow *color.Color = color.New(color.FgYellow)
var	BoldYellow *color.Color = Yellow.Add(color.Bold)
var	Green *color.Color = color.New(color.FgGreen)
var	BoldGreen *color.Color = Green.Add(color.Bold)


type Parameters struct {
	Mode string
	KeySize int
	Key []byte
	Plus int
	Minus int
	RotValue uint
	Checksum int64
}


var BANNER string = `
__________.__  __ __________                   .___            
\______   \__|/  |\______   \ ____   ____    __| _/___________ 
 |    |  _/  \   __\    |  _// __ \ /    \  / __ |/ __ \_  __ \
 |    |   \  ||  | |    |   \  ___/|   |  \/ /_/ \  ___/|  | \/
 |______  /__||__| |______  /\___  >___|  /\____ |\___  >__|   
        \/                \/     \/     \/      \/    \/       

`

var HELP string =`
USAGE: 
	BitBender [options] <file> 
OPTIONS:
  	^	<KeySize>		Make XOR operation with a randomly generated key (Max:~/Min:1)
	^=	<Key>			Make a XOR operation with given key 
	+	<IncrementValue>  	Increment each byte of the file with given value (Max:255/Min:1)
	-	<DecrementValue>	Decrement each byte of the file with given value (Max:255/Min:1)
	!		-		Make a logical NOT operation to each byte of the file
	ror	<RotationValue>		Rotate eache byte of the file to right with given value
	rol	<RotationValue>		Rotate eache byte of the file to left with given value
	= 		-		Calculate the checksum of the given file 
	-h, 	--help 			Print this message 					
EXAMPLE:
	BitBender ^ 12 file
	BitBender ^= topsecretkey file
	BitBender + 4 file
	BitBender - 5 file
	BitBender ! file


`
