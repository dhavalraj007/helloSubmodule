package farewells

import "fmt"
import "rsc.io/quote"

func Bye(name string) string {
	message:= fmt.Sprintf("Bye %v, a quote for you: %v",name,quote.Go())
	return message
}