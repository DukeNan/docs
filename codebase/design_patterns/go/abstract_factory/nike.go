/*
 * @Author: shaun
 * @Date: 2021-06-26 19:29:45
 * @Last Modified by: shaun
 * @Last Modified time: 2021-06-26 19:38:50
 * 具体工厂
 */

package main

type nike struct {
}

func (n *nike) makeShoe() iShoe {
	return &nikeShoe{
		shoe: shoe{
			logo: "nike",
			size: 14,
		},
	}
}

func (n *nike) makeShirt() iShirt {
	return &nikeShirt{
		shirt: shirt{
			logo: "nike",
			size: 14,
		},
	}
}
