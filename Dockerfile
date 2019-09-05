FROM ubuntu:16.04
ADD rpc_srv /root/rpc_srv
ADD rpc_check /root/rpc_check
RUN chmod +x /root/rpc_srv  && chmod +x /root/rpc_check
EXPOSE 19000
CMD /root/rpc_srv -p 19000
