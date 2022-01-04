/*
 * @Author: shaun
 * @Date: 2022-01-04 18:06:49
 * @Last Modified by: shaun
 * @Last Modified time: 2022-01-04 18:23:54
 */
package main

func main() {

	client := &client{}
	mac := &mac{}

	client.insertLightningConnectorIntoComputer(mac)

	windowsMachine := &windows{}
	windowsMachineAdapter := &windowsAdapter{
		windowMachine: windowsMachine,
	}

	client.insertLightningConnectorIntoComputer(windowsMachineAdapter)
}
