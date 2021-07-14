package heap

import "unicode/utf16"

var internedStrings = map[string]*Object{}

func JString(loader *ClassLoader, goStr string) *Object {
	if internedStr, ok := internedStrings[goStr]; ok {
		return internedStr
	}
	chars := stringToUtf16(goStr)
	jChars := &Object{loader.LoadClass("[C"), chars}
	jStr := loader.LoadClass("java/lang/String").NewObject()
	jStr.SetRefVar("value", "[C", jChars)
	internedStrings[goStr] = jStr
	return jStr
}

func GoString(jStr *Object) string {
	charAttr := jStr.GetRefVar("value", "[C")
	return uft16ToString(charAttr.Chars())
}

func uft16ToString(chars []uint16) string {
	runes := utf16.Decode(chars) //utf8
	return string(runes)
}

func stringToUtf16(str string) interface{} {
	runes := []rune(str) //utf32
	return utf16.Encode(runes)
}
