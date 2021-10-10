package dumps

import "unicode"

func (f *Creator) DeleteNulls(data []byte) []byte {
	var cleanData []byte
	for _, ch := range data {
		if ch == 0 {
			continue
		}
		cleanData = append(cleanData, ch)
	}
	return cleanData
}

func (f *Creator) DeleteUnprintableCharacters(data []byte) string {
	dataStr := string(data)
	r := []rune(dataStr)
	var cleanUni []rune
	for _, ch := range r {
		if unicode.IsGraphic(ch) {
			/*if ch == 0xFFFD {
				ch = 0x000A
			}*/
			cleanUni = append(cleanUni, ch)
		}
	}
	return string(cleanUni)
}
