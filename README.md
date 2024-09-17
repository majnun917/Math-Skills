# MATH-SKILLS

**Purpose**:
This Go program calculates the average, median, variance, and standard deviation of a dataset from a text file.

**Input**:
A file containing a list of numbers

**Ouput**:

Average    
Median     
Variance   
Standard Deviation     

## Usage

1 - Clone the repository: 
```
git clone [your-repository-url]
```

2 - Open your terminal and run the following command:
```
go run main.go data.txt
```
(Replace data.txt with your actual file name)

*Example Output*:

==== Calculation of statistical measures ====       
Average: 150         
Median: 152              
Variance: 855           
Standard Deviation: 29         

### Limitations:

The program assumes the input file contains valid numbers, it displays an error message if empty lines and non-numerical data.

## Algorithm

This program utilizes two main functions:

### WelfordAlgo: 
This function implements the Welford online algorithm for calculating statistics (average, variance, standard deviation) in a single pass through the data.     
The Welford online algorithm calculates the mean and variance by going through the data stream just once, using a minimum amount of memory and avoiding the rounding errors that can occur when calculating the variance from the mean.

### MedianCalcul: 
This function sorts the data and calculates the median based on the number of elements.

The program reads the data from a file line by line, handling potential errors during file opening and data parsing.

**Testing**:
The program can be tested by providing a sample data file and comparing the output with manually calculated results.        
    After downloading the [file](https://assets.01-edu.org/stats-projects/stat-bin-dockerized.zip).     
Run the script with:
     ```
     ./bin/math-skills
     ``` 
     or 
     ```
     ./run.sh math-skills
     ```
Then run the program with the created file (data.txt) by the previous command.

**Languages**:
This specific code is written in Go. You can adapt the logic to other languages.
