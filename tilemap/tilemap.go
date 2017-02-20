package tilemap

const (
    tileWidth   = 16
    tileHeight  = 16
)

type Tile struct {
    solid bool
    texture string // path. could also be an interface.
    // map generation will be renderer agnostic
    // *ebiten.Image texture 
}

// TODO "3D" tilemap

type Region struct {
    x int32
    y int32 
    width float32
    height float32
}

type Map struct {
    sizeX int
    sizeY int
    tiles [][]*Tile     // 2D tile array TODO chunks
    regions []*Region   // stepping on a region should trigger an action, e.g entering a building
}

func NewMap(sizeX int, sizeY int) *Map {
    var tiles       [][]*Tile
    var regions     []*Region

    for x := 0; x < sizeX; x++ {
        for y := 0; y < sizeY; y++ {
            tiles[x][y] = &Tile{ false, "../_resources/tarmac.png" }
        }
    }

    m := Map{ sizeX, sizeY, tiles, regions }
    return &m
}

func generateBasicMap() *Map {
    return NewMap(32, 32)
}
