Test 1: 
    Step 1: add a piece of art
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

        Input:
            curl.exe 127.0.0.1:9090 -X POST -d '{\"creator\":\"Conner\"}'  
        Input:
            curl.exe 127.0.0.1:9090 | jq   
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
              },
              {
                "id": 3,
                "format": "",
                "creator": "Conner"
              }
            ]