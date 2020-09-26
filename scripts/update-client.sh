#!/usr/bin/env bash

oapi-codegen -generate types,client scaleway/client/scaleway-dns.yml > scaleway/client/client.go
