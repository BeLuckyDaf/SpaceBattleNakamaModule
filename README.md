# Development Environment

## Prerequisites
1. Install Visual Studio Code to use with WSL: https://code.visualstudio.com/
  a. Install the Remote WSL extension: https://marketplace.visualstudio.com/items?itemName=ms-vscode-remote.remote-wsl
  b. Probably Golang extension wouldn’t hurt too: https://marketplace.visualstudio.com/items?itemName=ms-vscode.Go
2. Install WSL: https://docs.microsoft.com/en-us/windows/wsl/install-win10
3. Install a distro for WSL from Microsoft Store, e.g. Ubuntu
4. Install the correct version of Go into WSL: https://golang.org/doc/install#install
5. Install CockroachDB: https://www.cockroachlabs.com/docs/stable/install-cockroachdb-linux.html

## Setting up Nakama & Cockroach

1. Make a directory for the local Nakama server
```bash
mkdir nakama-server
cd nakama-server
```

2. Download and extract Nakama binary

```bash
wget https://github.com/heroiclabs/nakama/releases/download/v2.12.0/nakama-2.12.0-linux-amd64.tar.gz
tar -xzvf nakama-2.12.0-linux-amd64.tar.gz
# Here you might want to remove everything except for the binary itself.
```

3. Create a directory for the local database & start it there

```bash
mkdir db
cd db
cockroach start-single-node --insecure
# Now go for another terminal, leave the db working here.
4. Start Nakama and make sure everything is fine
./nakama migrate up
./nakama
# If you see that everything went well, shut it down for now.
```

## Get it running

1. Clone the repo
```bash
git clone git@github.com:BeLuckyDaf/SpaceBattleNakamaModule
```

2. Go to `scripts/deploy.sh` and change the environment variables to point
to the place where your nakama runtime is set up.
```bash
# For example if you've installed it in the home directory.
# This would mean that after you run deploy.sh, the module
# named spacebattle.so would to be put to 
# .../nakama-server/data/modules/
export NAKAMA_RUNTIME_PATH=/home/<username>/nakama-server/
export NAKAMA_PATH=/home/<username>/nakama-server/
export MY_NAKAMA_MODULE_NAME=spacebattle
```

3. Run cockroachdb
```bash
./scripts/startdb.sh
```

4. Run nakama from another terminal
```bash
./scripts/deploy.sh
```
