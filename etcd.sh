#!/bin/sh
tip=http://$(docker inspect --format='{{range .NetworkSettings.Networks}}{{.IPAddress}}{{end}}' rapis_etcd_1):2379
echo $tip
entry_1=http://$(docker inspect --format='{{range .NetworkSettings.Networks}}{{.IPAddress}}{{end}}' rapis_test1_1):80
entry_2=http://$(docker inspect --format='{{range .NetworkSettings.Networks}}{{.IPAddress}}{{end}}' rapis_test2_1):80

echo $entry_1
echo $entry_2

curl -XPUT $tip/v2/keys/traefik/backends/backend1/circuitbreaker/expression -d value='NetworkErrorRatio() > 0.5'

curl -XPUT $tip/v2/keys/traefik/backends/backend1/servers/server1/url -d value=$entry_1
curl -XPUT $tip/v2/keys/traefik/backends/backend1/servers/server2/url -d value=$entry_2

curl -XPUT $tip/v2/keys/traefik/backends/backend1/servers/server1/weight -d value='4'
curl -XPUT $tip/v2/keys/traefik/backends/backend1/servers/server2/weight -d value='6'

curl -XPUT $tip/v2/keys/traefik/frontends/frontend1/backend -d value="backend1"
curl -XPUT $tip/v2/keys/traefik/frontends/frontend1/routes/test_1/rule -d value="Host:test.localhost"
curl -XPUT $tip/v2/keys/traefik/frontends/frontend1/entrypoints -d value="http"


