package origin

import (
	"fmt"
	"crypto/sha256"
	"encoding/binary"
	"bytes"
	"math"
	"strconv"
)

var m uint32 = 4096
var n = 256
var k = int((float64(m) / float64(n)) * math.Log(2))
var checkCount = 1000000

func Filter() {
	fmt.Println("m = " + strconv.Itoa(int(m)))
	fmt.Println("n = " + strconv.Itoa(n))
	fmt.Println("k = " + strconv.Itoa(k))

	fmt.Println("ideal Rate       = " + strconv.FormatFloat(math.Pow(0.5, float64(k)),'f', 6, 64))
	fmt.Println("Rate by formula  = " + strconv.FormatFloat(math.Pow(0.6185, float64(m) / float64(n)), 'f', 6, 64))
	rate := math.Pow(1 - math.Pow(float64(1 - (float32(1) / float32(m))), float64(k * n)), float64(k))
	fmt.Println("culculated Rate  = " + strconv.FormatFloat(rate, 'f', 6, 64))


	var filter 			= make([]bool, m)
	var addingStrings   = generateText(n, "adding_word")
	var checkingStrings = generateText(checkCount, "checking_word")
	var hashFuncs       = initializeHashFuncs(k)

	for _, str := range addingStrings {
		for _, hashFunc := range hashFuncs {
			hash := hashFunc(str)

			var i uint32
			buf := bytes.NewReader(hash[:])
			err := binary.Read(buf, binary.LittleEndian, &i)
			if err != nil {
				fmt.Println("binary.Read failded:", err)
			}

			filter[i % m] = true
		}
	}

	fmt.Println("----- add text done! -----")

	var countFalsePositive int
	for _, str := range checkingStrings {
		rslt := isExist(hashFuncs, filter[:], str)
		if rslt {
			countFalsePositive++
		}
	}

	fmt.Println("count of FalsePositives is " + strconv.Itoa(countFalsePositive))
	fmt.Println("False Positive Rate = " + strconv.FormatFloat(float64(countFalsePositive) / float64(checkCount), 'f', 4, 64))
	fmt.Println("----- end! -----")
}


func generateText(num int, word string) []string {
	var addingStrings = make([]string, num)
	for i := 0; i < num; i++ {
		addingStrings[i] = word + strconv.Itoa(i)
	}
	return addingStrings
}

func isExist(hashFuncs []func(word string) [sha256.Size]byte, filter []bool, text string) bool {
	for _, hashFunc := range hashFuncs {
		hash := hashFunc(text)

		var i uint32
		buf := bytes.NewReader(hash[:])
		err := binary.Read(buf, binary.LittleEndian, &i)
		if err != nil {
			fmt.Println("binary.Read failded:", err)
		}
		if !filter[i % m] {
			return false
		}
	}
	return true
}



func initializeHashFuncs(k int) []func(word string) [sha256.Size]byte {
	var hashFuncs = make([]func(word string) [sha256.Size]byte, k)
	for i := 0; i < k; i++ {
		hashFuncs[i] = makeHashFunc("function_seed_" + strconv.Itoa(i))
	}
	return hashFuncs[:]
}

func makeHashFunc(salt string) func(word string) [sha256.Size]byte {
	return func(word string) [sha256.Size]byte {
		//return sha256.Sum256([]byte(word + salt))
		h := sha256.Sum256([]byte(word + salt))
		return sha256.Sum256(h[:])
	}
}
