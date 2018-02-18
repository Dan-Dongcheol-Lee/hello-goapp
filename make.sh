#!/bin/bash

workspace=$(pwd)
modulename="personmgmt"
parent_dir=$(dirname ${workspace})

export GOPATH=${workspace}:${workspace}/vendor
echo "GOPATH=$GOPATH"

deploy() {
    echo "* Deploying application to the project ($1)"
    gcloud -q app deploy --project ${1} ${workspace}/src/${modulename}/app.yaml
    echo
}

serve() {
    echo "* Serving for local test"
    goapp serve ${workspace}/src/${modulename}/app.yaml
}

test() {
    echo "Running unit tests..."
    (cd ${workspace}/src/${modulename} && go test -v ./...)
}

project_id="pc-dataflow-gdg"
command="help"
if [ "$#" -ge 1 ]; then
    command=$1
fi
if [ "$#" -ge 2 ]; then
    project_id=$2
fi

cat << EOF
-----------------------------------------------------
 * workspace: ${workspace}
 * command: ${command}
 * project_id: ${project_id}
-----------------------------------------------------
EOF

case ${command} in
    test|serve|deploy)
        ${command} ${project_id}
        ;;
    *)
        ;;
esac
