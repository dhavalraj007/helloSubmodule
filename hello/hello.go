package hello

import "fmt"
import "rsc.io/quote"

func Hello(name string) string {
	message:= fmt.Sprintf("Hi %v, a quote for you: %v",name,quote.Go())
	return message
}