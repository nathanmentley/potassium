/*
Copyright 2019 Nathan Mentley

Permission is hereby granted, free of charge, to any person obtaining a copy of this software and associated documentation files (the "Software"), to deal in the Software without restriction, including without limitation the rights to use, copy, modify, merge, publish, distribute, sublicense, and/or sell copies of the Software, and to permit persons to whom the Software is furnished to do so, subject to the following conditions:

The above copyright notice and this permission notice shall be included in all copies or substantial portions of the Software.

THE SOFTWARE IS PROVIDED "AS IS", WITHOUT WARRANTY OF ANY KIND, EXPRESS OR IMPLIED, INCLUDING BUT NOT LIMITED TO THE WARRANTIES OF MERCHANTABILITY, FITNESS FOR A PARTICULAR PURPOSE AND NONINFRINGEMENT. IN NO EVENT SHALL THE AUTHORS OR COPYRIGHT HOLDERS BE LIABLE FOR ANY CLAIM, DAMAGES OR OTHER LIABILITY, WHETHER IN AN ACTION OF CONTRACT, TORT OR OTHERWISE, ARISING FROM, OUT OF OR IN CONNECTION WITH THE SOFTWARE OR THE USE OR OTHER DEALINGS IN THE SOFTWARE.
*/

package potassium

import (
    "strconv"
    "reflect"
    "runtime"
)

type Component struct {
    parent IComponentProcessor
    componentCache map[string]IComponentProcessor
    elementCounter int
}
func NewComponent(parent IComponentProcessor) Component {
    return Component{parent, make(map[string]IComponentProcessor), 0}
}

func (c *Component) SetInitialState(props map[string]interface{}) IState {
    return EmptyState{}
}
func (c *Component) ShouldComponentUpdate(processor IComponentProcessor) bool {
    return true
}
func (c *Component) ComponentDidMount(processor IComponentProcessor) {
    c.elementCounter = 0
}
func (c *Component) ComponentWillUpdate(processor IComponentProcessor) {
    c.elementCounter = 0
}
func (c *Component) ComponentDidUpdate(processor IComponentProcessor) {}
func (c *Component) ComponentWillUnmount(processor IComponentProcessor) {}
func (c *Component) getParent() IComponentProcessor {
    return c.parent
}

func (c *Component) CreateElement(componentBuilder func(IComponentProcessor) IComponent, props map[string]interface{}, children []IComponentProcessor) IComponentProcessor {
    c.elementCounter++

    key := newComponentKey(getFunctionName(componentBuilder) + strconv.Itoa(c.elementCounter))
    if val, ok := props["key"]; ok {
        if str, ok := val.(string); ok {
            key = newComponentKey(str)
        }
    }

    if component, ok := c.componentCache[key.String()]; ok {
        component.setProps(props)
        component.updateChildren(children)
        return component
    }

    component := newComponentWrapper(key, componentBuilder, props, children)
    c.componentCache[key.String()] = component

    return component
}
func (c *Component) clearComponentFromCache(key componentKey) {
    delete(c.componentCache, key.String())
}


func getFunctionName(i interface{}) string {
    return runtime.FuncForPC(reflect.ValueOf(i).Pointer()).Name()
}
