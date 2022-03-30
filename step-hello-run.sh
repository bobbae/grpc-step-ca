#!/bin/sh
set -x

echo 172.17.0.2 us-1 >> /etc/hosts

#CA_FINGERPRINT=$(step certificate fingerprint certs/root_ca.crt)
step ca root -f --ca-url us-1:9000 root_ca.crt --fingerprint $CA_FINGERPRINT
step ca bootstrap -f --ca-url us-1:9000 --fingerprint $CA_FINGERPRINT
step certificate install root_ca.crt

step ca certificate -f --ca-url us-1:9000 --token $CA_PASSWORD --san 172.17.0.3 --san helloserver localhost srv.crt srv.key
/app/hello-server
