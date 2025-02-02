from "ubuntu:latest"

after { flatten }

skip { run "apt update -qq" }
run "apt install ca-certificates -y"
copy "tinyci-#{getenv("VERSION")}.tar.gz", "/tinyci-release.tar.gz"
run "touch /tinyci-#{getenv("VERSION")}"
run "tar --no-same-owner --strip-components 1 -xvz -C /usr/local/bin -f /tinyci-release.tar.gz"
run "rm /tinyci-release.tar.gz"
