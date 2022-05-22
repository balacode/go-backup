// -----------------------------------------------------------------------------
// go-backup/get_args_test.go

package main

import (
	"reflect"
	"testing"
)

func Test_extractNamedArg_(t *testing.T) {
	{
		argsIn := []string{"arg1", "arg2", "arg3"}
		v, argsOut := extractNamedArg(argsIn, "--named")
		if v != "" {
			t.Errorf("0xEE8DD3")
		}
		if !reflect.DeepEqual(argsOut, []string{"arg1", "arg2", "arg3"}) {
			t.Errorf("0xEE5E92")
		}
	}
	{
		argsIn := []string{"arg1", "--named1", "-value", "arg2", "arg3"}
		v, argsOut := extractNamedArg(argsIn, "--named1")
		if v != "-value" {
			t.Errorf("0xEE66FD")
		}
		if !reflect.DeepEqual(argsOut, []string{"arg1", "arg2", "arg3"}) {
			t.Errorf("0xEE6FD6")
		}
	}
	{
		argsIn := []string{"arg1", "--named1", "v1", "named2", "v2"}
		v, argsOut := extractNamedArg(argsIn, "named1", "named2")
		if v != "v1" {
			t.Errorf("0xEE5F60")
		}
		if !reflect.DeepEqual(argsOut, []string{"arg1", "named2", "v2"}) {
			t.Errorf("0xEE89AA")
		}
	}
}

func Test_extractNextArg_(t *testing.T) {
	{
		v, argsOut := extractNextArg([]string{})
		if v != "" {
			t.Errorf("0xEE0C65")
		}
		if !reflect.DeepEqual(argsOut, []string{}) {
			t.Errorf("0xEE00DC")
		}
	}
	{
		v, argsOut := extractNextArg([]string{"123", "456", "789"})
		if v != "123" {
			t.Errorf("0xEE5AF1")
		}
		if !reflect.DeepEqual(argsOut, []string{"456", "789"}) {
			t.Errorf("0xEE0FA7")
		}
	}
}

// end
