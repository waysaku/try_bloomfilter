package main

import (
	"fmt"
	"crypto/sha256"
	"encoding/binary"
	"bytes"
)


func main() {

	// 32bit
	var filter [4294967295]bool
	addText 	:= "apple"

	for _, hashFunc := range initializeHashFuncs() {
		hash := hashFunc(addText)

		var i uint32
		buf := bytes.NewReader(hash[:])
		err := binary.Read(buf, binary.LittleEndian, &i)
		if err != nil {
			fmt.Println("binary.Read failded:", err)
		}

		filter[i] = true
	}

	fmt.Println(isExist(filter[:], "apple"))

}

func isExist(filter []bool, text string) bool {
	var exist bool
	for _, hashFunc := range initializeHashFuncs() {
		hash := hashFunc(text)

		var i uint32
		buf := bytes.NewReader(hash[:])
		err := binary.Read(buf, binary.LittleEndian, &i)
		if err != nil {
			fmt.Println("binary.Read failded:", err)
		}
		if filter[i] {
			exist = true
		}
	}
	return exist
}



func initializeHashFuncs() []func(word string) [sha256.Size]byte {
	return []func(word string) [sha256.Size]byte {
								makeHashFunc("hoge0"),
								makeHashFunc("hoge1"),
								makeHashFunc("hoge2"),
								makeHashFunc("hoge3"),
								makeHashFunc("hoge4"),
								makeHashFunc("hoge5"),
								makeHashFunc("hoge6"),
								makeHashFunc("hoge7"),
								makeHashFunc("hoge8"),
								}
}
func makeHashFunc(salt string) func(word string) [sha256.Size]byte {
	return func(word string) [sha256.Size]byte {
		return sha256.Sum256([]byte(word + salt))
	}
}
