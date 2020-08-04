package e

type Error interface {
	error
	Code() int
}

type CodeMsg struct {
	code int
	msg  string
}

func (cm *CodeMsg) Error() string {
	return cm.msg
}

func (cm *CodeMsg) Code() int {
	return cm.code
}

func ErrNewCode(code int) Error {
	return &CodeMsg{code, GetMsg(code)}
}

func ErrNewErr(err error) Error {
	if err == nil {
		return nil
	}
	return &CodeMsg{Fail, err.Error()}
}
