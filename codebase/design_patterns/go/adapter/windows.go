/*
 * @Author: shaun
 * @Date: 2022-01-04 18:01:50
 * @Last Modified by: shaun
 * @Last Modified time: 2022-01-04 18:25:29
 */
package main

import "fmt"

type windows struct{}

func (w *windows) insertIntoUSBPort() {
	fmt.Println("USB connector is plugged into windows machine.")
}
