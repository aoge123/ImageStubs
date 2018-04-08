//http://tech.nitoyon.com/en/blog/2015/12/31/go-image-gen/
package main

import (
    "fmt"
    "image"
    "image/color"
    "image/jpeg"
    "math/rand"
    "os"
    "strings"
    "strconv"
    "C"
)

const w, h int =  750, 1334
func generateSolid(name int) {

    colorValue := make([]uint8, 256)
    for i := range colorValue {
        colorValue[i] = uint8(i)
    }



    var r uint8 = colorValue[rand.Intn(256)]
    var g uint8 = colorValue[rand.Intn(256)]
    var b uint8 = colorValue[rand.Intn(256)]

    if r == g && g == b  {
        for r < 40 || r > 224 {

            if r < 40 {
                r, g, b = 40, 40, 40
            }
    
            if r > 224 {
                r, g, b = 224, 224, 224
            }

        }
       
    }

    m := image.NewRGBA(image.Rect(0, 0, w, h))
    for x := 0; x < w; x++ {
        for y := 0; y < h; y++ {
            c := color.RGBA{
                r,
                g,
                b,
                255,
            }
            m.Set(x, y, c)
        }
    }

    output := strings.Join([]string{strconv.Itoa(name), ".jpeg"}, "")

    f, err := os.OpenFile(output, os.O_WRONLY|os.O_CREATE, 0600)
    if err != nil {
        fmt.Println(err)
        return
    }
    defer f.Close()
    jpeg.Encode(f, m, nil)

}

// export
func createNImageStubs(n int) {
    options := []string{"Solid"}
    optionlen := len(options) //investigate

    imagesToCreate := make([]string, n) //Investigate
    for i := 0; i < n; i++ {
        var r int
        if (optionlen - 1) == 0 {
            r = 0
        }else {
            r = rand.Intn(optionlen - 1)
        }
        
        imagesToCreate[i] = options[r]
    }

    for index, imageType := range imagesToCreate {
        switch imageType {
        case "Solid":
            generateSolid(index)
        }
    }

}

func main() {
    n, err := strconv.Atoi(os.Args[1])
    if err != nil {
        fmt.Println(err)
        os.Exit(2)
    }

    createNImageStubs(n)
}