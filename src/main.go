package main


import (

	"fmt"
	"github.com/gedex/bp3d"
	"log"

)


func main() {


	fmt.Println("Testing a quick pack")
	p := bp3d.NewPacker()

	// Add bins.
	p.AddBin(bp3d.NewBin("Small Bin", 10, 15, 20, 10))
	p.AddBin(bp3d.NewBin("Medium Bin", 100, 150, 200, 20))

	// Add items.
	p.AddItem(bp3d.NewItem("Chicken", 2, 2, 1, 2))
	p.AddItem(bp3d.NewItem("Rice", 1, 3, 2, 3))
	p.AddItem(bp3d.NewItem("Rice", 1, 3, 2, 3))
	p.AddItem(bp3d.NewItem("Rice", 1, 3, 2, 3))
	p.AddItem(bp3d.NewItem("Rice", 1, 3, 2, 3))
	p.AddItem(bp3d.NewItem("Rice", 1, 3, 2, 3))

	// Pack items to bins.
	if err := p.Pack(); err != nil {
		log.Fatal(err)
	}

	// Each bin, b, in p.Bins might have packed items in b.Items
	displayPacked(p.Bins)

}

func displayPacked(bins []*bp3d.Bin) {
	for _, b := range bins {
		fmt.Println(b)
		fmt.Println(" packed items:")
		for _, i := range b.Items {
			fmt.Println("  ", i)
		}
	}
}



