package main

type reflectorSettings struct{
  country string
  protocol []string
  latest int
}

type stowSettings struct{
  target string
  folders []string
}

type dotfileSettings struct{
  url string
  location string
  useStow bool
  stow stowSettings
}

type nvidiaDriverType int8
const (
  Nvidia nvidiaDriverType = iota
  Nouveau
  DKMS
)

type packages struct{
  pacman []string
  yay []string
  useNvidia bool
  nvidiaType nvidiaDriverType
  userGroups []string
  services []string
}

type systemSettings struct{
  zoneInfo string
  locales []string
  hostname string
  keymap string
  userDirs []string
  packages packages
  reflector reflectorSettings
}

type Settings struct{
  doRefreshKeys bool
  userName string
  userPass string
  rootPass string
  system systemSettings
  dotfile dotfileSettings
}
