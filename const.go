package base64Captcha

const idLen = 20

// idChars are characters allowed in captcha id.
var idChars = []byte(TxtNumbers + TxtAlphabet)

const (
	imageStringDpi = 72.0
	//TxtNumbers chacters for numbers.
	TxtNumbers = "012346789"
	//TxtAlphabet characters for alphabet.
	TxtAlphabet = "ABCDEFGHJKMNOQRSTUVXYZabcdefghjkmnoqrstuvxyz"
	//TxtSimpleCharaters simple numbers and alphabet
	TxtSimpleCharaters = "13467ertyiadfhjkxcvbnERTYADFGHJKXCVBN"

	//MimeTypeImage output base64 mine-type.
	MimeTypeImage = "image/png"
	//Emoji is a source string for randTxt
	Emoji = "😀😃💯😄🤖😻😅🤣😂🧑🙃😉😊😇😍👴🤩😘😗☺👽♀😙♂😋😛🎨😜🤪😝🤑🤗🤭🤫🤔🤐🤨😐🙉😶😏💗🙄😬🤥😌😪🤤😷🤢🤮🤯😵🤠😎🧐😨😰😱😭😖😡🤬👿☠💀💥💢"
)

//var cjkFontFamilies = readCJKFonts()

const (
	//OptionShowHollowLine shows hollow line
	OptionShowHollowLine = 2
	//OptionShowSlimeLine shows slime line
	OptionShowSlimeLine = 4
	//OptionShowSineLine shows sine line
	OptionShowSineLine = 8
)
