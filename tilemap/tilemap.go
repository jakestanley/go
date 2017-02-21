package tilemap

const (
    DEFAULT_MAP_SX = 32
    DEFAULT_MAP_SY = 32
    DEFAULT_MAP_SF = 16
)

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

type ThreeDimensionalMap struct {
    sizeX int
    sizeY int
    floors int
    tiles [][][]*Tile
    regions []*Region
}

func NewMap(sizeX int, sizeY int) *Map {
    var tiles       [][]*Tile
    var regions     []*Region

    for x := 0; x < sizeX; x++ {
        for y := 0; y < sizeY; y++ {
            tiles[x][y] = GetVoidTile();
        }
    }

    m := Map{ sizeX, sizeY, tiles, regions }
    return &m
}

func Default3DMap() *ThreeDimensionalMap {
    return New3DMap(DEFAULT_MAP_SX, DEFAULT_MAP_SY, DEFAULT_MAP_SF);
}

func New3DMap(sizeX, sizeY, floors int) *ThreeDimensionalMap {
    var tiles       [floors][sizeX][sizeY]*Tile
    var regions     []*Region

    for f := 0; f < floors; f++ {
        for x := 0; x < sizeX; x++ {
            for y := 0; y < sizeY; y++ {
                tiles[f][x][y] = GetVoidTile();
            }
        }
    }

    m := &ThreeDimensionalMap{sizeX, sizeY, floors, tiles, regions};
    return m;
}
