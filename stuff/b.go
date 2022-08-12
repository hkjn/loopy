package stuff

import "hkjn.me/antiloop"

func B() string {
	return "weeeee! I'm in a loop! " + antiloop.A()
}
