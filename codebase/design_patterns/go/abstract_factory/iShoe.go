/*
 * @Author: shaun
 * @Date: 2021-06-26 18:18:24
 * @Last Modified by: shaun
 * @Last Modified time: 2021-06-26 18:21:13
 * 抽象产品
 */
package main

type iShoe interface {
	setLogo(logo string)
	setSize(size int)
	getLogo() string
	getSize() int
}
type shoe struct {
	logo string
	size int
}

func (s *shoe) setLogo(logo string) {
	s.logo = logo
}

func (s *shoe) setSize(size int) {
	s.size = size
}

func (s *shoe) getLogo() string {
	return s.logo
}

func (s *shoe) getSize() int {
	return s.size
}
