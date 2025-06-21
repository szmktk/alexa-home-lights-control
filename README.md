# StudyLight CLI

A tiny Go-based command-line tool for controlling a TP-Link Kasa smart bulb (KL110) via **Alexa routines** and **VoiceMonkey.io**.

Use it to set brightness from your terminal — cross-platform and zero dependencies.


## ✨ Features

- Set brightness to **25%, 50%, 75%, or 100%**
- Works anywhere Alexa is configured (via VoiceMonkey trigger)
- Cross-compiles for Linux, macOS (Apple Silicon), Windows, Raspberry Pi


## 🔧 Requirements

- A TP-Link Kasa bulb set up via the **Kasa app**
- An Amazon Alexa device or app (linked to the Kasa skill)
- A free [VoiceMonkey.io](https://voicemonkey.io) account


## 🚀 Setup Instructions

### 1. Set up VoiceMonkey

1. Go to [https://voicemonkey.io](https://voicemonkey.io) and create a free account.
2. Link VoiceMonkey to your Alexa account.
3. In VoiceMonkey:
   - Create "monkeys" (virtual Alexa triggers) for each routine you want:  
     `brightness_10`, `brightness_25`, `brightness_50`, etc.
4. In the **Alexa app**:
   - Create a new Routine
   - Set "When this happens" → **Smart Home** → select your VoiceMonkey trigger
   - Add the desired action (set brightness)


### 2. Build the CLI Tool

#### 🔨 Compile for your system

Run:

```bash
export VOICEMONKEY_TOKEN="your_token"
make TOKEN=$VOICEMONKEY_TOKEN
```

This generates binaries for:
  • studylight-linux-amd64
  • studylight-macos-arm64
  • studylight-windows-amd64.exe
  • studylight-rpi-arm (for Raspberry Pi with 32-bit OS)


💡 You can also build just for your own OS:

```
make linux-amd64 TOKEN=$VOICEMONKEY_TOKEN
make macos-arm64 TOKEN=$VOICEMONKEY_TOKEN
make rpi-arm TOKEN=$VOICEMONKEY_TOKEN
make windows TOKEN=$VOICEMONKEY_TOKEN
```


🧪 Test without compiling (env-based)

```
export VOICEMONKEY_TOKEN="your_token"
go run main.go -b 50
```


🧰 Make Targets

Command            | Description
-------------------|-------------------
`make`             | Build all platforms
`make clean`       | Remove dist/ directory
`make linux-amd64` | Build Linux AMD64 binary
`make macos-arm64` | Build for macOS ARM64
`make rpi-arm`     | Build for Raspberry Pi (32-bit)
`make windows`     | Build for Windows 64-bit


## Usage

🖥 From your terminal:

```
./studylight --brightness 75   # Set brightness
./studylight -b 25             # Short form
```


⚙️ Environment Variable

If running with `go run ./main.go`, the token must be set:

```
export VOICEMONKEY_TOKEN="your_token"
```

If you’re using a compiled binary, the token is baked in via:

```
make TOKEN="your_token"
```


---
### Security Note

🔐 Avoid checking your TOKEN into version control. Either export it temporarily or store it in a .env or secrets manager if you expand the project.


### License

📄 MIT — Do what you like, just don’t blame me if your lights don’t cooperate.


### Credits

- VoiceMonkey.io for making Alexa routines remotely triggerable
- TP-Link for solid hardware (even if they don’t open APIs 🙃)
