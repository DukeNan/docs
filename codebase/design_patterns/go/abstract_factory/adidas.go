/*
 * @Author: shaun
 * @Date: 2021-06-26 18:15:12
 * @Last Modified by: shaun
 * @Last Modified time: 2021-06-26 19:27:14
 * 具体工厂
 */

package main

type adidas struct {
}

func (a *adidas) makeShoe() iShoe {
	return &adidasShoe{
		shoe: shoe{
			logo: "adidas",
			size: 14,
		},
	}
}

func (a *adidas) makeShirt() iShirt {
	return &adidasShirt{
		shirt: shirt{
			logo: "adidas",
			size: 14,
		},
	}
}
