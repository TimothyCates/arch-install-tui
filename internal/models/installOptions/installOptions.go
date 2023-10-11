package installOptions

type credentials struct {
    userName string
    userPass string
    rootPass string
}

type reflector struct {
    country  string
    protocol []string
    latest   int
}

type packages struct {
    useNvidia  bool
    nvidiaType NvidiaDriverType
    pacman     []string
    yay        []string
    userGroups []string
    services   []string
}

type system struct {
    zoneInfo string
    hostname string
    keymap   string
    locales  []string
    userDirs []string
    kernels []string
    reflector
    packages
}

type dotfile struct {
    url      string
    location string
    useStow  bool
    stow     stow
}

type stow struct {
    target  string
    folders []string
}

type NvidiaDriverType int8
const (
    Nvidia  NvidiaDriverType = iota
    Nouveau
    DKMS
)

type Options struct {
    doRefreshKeys bool
    credentials
    system
    dotfile
}
