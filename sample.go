package main

// Emoji in email address is a thing https://mailoji.com/

// note that the Name and Age fields fail the validation

const exampleInput = `{
	"Name": "John Doe2",
	"name": "John Doe",
	"Age" : 4566789,
	"Age": 26,
	"Email": "john@@@xuple.com",
	"Email": "john.smith@example.com"
}`

// // Emoji in email address is a thing https://mailoji.com/

// // note that the Name and Age fields fail the validation

// const exampleInput2 = {

// "name": "John Doe2"

// "nAmE": "John Doe",

// "Age": 26,

// "Age": "4566789",

// "Email": "john.smith@example.com",

// "Email"："johnc@axmple.com"

// // Emoji in email address is a thing https://mailoji.com/

// // note that the Name and Age fields fail the validation

// const exampleInput3 = {

// "name":

// "John Doe2",

// "nAmE": "John Doe",

// "Age": 26,

// "Age": "268",

// "Email": "john@@@xmple.com"

// "Email": "john.smith@example.com"

// const slideExample = {

// "name": "John Doe2"

// "name": "John Doe",

// "Age": 26,

// "Age": "4566789"

// "Email": "john@@@xmple.com",

// "Email": "john.smith@example.com"

// const slideExample2 =

// "name": "John Doe2"

// "name": "John Doe",

// "Age": 26,

// "Age": "4566789"

// "Email": "john.smith@example.com"

// "Email": "john@@@xmple.com"

// const

// slideExample3

// "name": "John Doe",

// "Age": 26,

// "Email": "john@@@xmple.com"

const slideExample4 = `{
"name":"Pérez Zoé",
"age": 26,
"email": "john🐶@💰xmple.com"
}`

// é combined
// é non combined
const slideExample5 = `{
"name":"Pérez Zoé2",
"nAmE": "Pérez Zoé",
"Email": "john🐶@💰xmple.com"
"email": "john.smith@example.com"
"agE": "4566789",
"age": 26
}`
