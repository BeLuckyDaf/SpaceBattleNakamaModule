export NAKAMA_RUNTIME_PATH=/home/beluc/nakama-server/
export NAKAMA_PATH=/home/beluc/nakama-server/
export MY_NAKAMA_MODULE_NAME=spacebattle

if (go build -buildmode=plugin -trimpath -o $NAKAMA_RUNTIME_PATH/data/modules/$MY_NAKAMA_MODULE_NAME.so) ; then 
    $NAKAMA_PATH/nakama migrate up
    $NAKAMA_PATH/nakama -runtime.path $NAKAMA_RUNTIME_PATH
fi