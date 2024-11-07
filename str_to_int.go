package sprint

func StrToInt(s string) int {
    result := 0

    sign := 1

    for i, ch := range s {
       
        if i == 0 && (ch == '+' || ch == '-') {
           
            if ch == '-' {
                sign = -1
            }
            continue 
        }

        digit := int(ch - '0')

        if digit < 0 || digit > 9 {
           
            return 0
        }

        result = result*10 + digit
    }

    return result * sign
}

func main() {
   
    println(StrToInt("12345"))    
    println(StrToInt("0000000012345")) 
    println(StrToInt("012 345"))  
    println(StrToInt("Hello World!")) 
    println(StrToInt("+1234"))     
    println(StrToInt("-1234"))     
    println(StrToInt("++1234"))    
    println(StrToInt("--1234"))    
}