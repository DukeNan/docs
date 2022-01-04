/*
 * @Author: shaun
 * @Date: 2022-01-04 17:48:21
 * @Last Modified by: shaun
 * @Last Modified time: 2022-01-04 18:24:45
 */
package main

import "fmt"

type client struct {
}

func (c *client) insertLightningConnectorIntoComputer(com computer) {
	fmt.Println("Client inserts Lightning connector into computer.")
	com.insertIntoLightningPort()
}
