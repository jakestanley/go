package deltatime

import "time"

const (
    // time
    SECOND = 1000000000;
    HALF_SECOND = 500000000;
)

var prevTime time.Time;
var delta float64;
var oldDelta bool;

func InitDelta() {
    oldDelta = false;
    prevTime = time.Now();
    delta = 1;
}

func UpdateDelta() {
    now := time.Now();
    timediff := now.Sub(prevTime);
    delta = (float64(timediff.Nanoseconds()) / SECOND) * 60;
    prevTime = now;
}

// TODO tests

func GetDelta() float64 {
    return delta;
}

// TODO add a function that prevents functions from using an old delta value