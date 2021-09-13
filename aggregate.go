// Copyright 2021 gotomicro
//
// Licensed under the Apache License, Version 2.0 (the "License");
// you may not use this file except in compliance with the License.
// You may obtain a copy of the License at
//
// http://www.apache.org/licenses/LICENSE-2.0
//
// Unless required by applicable law or agreed to in writing, software
// distributed under the License is distributed on an "AS IS" BASIS,
// WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
// See the License for the specific language governing permissions and
// limitations under the License.

package eql

// Aggregate represents aggregate expression, including AVG, MAX, MIN...
type Aggregate struct {
	funcCall
	alias string
}

// As specifies the alias
func (Aggregate) As(alias string) Aggregate {
	panic("implement me")
}

// Avg represents AVG
func Avg(c string) Aggregate {
	panic("implement")
}

func (a Aggregate) selected() {}