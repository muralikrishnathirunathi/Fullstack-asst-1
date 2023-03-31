package main

import (
	"fmt"
	"log"
	"os"

	"github.com/go-gota/gota/dataframe"
	"github.com/sajari/regression"
)

func main() {
	// Open the CSV file.
	breastCancer, err := os.Open("breast_cancer.csv")
	if err != nil {
		log.Fatal(err)
	}
	defer breastCancer.Close()

	// Create a dataframe from the CSV file.
	breastCDF := dataframe.ReadCSV(breastCancer)

	// Use the Describe method to calculate summary statistics
	// for all of the columns in one shot.
	breastSummary := breastCDF.Describe()

	// Output the summary statistics.
	fmt.Println(breastSummary)

	// Print the complete dataframe.
	fmt.Println("Complete Dataframe:")
	fmt.Println(breastCDF)

	// Create a new regression instance.
	r := new(regression.Regression)

	// Add the independent variables to the regression.
	for _, colName := range breastCDF.Names() {
		if colName == "id" || colName == "diagnosis" {
			continue
		}
		colData := breastCDF.Col(colName).Float()
		dp := make([]regression.DataPoint, len(colData))
		for i, f := range colData {
			dp[i].Observed = f
			dp[i].Variables = []float64{breastCDF.Col("radius_mean").Float()[i]}
		}
		r.Train(dp)
	}

	// Add the dependent variable to the regression.
	diagnosis := breastCDF.Col("diagnosis").Float()
	dp := make([]regression.DataPoint, len(diagnosis))
	for i, f := range diagnosis {
		dp[i].Observed = f
	}
	r.Train(dp)

	// Fit the regression using the normal equation.
	r.Run()

	// Output the regression formula.
	formula := "diagnosis = "
	for i, coef := range r.Weights {
		if i == 0 {
			formula += fmt.Sprintf("%.2f", coef)
		} else {
			formula += fmt.Sprintf(" + %.2f * %s", coef, breastCDF.Names()[i-1])
		}
	}
	fmt.Println("Regression Formula:", formula)

	// Output the regression statistics.
	fmt.Println("R²:", r.R2)
	fmt.Println("Coefficients:", r.Weights)
	fmt.Println("Intercept:", r.Intercept)

	// Print the complete dataframe.
	fmt.Println("Complete Dataframe:")
	fmt.Println(breastCDF)
}
