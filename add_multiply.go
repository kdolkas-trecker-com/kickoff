package main
import "fmt"
import "strconv"
import "os"
import "bufio"
import "strings"

func checkError(err error){
    if err != nil {
        // handle error
        fmt.Println(err)
        os.Exit(2)
    }
}

func multiply(x int, y int) int {
    return x * y
}

func add(x int, y int) int {
    return x + y
}


func getinput() string {
    fmt.Println("Enter the calculation you want")
    reader := bufio.NewReader(os.Stdin)
    input, _ := reader.ReadString('\n')
    return strings.TrimRight(input, "\n")
}

func main(){
    var first int ;
    var second int ;
    var number string ;
    var err error;
    var operation string;

    // Get nput from the user
    var calculation = getinput();
 
    // Iterate through each character of the string  until you find the calculation sign
 	for  i := 0; i < len(calculation); i++ {
         yo := string([]rune(calculation)[i])
         if(yo == "*" || yo == "+" || yo == "-" || yo == "/") {
             operation = yo ;
             //move all the characters till now to first 
             first,err = strconv.Atoi(number);
             checkError(err)
             number = "";
         } else {
             number = number + yo
         }     
     }
    //Since the string ended and no calculation sign was found move the remaining string to second
    second,err = strconv.Atoi(number);
    checkError(err)
     
     fmt.Print("So the calculation is: ")
     fmt.Print(first, operation, second,"=")

    if (operation == "*"){
         fmt.Println(multiply(first,second))
     }
    if (operation == "+"){
         fmt.Println(add(first,second))
     }

}
