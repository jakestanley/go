package renderers

"github.com/hajimehoshi/ebiten"

const DEFAULT_TILE_SIZE = 16;

// using ebiten ImageParts helps to reduce draw calls
// TODO make these dynamic. they're pretty static now... can't remove tiles easily
type TileSheet struct {
    config *TileRenderConfig;
    sheet *ebiten.Image;
    textureRects map[string]image.Rectangle; // refs references this
    count int;
    textureRefs []string;
    textureRectsArray []image.Rectangle;
}

// TODO don't render tiles if they're outside of the screen area, assuming 
// ebiten doesn't do that automatically

type TileRenderConfig struct {
    tileSize int;
}

func NewTileSheet(path string) *TileSheet {
    ts := TileSheet{};
    sheet, err := common.LoadImage(path);
    ts.textureRects = make(map[string]image.Rectangle);
    if(err != nil){
        // TODO do something
    }
    ts.sheet = sheet;
    return &ts;
}

func (ts *TileSheet) Add(ref string, x, y int) { // TODO return error here
    rect := NewTile(x, y);
    ts.textureRects[ref] = rect; // image.Rect(0, 0, 15, 15);
    ts.textureRefs = append(ts.textureRefs, ref);
    ts.textureRectsArray = append(ts.textureRectsArray, rect);
    ts.count = ts.count + 1;
}

// creates a rect based on tile size and tile coordinates, e.g 1,1, should return the rectangle 
// 16, 16, 31, 31 // TODO test
func NewTileCustom(x, y, size int) image.Rectangle {
    x0 := x * size;
    y0 := y * size;
    x1 := x0 + size;
    y1 := y0 + size;
    r := image.Rect(x0, y0, x1, y1);
    return r;
}

func DefaultTileRenderConfig() *TileRenderConfig {
    return &TileRenderConfig{DEFAULT_TILE_SIZE};
}

func LoadTileSheet(imgPath string)

func LoadTileSheet(imgPath string, tileSize int) {



}

// TODO
// func TileRenderConfig() *TileRenderConfig {
// 
// }