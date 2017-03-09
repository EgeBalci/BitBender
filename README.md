# BitBender
BitBender is a byte manipulation tool



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
