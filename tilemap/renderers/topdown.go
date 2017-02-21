package renderers

import (
    // "fmt"
    "image"
    "time"
    "math/rand"
    "github.com/hajimehoshi/ebiten"
    // common      "github.com/madstanners/go/common"
    gg_cam      "github.com/madstanners/go/camera"
    gg_tiles    "github.com/madstanners/go/tilemap"
)

var ts TileSheet;
var tileArranger *TileArranger;

// calls NewTile(x, y, size) with the DEFAULT_TILE_SIZE const
func NewTile(x, y int) image.Rectangle {
    r := NewTileCustom(x, y, DEFAULT_TILE_SIZE);
    return r;
}

// tile arranger
type TileArranger struct {
    sheet *TileSheet;
    count int;
    tile []image.Rectangle;
    point []image.Point;
}

func NewTileArranger(sheet TileSheet) *TileArranger {
    ta := TileArranger{};
    ta.sheet = &sheet;
    ta.count = 0;
    return &ta;
}

func (ta *TileArranger) Len() (int) {
    return ta.count;
}

func (ta *TileArranger) Src(i int) (x0, y0, x1, y1 int) {
    t := ta.tile[i];
    return t.Min.X, t.Min.Y, t.Max.X, t.Max.Y;
}

func (ta *TileArranger) Dst(i int) (x0, y0, x1, y1 int) {
    // point at which to start drawing the tile
    p := ta.point[i]
    x0 = p.X;
    y0 = p.Y;

    // point at which the tile should end
    t := ta.tile[i];
    x1 = x0 + t.Max.X - t.Min.X;
    y1 = y0 + t.Max.Y - t.Min.Y;

    return x0, y0, x1, y1;
}

// TODO add a user defined string seed and default to time.Now().UnixNano() if empty or nil
func (ta *TileArranger) Randomize(widthInTiles, heightInTiles int) {
    rand.Seed(time.Now().UnixNano()); // doesn't work with the same number FIXME 
    for x := 0; x < widthInTiles; x++ {
        for y := 0; y < heightInTiles; y++ {
            adjX := x * DEFAULT_TILE_SIZE;
            adjY := y * DEFAULT_TILE_SIZE;
            texRef := ta.sheet.textureRefs[rand.Intn(ta.sheet.count)];
            ta.Add(texRef, image.Point{adjX, adjY});
        }
    }
}

func (ta *TileArranger) Add(texture string, xy image.Point) {
    ta.tile = append(ta.tile, ta.sheet.textureRects[texture]);
    ta.point = append(ta.point, xy);
    ta.count = ta.count + 1;
}

func InitTiles() {
    // load ground textures
    // ts = NewTileSheet("_resources/tiles.png");

    // create texture rectangles
    ts.Add("tx1", 0, 0); // image.Rect(0, 0, 15, 15);
    ts.Add("tx2", 1, 0); // image.Rect(16, 0, 31, 15);
    ts.Add("tx3", 0, 1); // image.Rect(0, 16, 15, 31);
    ts.Add("tx4", 1, 1); // image.Rect(16, 16, 31, 31);

    // create tile arranger
    tileArranger = NewTileArranger(ts);
    tileArranger.Randomize(8, 8);

    // if(err != nil){
    //     fmt.Println("error occurred adding to the tile arranger");
    // }

}

func DrawTiles(screen *ebiten.Image, camera *gg_cam.Camera) { // TODO handle multiple tilesheets loaded into memory?
    imgx := ts.sheet;
    options := &ebiten.DrawImageOptions{};
    options.ImageParts = tileArranger;

    // apply view transformation
    camera.ApplyTransformation(&options.GeoM);

    screen.DrawImage(imgx, options);
}

/**
 * @brief      Renders a map onto an ebiten image
 *
 * @param      ebitenImage  The target image
 * @param      Map          The map
 *
 * @return     ebiten image with the map rendered onto it
 */
func RenderMap(img *ebiten.Image, m *gg_tiles.Map) *ebiten.Image {
    // TODO metadata, e.g offset, other shit
    return nil
}