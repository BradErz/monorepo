FROM envoyproxy/envoy:v1.17.0

COPY services/envoy/envoy.yaml /etc/envoy/envoy.yaml

CMD /usr/local/bin/envoy -c /etc/envoy/envoy.yaml -l trace