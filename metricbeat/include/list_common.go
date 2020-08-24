// Licensed to Elasticsearch B.V. under one or more contributor
// license agreements. See the NOTICE file distributed with
// this work for additional information regarding copyright
// ownership. Elasticsearch B.V. licenses this file to you under
// the Apache License, Version 2.0 (the "License"); you may
// not use this file except in compliance with the License.
// You may obtain a copy of the License at
//
//     http://www.apache.org/licenses/LICENSE-2.0
//
// Unless required by applicable law or agreed to in writing,
// software distributed under the License is distributed on an
// "AS IS" BASIS, WITHOUT WARRANTIES OR CONDITIONS OF ANY
// KIND, either express or implied.  See the License for the
// specific language governing permissions and limitations
// under the License.

// Code generated by beats/dev-tools/cmd/module_include_list/module_include_list.go - DO NOT EDIT.

package include

import (
	// Import packages that need to register themselves.
	_ "github.com/snappyflow/beats/v7/metricbeat/module/aerospike"
	_ "github.com/snappyflow/beats/v7/metricbeat/module/aerospike/namespace"
	_ "github.com/snappyflow/beats/v7/metricbeat/module/apache"
	_ "github.com/snappyflow/beats/v7/metricbeat/module/apache/status"
	_ "github.com/snappyflow/beats/v7/metricbeat/module/beat"
	_ "github.com/snappyflow/beats/v7/metricbeat/module/beat/state"
	_ "github.com/snappyflow/beats/v7/metricbeat/module/beat/stats"
	_ "github.com/snappyflow/beats/v7/metricbeat/module/ceph"
	_ "github.com/snappyflow/beats/v7/metricbeat/module/ceph/cluster_disk"
	_ "github.com/snappyflow/beats/v7/metricbeat/module/ceph/cluster_health"
	_ "github.com/snappyflow/beats/v7/metricbeat/module/ceph/cluster_status"
	_ "github.com/snappyflow/beats/v7/metricbeat/module/ceph/mgr_cluster_disk"
	_ "github.com/snappyflow/beats/v7/metricbeat/module/ceph/mgr_cluster_health"
	_ "github.com/snappyflow/beats/v7/metricbeat/module/ceph/mgr_osd_perf"
	_ "github.com/snappyflow/beats/v7/metricbeat/module/ceph/mgr_osd_pool_stats"
	_ "github.com/snappyflow/beats/v7/metricbeat/module/ceph/mgr_osd_tree"
	_ "github.com/snappyflow/beats/v7/metricbeat/module/ceph/mgr_pool_disk"
	_ "github.com/snappyflow/beats/v7/metricbeat/module/ceph/monitor_health"
	_ "github.com/snappyflow/beats/v7/metricbeat/module/ceph/osd_df"
	_ "github.com/snappyflow/beats/v7/metricbeat/module/ceph/osd_tree"
	_ "github.com/snappyflow/beats/v7/metricbeat/module/ceph/pool_disk"
	_ "github.com/snappyflow/beats/v7/metricbeat/module/consul"
	_ "github.com/snappyflow/beats/v7/metricbeat/module/consul/agent"
	_ "github.com/snappyflow/beats/v7/metricbeat/module/couchbase"
	_ "github.com/snappyflow/beats/v7/metricbeat/module/couchbase/bucket"
	_ "github.com/snappyflow/beats/v7/metricbeat/module/couchbase/cluster"
	_ "github.com/snappyflow/beats/v7/metricbeat/module/couchbase/node"
	_ "github.com/snappyflow/beats/v7/metricbeat/module/couchdb"
	_ "github.com/snappyflow/beats/v7/metricbeat/module/couchdb/server"
	_ "github.com/snappyflow/beats/v7/metricbeat/module/dropwizard"
	_ "github.com/snappyflow/beats/v7/metricbeat/module/dropwizard/collector"
	_ "github.com/snappyflow/beats/v7/metricbeat/module/elasticsearch"
	_ "github.com/snappyflow/beats/v7/metricbeat/module/elasticsearch/ccr"
	_ "github.com/snappyflow/beats/v7/metricbeat/module/elasticsearch/cluster_stats"
	_ "github.com/snappyflow/beats/v7/metricbeat/module/elasticsearch/enrich"
	_ "github.com/snappyflow/beats/v7/metricbeat/module/elasticsearch/index"
	_ "github.com/snappyflow/beats/v7/metricbeat/module/elasticsearch/index_recovery"
	_ "github.com/snappyflow/beats/v7/metricbeat/module/elasticsearch/index_summary"
	_ "github.com/snappyflow/beats/v7/metricbeat/module/elasticsearch/ml_job"
	_ "github.com/snappyflow/beats/v7/metricbeat/module/elasticsearch/node"
	_ "github.com/snappyflow/beats/v7/metricbeat/module/elasticsearch/node_stats"
	_ "github.com/snappyflow/beats/v7/metricbeat/module/elasticsearch/pending_tasks"
	_ "github.com/snappyflow/beats/v7/metricbeat/module/elasticsearch/shard"
	_ "github.com/snappyflow/beats/v7/metricbeat/module/envoyproxy"
	_ "github.com/snappyflow/beats/v7/metricbeat/module/envoyproxy/server"
	_ "github.com/snappyflow/beats/v7/metricbeat/module/etcd"
	_ "github.com/snappyflow/beats/v7/metricbeat/module/etcd/leader"
	_ "github.com/snappyflow/beats/v7/metricbeat/module/etcd/metrics"
	_ "github.com/snappyflow/beats/v7/metricbeat/module/etcd/self"
	_ "github.com/snappyflow/beats/v7/metricbeat/module/etcd/store"
	_ "github.com/snappyflow/beats/v7/metricbeat/module/golang"
	_ "github.com/snappyflow/beats/v7/metricbeat/module/golang/expvar"
	_ "github.com/snappyflow/beats/v7/metricbeat/module/golang/heap"
	_ "github.com/snappyflow/beats/v7/metricbeat/module/graphite"
	_ "github.com/snappyflow/beats/v7/metricbeat/module/graphite/server"
	_ "github.com/snappyflow/beats/v7/metricbeat/module/haproxy"
	_ "github.com/snappyflow/beats/v7/metricbeat/module/haproxy/info"
	_ "github.com/snappyflow/beats/v7/metricbeat/module/haproxy/stat"
	_ "github.com/snappyflow/beats/v7/metricbeat/module/http"
	_ "github.com/snappyflow/beats/v7/metricbeat/module/http/json"
	_ "github.com/snappyflow/beats/v7/metricbeat/module/http/server"
	_ "github.com/snappyflow/beats/v7/metricbeat/module/jolokia"
	_ "github.com/snappyflow/beats/v7/metricbeat/module/jolokia/jmx"
	_ "github.com/snappyflow/beats/v7/metricbeat/module/kafka"
	_ "github.com/snappyflow/beats/v7/metricbeat/module/kafka/consumergroup"
	_ "github.com/snappyflow/beats/v7/metricbeat/module/kafka/partition"
	_ "github.com/snappyflow/beats/v7/metricbeat/module/kibana"
	_ "github.com/snappyflow/beats/v7/metricbeat/module/kibana/stats"
	_ "github.com/snappyflow/beats/v7/metricbeat/module/kibana/status"
	_ "github.com/snappyflow/beats/v7/metricbeat/module/kvm"
	_ "github.com/snappyflow/beats/v7/metricbeat/module/kvm/dommemstat"
	_ "github.com/snappyflow/beats/v7/metricbeat/module/kvm/status"
	_ "github.com/snappyflow/beats/v7/metricbeat/module/linux"
	_ "github.com/snappyflow/beats/v7/metricbeat/module/linux/conntrack"
	_ "github.com/snappyflow/beats/v7/metricbeat/module/linux/ksm"
	_ "github.com/snappyflow/beats/v7/metricbeat/module/linux/pageinfo"
	_ "github.com/snappyflow/beats/v7/metricbeat/module/logstash"
	_ "github.com/snappyflow/beats/v7/metricbeat/module/logstash/node"
	_ "github.com/snappyflow/beats/v7/metricbeat/module/logstash/node_stats"
	_ "github.com/snappyflow/beats/v7/metricbeat/module/memcached"
	_ "github.com/snappyflow/beats/v7/metricbeat/module/memcached/stats"
	_ "github.com/snappyflow/beats/v7/metricbeat/module/mongodb"
	_ "github.com/snappyflow/beats/v7/metricbeat/module/mongodb/collstats"
	_ "github.com/snappyflow/beats/v7/metricbeat/module/mongodb/dbstats"
	_ "github.com/snappyflow/beats/v7/metricbeat/module/mongodb/metrics"
	_ "github.com/snappyflow/beats/v7/metricbeat/module/mongodb/replstatus"
	_ "github.com/snappyflow/beats/v7/metricbeat/module/mongodb/status"
	_ "github.com/snappyflow/beats/v7/metricbeat/module/munin"
	_ "github.com/snappyflow/beats/v7/metricbeat/module/munin/node"
	_ "github.com/snappyflow/beats/v7/metricbeat/module/mysql"
	_ "github.com/snappyflow/beats/v7/metricbeat/module/mysql/galera_status"
	_ "github.com/snappyflow/beats/v7/metricbeat/module/mysql/query"
	_ "github.com/snappyflow/beats/v7/metricbeat/module/mysql/status"
	_ "github.com/snappyflow/beats/v7/metricbeat/module/nats"
	_ "github.com/snappyflow/beats/v7/metricbeat/module/nats/connections"
	_ "github.com/snappyflow/beats/v7/metricbeat/module/nats/routes"
	_ "github.com/snappyflow/beats/v7/metricbeat/module/nats/stats"
	_ "github.com/snappyflow/beats/v7/metricbeat/module/nats/subscriptions"
	_ "github.com/snappyflow/beats/v7/metricbeat/module/nginx"
	_ "github.com/snappyflow/beats/v7/metricbeat/module/nginx/stubstatus"
	_ "github.com/snappyflow/beats/v7/metricbeat/module/php_fpm"
	_ "github.com/snappyflow/beats/v7/metricbeat/module/php_fpm/pool"
	_ "github.com/snappyflow/beats/v7/metricbeat/module/php_fpm/process"
	_ "github.com/snappyflow/beats/v7/metricbeat/module/postgresql"
	_ "github.com/snappyflow/beats/v7/metricbeat/module/postgresql/activity"
	_ "github.com/snappyflow/beats/v7/metricbeat/module/postgresql/bgwriter"
	_ "github.com/snappyflow/beats/v7/metricbeat/module/postgresql/database"
	_ "github.com/snappyflow/beats/v7/metricbeat/module/postgresql/statement"
	_ "github.com/snappyflow/beats/v7/metricbeat/module/prometheus"
	_ "github.com/snappyflow/beats/v7/metricbeat/module/prometheus/collector"
	_ "github.com/snappyflow/beats/v7/metricbeat/module/prometheus/query"
	_ "github.com/snappyflow/beats/v7/metricbeat/module/prometheus/remote_write"
	_ "github.com/snappyflow/beats/v7/metricbeat/module/rabbitmq"
	_ "github.com/snappyflow/beats/v7/metricbeat/module/rabbitmq/connection"
	_ "github.com/snappyflow/beats/v7/metricbeat/module/rabbitmq/exchange"
	_ "github.com/snappyflow/beats/v7/metricbeat/module/rabbitmq/node"
	_ "github.com/snappyflow/beats/v7/metricbeat/module/rabbitmq/queue"
	_ "github.com/snappyflow/beats/v7/metricbeat/module/redis"
	_ "github.com/snappyflow/beats/v7/metricbeat/module/redis/info"
	_ "github.com/snappyflow/beats/v7/metricbeat/module/redis/key"
	_ "github.com/snappyflow/beats/v7/metricbeat/module/redis/keyspace"
	_ "github.com/snappyflow/beats/v7/metricbeat/module/system"
	_ "github.com/snappyflow/beats/v7/metricbeat/module/system/core"
	_ "github.com/snappyflow/beats/v7/metricbeat/module/system/cpu"
	_ "github.com/snappyflow/beats/v7/metricbeat/module/system/diskio"
	_ "github.com/snappyflow/beats/v7/metricbeat/module/system/entropy"
	_ "github.com/snappyflow/beats/v7/metricbeat/module/system/filesystem"
	_ "github.com/snappyflow/beats/v7/metricbeat/module/system/fsstat"
	_ "github.com/snappyflow/beats/v7/metricbeat/module/system/load"
	_ "github.com/snappyflow/beats/v7/metricbeat/module/system/memory"
	_ "github.com/snappyflow/beats/v7/metricbeat/module/system/network"
	_ "github.com/snappyflow/beats/v7/metricbeat/module/system/network_summary"
	_ "github.com/snappyflow/beats/v7/metricbeat/module/system/process"
	_ "github.com/snappyflow/beats/v7/metricbeat/module/system/process_summary"
	_ "github.com/snappyflow/beats/v7/metricbeat/module/system/raid"
	_ "github.com/snappyflow/beats/v7/metricbeat/module/system/service"
	_ "github.com/snappyflow/beats/v7/metricbeat/module/system/socket"
	_ "github.com/snappyflow/beats/v7/metricbeat/module/system/socket_summary"
	_ "github.com/snappyflow/beats/v7/metricbeat/module/system/uptime"
	_ "github.com/snappyflow/beats/v7/metricbeat/module/system/users"
	_ "github.com/snappyflow/beats/v7/metricbeat/module/traefik"
	_ "github.com/snappyflow/beats/v7/metricbeat/module/traefik/health"
	_ "github.com/snappyflow/beats/v7/metricbeat/module/uwsgi"
	_ "github.com/snappyflow/beats/v7/metricbeat/module/uwsgi/status"
	_ "github.com/snappyflow/beats/v7/metricbeat/module/vsphere"
	_ "github.com/snappyflow/beats/v7/metricbeat/module/vsphere/datastore"
	_ "github.com/snappyflow/beats/v7/metricbeat/module/vsphere/host"
	_ "github.com/snappyflow/beats/v7/metricbeat/module/vsphere/virtualmachine"
	_ "github.com/snappyflow/beats/v7/metricbeat/module/windows"
	_ "github.com/snappyflow/beats/v7/metricbeat/module/windows/perfmon"
	_ "github.com/snappyflow/beats/v7/metricbeat/module/windows/service"
	_ "github.com/snappyflow/beats/v7/metricbeat/module/zookeeper"
	_ "github.com/snappyflow/beats/v7/metricbeat/module/zookeeper/connection"
	_ "github.com/snappyflow/beats/v7/metricbeat/module/zookeeper/mntr"
	_ "github.com/snappyflow/beats/v7/metricbeat/module/zookeeper/server"
)
