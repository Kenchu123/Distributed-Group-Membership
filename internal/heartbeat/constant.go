package heartbeat

import "time"

var HEARTBEAT_PORT = "7140"
var HEARTBEAT_INTERVAL = 1000 * time.Millisecond
var FAILURE_DETECT_INTERVAL = 1000 * time.Millisecond
var FAILURE_DETECT_TIMEOUT = 3000 * time.Millisecond
var MEMBER_CLEANUP_INTERVAL = 1000 * time.Millisecond
var MEMBER_CLEANUP_TIMEOUT = 3000 * time.Millisecond
