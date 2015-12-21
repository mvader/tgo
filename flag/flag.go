// Copyright 2015 trivago GmbH
//
// Licensed under the Apache License, Version 2.0 (the "License");
// you may not use this file except in compliance with the License.
// You may obtain a copy of the License at
//
//     http://www.apache.org/licenses/LICENSE-2.0
//
// Unless required by applicable law or agreed to in writing, software
// distributed under the License is distributed on an "AS IS" BASIS,
// WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
// See the License for the specific language governing permissions and
// limitations under the License.

package flag

import (
	"flag"
	"fmt"
)

type flagKeyDescription struct {
	flagKey string
	usage   string
	value   string
}

var descriptions = []flagKeyDescription{}
var tabWidth = []int{0, 0, 0}

const (
	tabKey   = iota
	tabUsage = iota
	tabValue = iota
)

func addKeyDescription(short string, long string, value interface{}, usage string) {
	desc := flagKeyDescription{
		flagKey: fmt.Sprintf("-%s, --%s", short, long),
		usage:   usage,
		value:   fmt.Sprintf("%v", value),
	}

	if tabWidth[tabKey] < len(desc.flagKey) {
		tabWidth[tabKey] = len(desc.flagKey)
	}
	if tabWidth[tabValue] < len(desc.value) {
		tabWidth[tabValue] = len(desc.value)
	}
	if tabWidth[tabUsage] < len(desc.usage) {
		tabWidth[tabUsage] = len(desc.usage)
	}

	descriptions = append(descriptions, desc)
}

func SwitchVar(flagVar *bool, short string, long string, usage string) *bool {
	return BoolVar(flagVar, short, long, false, usage)
}

func BoolVar(flagVar *bool, short string, long string, value bool, usage string) *bool {
	flag.BoolVar(flagVar, short, value, usage)
	flag.BoolVar(flagVar, "-"+long, value, usage)
	addKeyDescription(short, long, value, usage)
	return flagVar
}

func IntVar(flagVar *int, short string, long string, value int, usage string) *int {
	flag.IntVar(flagVar, short, value, usage)
	flag.IntVar(flagVar, "-"+long, value, usage)
	addKeyDescription(short, long, value, usage)
	return flagVar
}

func Int64Var(flagVar *int64, short string, long string, value int64, usage string) *int64 {
	flag.Int64Var(flagVar, short, value, usage)
	flag.Int64Var(flagVar, "-"+long, value, usage)
	addKeyDescription(short, long, value, usage)
	return flagVar
}

func Float64Var(flagVar *float64, short string, long string, value float64, usage string) *float64 {
	flag.Float64Var(flagVar, short, value, usage)
	flag.Float64Var(flagVar, "-"+long, value, usage)
	addKeyDescription(short, long, value, usage)
	return flagVar
}

func StringVar(flagVar *string, short string, long string, value string, usage string) *string {
	flag.StringVar(flagVar, short, value, usage)
	flag.StringVar(flagVar, "-"+long, value, usage)
	addKeyDescription(short, long, value, usage)
	return flagVar
}

func Switch(short string, long string, usage string) *bool {
	var flagVar bool
	return SwitchVar(&flagVar, short, long, usage)
}

func Bool(short string, long string, value bool, usage string) *bool {
	flagVar := value
	return BoolVar(&flagVar, short, long, value, usage)
}

func Int(short string, long string, value int, usage string) *int {
	flagVar := value
	return IntVar(&flagVar, short, long, value, usage)
}

func Int64(short string, long string, value int64, usage string) *int64 {
	flagVar := value
	return Int64Var(&flagVar, short, long, value, usage)
}

func Float64(short string, long string, value float64, usage string) *float64 {
	flagVar := value
	return Float64Var(&flagVar, short, long, value, usage)
}

func String(short string, long string, value string, usage string) *string {
	flagVar := value
	return StringVar(&flagVar, short, long, value, usage)
}

func Parse() {
	flag.Parse()
}

func PrintFlags(header string) {
	fmt.Println(header)

	valueFmt := fmt.Sprintf("%%-%ds\tdefault: %%-%ds\t%%-%ds\n", tabWidth[tabKey], tabWidth[tabValue], tabWidth[tabUsage])

	for _, desc := range descriptions {
		fmt.Printf(valueFmt, desc.flagKey, desc.value, desc.usage)
	}
}