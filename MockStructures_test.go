/*
Copyright 2019 Nathan Mentley

Permission is hereby granted, free of charge, to any person obtaining a copy of this software and associated documentation files (the "Software"), to deal in the Software without restriction, including without limitation the rights to use, copy, modify, merge, publish, distribute, sublicense, and/or sell copies of the Software, and to permit persons to whom the Software is furnished to do so, subject to the following conditions:

The above copyright notice and this permission notice shall be included in all copies or substantial portions of the Software.

THE SOFTWARE IS PROVIDED "AS IS", WITHOUT WARRANTY OF ANY KIND, EXPRESS OR IMPLIED, INCLUDING BUT NOT LIMITED TO THE WARRANTIES OF MERCHANTABILITY, FITNESS FOR A PARTICULAR PURPOSE AND NONINFRINGEMENT. IN NO EVENT SHALL THE AUTHORS OR COPYRIGHT HOLDERS BE LIABLE FOR ANY CLAIM, DAMAGES OR OTHER LIABILITY, WHETHER IN AN ACTION OF CONTRACT, TORT OR OTHERWISE, ARISING FROM, OUT OF OR IN CONNECTION WITH THE SOFTWARE OR THE USE OR OTHER DEALINGS IN THE SOFTWARE.
*/

package potassium

//Mock Classes
type MockState struct {}

type MockComponentProcessor struct {
    key string
}
func (m *MockComponentProcessor) SetState(state IState) {}
func (m *MockComponentProcessor) GetState() IState { return nil }
func (m *MockComponentProcessor) GetProps() map[string]interface{} { return nil }
func (m *MockComponentProcessor) GetChildren() []IComponentProcessor { return nil }
func (m *MockComponentProcessor) GetParent() IComponentProcessor { return nil }
func (m *MockComponentProcessor) GetComponent() IComponent { return nil }

func (m *MockComponentProcessor) getKey() componentKey { return newComponentKey(m.key) }
func (m *MockComponentProcessor) setProps(props map[string]interface{}) {}
func (m *MockComponentProcessor) updateChildren(children []IComponentProcessor) {}

func (m *MockComponentProcessor) mount(parent IComponentProcessor, toolkit IAppToolkit) bool { return true }
func (m *MockComponentProcessor) render() *RenderResult { return nil }
func (m *MockComponentProcessor) unmount() {}

type MockComponent struct {
    Component
}
func (m *MockComponent) SetInitialState(props map[string]interface{}) IState { return &MockState{} }
func (m *MockComponent) Render(processor IComponentProcessor) *RenderResult { return &RenderResult{[]IComponentProcessor{}} }
func NewMockComponent(parent IComponentProcessor) IComponent {
    return &MockComponent{NewComponent(parent)}
}
