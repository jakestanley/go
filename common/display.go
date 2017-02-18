package common;

const (
    DEFAULT_WIDTH   = 640;
    DEFAULT_HEIGHT  = 480;
    DEFAULT_SCALE   = 1;
)

type DisplayConfig struct {
    width int;
    height int;
    scale float64;
}

func DefaultDisplay() *DisplayConfig {
    return &DisplayConfig{ DEFAULT_WIDTH, DEFAULT_HEIGHT, DEFAULT_SCALE }
}

func NewDisplay(width, height int, scale float64) *DisplayConfig {
    return &DisplayConfig{ width, height, scale }
}

func (d *DisplayConfig) GetScreenWidth() int {
    return d.width;
}

func (d *DisplayConfig) GetScreenHeight() int {
    return d.height;
}

func (d *DisplayConfig) GetScale() float64 {
    return d.scale;
}