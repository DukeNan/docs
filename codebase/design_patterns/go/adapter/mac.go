/*
 * @Author: shaun
 * @Date: 2022-01-04 17:59:26
 * @Last Modified by: shaun
 * @Last Modified time: 2022-01-04 18:25:16
 */
package main

import "fmt"

type mac struct {
}

func (m *mac) insertIntoLightningPort() {
	fmt.Println("Lightning connector is plugged into mac machine.")
}
