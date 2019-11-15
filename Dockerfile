FROM scratch
COPY leetcode-cli /
ENTRYPOINT ["/leetcode-cli"]