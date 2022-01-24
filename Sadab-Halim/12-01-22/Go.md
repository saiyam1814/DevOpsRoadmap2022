## Continued Learning Go from [Golang Tutorial for Beginners](https://www.youtube.com/watch?v=yyUHQIec83I)

### Conditionals
| Operator | Description             |
| -------- | ----------------------- |
| ==       | Compares two things     |
| !=       | Not equals              |
| <        | Less than               |
| <=       | Less than equal to      |
| >        | Greater than            |
| >=       | Greater than equal to   |

**If Else-if Else Statement**

```
if condition {
   // code to be executed if condition is true
} else if condition {
  // code to be executed if condition is true
else {
  // code to be executed if both conditions are false
}
```

**Boolean**

```
func main() {
    var bVal bool   // default is false
    fmt.Printf("bVal: %v\n", bVal)
}
```

**Switch Statement**

```
func main() {
	
	switch day:=4; day{
	case 1:
	fmt.Println("Monday")
	case 2:
	fmt.Println("Tuesday")
	case 3:
	fmt.Println("Wednesday")
	case 4:
	fmt.Println("Thursday")
	case 5:
	fmt.Println("Friday")
	case 6:
	fmt.Println("Saturday")
	case 7:
	fmt.Println("Sunday")
	default:
	fmt.Println("Invalid")
  }
}
```

### User Input Validation
