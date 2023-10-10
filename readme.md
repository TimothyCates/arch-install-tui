# Arch Installer
### A Terminal-based Interactive Arch Linux Installer

Arch Installer offers an intuitive terminal user interface (TUI) for installing a tailored Arch system. While this installer has been crafted based on specific preferences, it aims to provide a streamlined experience, though it might not encompass all potential packages or window-manager/desktop environments.

## Why

Existing methods for installing Arch Linux, such as various `install_arch.sh` scripts, YouTube tutorials, and the Arch wiki, can be comprehensive but may lack interactivity. Arch Installer was born out of the desire to transform a personal installation script into an interactive TUI that's both visually appealing and functional.

## Usage

Follow these steps to use Arch Installer:

1.  Boot into the Arch ISO.
2.  Update the package list and install git: 

    `pacman -Sy && pacman -S git`
    
3. Clone the repository and initiate the installation:

    `git clone https://github.com/timothycates/arch-install-tui && ./arch-install-tui/bin/install` 
    

> **Warning**: This tool has the potential to format and erase data from your drives, leading to data loss. Always exercise caution when using tools that perform destructive actions. It's advisable to review the code and understand its workings before usage. While the logic behind this tool is designed to be robust and is used personally by the developer, always remember to backup crucial data and use at your own discretion.

## Features

### System Setup

-   Drive formatting and mounting
-   Installation of linux, linux-lts, or linux-zen kernels
-   Base system configuration:
    -   `timedatectl`
    -   `genfstab`
    -   Locales, keymaps, and timezone
    -   Hosts file, root password, and admin user setup
    -   `pacman` and `mkpkg` config customization
    -   Bootloader setup

### Software & Drivers

-   Graphics drivers setup
-   Installation of `yay` (AUR helper)
-   Customizable software group installations:
    -   Fonts
    -   Shells
    -   Desktops / Window Manager / Display managers
    -   Gaming software
    -   Media applications
    -   Networking, devices, and console utilities
    -   Office programs
    -   Web browsers
    -   The TUI allows adding custom software to this list

### Dotfile Management

-   Facilitates the installation of dotfiles from a git repository

## Contributing

Contributions are welcome! If you wish to add a feature, support a new desktop environment, or rectify any issues, please submit a pull request. Feedback and improvements are always appreciated.

## License

This project is licensed under the MIT License. See the `LICENSE` file for more details.
