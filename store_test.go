package main

import (
	"bytes"
	"fmt"
	"io/ioutil"
	"testing"
)

func TestPathTransformFunc(t *testing.T) {
	key := "bestfamilypicture"
	pathKey := CASPathTransformFunc(key)
	expectedOriginalKey := "50b718784409c1cb81893557aaf3bba7e0b1ef2e"
	expectedPathName := "50b71/87844/09c1c/b8189/3557a/af3bb/a7e0b/1ef2e"
	if pathKey.PathName != expectedPathName {
		t.Errorf("have %q want %q", pathKey.PathName, expectedPathName)
	}
	if pathKey.Filename != expectedOriginalKey {
		t.Errorf("have %q want %q", pathKey.Filename, expectedOriginalKey)
	}
}

func TestStoreDeleteKey(t *testing.T) {
	opts := StoreOpts{
		PathTransformFunc: CASPathTransformFunc,
	}
	s := NewStore(opts)
	key := "myspecials"
	data := []byte("some jpg bytes")

	if err := s.writeStream(key, bytes.NewReader(data)); err != nil {
		t.Error(err)
	}

	if err := s.Delete(key); err != nil {
		t.Error(err)
	}
}

func TestStore(t *testing.T) {
	opts := StoreOpts{
		PathTransformFunc: CASPathTransformFunc,
	}
	s := NewStore(opts)
	key := "myspecials"
	data := []byte("some jpg bytes")

	if err := s.writeStream(key, bytes.NewReader(data)); err != nil {
		t.Error(err)
	}

	if ok := s.Has(key); !ok {
		t.Errorf("expected to have key %q", key)
	}

	r, err := s.Read(key)
	if err != nil {
		t.Error(err)
	}

	b, _ := ioutil.ReadAll(r)
	if string(b) != string(data) {
		t.Errorf("want %s have %s", data, b)
	}
	fmt.Println(string(b))

	s.Delete(key)
}
