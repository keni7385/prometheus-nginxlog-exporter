#install builds dependencies

python -m venv venv/.
source ./venv/bin/activate
pip install -r venv/requirements.txt

GO111MODULE=off go get github.com/kisielk/errcheck
GO111MODULE=on go mod download
GO111MODULE=on go mod vendor
