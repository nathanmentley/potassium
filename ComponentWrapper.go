/*
Copyright 2019 Nathan Mentley

Permission is hereby granted, free of charge, to any person obtaining a copy of this software and associated documentation files (the "Software"), to deal in the Software without restriction, including without limitation the rights to use, copy, modify, merge, publish, distribute, sublicense, and/or sell copies of the Software, and to permit persons to whom the Software is furnished to do so, subject to the following conditions:

The above copyright notice and this permission notice shall be included in all copies or substantial portions of the Software.

THE SOFTWARE IS PROVIDED "AS IS", WITHOUT WARRANTY OF ANY KIND, EXPRESS OR IMPLIED, INCLUDING BUT NOT LIMITED TO THE WARRANTIES OF MERCHANTABILITY, FITNESS FOR A PARTICULAR PURPOSE AND NONINFRINGEMENT. IN NO EVENT SHALL THE AUTHORS OR COPYRIGHT HOLDERS BE LIABLE FOR ANY CLAIM, DAMAGES OR OTHER LIABILITY, WHETHER IN AN ACTION OF CONTRACT, TORT OR OTHERWISE, ARISING FROM, OUT OF OR IN CONNECTION WITH THE SOFTWARE OR THE USE OR OTHER DEALINGS IN THE SOFTWARE.
*/

package potassium

type ComponentWrapper struct {
    key ComponentKey                                        //unique random id (used to keep track in dom)
    componentBuilder func(IComponentProcessor) IComponent   //logic to create instances of the component
    props IProps                                            //current props value
    children []IComponentProcessor                          //all children instances

    toolkit IAppToolkit                                     //a reference to the toolkit being used to build the gui
    component IComponent                                    //component instance
    state IState                                            //current state value
    previousResult *RenderResult                            //last result (used to return the previous result if there is no reason to update)
}

func NewComponentWrapper(key ComponentKey, componentBuilder func(IComponentProcessor) IComponent, props IProps, children []IComponentProcessor) IComponentProcessor {
    return &ComponentWrapper{key, componentBuilder, props, children, nil, nil, nil, nil}
}

//Since we don't build the component instance when the wrapper is defined. This logic will do that. once we're actaully mounting the component.
func (c *ComponentWrapper) init(parent IComponentProcessor, toolkit IAppToolkit) {
    c.toolkit = toolkit
    c.component = c.componentBuilder(parent)
    c.state = c.component.SetInitialState(c.props)
}

//IComponentProcessor implemented on the pointer
func (c *ComponentWrapper) SetState(state IState) {
    c.state = state
    c.render()
}
func (c *ComponentWrapper) GetState() IState {
    return c.state
}
func (c *ComponentWrapper) GetProps() IProps {
    return c.props
}
func (c *ComponentWrapper) GetChildren() []IComponentProcessor {
    return c.children
}
func (c *ComponentWrapper) GetParent() IComponentProcessor {
    if c.component != nil {
        return c.component.getParent()
    }
    return nil
}
func (c *ComponentWrapper) GetComponent() IComponent {
    return c.component
}
func (c *ComponentWrapper) GetKey() ComponentKey {
    return c.key
}

func (c *ComponentWrapper) setProps(props IProps) {
    c.props = props
}

func isValueInList(value IComponentProcessor, list []IComponentProcessor) bool {
    for _, v := range list {
        if v.GetKey().String() == value.GetKey().String() {
            return true
        }
    }
    return false
}
func (c *ComponentWrapper) updateChildren(children []IComponentProcessor) {
    //unmount missing components.
    toUnmount := []IComponentProcessor{}

    //add missing components from new slice to unmount slice
    for _, child := range c.children {
        if !isValueInList(child, children) {
            toUnmount = append(toUnmount, child)
        }
    }

    for _, child := range toUnmount {
        child.unmount()
        c.component.clearComponentFromCache(child.GetKey())

        if c.toolkit != nil {
            c.toolkit.Unmount(c, child)
        }
    }

    //replace array
    c.children = children
}


func (c *ComponentWrapper) mount(parent IComponentProcessor, toolkit IAppToolkit) bool {
    if c.component == nil {
        //if component isn't created let's init it.
        c.init(parent, toolkit)
        c.component.ComponentDidMount(c)

        return true
    }

    return false
}

func (c *ComponentWrapper) render() *RenderResult {
    if c.previousResult == nil || c.component.ShouldComponentUpdate(c) {
        c.component.ComponentWillUpdate(c)
        ret := c.component.Render(c)

        for index, child := range ret.Children {
            didMount := child.mount(c, c.toolkit)
            child.render()

            if didMount {
                if c.toolkit != nil {
                    c.toolkit.Mount(c, child, index)
                }
            } else {
                if c.toolkit != nil {
                    c.toolkit.EnsureMount(c, child, index)
                }
            }
        }

        c.component.ComponentDidUpdate(c)

        //clean up components that no longer exist from previous result.
        if c.previousResult != nil {
            toUnmount := []IComponentProcessor{}

            //add missing components from new slice to unmount slice
            for _, child := range c.previousResult.Children {
                if !isValueInList(child, ret.Children) {
                    toUnmount = append(toUnmount, child)
                }
            }

            for _, child := range toUnmount {
                child.unmount()
                c.component.clearComponentFromCache(child.GetKey())

                if c.toolkit != nil {
                    c.toolkit.Unmount(c, child)
                }
            }
        }

        c.previousResult = ret
        return ret 
    }

    return c.previousResult
}

func (c *ComponentWrapper) unmount() {
    if c.previousResult != nil {
        for _, child := range c.previousResult.Children {
            child.unmount()
            c.component.clearComponentFromCache(child.GetKey())

            if c.toolkit != nil {
                c.toolkit.Unmount(c, child)
            }
        }
    }

    c.component.ComponentWillUnmount(c)
}
