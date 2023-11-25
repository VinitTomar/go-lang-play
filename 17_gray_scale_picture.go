package main

import "golang.org/x/tour/pic"

func Pic(dx, dy int) [][]uint8 {
	var myPic [][]uint8 = make([][]uint8, dy);
	
	for i :=0; i < dy; i++ {
		myPic[i] = make([]uint8, dx);
		
		for j :=0; j < dx; j++ {
			myPic[i][j] = uint8(j * i);
		}
	}	
	
	return myPic;
}

func showMyPic() {
	pic.Show(Pic)
}
