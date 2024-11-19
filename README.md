# Pomodoro

Just another pomodoro timer with music playing support

https://github.com/user-attachments/assets/35afd251-d41e-4e2c-a43a-d8d6190ad850

## Installation

Currently, the App is not pre-packaged. To run it, you need to build it from source
but don't worry, there is not much to do, follow the steps below

1. Download the GitHub repo 

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

## Thanks

Thanks to [ashtron](https://github.com/ashtron/tomatillo) for bringing the idea of pairing a pomodoro timer with a music player

## TODO

- [ ] Add support for timer `reset` feature.
- [ ] Add support for more player features (pause/resume track, next/prev track).
- [ ] Add support for breaks management (long breaks, short breaks, ...).
- [ ] Add support for playing *.mp4, *.webm files.
