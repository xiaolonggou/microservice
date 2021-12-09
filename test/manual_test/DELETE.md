step:
    input:
        curl 127.0.0.1:9090 | jq
    output:
        [
        {
            "id": 1,
            "format": "oil on canvas",
            "creator": "Vincent Van Gogh",
            "price": 0,
            "description": "young fisherman on the beach"
        },
        {
            "id": 2,
            "format": "song",
            "creator": "Hildegard Knef",
            "price": 0,
            "description": "es soll für mich rote Rosen regnen"
        }
        ]
    input:
        curl 127.0.0.1:9090/art/1 -X DELETE | jq
    input
        curl 127.0.0.1:9090 | jq
    output:
        [
        {
            "id": 2,
            "format": "song",
            "creator": "Hildegard Knef",
            "price": 0,
            "description": "es soll für mich rote Rosen regnen"
        }
        ]