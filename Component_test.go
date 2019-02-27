/*
Copyright 2019 Nathan Mentley

Permission is hereby granted, free of charge, to any person obtaining a copy of this software and associated documentation files (the "Software"), to deal in the Software without restriction, including without limitation the rights to use, copy, modify, merge, publish, distribute, sublicense, and/or sell copies of the Software, and to permit persons to whom the Software is furnished to do so, subject to the following conditions:

The above copyright notice and this permission notice shall be included in all copies or substantial portions of the Software.

THE SOFTWARE IS PROVIDED "AS IS", WITHOUT WARRANTY OF ANY KIND, EXPRESS OR IMPLIED, INCLUDING BUT NOT LIMITED TO THE WARRANTIES OF MERCHANTABILITY, FITNESS FOR A PARTICULAR PURPOSE AND NONINFRINGEMENT. IN NO EVENT SHALL THE AUTHORS OR COPYRIGHT HOLDERS BE LIABLE FOR ANY CLAIM, DAMAGES OR OTHER LIABILITY, WHETHER IN AN ACTION OF CONTRACT, TORT OR OTHERWISE, ARISING FROM, OUT OF OR IN CONNECTION WITH THE SOFTWARE OR THE USE OR OTHER DEALINGS IN THE SOFTWARE.
*/

package potassium

import (
    "testing"
)

//Mock Classes

type MockProps struct {}
type MockState struct {}

type MockComponentProcessor struct {
    key string
}
func (m *MockComponentProcessor) SetState(state IState) {}
func (m *MockComponentProcessor) GetState() IState { return nil }
func (m *MockComponentProcessor) GetProps() IProps { return nil }
func (m *MockComponentProcessor) GetChildren() []IComponentProcessor { return nil }
func (m *MockComponentProcessor) GetParent() IComponentProcessor { return nil }
func (m *MockComponentProcessor) GetComponent() IComponent { return nil }
func (m *MockComponentProcessor) GetKey() ComponentKey { return NewComponentKey(m.key) }

func (m *MockComponentProcessor) setProps(props IProps) {}
func (m *MockComponentProcessor) updateChildren(children []IComponentProcessor) {}

func (m *MockComponentProcessor) mount(parent IComponentProcessor, toolkit IAppToolkit) bool { return true }
func (m *MockComponentProcessor) render() *RenderResult { return nil }
func (m *MockComponentProcessor) unmount() {}

type MockComponent struct {
    Component
}
func (m *MockComponent) SetInitialState(props IProps) IState { return &MockState{} }
func (m *MockComponent) Render(processor IComponentProcessor) *RenderResult { return nil }
func NewMockComponet(parent IComponentProcessor) IComponent {
    return &MockComponent{NewComponent(parent)}
}


func TestComponentGetParent(t *testing.T) {
    if component := NewComponent(nil); component.getParent() != nil {
        t.Errorf("Expected component.getParent() to return nil if there is no parent.")
    }

    parent := &MockComponentProcessor{"parent"}
    if component := NewComponent(parent); component.getParent() != parent {
        t.Errorf("Expected component.getParent() to match passed in parent processor.")
    }
}


func TestComponentCreateElementCache(t *testing.T) {
    component := NewComponent(nil)

    childComponent1 := component.CreateElement(NewComponentKey("TestKey 1"), NewMockComponet, &MockProps{}, []IComponentProcessor{})
    childComponent2 := component.CreateElement(NewComponentKey("TestKey 2"), NewMockComponet, &MockProps{}, []IComponentProcessor{})

    childComponent1Expected := component.CreateElement(NewComponentKey("TestKey 1"), NewMockComponet, &MockProps{}, []IComponentProcessor{})
    childComponent2Expected := component.CreateElement(NewComponentKey("TestKey 2"), NewMockComponet, &MockProps{}, []IComponentProcessor{})

    if childComponent1 != childComponent1Expected {
        t.Errorf("Component.CreateElement should recycle components with the same key")
    }
    if childComponent2 != childComponent2Expected {
        t.Errorf("Component.CreateElement should recycle components with the same key")
    }
    if childComponent2 == childComponent1Expected {
        t.Errorf("Component.CreateElement should recycle components only with components of the same key")
    }
    if childComponent1 == childComponent2Expected {
        t.Errorf("Component.CreateElement should recycle components only with components of the same key")
    }
}
