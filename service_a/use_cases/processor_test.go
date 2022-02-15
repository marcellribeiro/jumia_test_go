package use_cases

import (
	"github.com/marcellribeiro/jumia_test_go/service_a/models"
	"reflect"
	"testing"
)

func gotWantProccess(want []models.Order, t *testing.T) {
	got, _ := ProcessFile("processor_test.csv")

	if !reflect.DeepEqual(got, want) {
		t.Errorf("not equal")
	}
}

func TestProcessWithSuccess(t *testing.T) {
	tmpOrder := []models.Order{
		models.Order{
			ID:     "1",
			Email:  "email1@email.com",
			Phone:  "237 209993809",
			Weight: "24.45",
		},
		models.Order{
			ID:     "2",
			Email:  "email2@email.com",
			Phone:  "258 852828436",
			Weight: "1.33",
		},
		models.Order{
			ID:     "3",
			Email:  "email3@email.com",
			Phone:  "256 217813782",
			Weight: "15.16",
		},
	}
	gotWantProccess(tmpOrder, t)
}

//func TestProcessWithError(t *testing.T) {
//	gotWantProccess("5587609808", t)
//}
