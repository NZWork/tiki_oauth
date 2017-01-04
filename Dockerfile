FROM golang:1.6-onbuild

CMD ["./tiki_oauth", "-c", "product"]