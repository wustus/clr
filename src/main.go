package main

import (
	"fmt"
	"os"
	"strconv"
)

func OctToHex(val *string) {

    oct, err := strconv.ParseUint(*val, 10, 8);

    if err != nil {
        fmt.Printf("Could not parse %s: %v", *val, err);
        os.Exit(1);
    }

    fmt.Printf("%02x", oct);
}

func RGBToHex(r *string, g *string, b *string) {

    OctToHex(r);
    OctToHex(g);
    OctToHex(b);
}

func main() {

    if len(os.Args) < 2 {
        fmt.Print("Please enter a value.")
        os.Exit(1);
    }

    if len(os.Args) == 2 {
        OctToHex(&os.Args[1]);
        fmt.Print("\n");
        os.Exit(0);
    }

    if len(os.Args) == 4 {
        RGBToHex(&os.Args[1], &os.Args[2], &os.Args[3]);
        fmt.Print("\n");
        os.Exit(0);
    }
}
