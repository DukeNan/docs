/*
 * @Author: shaun
 * @Date: 2021-06-30 11:31:17
 * @Last Modified by: shaun
 * @Last Modified time: 2021-06-30 12:29:28
 * 生成器接口
 */
package main

type iBuilder interface {
	setWindowType()
	setDoorType()
	setNumFloor()
	getHouse() house
}

func getBuilder(builderType string) iBuilder {
	if builderType == "normal" {
		return &normalBuilder{}
	}
	if builderType == "igloo" {
		return &iglooBuilder{}
	}
	return nil
}
