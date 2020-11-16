#Issues faced during insertion

panic: elastic: Error 400 (Bad Request): Limit of total fields [1000] has been exceeded [type=illegal_argument_exception]
goroutine 1 [running]:

#Solution found
1) curl -XPUT 'localhost:9200/trading/_settings' -H 'Content-Type: application/json' -d' {"index" : {"mapping" : {"total_fields" : {"limit" : "100000"}}}}'

2)verfication: curl -XGET 'localhost:9200/trading/_settings?pretty'

