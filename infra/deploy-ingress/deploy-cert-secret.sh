# delete secret
kubectl delete secret aks-ingress-tls

# create tls secret
kubectl create secret tls aks-ingress-tls \
--key ./deploy-ingress/certs/guestbook.kainiindustries.net.private.key \
--cert ./deploy-ingress/certs/guestbook.kainiindustries.net.certificate.crt 