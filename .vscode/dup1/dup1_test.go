package main

import (
	"bytes"
	"io"
	"log"
	"os"
	"strings"
	"testing"

	"github.com/stretchr/testify/require"
)

func TestMain(m *testing.M) {

	exitVal := m.Run()
	log.Println(exitVal)

	os.Exit(exitVal)
}

func TestA(t *testing.T) {

	inputs := [5]string { "test", "test1", "test2", "test", "test2"}

	buf := bytes.NewBufferString(strings.Join(inputs[:], "\n"))
	os.Stdin.Write(buf.Bytes())
	os.Stdin.Close()

	old := os.Stdout
	r, w, _ := os.Pipe()
	os.Stdout = w

	CountDups()

	outC := make(chan string)
	go func() {
		var buf bytes.Buffer
		io.Copy(&buf, r)
		outC <- buf.String()
	}()

	w.Close()
	os.Stdout = old
	out := <-outC
	
	require.Equal(t, out, "2\ttest")

	log.Println(out)
}