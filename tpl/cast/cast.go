// Copyright 2017 The Hugo Authors. All rights reserved.
//
// Licensed under the Apache License, Version 2.0 (the "License");
// you may not use this file except in compliance with the License.
// You may obtain a copy of the License at
// http://www.apache.org/licenses/LICENSE-2.0
//
// Unless required by applicable law or agreed to in writing, software
// distributed under the License is distributed on an "AS IS" BASIS,
// WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
// See the License for the specific language governing permissions and
// limitations under the License.

// Package cast provides template functions for data type conversions.
package cast

import (
	"html/template"

	_cast "github.com/spf13/cast"
)

// New returns a new instance of the cast-namespaced template functions.
func New() *Namespace {
	return &Namespace{}
}

// Namespace provides template functions for the "cast" namespace.
type Namespace struct {
}

// ToInt converts v to an int.
func (ns *Namespace) ToInt(v any) (int, error) {
	v = convertTemplateToString(v)
	return _cast.ToIntE(v)
}

// ToString converts v to a string.
func (ns *Namespace) ToString(v any) (string, error) {
	return _cast.ToStringE(v)
}

// ToFloat converts v to a float.
func (ns *Namespace) ToFloat(v any) (float64, error) {
	v = convertTemplateToString(v)
	return _cast.ToFloat64E(v)
}

// ToBool converts v to a boolean.
func (ns *Namespace) ToBool(v any) (bool, error) {
	v = convertTemplateToString(v)
	return _cast.ToBoolE(v)
}

// ToTruth yields the same behavior as ToBool when possible.
// If the cast is unsuccessful, ToTruth converts v to a boolean using the JavaScript [definition of truthy](https://developer.mozilla.org/en-US/docs/Glossary/Truthy).
// Accordingly, it never yields an error, but maintains the signature of other cast methods for consistency.
func (ns *Namespace) ToTruth(v any) (bool, error) {
	result, err := ns.ToBool(v)
	if err != nil {
		switch v {
		case "", "nil", "null", "undefined", "NaN":
			return false, nil
		default:
			return true, nil
		}
	}
	return result, nil
}

func convertTemplateToString(v any) any {
	switch vv := v.(type) {
	case template.HTML:
		v = string(vv)
	case template.CSS:
		v = string(vv)
	case template.HTMLAttr:
		v = string(vv)
	case template.JS:
		v = string(vv)
	case template.JSStr:
		v = string(vv)
	}
	return v
}
