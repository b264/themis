package main

import (
	"fmt";
	"os";
	"github.com/cossacklabs/themis/gothemis/cell";
	"encoding/base64"
)

func main(){
	if 4 != len(os.Args) {
		fmt.Printf("usage %s <command> <password> <message>\n", os.Args[0])
		return
	}
	sc := cell.New([]byte(os.Args[2]), cell.CELL_MODE_SEAL)
	if "enc" == os.Args[1] {
		encData, _, err := sc.Protect([]byte(os.Args[3]), nil)
		if nil != err{
			fmt.Println("error encrypting message")
			return
		}
		fmt.Println(base64.StdEncoding.EncodeToString(encData))
	} else if "dec" == os.Args[1] {
		decoded_message, err := base64.StdEncoding.DecodeString(os.Args[3])
		if nil != err{
			fmt.Println("error decoding message")
			return
		}
		decData, err := sc.Unprotect(decoded_message, nil, nil)
		if nil != err{
			fmt.Println("error decrypting message")
			return
		}
		fmt.Println(string(decData[:]))
	}
}

