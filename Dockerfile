FROM quay.io/coreos/etcd:v3.5.2

RUN \
  -p 2379:2379 \
  -p 2380:2380 \
  --mount type=bind,source=/tmp/etcd-data.tmp,destination=/etcd-data \ #将文件挂载到容器中
  --name quay-etcd0 \   #容器名字
  /usr/local/bin/etcd \
  --name etcd0 \    #etcd服务名
  --data-dir /etcd-data \
  --listen-client-urls http://0.0.0.0:2379 \
  --advertise-client-urls http://106.53.97.81:2379 \  #对客户端提供服务的url 公网ip：port
  --listen-peer-urls http://0.0.0.0:2380 \..
  --initial-advertise-peer-urls http://10.0.12.12:2380 \   #对集群提供服务的url
  --initial-cluster etcd0=http://10.0.12.12:2380, http://10.0.12.12:12380, http://10.0.12.12:22380\  #用于引导的初始集群配置。
  --initial-cluster-token tkn \
  --initial-cluster-state new \
  --log-level info \
  --logger zap \
  --log-outputs stderr

FROM quay.io/coreos/etcd:v3.5.2

RUN \
  -p 2379:12379 \
  -p 2380:12380 \
  --mount type=bind,source=/tmp/etcd-data.tmp,destination=/etcd-data \ #将文件挂载到容器中
  --name quay-etcd1 \   #容器名字
  /usr/local/bin/etcd \
  --name etcd1 \    #etcd服务名
  --data-dir /etcd-data \
  --listen-client-urls http://0.0.0.0:2379 \
  --advertise-client-urls http://106.53.97.81:12379 \  #对客户端提供服务的url 公网ip：port
  --listen-peer-urls http://0.0.0.0:2380 \..
  --initial-advertise-peer-urls http://10.0.12.12:12380 \   #对集群提供服务的url
  --initial-cluster etcd0=http://10.0.12.12:2380, http://10.0.12.12:12380, http://10.0.12.12:22380\  #用于引导的初始集群配置。
  --initial-cluster-token tkn \
  --initial-cluster-state new \
  --log-level info \
  --logger zap \
  --log-outputs stderr

FROM quay.io/coreos/etcd:v3.5.2

RUN \
  -p 2379:22379 \
  -p 2380:22380 \
  --mount type=bind,source=/tmp/etcd-data.tmp,destination=/etcd-data \ #将文件挂载到容器中
  --name quay-etcd2 \   #容器名字
  /usr/local/bin/etcd \
  --name etcd2 \    #etcd服务名
  --data-dir /etcd-data \
  --listen-client-urls http://0.0.0.0:2379 \
  --advertise-client-urls http://106.53.97.81:22379 \  #对客户端提供服务的url 公网ip：port
  --listen-peer-urls http://0.0.0.0:2380 \..
  --initial-advertise-peer-urls http://10.0.12.12:22380 \   #对集群提供服务的url
  --initial-cluster etcd0=http://10.0.12.12:2380, http://10.0.12.12:12380, http://10.0.12.12:22380\  #用于引导的初始集群配置。
  --initial-cluster-token tkn \
  --initial-cluster-state new \
  --log-level info \
  --logger zap \
  --log-outputs stderr