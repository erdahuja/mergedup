package item

import (
	"context"
	"fmt"
	"testing"
	"time"

	"github.com/golang/mock/gomock"
)

func Test_Item(t *testing.T) {
	t.Run("feat", feat)
}

func feat(t *testing.T) {
	ctrl := gomock.NewController(t)
	defer ctrl.Finish()
	m := NewMockStorer(ctrl)
	now := time.Now()
	itm := Item{
		Name:        "one",
		Cost:        1,
		Quantity:    1,
		DateCreated: now,
		DateUpdated: now,
	}
	m.EXPECT().QueryByID(gomock.Any(), gomock.Any()).Return(itm, nil)
	c := NewCore(m)
	itm, err := c.QueryByID(context.TODO(), 1)
	if err != nil {
		t.Fail()
	}
	fmt.Println(itm)

}
