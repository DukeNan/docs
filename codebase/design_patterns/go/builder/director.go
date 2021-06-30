/*
 * @Author: shaun
 * @Date: 2021-06-30 12:30:00
 * @Last Modified by: shaun
 * @Last Modified time: 2021-06-30 12:36:34
 * 主管
 */
package main

type director struct {
	builder iBuilder
}

func newDirector(b iBuilder) *director {
	return &director{
		builder: b,
	}
}

func (d *director) setBuilder(b iBuilder) {
	d.builder = b

}

func (d *director) builderHouse() house {
	d.builder.setDoorType()
	d.builder.setWindowType()
	d.builder.setNumFloor()
	return d.builder.getHouse()
}
