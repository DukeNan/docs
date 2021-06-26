/*
 * @Author: shaun
 * @Date: 2021-06-26 19:15:44
 * @Last Modified by: shaun
 * @Last Modified time: 2021-06-26 19:18:30
 * 抽象产品
 */

package main

type iShirt interface {
	setLogo(logo string)
	setSize(size int)
	getLogo() string
	getSize() int
}

type shirt struct {
	logo string
	size int
}

func (s *shirt) setLogo(logo string) {
	s.logo = logo
}

func (s *shirt) setSize(size int) {
	s.size = size
}

func (s *shirt) getLogo() string {
	return s.logo
}

func (s *shirt) getSize() int {
	return s.size
}
