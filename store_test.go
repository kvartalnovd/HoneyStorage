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
	expectedFilename := "50b718784409c1cb81893557aaf3bba7e0b1ef2e"
	expectedPathName := "50b71/87844/09c1c/b8189/3557a/af3bb/a7e0b/1ef2e"
	if pathKey.PathName != expectedPathName {
		t.Errorf("have %q want %q", pathKey.PathName, expectedPathName)
	}
	if pathKey.Filename != expectedFilename {
		t.Errorf("have %q want %q", pathKey.Filename, expectedFilename)
	}
}

func TestStore(t *testing.T) {
	s := newStore()
	defer teardown(t, s)

	for i := 0; i < 50; i++ {
		key := fmt.Sprintf("foo_%d", i)
		data := []byte("some jpg bytes")

		if _, err := s.writeStream(key, bytes.NewReader(data)); err != nil {
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

		if err := s.Delete(key); err != nil {
			t.Error(err)
		}

		if ok := s.Has(key); ok {
			t.Errorf("expected to NOT have key %q", key)
		}
	}
}

func newStore() *Store {
	opts := StoreOpts{
		PathTransformFunc: CASPathTransformFunc,
	}
	return NewStore(opts)
}

func teardown(t *testing.T, s *Store) {
	if err := s.Clear(); err != nil {
		t.Error(err)
	}
}
