/*
 * @Author: shaun
 * @Date: 2022-01-04 18:03:18
 * @Last Modified by: shaun
 * @Last Modified time: 2022-01-04 18:25:42
 */
package main

import "fmt"

type windowsAdapter struct {
	windowMachine *windows
}

func (w *windowsAdapter) insertIntoLightningPort() {
	fmt.Println("Adapter converts Lightning signal to USB.")
	w.windowMachine.insertIntoUSBPort()
}
