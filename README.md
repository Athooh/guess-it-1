Guess It 1

Guess It 1 is a command-line tool written in Go that reads numbers from standard input, calculates the average, variance, and standard deviation, and determines a range based on these statistics. The range is calculated as mean ± 1.5 * standard deviation.
Table of Contents

    Installation
    Project Structure
    Usage
    Testing
    Contributing
    Author

Installation

To use this project, you need to have Go installed. Follow the steps below to get started:

    Clone the repository:

    git clone https://learn.zone01kisumu.ke/git/seodhiambo/guess-it-1.git
    cd guess-it-1

Project Structure

    student/main.go: The main entry point for the application.
    student/mathfuncs/average.go: Calculates the average of a slice of float64 numbers.
    student/mathfuncs/median.go: Calculates the median of a slice of float64 numbers.
    student/mathfuncs/range.go: Calculates the range based on mean and standard deviation.
    student/mathfuncs/std_deviation.go: Calculates the standard deviation from variance.
    student/mathfuncs/variance.go: Calculates the variance of a slice of float64 numbers.

Usage

    Change the permission of the script file to executable:

chmod +x ./student/script.sh

    To run the project, use the provided bash script:

./student/script.sh

Enter numbers one by one and press Enter. The program will calculate and print the range after each entry. To exit, simply press Enter without typing a number.
Testing

To test the project, follow these steps:

    Download the test zip file containing the tester.
    Extract the zip file in the root directory of the project.
    Move the student/ folder into the extracted directory.
    Follow the instructions provided in the tester's README file.

Contributing

    Fork the repository.
    Create a new branch (git checkout -b feature-branch).
    Commit your changes (git commit -am 'Add new feature).
    Push to the branch (git push origin feature-branch).
    Create a new Pull Request.

Author

   Seth Athooh
