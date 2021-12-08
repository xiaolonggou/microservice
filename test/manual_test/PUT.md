Test 1: 
    Step 1: display list of art pieces
        Input:
            PS C:\Users\D058009\lawrence> curl 127.0.0.1:9090 | jq
        Output:
                [
                {
                    "id": 1,
                    "format": "oil on canvas",
                    "creator": "Vincent Van Gogh"
                },
                {
                    "id": 2,
                    "format": "song",
                    "creator": "Hildegard Knef"
                }
                ]
    Step 2: update artpiece with id "1", as given in the path
        Input: 
            Windows:
                curl 127.0.0.1:9090/1 -XPUT -d '{\"Id\": 35, \"format\":\"sketch on paper\", \"creator\":\"Lawrence\" }' | jq
            Mac:
                curl 127.0.0.1:9090/1 -XPUT -d '{"Id": 35, "format":"sketch on paper", "creator":"Lawrence", "Description":"like it a lot" }'
        Output: [server]
            handle http PUT ArtPiece request
        Input:
            curl 127.0.0.1:9090 | jq
        Output: [client]
            [
              {
                "id": 1,
                "format": "sketch on paper",
                "creator": "Lawrence"
              },
              {
                "id": 2,
                "format": "song",
                "creator": "Hildegard Knef"
              }
            ]