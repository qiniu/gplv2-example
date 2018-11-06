#!/usr/bin/expect
set timeout 5
spawn go run hello.go
set expect_str "Hello, 世界!"
expect {
    default {
        send_user "✘ Output didn't match, expect '$expect_str'\n"
        exit 1
    }
    $expect_str
}

spawn go run destroy_earth.go
set expect_str "地球已被二向箔毁灭。"
expect {
    default {
        send_user "✘ Output didn't match, expect '$expect_str'\n"
        exit 1
    }
    $expect_str
}
exit
