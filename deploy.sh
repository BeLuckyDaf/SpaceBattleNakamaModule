export NAKAMA_RUNTIME_PATH=/home/beluc/nakama-server/
export NAKAMA_PATH=/home/beluc/nakama-server/

go build -buildmode=plugin -trimpath -o $NAKAMA_RUNTIME_PATH/data/modules/spacebattle.so
$NAKAMA_PATH/nakama -runtime.path $NAKAMA_RUNTIME_PATH
