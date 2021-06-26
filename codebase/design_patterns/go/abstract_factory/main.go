/*
 * @Author: shaun
 * @Date: 2021-06-26 19:40:58
 * @Last Modified by: shaun
 * @Last Modified time: 2021-06-26 19:48:55
 */
package main

import "fmt"

func main() {
	adidasFactory, _ := getSportsFactory("adidas")
	nikeFactory, _ := getSportsFactory("nike")

	nikeShoe := nikeFactory.makeShoe()
	nikeShirt := nikeFactory.makeShirt()

	adidasShoe := adidasFactory.makeShoe()
	adidasShirt := adidasFactory.makeShirt()

	printShoeDetails(nikeShoe)
	printShirtDetails(nikeShirt)

	printShoeDetails(adidasShoe)
	printShirtDetails(adidasShirt)

}

func printShoeDetails(s iShoe) {
	fmt.Printf("Shoe logo: %s\n", s.getLogo())
	fmt.Printf("Shoe size: %d\n", s.getSize())
}

func printShirtDetails(s iShirt) {
	fmt.Printf("Shirt logo: %s\n", s.getLogo())
	fmt.Printf("Shirt size: %d\n", s.getSize())
}
