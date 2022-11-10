func floodFill(image [][]int, sr int, sc int, color int) [][]int {
    ogColor := image[sr][sc]
    if ogColor == color {
        return image
    }
    floodFiller(image, sr, sc, color, ogColor)
    return image
}

func floodFiller(image [][]int, sr int, sc int, color int, ogColor int) {
    if image[sr][sc] == ogColor {
        image[sr][sc] = color
        if sr >= 1 {
            floodFiller(image, sr-1, sc, color, ogColor)
        }

        if sc >= 1 {
            floodFiller(image, sr, sc-1, color, ogColor)
        }

        if sr+1 < len(image) {
            floodFiller(image, sr+1, sc, color, ogColor)
        }

        if sc+1 < len(image[0]) {
            floodFiller(image, sr, sc+1, color, ogColor)
        }
    }
}
