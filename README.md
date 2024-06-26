# Bilinear gradient generator GUI

Uses bilinear-interpolation to generate random colour gradients. A random value is selected for each corner of the image, then bilinear-interpolation is used to interpolate the remaining pixels. This is repeated 4 times. These 4 separate pixel arrays are then combined into the red, green, blue and alpha channels to produce a final image.

Version 2 adds the ability to manually select the colour of each corner and have them be used to generate the gradient. Click on the gradient to get the exact colour of a pixel!

Based *very* heavily on [Bilinear gradient generator](https://github.com/odddollar/Bilinear-gradient-generator).

## Building

Built using the [Fyne](https://fyne.io/) GUI framework for Go, this program can be compiled to a single binary with the following commands:

```
git clone https://github.com/odddollar/Bilinear-gradient-generator-GUI.git
cd Bilinear-gradient-generator-GUI
go install fyne.io/fyne/v2/cmd/fyne@latest // installs the necessary Fyne tooling
fyne package --release
```

## Screenshots

![Demo 1](screenshots/demo1.gif)

![Image 1](screenshots/image1.png)

![Image 2](screenshots/image2.png)

![Image 3](screenshots/image3.png)

![Image 4](screenshots/image4.png)
