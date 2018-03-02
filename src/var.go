package main

import "github.com/fatih/color"

var Red *color.Color = color.New(color.FgRed)
var BoldRed *color.Color = Red.Add(color.Bold)
var Blue *color.Color = color.New(color.FgBlue)
var BoldBlue *color.Color = Blue.Add(color.Bold)
var Yellow *color.Color = color.New(color.FgYellow)
var BoldYellow *color.Color = Yellow.Add(color.Bold)
var Green *color.Color = color.New(color.FgGreen)
var BoldGreen *color.Color = Green.Add(color.Bold)

type Parameters struct {
	RC4 *bool
	XOR *bool
	INC *bool
	DEC *bool
	ROR *bool
	ROL *bool
	NOT *bool
	CHK *bool

	KeySize *int
	key     *string
}

var HELP string = `
# BitBender 
> Author: Ege Balcı
> Source: github.com/egebalci/BitBender

Usage of ./bib:
  -K string
    	Cipher key.
  -chk
    	Calculate the checksum.
  -dec
    	Cipher with logical DEC operation.
  -inc
    	Cipher with logical INC operation.
  -k int
    	Size of the random key. (default 1)
  -not
    	Cipher with logical NOT operation.
  -rc4
    	Cipher with RC4 algorithm.
  -rol
    	Cipher with logical ROL operation.
  -rot
    	Cipher with logical ROT operation.
  -xor
    	Cipher with logical XOR operation.

`
