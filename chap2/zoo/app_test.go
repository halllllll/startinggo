package main

import(
	"testing"
)

func Testappname(t *testing.T){	
	except:="Zoo Application"
	actual := AppName()
	if except!=actual{
		t.Error("%s != %s", except, actual)
	}
}