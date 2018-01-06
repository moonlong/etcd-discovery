#!/bin/sh

work_path=$(dirname $0)
cd ./${work_path}

#echo $(pwd)
#echo `date`



case $1 in

1)
echo -e "[1]start first server\n"
./etcd --name cd0 --initial-advertise-peer-urls http://127.0.0.1:2380 \
  --listen-peer-urls http://127.0.0.1:2380 \
  --listen-client-urls http://192.168.1.17:2379,http://127.0.0.1:2379 \
  --advertise-client-urls http://192.168.1.17:2379,http://127.0.0.1:2379 \
  --initial-cluster-token etcd-cluster-1 \
  --initial-cluster cd0=http://127.0.0.1:2380,cd1=http://127.0.0.1:2480,cd2=http://127.0.0.1:2580 \
  --initial-cluster-state new
  ;;
2)
echo -e "[2]start second  server\n"
./etcd --name cd1 --initial-advertise-peer-urls http://127.0.0.1:2480 \
  --listen-peer-urls http://127.0.0.1:2480 \
  --listen-client-urls http://192.168.1.17:2479,http://127.0.0.1:2479 \
  --advertise-client-urls http://192.168.1.17:2479,http://127.0.0.1:2479 \
  --initial-cluster-token etcd-cluster-1 \
  --initial-cluster cd0=http://127.0.0.1:2380,cd1=http://127.0.0.1:2480,cd2=http://127.0.0.1:2580 \
  --initial-cluster-state new
  ;;
3)
echo -e "[3]start third server\n"
./etcd --name cd2 --initial-advertise-peer-urls http://127.0.0.1:2580 \
  --listen-peer-urls http://127.0.0.1:2580 \
  --listen-client-urls http://192.168.1.17:2579,http://127.0.0.1:2579 \
  --advertise-client-urls http://192.168.1.17:2579,http://127.0.0.1:2579 \
  --initial-cluster-token etcd-cluster-1 \
  --initial-cluster cd0=http://127.0.0.1:2380,cd1=http://127.0.0.1:2480,cd2=http://127.0.0.1:2580 \
  --initial-cluster-state new
  ;;
*)
echo "error paramater"
;;
esac
