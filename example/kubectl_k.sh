## SSS_CUSTOM_SNIPPET: kubectl_k
## Status: true
k() {
    if [[ "$1" == "help" ]]
    then
        echo "============================================ KubeHack ============================================"
        echo "version\t\tkubectl hack plugin version."
        echo "server\t\tset env var.value is the kube api server endpoint. Example: https://127.0.0.1:6443"
        echo "token\t\tset env var.value is the service account or other JWT token. Example: eyJhxx.x.x"
        echo "*\t\texecute kubectl comamnd with server and token"
        echo "help\t\tThis help banner"
        echo "=================================================================================================="
        return
    fi

    echo "KubeHack Plugin is waiting your order. use "kube help" to get usage"

    if [[ "$1" == "server" ]]
    then
        export KUBE_HACK_PLUGIN_VAR_SERVER=$2
        echo "successfully set KUBE_HACK_PLUGIN_VAR_SERVER as $KUBE_HACK_PLUGIN_VAR_SERVER"
        return
    fi

    if [[ "$1" == "token" ]]
    then
        export KUBE_HACK_PLUGIN_VAR_TOKEN=$2
        echo "successfully set KUBE_HACK_PLUGIN_VAR_TOKEN as $KUBE_HACK_PLUGIN_VAR_TOKEN"
        return
    fi

    if [[ -z "$KUBE_HACK_PLUGIN_VAR_SERVER" ]] || [[ -z "$KUBE_HACK_PLUGIN_VAR_TOKEN" ]]
    then
        echo "KUBE_HACK_PLUGIN_VAR_SERVER or KUBE_HACK_PLUGIN_VAR_TOKEN is not set. "
        echo "use k as kubectl alias"
        kubectl $@
        return
    fi
    echo "Executing command: kubectl --token=${KUBE_HACK_PLUGIN_VAR_TOKEN:0:7}...${KUBE_HACK_PLUGIN_VAR_TOKEN: -7} --server=$KUBE_HACK_PLUGIN_VAR_SERVER --insecure-skip-tls-verify=true $@"
    kubectl --token=$KUBE_HACK_PLUGIN_VAR_TOKEN --server=$KUBE_HACK_PLUGIN_VAR_SERVER --insecure-skip-tls-verify=true $@
    return
}
