package main

import (
	"fmt"
	"log"
	"math"
	"strconv"
	"strings"
)

type Operation struct {
	Parameters OperationParameters
	Operator   string
}

type OperationParameters struct {
	NumberX float64
	NumberY float64
}

func getOperationParametersFromRawOperation(rawOperation, operator string) OperationParameters {
	numbers := strings.Split(rawOperation, operator)

	rawX := numbers[0]
	rawY := numbers[1]

	x, _ := strconv.ParseFloat(rawX, 64)
	y, _ := strconv.ParseFloat(rawY, 64)

	log.Print("[INFO] x: ", x, " y: ", y)

	return OperationParameters{
		NumberX: x,
		NumberY: y,
	}
}

type Mathematician interface {
	Calculate(params OperationParameters) (float64, error)
}

type Mult struct {
	Operation
}

func (m Mult) Calculate(params OperationParameters) (float64, error) {
	return params.NumberX * params.NumberY, nil
}

type Sum struct {
	Operation
}

func (m Sum) Calculate(params OperationParameters) (float64, error) {
	return params.NumberX + params.NumberY, nil
}

type Sub struct {
	Operation
}

func (m Sub) Calculate(params OperationParameters) (float64, error) {
	return params.NumberX - params.NumberY, nil
}

type Div struct {
	Operation
}

func (m Div) Calculate(params OperationParameters) (float64, error) {
	return params.NumberX / params.NumberY, nil
}

type Pow struct {
	Operation
}

func (m Pow) Calculate(params OperationParameters) (float64, error) {
	return math.Pow(params.NumberX, params.NumberY), nil
}

var ErrInvalidRoot = fmt.Errorf("Invalid root")

type Rot struct {
	Operation
}

func (m Rot) Calculate(params OperationParameters) (float64, error) {
	if params.NumberY == 2 {
		return math.Sqrt(params.NumberX), nil
	} else if params.NumberY == 3 {
		return math.Cbrt(params.NumberX), nil
	} else {
		return 0, ErrInvalidRoot
	}
}
