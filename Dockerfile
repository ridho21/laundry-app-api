FROM golang:alpine AS build
# Build Stage 
WORKDIR /app
# semua isi yang ada dalam project go-enigma-laundry akan disalin kedalam /app 
COPY . .

RUN go mod download
RUN go build -o go-enigma-laundry

# Final Stage 
FROM alpine 
WORKDIR /app
COPY --from=build /app/go-enigma-laundry /app/go-enigma-laundry

ENTRYPOINT ["/app/go-enigma-laundry"]