# goml

`goml` is a lightweight machine learning library implemented in Golang. Currently, it provides regression algorithms, and the aim is to expand its capabilities to include various machine learning algorithms in the future.

## Installation

To use `goml` in your Golang project, simply download it as follows:

```bash
go get github.com/thisdoraemon/goml
```

## Regression Algorithms

### Simple Linear Regression
Function: `CalculateSlope`
```go
func CalculateSlope(x []float64, meanX float64, y []float64, meanY float64) float64
```

Calculates the slope of the regression line for simple linear regression.

Parameters:

- `x:` Input feature values.
- `meanX:` Mean of the input feature values.
- `y:` Output values.
- `meanY:` Mean of the output values.

Returns:

- The slope of the regression line.

Function: `Mean`
```go
func Mean(data []float64) float64
```

Calculates the mean of a given slice of float64 values.

Parameters:

- `data:` Slice of float64 values.

Returns:

- The mean of the input data.

## Example Usage

```go
package main

import (
   "fmt"
   "github.com/thisdoraemon/goml"
)

func main() {
   // Example data
   x := []float64{1, 2, 3, 4, 5}
   y := []float64{2, 4, 5, 4, 5}

   // Calculate means
   meanX := goml.Mean(x)
   meanY := goml.Mean(y)

   // Calculate slope
   slope := goml.CalculateSlope(x, meanX, y, meanY)

   // Display result
   fmt.Printf("Slope: %.2f\n", slope)
}
```

## Contribution

Feel free to contribute to `goml` by submitting issues, feature requests, or pull requests. Your contributions are highly appreciated.

## License

This project is licensed under the MIT License - see the LICENSE file for details.

