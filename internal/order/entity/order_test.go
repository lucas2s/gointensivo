package entity

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestGivenAnEmptyID_WhenCreateOrderANewOrder_ThenShouldReceiveError(t *testing.T) {
	order := Order{}
	assert.Error(t, order.IsValid(), "invalid id")
}

func TestGivenAnEmptyPrice_WhenCreateOrderANewOrder_ThenShouldReceiveError(t *testing.T) {
	order := Order{ID: "123"}
	assert.Error(t, order.IsValid(), "invalid id")
}

func TestGivenAnEmptyTax_WhenCreateOrderANewOrder_ThenShouldReceiveError(t *testing.T) {
	order := Order{ID: "123", Price: 10}
	assert.Error(t, order.IsValid(), "invalid tax")
}

func TestGivenAValidParams_WhenICallNewOrder_ThenShouldReceivecreateOrderWithAllParams(t *testing.T) {
	order, err := NewOrder("123", 10.0, 2.0)
	assert.Nil(t, err)
	assert.Equal(t, "123", order.ID)
	assert.Equal(t, 10.0, order.Price)
	assert.Equal(t, 2.0, order.Tax)
}

func TestGivenAPriceAndTax_WhenICallCalculatePrice_ThenShouldSetFinalPrice(t *testing.T) {
	order, err := NewOrder("123", 10.0, 2.0)
	assert.Nil(t, err)
	assert.Nil(t, order.CalculateFinalPrice())
	assert.Equal(t, 12.0, order.FinalPrice)
}
