package ui

import (
	"fmt"
	"github.com/google/uuid"
	"github.com/hajimehoshi/ebiten/v2"
	"math"
)

type ComponentDataIteratorType int

const (
	NormalComponentDataIteratorType ComponentDataIteratorType = iota
	ReverseComponentDataIteratorType
)

type Page interface {
	OnDraw(*ebiten.Image)
	HandleClick(float64, float64) bool
}

type Component interface {
	OnDraw(*ebiten.Image)
	HandleClick(float64, float64) bool
	IsWithin(float64, float64) bool
	GetEndGeometryX() float64
	GetEndGeometryY() float64
	GetStartGeometryX() float64
	GetStartGeometryY() float64
	SetStartGeometryX(float64)
	SetStartGeometryY(float64)
}

type ComponentWithData interface {
	GetComponentData() *ComponentData
	SetComponentData(*ComponentData)
}

type ComponentData struct {
	data  map[string]Component
	order []string
}

func (componentData *ComponentData) Add(component Component) *ComponentData {
	componentData.AddById(component, componentData.generateUUID())

	return componentData
}

func (componentData *ComponentData) AddById(component Component, newID string) *ComponentData {
	if componentData.HasById(newID) {
		return componentData
	}

	componentData.data[newID] = component
	componentData.order = append(componentData.order, newID)

	return componentData
}

func (componentData *ComponentData) AddToBegin(component Component) *ComponentData {
	componentData.AddToBeginById(component, componentData.generateUUID())

	return componentData
}

func (componentData *ComponentData) AddToBeginById(component Component, newID string) *ComponentData {
	if componentData.HasById(newID) {
		return componentData
	}

	componentData.data[newID] = component
	componentData.order = append([]string{newID}, componentData.order...)

	return componentData
}

func (componentData *ComponentData) GetById(id string) (*Component, bool) {
	if componentData.HasById(id) {
		value, exists := componentData.data[id]

		return &value, exists
	}

	for _, component := range componentData.data {
		if compWithData, ok := component.(ComponentWithData); ok {
			value, exists := compWithData.GetComponentData().GetById(id)

			if exists {
				return value, exists
			}
		}
	}

	return nil, false
}

func (componentData *ComponentData) GetByIndex(index int) (*Component, bool) {
	return componentData.GetById(componentData.GetIdByIndex(index))
}

func (componentData *ComponentData) HasById(id string) bool {
	_, exists := componentData.data[id]

	return exists
}

func (componentData *ComponentData) HasByIndex(index int) bool {
	return componentData.HasById(componentData.GetIdByIndex(index))
}

func (componentData *ComponentData) RemoveByIndex(index int) *ComponentData {
	delete(componentData.data, componentData.GetIdByIndex(index))

	componentData.order = append(componentData.order[:index], componentData.order[index+1:]...)

	return componentData
}

func (componentData *ComponentData) RemoveById(id string) *ComponentData {
	for index, value := range componentData.order {
		if value == id {
			componentData.RemoveByIndex(index)

			break
		}
	}

	return componentData
}

func (componentData *ComponentData) UpdateByIndex(component Component, index int) *ComponentData {
	componentData.UpdateById(component, componentData.GetIdByIndex(index))

	return componentData
}

func (componentData *ComponentData) UpdateById(component Component, id string) *ComponentData {
	componentData.data[id] = component

	return componentData
}

func (componentData *ComponentData) GetTotal() int {
	return len(componentData.order)
}

func (componentData *ComponentData) GetIdByIndex(index int) string {
	componentDataTotal := componentData.GetTotal()

	if componentDataTotal <= index || index < 0 {
		panic(fmt.Sprintf("Index %d not found!", index))
	}

	return componentData.order[index]
}

func (componentData *ComponentData) GetNormalComponentDataIterator() ComponentDataIterator {
	return NewBuilderComponentDataIterator().
		SetIteratorType(NormalComponentDataIteratorType).
		SetComponentData(componentData).
		GetComponentDataIterator()
}

func (componentData *ComponentData) GetReverseComponentDataIterator() ComponentDataIterator {
	return NewBuilderComponentDataIterator().
		SetIteratorType(ReverseComponentDataIteratorType).
		SetComponentData(componentData).
		GetComponentDataIterator()
}

func (componentData *ComponentData) generateUUID() string {
	u, err := uuid.NewRandom()

	if err != nil {
		return componentData.generateUUID()
	}

	return u.String()
}

type ComponentDataIterator interface {
	HasNext() bool
	GetNext() *Component
	GetId() string
}

type NormalComponentDataIterator struct {
	index         int
	componentData *ComponentData
}

func (iterator *NormalComponentDataIterator) GetId() string {
	return iterator.componentData.GetIdByIndex(iterator.index - 1)

}

func (iterator *NormalComponentDataIterator) HasNext() bool {
	return iterator.index < iterator.componentData.GetTotal()

}

func (iterator *NormalComponentDataIterator) GetNext() *Component {
	if iterator.HasNext() {
		component, exists := iterator.componentData.GetByIndex(iterator.index)

		iterator.index++

		if exists {
			return component
		}
	}

	return nil
}

type ReverseComponentDataIterator struct {
	index         int
	componentData *ComponentData
}

func (iterator *ReverseComponentDataIterator) GetId() string {
	return iterator.componentData.GetIdByIndex(iterator.index - 1)

}

func (iterator *ReverseComponentDataIterator) HasNext() bool {
	return iterator.index >= 0

}

func (iterator *ReverseComponentDataIterator) GetNext() *Component {
	if iterator.HasNext() {
		component, exists := iterator.componentData.GetByIndex(iterator.index)

		iterator.index--

		if exists {
			return component
		}
	}

	return nil
}

type BuilderComponentDataIterator struct {
	index         int
	componentData *ComponentData
	iteratorType  ComponentDataIteratorType
}

func NewBuilderComponentDataIterator() *BuilderComponentDataIterator {
	return &BuilderComponentDataIterator{
		index:         math.MinInt,
		componentData: NewBuilderComponentData().GetComponentData(),
		iteratorType:  NormalComponentDataIteratorType,
	}
}

func (builder *BuilderComponentDataIterator) GetComponentDataIterator() ComponentDataIterator {
	switch builder.iteratorType {
	case ReverseComponentDataIteratorType:
		index := builder.componentData.GetTotal() - 1

		if builder.index != math.MinInt {
			index = builder.index
		}

		return &ReverseComponentDataIterator{
			index:         index,
			componentData: builder.componentData,
		}
	default:
		index := 0

		if builder.index != math.MinInt {
			index = builder.index
		}

		return &NormalComponentDataIterator{
			index:         index,
			componentData: builder.componentData,
		}
	}
}

func (builder *BuilderComponentDataIterator) SetIndex(index int) *BuilderComponentDataIterator {
	builder.index = index

	return builder
}

func (builder *BuilderComponentDataIterator) SetComponentData(componentData *ComponentData) *BuilderComponentDataIterator {
	builder.componentData = componentData

	return builder
}

func (builder *BuilderComponentDataIterator) SetIteratorType(iteratorType ComponentDataIteratorType) *BuilderComponentDataIterator {
	builder.iteratorType = iteratorType

	return builder
}

type BuilderComponentData struct {
	componentData *ComponentData
}

func NewBuilderComponentData() *BuilderComponentData {
	return &BuilderComponentData{
		componentData: &ComponentData{
			data:  make(map[string]Component),
			order: []string{},
		},
	}
}

func (builder *BuilderComponentData) GetComponentData() *ComponentData {
	return builder.componentData
}

func (builder *BuilderComponentData) SetComponentData(componentData *ComponentData) *BuilderComponentData {
	builder.componentData = componentData

	return builder
}

func (builder *BuilderComponentData) AddComponent(component Component) *BuilderComponentData {
	builder.componentData.Add(component)

	return builder
}

func (builder *BuilderComponentData) AddComponentById(component Component, newID string) *BuilderComponentData {
	builder.componentData.AddById(component, newID)

	return builder
}

func (builder *BuilderComponentData) AddComponentToBegin(component Component) *BuilderComponentData {
	builder.componentData.AddToBegin(component)

	return builder
}

func (builder *BuilderComponentData) AddComponentToBeginById(component Component, newID string) *BuilderComponentData {
	builder.componentData.AddToBeginById(component, newID)

	return builder
}
