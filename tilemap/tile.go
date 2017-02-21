package tilemap

var voidTile *Tile;

type Tile struct {
    void bool
    solid bool
    // texture string // path. could also be an interface. can remove
    // map generation will be renderer agnostic
    // *ebiten.Image texture 
}

func init() {
    
}

/**
 * @brief      Determines if a void tile.
 *
 * @return     True if void, False otherwise. Do not attempt to render this 
 *             tile.
 */
func (t *Tile) IsVoid() bool {
    return t.void;
}

/**
 * @brief      Returns the void tile.
 *
 * @return     Returns a static tile that isn't solid and shouldn't be drawn. 
 *             No other attributes will be defined. It's supposed to be used as 
 *             a placeholder to prevent null pointer exceptions meaningfully.
 */
func GetVoidTile() *Tile {
    if(voidTile == nil){
        voidTile = &Tile{ true, false }
    }
    return voidTile;
}
