                                                                                                       
                                                                                                         
---------------------------------------------------------------------------------------------------------
-> Elasticsearch security features have been automatically configured!                                   
-> Authentication is enabled and cluster connections are encrypted.                                                                                                                                               
->  Password for the elastic user (reset with `bin/elasticsearch-reset-password -u elastic`):            
  Xef5xd846u2XOlygMlCH                                                                                   
                                                                                                         
->  HTTP CA certificate SHA-256 fingerprint:                                                             
  8baafea947a2002112806ffc35497007a4228959f3eb5ff03342871fd5f794ed


-> Configure other nodes to join this cluster:
* Copy the following enrollment token and start new Elasticsearch nodes with `bin/elasticsearch --enrollm
ent-token <token>` 

                    
  If you're running in Docker, copy the enrollment token and run:
  `docker run -e "ENROLLMENT_TOKEN=<token>" docker.elastic.co/elasticsearch/elasticsearch:8.1.1`
---------------------------------------------------------------------------------------------------------
       