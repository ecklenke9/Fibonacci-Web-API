package controller

import (
	"Fibonacci-Web-API/api/model"
	"github.com/gin-gonic/gin"
	"reflect"
	"testing"
)

func TestDeleteAllFibonacci(t *testing.T) {
	type args struct {
		in0 *gin.Context
	}
	tests := []struct {
		name string
		args args
	}{
		{
			name: "Successful Delete",
			args: args{},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
		})
	}
}

func TestGetAllFibonacci(t *testing.T) {
	type args struct {
		c *gin.Context
	}
	tests := []struct {
		name string
		args args
	}{
		{
			name: "Successful Get All Fibonacci",
			args: args{},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
		})
	}
}

func TestGetFibonacci(t *testing.T) {
	type args struct {
		c *gin.Context
	}
	tests := []struct {
		name string
		args args
	}{
		{
			name: "Successful Get Fibonacci",
			args: args{},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
		})
	}
}

func TestGetMemoizedResults(t *testing.T) {
	type args struct {
		c *gin.Context
	}
	tests := []struct {
		name string
		args args
	}{
		{
			name: "Successful Get Memoized Results",
			args: args{},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
		})
	}
}

func Test_fibonacciCalculation(t *testing.T) {
	type args struct {
		ordinal int
	}
	tests := []struct {
		name     string
		args     args
		want     int
		fibModel []model.Fibonacci
	}{
		{
			name: "Successful Fibonacci Calculation",
			args: args{4},
			want: 3,
			fibModel: []model.Fibonacci{
				{
					Ordinal: 0,
					FibNum:  1,
				},
				{
					Ordinal: 1,
					FibNum:  1,
				},
				{
					Ordinal: 2,
					FibNum:  1,
				},
				{
					Ordinal: 3,
					FibNum:  2,
				},
				{
					Ordinal: 4,
					FibNum:  3,
				},
			},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			fib, fibArray := fibonacciCalculation(tt.args.ordinal)
			if fib != tt.want {
				t.Errorf("fibonacciCalculation() got = %v, want %v", fib, tt.want)
			}
			if !reflect.DeepEqual(fibArray, tt.fibModel) {
				t.Errorf("fibonacciCalculation() got1 = %v, want %v", fibArray, tt.fibModel)
			}
		})
	}
}

func Test_upsertFibonacci(t *testing.T) {
	type args struct {
		fibonacciArray []model.Fibonacci
	}
	tests := []struct {
		name string
		args args
	}{
		{
			name: "Successful Upsert Fibonacci",
			args: args{},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
		})
	}
}
