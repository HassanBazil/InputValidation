package main

import (
	"fmt"
	"unicode/utf8"

	validator "github.com/go-playground/validator/v10"
	"golang.org/x/text/unicode/norm"

	StockJSON "encoding/json"

	GoJay "github.com/francoispqt/gojay"
	goccyJSON "github.com/goccy/go-json"
	JSONiter "github.com/json-iterator/go"
	SegmentIoJSON "github.com/segmentio/encoding/json"
	//SIMD doesn't work on apple silicon
	//jettison doesn't decode
	//easyJSON is a parser code generator
)

const NumberOfParsers = 6

// Yes you can use regex, but should you?
type User struct {
	Name  string `validate:"required" json:"name"`
	Age   int    `validate:"required,gte=0,lte=130" json:"age"`
	Email string `validate:"required,email" json:"email"`
}

func main() {

	users := make([]User, NumberOfParsers, NumberOfParsers)

	// Origin - ip address, access key (token)
	input := []byte(slideExample4) //Hide input for fun

	// Size - Is message reasonably long? Are all fields reasonably long as well?
	if len(input) > 1000 {
		fmt.Println("@ Input is larger than should be.")
	} else {
		fmt.Println("/ Input is size is okay!")
	}

	//  Lexical validation - Content uses right characters and encoding?
	// -> Normalise encoding, reject invalid chars.
	// -> Allow only certain character categories (i.e.: Only latin chars, uppercase letters etc.) // -> Allow only certain chars i.e.: ' for names for our Irish friends but not .
	if !utf8.Valid(input) {
		fmt.Println("@ Incorrect character encoding!")
	} else {
		fmt.Println(" Character encoding is okay!")
	}

	//Normalise unicode bytes using composing representation for more compact representation.
	input = norm.NFC.Bytes(input)

	//Syntax validation - Is the format is right? (JSON double property etc. )
	// ->JSON double property
	// ->JSON extra property during deserialization
	// -- >email address should adhere to standard
	if !StockJSON.Valid([]byte(input)) {
		fmt.Println("@ Incorrect JSON")
	} else {
		fmt.Println(" JSON is valid.")
	}

	// Parsing
	// Implements JSON unmarshal using libraries as shown in table here: https://github.com/goccy/go-j

	var parsers = map[int]string{
		0: "Stock",
		1: "goccy",
		2: "JSONiter",
		3: "JSONiter compatible mode",
		4: "GoJay",
		5: "SegmentIOJSON",
	}

	StockJSON.Unmarshal(input, &users[0])
	goccyJSON.Unmarshal(input, &users[1])
	JSONiter.Unmarshal(input, &users[2])
	var JSONIterCompatibilityMode = JSONiter.ConfigCompatibleWithStandardLibrary
	JSONIterCompatibilityMode.Unmarshal(input, &users[3])
	GoJay.UnmarshalJSONObject(input, &users[4])
	SegmentIoJSON.Unmarshal(input, &users[5])
	//SIMDUnmarshal(input) doesn't work on mac .... : )

	for i := 0; i < NumberOfParsers; i++ {
		fmt.Println(parsers[i], users[i])
	}

	fmt.Println("Let's validate")

	validate := validator.New()
	for i := 0; i < NumberOfParsers; i++ {
		err := validate.Struct(users[i])
		if err != nil {
			fmt.Println("@", parsers[i], " struct did not pass validation.")
		} else {
			fmt.Println("", parsers[i], " struct did passed validation.")
		}
		fmt.Println(users[i])
	}

	fmt.Println("")
	// Semantic validation - Does it make sense?
	// ->Is price of good is positive?
	// ->Is this address actually exist?
	//->Isthis filetype make sense?
	// ->Is this email controlled by the user?
	// ->Is cryptographic controls applied strong and passes validation?
	// ->Is this url points to public or private resource?
	// ->Is this user exist in the DB?
	// -> uniqness check
	// -> dereferencing
	// -> TOCTOU locking
	for i := 0; i < NumberOfParsers; i++ {
		if users[i].Name != norm.NFC.String("Pérez Zoé") {
			fmt.Println("Not the Zoé", parsers[i], "expected. This is ", users[i].Name, ".")
		} else {
			fmt.Println(" Hello Zoé", parsers[i], "expected you.")
		}
	}
}
