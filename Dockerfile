FROM alpine
ADD ./settimeout /bin/settimeout
ADD assets /var/assets
CMD ["/bin/settimeout"]
EXPOSE 51004
