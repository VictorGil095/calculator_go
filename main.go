go
package main

import (
 "bufio"
 "fmt"
 "os"
 "strconv"
 "strings"
)

func main() {
 reader := bufio.NewReader(os.Stdin)

 fmt.Print("Enter the expression: ")
 expression, _ := reader.ReadString('\n')
 expression = strings.TrimSpace(expression)

 result, err := calculate(expression)
 if err != nil {
  fmt.Println(err)
  return
 }

 fmt.Println("Result:", result)
}

func calculate(expression string) (string, error) {
 tokens := strings.Split(expression, " ")
 if len(tokens) != 3 {
  return "", fmt.Errorf("Invalid expression")
 }

 a, err := parseNumber(tokens[0])
 if err != nil {
  return "", err
 }

 operator := tokens[1]

 b, err := parseNumber(tokens[2])
 if err != nil {
  return "", err
 }

 var result int
 switch operator {
 case "+":
  result = a + b
 case "-":
  result = a - b
 case "*":
  result = a * b
 case "/":
  if b == 0 {
   return "", fmt.Errorf("Division by zero")
  }
  result = a / b
 default:
  return "", fmt.Errorf("Invalid operator")
 }

 return formatResult(result), nil
}

func parseNumber(str string) (int, error) {
 num, err := strconv.Atoi(str)
 if err == nil {
  if num < 1 || num > 10 {
   return 0, fmt.Errorf("Number out of range")
  }
  return num, nil
 }

 romanNumeral := strings.ToUpper(str)
 if !isValidRomanNumeral(romanNumeral) {
  return 0, fmt.Errorf("Invalid Roman numeral")
 }

 return convertRomanToArabic(romanNumeral), nil
}

func isValidRomanNumeral(romanNumeral string) bool {
 return true
}

func convertRomanToArabic(romanNumeral string) int {
 return 0
}

func formatResult(result int) string {
 if result < 0 {
  return strconv.Itoa(result)
 }

 return convertArabicToRoman(result)
}

func convertArabicToRoman(number int) string {
 return ""
}