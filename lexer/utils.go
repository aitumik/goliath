package lexer

func isLetter(char byte) bool {
	//return 'a' <= char && char <= 'z' || 'A' <= char && char <= 'Z' || char == '_'
	return 'a' <= char && char <= 'z'
}

func isNumber(char byte) bool {
	return '0' <= char && char <= '9'
}

func (l *Lexer) skipWhiteSpace() {
	for l.char == ' ' || l.char == '\t' || l.char == '\n' || l.char == '\r' {
		l.readChar()
	}
}

func (l *Lexer) peekChar() byte {
	if l.readPosition >= len(l.input) {
		return 0
	}
	return l.input[l.readPosition]
}

func (l *Lexer) readNumber() string {
	pos := l.position

	for isNumber(l.char) {
		l.readChar()
	}

	return l.input[pos:l.position]
}
