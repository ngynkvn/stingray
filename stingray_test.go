package stingray_test

import (
	"fmt"
	"os"
	"testing"

	"github.com/ngynkvn/stingray"
)

func TestBasicParsing(t *testing.T) {
	f, err := os.Open("./testdata/10078705.dem")
	// f, err := os.Open("testdata/534870.dem")
	if err != nil {
		t.Log(err)
		t.FailNow()
	}
	defer f.Close()

	p, err := stingray.NewStreamParser(f, stingray.Logger)
	if err != nil {
		t.Log(err)
		t.FailNow()
	}
	p.Callbacks.OnCDemoFileHeader(dump)
	// p.Callbacks.OnCDemoSignonPacket(dump)
	err = p.Start()
	if err != nil {
		t.Log(err)
		t.FailNow()
	}

	t.Log("Parse complete!")
}

func dump[T fmt.Stringer](pkt T) error {
	println(pkt.String())
	return nil
}
