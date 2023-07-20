package rest

import (
	"bytes"
	"errors"
	"log"
	"net/http"
)

// UnauthorizedError retorna erro de autenticação inválida.
func UnauthorizedError() error {
	return &Error{Status: http.StatusUnauthorized, Code: "unauthorized", Message: "Autenticação inválida."}
}

// ForbiddenError retorna erro de usuário autorizado mas sem acesso a um recurso específico.
func ForbiddenError() error {
	return &Error{Status: http.StatusUnauthorized, Code: "forbidden", Message: "Acesso negado."}
}

// NotFoundError retorna erro de recurso não encontrado.
func NotFoundError() error {
	return &Error{Status: http.StatusNotFound, Code: "not_found", Message: "Recurso não encontrado."}
}

// ConflictError retorna erro de recurso já cadastrado.
func ConflictError() error {
	return &Error{Status: http.StatusConflict, Code: "conflict", Message: "Recurso já existe."}
}

// UnexpectedError retorna erro inesperado.
func UnexpectedError() error {
	return &Error{Status: http.StatusInternalServerError, Code: "unexpected", Message: "Ocorreu um erro inesperado."}
}

// JSONSyntaxError retorna erro de sintaxe JSON incorreta.
func JSONSyntaxError() error {
	return &Error{Status: http.StatusBadRequest, Code: "json_syntax", Message: "Corpo de requisição com sintaxe incorreta."}
}

// Error representa uma mensagem de erro que será enviada para o cliente.
type Error struct {
	Status   int                    `json:"status"`
	Code     string                 `json:"code"`
	Message  string                 `json:"message"`
	Detail   map[string]interface{} `json:"detail,omitempty"`
	Internal error                  `json:"-"`
}

// Unwrap retorna o erro interno que originou o atual.
func (e *Error) Unwrap() error {
	return e.Internal
}

func (e *Error) Error() string {
	return e.Message
}

func (e *Error) Is(err error) bool {
	typedError, ok := err.(*Error)
	return ok && typedError.Code == e.Code
}

// ErrorMap representa um mapa de erros que será enviado para o cliente.
type ErrorMap map[string]error

func (e ErrorMap) Error() string {
	var result bytes.Buffer
	for k, v := range e {
		result.WriteString(k + ": " + v.Error() + "\n")
	}
	return result.String()
}

// LogError escreve detalhes de um erro em log.
func LogError(err error, v ...interface{}) error {
	if err == nil {
		return nil
	}
	var e *Error
	if errors.As(err, &e) {
		if len(v) > 0 {
			internErr := errors.Unwrap(e)
			if internErr == nil {
				internErr = e
			}
			v = append(v, ":", internErr)
			log.Println(v...)
		}
		return err
	}
	if len(v) > 0 {
		v = append(v, ":", err)
		log.Println(v...)
	}
	return &Error{Status: http.StatusInternalServerError, Code: "unexpected", Message: "Ocorreu um erro inesperado.", Internal: err}
}
