# Convert an image to a console image

Using the library from [gizak/termui](https://github.com/gizak/termui/blob/master/_examples/image.go).

## Instructions

1. Build the binary with `go build`
2. There are two ways to load images. The images can be png, jpg, or gif.
    * have an image file in the same directory as the binary called `veba_otto.png`
    * or, add arguments when you run the binary to the paths. So if you have images in the parent directory, you can do something like `./consoleImage ../veba_basic_320.png ../veba_basic.png ../veba_orca6.jpeg`.
3. Rotate through the images using right and left arrow keys.
4. You can change the image size by pressing `x` to increase the width, `z` to decrease the width, `y` to increase the height, and `t` to decrease to height.
5. If the image is not in monochrome, you can set it be monochrome with `<Enter>`.
6. If the image is in monochrome, you can adjust the monochrome threshold with the up and down arrows.
7. If the image is in monochrome, you can invert the monochrome with `<Tab>`.
8. And when you’re done, press `q`.
