package sprint
func SimpleStrToInt(s string) int {
   
    result := 0

    for _, ch := range s {
        
        digit := int(ch - '0')

        if digit < 0 || digit > 9 {
            
            return 0
        }

        result = result*10 + digit
    }

    return result
}

func main() {
    
    println(SimpleStrToInt("0000000012345")) 
}