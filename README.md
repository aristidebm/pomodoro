# Pomodoro

Just another pomodoro timer with music playing support

![demo](demo.mp4)

## Installation

Currently, the App is not pre-packaged. To run it, you need to build it from source
but don't worry, there is not much to do, follow the steps below

1. Download the Github repo 

```bash
$)> git clone  https://gitlab.com/aristidebm/pomodoro.git \<directory\>
```

2. Go into the directory above

```bash
$)> cd \<directory\>
```

3. Build the app

You can build the app by executing the command below. It will create a file named `pomodoro` inside `/tmp` folder

```bash
make build
```

4. Copy the file into local bin (optional)

```bash
make deploy
```

After these steps, you are ready to go!

## Usage

```bash
$)> pomodoro
```

Need to know available CLI options ?

```bash
$)> pomodoro --help
```

## TODO

- [ ] Add `reset` feature
- [ ] Add the ability to pause the music player when pausing the timer and make it optional
- [ ] Add support for break managements (long breaks, short breaks, ...)
