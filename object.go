package gsonic

import (
	"bytes"
	"fmt"
)

type Object struct {
	Collection string `json:"collection,omitempty"`
	Bucket     string `json:"bucket,omitempty"`
	Object     string `json:"object,omitempty"`
	Text       string `json:"text,omitempty"`
}

func NewObject(collection, bucket, object, text string) *Object {
	return &Object{
		Collection: collection,
		Bucket:     bucket,
		Object:     object,
		Text:       text,
	}
}

func (i *Object) Prepare(op string) []byte {
	var b bytes.Buffer
	fmt.Fprintf(&b, "%s %s %s %s", op, i.Collection, i.Bucket, i.Object)
	if i.Text != "" {
		fmt.Fprintf(&b, " \"%s\"", i.Text)
	}
	return b.Bytes()
}

type ObjectBuilder struct {
	o *Object
}

func NewObjectBuilder() *ObjectBuilder {
	return &ObjectBuilder{o: &Object{}}
}

func (b *ObjectBuilder) Collection(collection string) *ObjectBuilder {
	b.o.Collection = collection
	return b
}

func (b *ObjectBuilder) Bucket(bucket string) *ObjectBuilder {
	b.o.Bucket = bucket
	return b
}

func (b *ObjectBuilder) Object(object string) *ObjectBuilder {
	b.o.Object = object
	return b
}

func (b *ObjectBuilder) Text(text string) *ObjectBuilder {
	b.o.Text = text
	return b
}

func (b *ObjectBuilder) Build() *Object {
	return b.o
}
