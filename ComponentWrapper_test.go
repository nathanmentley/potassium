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

func TestComponentWrapperGetParent(t *testing.T) {
    component := NewComponent(nil)
    props := make(map[string]interface{})
    componentWrapper := component.CreateElement(NewMockComponent, props, []IComponentProcessor{})

    if componentWrapper.GetParent() != nil {
        t.Errorf("Expected component.getParent() to match passed in parent processor.")
    }
}
/*
func TestComponentWrapperGetProps(t *testing.T) {
    component := NewComponent(nil)
    props := map[string]interface{}{}
    componentWrapper := component.CreateElement(NewMockComponent, props, []IComponentProcessor{})

    if reflect.DeepEqual(componentWrapper.GetProps(), props) {
        t.Errorf("Expected component.GetProps() should get the props.")
    }

    props2 := map[string]interface{}{}
    componentWrapper.setProps(props2)
    if !reflect.DeepEqual(componentWrapper.GetProps(), props2) {
        t.Errorf("Expected component.GetProps() should update the props.")
    }
}
*/
