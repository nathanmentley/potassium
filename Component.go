/*
Copyright 2019 Nathan Mentley

Permission is hereby granted, free of charge, to any person obtaining a copy of this software and associated documentation files (the "Software"), to deal in the Software without restriction, including without limitation the rights to use, copy, modify, merge, publish, distribute, sublicense, and/or sell copies of the Software, and to permit persons to whom the Software is furnished to do so, subject to the following conditions:

The above copyright notice and this permission notice shall be included in all copies or substantial portions of the Software.

THE SOFTWARE IS PROVIDED "AS IS", WITHOUT WARRANTY OF ANY KIND, EXPRESS OR IMPLIED, INCLUDING BUT NOT LIMITED TO THE WARRANTIES OF MERCHANTABILITY, FITNESS FOR A PARTICULAR PURPOSE AND NONINFRINGEMENT. IN NO EVENT SHALL THE AUTHORS OR COPYRIGHT HOLDERS BE LIABLE FOR ANY CLAIM, DAMAGES OR OTHER LIABILITY, WHETHER IN AN ACTION OF CONTRACT, TORT OR OTHERWISE, ARISING FROM, OUT OF OR IN CONNECTION WITH THE SOFTWARE OR THE USE OR OTHER DEALINGS IN THE SOFTWARE.
*/

package potassium

type Component struct {
    parent IComponentProcessor
    componentCache map[string]IComponentProcessor
}
func NewComponent(parent IComponentProcessor) Component {
    return Component{parent, make(map[string]IComponentProcessor)}
}

func (c *Component) SetInitialState(props map[string]interface{}) IState {
    return EmptyState{}
}
func (c *Component) ShouldComponentUpdate(processor IComponentProcessor) bool {
    return true
}
func (c *Component) ComponentDidMount(processor IComponentProcessor) {}
func (c *Component) ComponentWillUpdate(processor IComponentProcessor) {}
func (c *Component) ComponentDidUpdate(processor IComponentProcessor) {}
func (c *Component) ComponentWillUnmount(processor IComponentProcessor) {}
func (c *Component) getParent() IComponentProcessor {
    return c.parent
}

func (c *Component) CreateElement(key ComponentKey, componentBuilder func(IComponentProcessor) IComponent, props map[string]interface{}, children []IComponentProcessor) IComponentProcessor {
    if component, ok := c.componentCache[key.String()]; ok {
        component.setProps(props)
        component.updateChildren(children)
        return component
    }
        
    component := NewComponentWrapper(key, componentBuilder, props, children)
    c.componentCache[key.String()] = component

    return component
}
func (c *Component) clearComponentFromCache(key ComponentKey) {
    delete(c.componentCache, key.String())
}
