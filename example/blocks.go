/*
Copyright 2014 Hajime Hoshi

Licensed under the Apache License, Version 2.0 (the "License");
you may not use this file except in compliance with the License.
You may obtain a copy of the License at

    http://www.apache.org/licenses/LICENSE-2.0

Unless required by applicable law or agreed to in writing, software
distributed under the License is distributed on an "AS IS" BASIS,
WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
See the License for the specific language governing permissions and
limitations under the License.
*/

package main

import (
	"flag"
	"github.com/hajimehoshi/ebiten/example/blocks"
	"github.com/hajimehoshi/ebiten/runner"
	"log"
	"os"
	"runtime"
	"runtime/pprof"
)

func init() {
	runtime.LockOSThread()
}

var cpuProfile = flag.String("cpuprofile", "", "write cpu profile to file")

func main() {
	flag.Parse()
	if *cpuProfile != "" {
		f, err := os.Create(*cpuProfile)
		if err != nil {
			log.Fatal(err)
		}
		pprof.StartCPUProfile(f)
		defer pprof.StopCPUProfile()
	}

	game := blocks.NewGame()
	if err := runner.Run(game, blocks.ScreenWidth, blocks.ScreenHeight, 2, "Blocks (Ebiten Demo)", 60); err != nil {
		log.Fatal(err)
	}
}
