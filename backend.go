package smtp

import (
	"io"
)

var (
	ErrAutFailed = &SMTPError {
		Code:        "",
	EnhancedCode:  EnhancedCode{}, 	
	Message:         "Authentication failed",
	}
)

ErrAutRequired = &SMTPError {
	Code:        "",
	EnhancedCode:  EnhancedCode{}, 	
	Message:         "Authentication not supported",
}

type Backend interface{
	NewSession(c *Conn) (Session, error)
}

type Session interface{
	Reset()

	Logout()

	AuthPlain(username string, password string) error
   
	Mail(from string, opts *MailOptions) error

	Rcpt(to string, opts *RcptOptions) error

	Data(r io.Reader) error


}


type LMTPSession interface{
	LMTPData(r io.Reader, status StatusCollector) error
}

type StatusCollector interface{
	StatusCollector(rpctTo string, err error)
}
