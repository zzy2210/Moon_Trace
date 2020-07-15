package subdomain

import (
	"Moon_Trace/subdomain/modules"
	"fmt"
	"github.com/fatih/color"
)

func FindSubdomain(tg string){ // use function to find subdomain and organize data
	color.Yellow("Just standby ")
	c := color.New(color.FgCyan)
	subdomain := modules.Find(tg) // subdomain is []string
	subdomain = unique(subdomain)


	for n,_ := range subdomain{
		c.Print("[Sub]:")
		fmt.Println(subdomain[n])
	}
}

	func unique(ataxic []string) []string{
		//I want to use simHash
		//But now,I don't know how to implement the code
		//So only index of
		var unique []string
		for _,value := range ataxic {
			if !indexOf(value,unique){
				unique = append(unique,value)
			}
		}

		return unique
	}

	func indexOf(atom string,array []string) bool{
		// Did atom in array?
		for _,value := range array {
			if atom == value {
				return true
			}
		}
		return false
	}