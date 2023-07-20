package core

import (
	"fmt"
	"net/http"
	"net/http/httptest"
	"testing"
)

func TestGetIndexPostCards(t *testing.T) {
	r := httptest.NewRequest(http.MethodGet, "/", nil)
	ctx := r.Context()
	manager := NewPostManager()
	sendManager, err := manager.GetIndexPostCards(ctx)
	if err != nil {
		t.Errorf("aplication failed test:%v", err)
	} else {
		fmt.Println(sendManager)
	}

}
