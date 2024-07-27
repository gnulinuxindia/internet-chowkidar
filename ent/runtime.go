// Code generated by ent, DO NOT EDIT.

package ent

import (
	"github.com/gnulinuxindia/internet-chowkidar/ent/counter"
	"github.com/gnulinuxindia/internet-chowkidar/ent/schema"
)

// The init function reads all schema descriptors with runtime code
// (default values, validators, hooks and policies) and stitches it
// to their package variables.
func init() {
	counterFields := schema.Counter{}.Fields()
	_ = counterFields
	// counterDescCount is the schema descriptor for count field.
	counterDescCount := counterFields[0].Descriptor()
	// counter.DefaultCount holds the default value on creation for the count field.
	counter.DefaultCount = counterDescCount.Default.(int)
}