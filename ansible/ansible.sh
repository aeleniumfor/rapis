#!/bin/bash
TOKEN=$(curl https://discovery.etcd.io/new)
echo "" > /root/.ssh/known_hosts
export ANSIBLE_INVENTORY=$(pwd)/hosts
export ANSIBLE_HOST_KEY_CHECKING=False
ansible master -i ./hosts -m ping
ansible-playbook -i ./hosts ./playbook.yml --extra-vars TOKEN=$TOKEN
