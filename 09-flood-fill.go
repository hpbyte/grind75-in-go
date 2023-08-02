/**
 * DFS approach
 * return early if the start color is already the new color
 */
func floodFill(image [][]int, sr int, sc int, color int) [][]int {
  startColor := image[sr][sc]

  if startColor == color {
    return image
  }

  m := len(image)
  n := len(image[0])

  var fill func(i int, j int)

  fill = func(i int, j int) {
    if i < 0 || i >= m || j < 0 || j >= n || image[i][j] != startColor {
      return
    }

    image[i][j] = color

    fill(i-1, j)
    fill(i+1, j)
    fill(i, j-1)
    fill(i, j+1)
  }

  fill(sr, sc)

  return image
}
