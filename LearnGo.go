package main
import ("fmt")

func myMessage() {
  fmt.Println("myMessage()")
}

// Struct
type Person struct {
  name string
  age int
  job string
  salary int
}

func main() {

  // Variables
  var i int = 1
  var s string = "s"
  var b bool = true
  fmt.Println(i)
  fmt.Println(s)
  fmt.Println(b)
  
  
  // Arrays
  var arr1 = [3]int{1,2,3}
  fmt.Println(arr1)
  
  // If else
  if 20 > 18 {
    fmt.Println("20 is greater than 18")
  }
  
  // Switch
  day := 4
  switch day {
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
  }
  
  // For loop
  for i:=0; i < 5; i++ {
    fmt.Println(i)
  }
  
  // Function
  myMessage()
  
  // Maps
  var a = map[string]string{"brand": "Ford", "model": "Mustang", "year": "1964"}
  b := map[string]int{"Oslo": 1, "Bergen": 2, "Trondheim": 3, "Stavanger": 4}
  fmt.Printf("a\t%v\n", a)
  fmt.Printf("b\t%v\n", b)
}