package app

import (
	"github.com/cybreem/sequis-test/app/domains/types"
	"github.com/h4lim/go-sdk/caching"
	"github.com/h4lim/go-sdk/logging"
)

var Log = logging.MustGetLogger("sequis-test")
var SimpleCache caching.SimpleCache
var HeaderContent types.HeaderContent
var Username string
var LogID string
var RoleID uint

const (
	SERVICE_NAME = "sequis-test"
	MESSAGE_EN   = "EN"
	MESSAGE_ID   = "ID"

	CACHE_LOG = "log_id"
)
