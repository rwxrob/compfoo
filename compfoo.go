/*

Package compfoo is a completion driver for Bonzai command trees and
fulfills the bonzai.Completer package interface by returning "foo" and
"bar". It is conventional to begin Completer packages with "comp".
*/
package compfoo

import "github.com/rwxrob/bonzai"

// fulfills the bonzai.Completer interface
type Completer struct{}

// Complete completes everything with the word "foo" and "bar"
func (Completer) Complete(x bonzai.Command, args ...string) []string {
	return []string{"foo", "bar"}
}
