#!/bin/bash
export PATH=/root/local/bin:$PATH
kubectl -n <replace-your-k8s-namespace> get pods | grep -q mmm-newdemo
if [ "$?" == "1" ];then
	kubectl create -f /tmp/dev/mmm-newdemo/mmm-newdemo.yaml --record
	kubectl -n <replace-your-k8s-namespace> get svc | grep -q mmm-newdemo
	if [ "$?" == "0" ];then
		echo "mmm-newdemo service install success!"
	else
		echo "mmm-newdemo service install fail!"
	fi
	kubectl -n <replace-your-k8s-namespace> get pods | grep -q mmm-newdemo
	if [ "$?" == "0" ];then
		echo "mmm-newdemo deployment install success!"
	else
		echo "mmm-newdemo deployment install fail!"
	fi
else
	kubectl delete -f /tmp/dev/mmm-newdemo/mmm-newdemo.yaml
	kubectl -n <replace-your-k8s-namespace> get svc | grep -q mmm-newdemo
	while [ "$?" == "0" ]
	do
	kubectl -n <replace-your-k8s-namespace> get svc | grep -q mmm-newdemo
	done
	kubectl -n <replace-your-k8s-namespace> get pods | grep -q mmm-newdemo
	while [ "$?" == "0" ]
	do
	kubectl -n <replace-your-k8s-namespace> get pods | grep -q mmm-newdemo
	done
	kubectl create -f /tmp/dev/mmm-newdemo/mmm-newdemo.yaml --record
	kubectl -n <replace-your-k8s-namespace> get svc | grep -q mmm-newdemo
	if [ "$?" == "0" ];then
		echo "mmm-newdemo service update success!"
	else
		echo "mmm-newdemo service update fail!"
	fi
	kubectl -n <replace-your-k8s-namespace> get pods | grep -q mmm-newdemo
	if [ "$?" == "0" ];then
		echo "mmm-newdemo deployment update success!"
	else
		echo "mmm-newdemo deployment update fail!"
	fi
fi