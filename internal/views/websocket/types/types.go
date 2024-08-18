package types

var State = struct {
    Online  string
    Offline string
}{
    Online:  "online",
    Offline: "offline",
}


var List = struct {
    Following string
    DirectMessages string
}{
    Following: "Following",
    DirectMessages: "Direct Messages",
}
