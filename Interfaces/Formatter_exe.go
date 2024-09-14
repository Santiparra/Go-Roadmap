package main

type Formatter interface {
	Format() string
}

type PlainText struct {
	message string
}

type Bold struct {
	message string
}

type Code struct {
	message string
}

func (m PlainText) Format() string {
	return m.message	
}

func (m Bold) Format() string {
	return "**" + m.message + "**"
}

func (m Code) Format() string {
	return "`" + m.message + "`"
}

func SendMessage(formatter Formatter) string {
	return formatter.Format()
}