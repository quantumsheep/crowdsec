#!/usr/bin/env bats
# vim: ft=bats:list:ts=8:sts=4:sw=4:et:ai:si:

set -u

setup_file() {
    load "../lib/setup_file.sh"
    ./instance-data load
    HUB_DIR=$(config_get '.config_paths.hub_dir')
    export HUB_DIR
    INDEX_PATH=$(config_get '.config_paths.index_path')
    export INDEX_PATH
    CONFIG_DIR=$(config_get '.config_paths.config_dir')
    export CONFIG_DIR
}

teardown_file() {
    load "../lib/teardown_file.sh"
}

setup() {
    load "../lib/setup.sh"
    load "../lib/bats-file/load.bash"
    ./instance-data load
    hub_strip_index
}

teardown() {
    ./instance-crowdsec stop
}

#----------
#
# Tests that don't need to be repeated for each hub type
#

@test "hub versions are correctly sorted during sync" {
    # hash of an empty file
    sha256_empty="e3b0c44298fc1c149afbf4c8996fb92427ae41e4649b934ca495991b7852b855"

    # add two versions with the same hash, that don't sort the same way
    # in a lexical vs semver sort. CrowdSec should report the latest version

    new_hub=$( \
        jq --arg DIGEST "$sha256_empty" <"$INDEX_PATH" \
        '. * {collections:{"crowdsecurity/sshd":{"versions":{"1.2":{"digest":$DIGEST, "deprecated": false}, "1.10": {"digest":$DIGEST, "deprecated": false}}}}}' \
    )
    echo "$new_hub" >"$INDEX_PATH"
 
    rune -0 cscli collections install crowdsecurity/sshd

    truncate -s 0 "$CONFIG_DIR/collections/sshd.yaml"

    rune -0 cscli collections inspect crowdsecurity/sshd -o json
    # XXX: is this supposed to be tainted or up to date?
    rune -0 jq -c '[.local_version,.up_to_date,.tainted]' <(output)
    assert_json '["1.10",false,false]'
}

@test "hub index with invalid (non semver) version numbers" {
    rune -0 cscli collections remove crowdsecurity/sshd --purge

    new_hub=$( \
        jq <"$INDEX_PATH" \
        '. * {collections:{"crowdsecurity/sshd":{"versions":{"1.2.3.4":{"digest":"foo", "deprecated": false}}}}}' \
    )
    echo "$new_hub" >"$INDEX_PATH"
 
    rune -0 cscli collections install crowdsecurity/sshd
    rune -1 cscli collections inspect crowdsecurity/sshd --no-metrics -o json
    # XXX: we are on the verbose side here...
    rune -0 jq -r ".msg" <(stderr)
    assert_output "failed to read Hub index: failed to sync items: failed to scan $CONFIG_DIR: while syncing collections sshd.yaml: 1.2.3.4: Invalid Semantic Version. Run 'sudo cscli hub update' to download the index again"
}
