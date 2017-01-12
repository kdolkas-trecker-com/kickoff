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

func main(){
    var first int ;
    var second int ;
    var number string ;
    var err error;
    var operation string;

    
    fmt.Println("Enter the calculation you want")
    reader := bufio.NewReader(os.Stdin)
    input, _ := reader.ReadString('\n')
    var calculation = strings.TrimRight(input, "\n")
 
 	for  i := 0; i < len(calculation); i++ {
         yo := string([]rune(calculation)[i])
         fmt.Println(yo)
         if(yo == "*" || yo == "+" || yo == "-" || yo == "/") {
             operation = yo ;
             first,err = strconv.Atoi(number);
             number = "";
             checkError(err)
         }else{
             number = number + yo
         }     
     }
     second,err = strconv.Atoi(number);
     checkError(err)
     fmt.Println(" So the result is:")
     fmt.Print(first, operation, second,"=")
     if (operation == "*"){
         fmt.Println(multiply(first,second))
     }
     if (operation == "+"){
         fmt.Println(add(first,second))
     }
}
