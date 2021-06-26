/*
 * @Author: shaun
 * @Date: 2021-06-26 18:08:57
 * @Last Modified by: shaun
 * @Last Modified time: 2021-06-26 19:40:44
 * 抽象工厂接口
 */
package main

import "fmt"

type iSportsFactory interface {
	makeShoe() iShoe
	makeShirt() iShirt
}

func getSportsFactory(brand string) (iSportsFactory, error) {
	if brand == "adidas" {
		return &adidas{}, nil
	}
	if brand == "nike" {
		return &nike{}, nil
	}
	return nil, fmt.Errorf("wrong brand type passed")
}
