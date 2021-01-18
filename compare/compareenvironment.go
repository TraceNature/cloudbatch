package compare

import (
    "github.com/go-redis/redis/v7"
)

type CompoareEnvironment struct {
    Sclinet *redis.Client
    Tclient *redis.Client
}

func (compare *CompoareEnvironment) DiffParameters() map[string][]string {

    m := make(map[string][]string)
    sparameters := GetInstanceParameters(compare.Sclinet.Conn())
    tparameters := GetInstanceParameters(compare.Tclient.Conn())

    for sk, sv := range sparameters {
        val := []string{}
        if tparameters[sk] != sv {
            val = append(val, sv)
            val = append(val, tparameters[sk])
            m[sk] = val
        }
    }

    for tk, tv := range tparameters {
        val := []string{}
        if sparameters[tk] != tv {
            if len(m[tk]) == 0 {
                val = append(val, sparameters[tk])
                val = append(val, tv)
                m[tk] = val
            }
        }

    }

    return m
}

func GetInstanceParameters(conn *redis.Conn) map[string]string {
    parameters := GetAllParametersNames()
    m := make(map[string]string)
    for _, v := range parameters {
        var s string = ""
        value := conn.ConfigGet(v).Val()

        if len(value) != 0 {
            for i := 0; i < len(value); i++ {
                if i == 0 {
                    continue
                }
                if i == len(value)-1 {
                    s = s + value[i].(string)
                } else {
                    s = s + value[i].(string) + " "
                }
            }
            //fmt.Println(s)
            m[value[0].(string)] = s
        }
    }
    return m
}

func GetAllParametersNames() []string {
    return []string{"aclfile",
        "acllog-max-len",
        "activedefrag",
        "active-defrag-cycle-max",
        "active-defrag-cycle-min",
        "active-defrag-ignore-bytes",
        "active-defrag-max-scan-fields",
        "active-defrag-threshold-lower",
        "active-defrag-threshold-upper",
        "active-expire-effort",
        "activerehashing",
        "always-show-logo",
        "aof-load-truncated",
        "aof_rewrite_cpulist",
        "aof-rewrite-incremental-fsync",
        "aof-use-rdb-preamble",
        "appendfilename",
        "appendfsync",
        "appendonly",
        "auto-aof-rewrite-min-size",
        "auto-aof-rewrite-percentage",
        "bgsave_cpulist",
        "bind",
        "bio_cpulist",
        "client-output-buffer-limit",
        "client-query-buffer-limit",
        "cluster-allow-reads-when-down",
        "cluster-announce-bus-port",
        "cluster-announce-ip",
        "cluster-announce-port",
        "cluster-config-file",
        "cluster-enabled",
        "cluster-migration-barrier",
        "cluster-node-timeout",
        "cluster-replica-no-failover",
        "cluster-replica-validity-factor",
        "cluster-require-full-coverage",
        "cluster-slave-no-failover",
        "cluster-slave-validity-factor",
        "daemonize",
        "databases",
        "dbfilename",
        "dir",
        "dynamic-hz",
        "gopher-enabled",
        "hash-max-ziplist-entries",
        "hash-max-ziplist-value",
        "hll-sparse-max-bytes",
        "hz",
        "include",
        "io-threads",
        "io-threads-do-reads",
        "jemalloc-bg-thread",
        "latency-monitor-threshold",
        "lazyfree-lazy-eviction",
        "lazyfree-lazy-expire",
        "lazyfree-lazy-server-del",
        "lazyfree-lazy-user-del",
        "lfu-decay-time",
        "lfu-log-factor",
        "list-compress-depth",
        "list-max-ziplist-entries",
        "list-max-ziplist-size",
        "list-max-ziplist-value",
        "loadmodule",
        "logfile",
        "loglevel",
        "lua-time-limit",
        "masterauth",
        "masteruser",
        "maxclients",
        "maxmemory",
        "maxmemory-policy",
        "maxmemory-samples",
        "min-replicas-max-lag",
        "min-replicas-to-write",
        "min-slaves-max-lag",
        "min-slaves-to-write",
        "no-appendfsync-on-rewrite",
        "notify-keyspace-events",
        "oom-score-adj",
        "oom-score-adj-values",
        "otify-keyspace-events",
        "pidfile",
        "port",
        "protected-mode",
        "proto-max-bulk-len",
        "rdbchecksum",
        "rdbcompression",
        "rdb-del-sync-files",
        "rdb-save-incremental-fsync",
        "rename-command",
        "repl-backlog-size",
        "repl-backlog-ttl",
        "repl-disable-tcp-nodelay",
        "repl-diskless-load",
        "repl-diskless-sync",
        "repl-diskless-sync-delay",
        "replica-announce-ip",
        "replica-announce-port",
        "replica-ignore-maxmemory",
        "replica-lazy-flush",
        "replicaof",
        "replica-priority",
        "replica-read-only",
        "replica-serve-stale-data",
        "repl-ping-replica-period",
        "repl-ping-slave-period",
        "repl-timeout",
        "requirepass",
        "save",
        "server_cpulist",
        "set-max-intset-entries",
        "slave-announce-ip",
        "slave-announce-port",
        "slave-lazy-flush",
        "slaveof",
        "slave-priority",
        "slave-read-only",
        "slave-serve-stale-data",
        "slowlog-max-len",
        "stop-writes-on-bgsave-error",
        "stream-node-max-bytes",
        "stream-node-max-entries",
        "supervised",
        "syslog-enabled",
        "syslog-facility",
        "syslog-ident",
        "tcp-backlog",
        "tcp-keepalive",
        "timeout",
        "tls-auth-clients",
        "tls-ca-cert-dir",
        "tls-ca-cert-file",
        "tls-cert-file",
        "tls-ciphers",
        "tls-ciphersuites",
        "tls-cluster",
        "tls-dh-params-file",
        "tls-key-file",
        "tls-port",
        "tls-prefer-server-ciphers",
        "tls-protocols",
        "tls-replication",
        "tls-session-cache-size",
        "tls-session-cache-timeout",
        "tls-session-caching",
        "tracking-table-max-keys",
        "unixsocket",
        "unixsocketperm",
        "zset-max-ziplist-entries",
        "zset-max-ziplist-value"}

}
