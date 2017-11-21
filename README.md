# VPM

Vim + Plugin Manager = VPM

VPM is a plugin manager for vim8 new native package
loading implemented in go. List, install and remove
your vim plugins.

## Installation

#### MacOS

`git clone https://github.com/Waasi/vpm.git`

`make install-macos`

#### Linux

`git clone https://github.com/Waasi/vpm.git`

`make install-linux`

## Usage

First initialize your plugins container with:

`vpm init`

Then look for the https url of your
plugin in github and install it with:

`vpm add <plugin_url>`

To remove a plugin run:

`vpm remove <plugin_name>`

To list all installed plugins run:

`vpm list`

## Example

In this example we will install nerdtree vim plugin.

1. Install the plugin:

  `> vpm add https://github.com/scrooloose/nerdtree.git`

  `> added plugin nerdtree`

2. See if the plugin was installed:

  `> vpm list`

  `> nerdtree`

3. Remove the plugin:

  `> vpm remove nerdtree`

  `> removed plugin nerdtree`

## Contributing

1. Fork it ( https://github.com/[my-github-username]/vpm/fork )
2. Create your feature branch (git checkout -b feature/my_new_feature)
3. Commit your changes (git commit -am 'Add some feature')
4. Push to the branch (git push origin my-new-feature)
5. Create a new Pull Request
